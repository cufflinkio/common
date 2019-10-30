// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"context"
	"math/big"
	"time"

	"github.com/zeebo/errs"

	"storj.io/storj/satellite/payments/coinpayments"
	"storj.io/storj/satellite/payments/stripecoinpayments"
	dbx "storj.io/storj/satellite/satellitedb/dbx"
)

// ensure that coinpaymentsTransaction implements stripecoinpayments.TransactionsDB.
var _ stripecoinpayments.TransactionsDB = (*coinpaymentsTransactions)(nil)

// applyBalanceIntentState defines states of the apply balance intents.
type applyBalanceIntentState int

const (
	// apply balance intent waits to be applied.
	applyBalanceIntentStateUnapplied applyBalanceIntentState = 0
	// transaction which balance intent points to has been consumed.
	applyBalanceIntentStateConsumed applyBalanceIntentState = 1
)

// Int returns intent state as int.
func (intent applyBalanceIntentState) Int() int {
	return int(intent)
}

// coinpaymentsTransactions is Coinpayments transactions DB.
//
// architecture: Database
type coinpaymentsTransactions struct {
	db *dbx.DB
}

// Insert inserts new coinpayments transaction into DB.
func (db *coinpaymentsTransactions) Insert(ctx context.Context, tx stripecoinpayments.Transaction) (*stripecoinpayments.Transaction, error) {
	amount, err := tx.Amount.GobEncode()
	if err != nil {
		return nil, errs.Wrap(err)
	}
	received, err := tx.Received.GobEncode()
	if err != nil {
		return nil, errs.Wrap(err)
	}

	dbxCPTX, err := db.db.Create_CoinpaymentsTransaction(ctx,
		dbx.CoinpaymentsTransaction_Id(tx.ID.String()),
		dbx.CoinpaymentsTransaction_UserId(tx.AccountID[:]),
		dbx.CoinpaymentsTransaction_Address(tx.Address),
		dbx.CoinpaymentsTransaction_Amount(amount),
		dbx.CoinpaymentsTransaction_Received(received),
		dbx.CoinpaymentsTransaction_Status(tx.Status.Int()),
		dbx.CoinpaymentsTransaction_Key(tx.Key),
	)
	if err != nil {
		return nil, err
	}

	return fromDBXCoinpaymentsTransaction(dbxCPTX)
}

// Update updates status and received for set of transactions.
func (db *coinpaymentsTransactions) Update(ctx context.Context, updates []stripecoinpayments.TransactionUpdate, applies coinpayments.TransactionIDList) error {
	if len(updates) == 0 {
		return nil
	}

	return db.db.WithTx(ctx, func(ctx context.Context, tx *dbx.Tx) error {
		for _, update := range updates {
			received, err := update.Received.GobEncode()
			if err != nil {
				return errs.Wrap(err)
			}

			_, err = tx.Update_CoinpaymentsTransaction_By_Id(ctx,
				dbx.CoinpaymentsTransaction_Id(update.TransactionID.String()),
				dbx.CoinpaymentsTransaction_Update_Fields{
					Received: dbx.CoinpaymentsTransaction_Received(received),
					Status:   dbx.CoinpaymentsTransaction_Status(update.Status.Int()),
				},
			)
			if err != nil {
				return err
			}
		}

		for _, txID := range applies {
			_, err := tx.Create_StripecoinpaymentsApplyBalanceIntent(ctx,
				dbx.StripecoinpaymentsApplyBalanceIntent_TxId(txID.String()),
				dbx.StripecoinpaymentsApplyBalanceIntent_State(applyBalanceIntentStateUnapplied.Int()))
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// Consume marks transaction as consumed, so it won't participate in apply account balance loop.
func (db *coinpaymentsTransactions) Consume(ctx context.Context, id coinpayments.TransactionID) error {
	_, err := db.db.Update_StripecoinpaymentsApplyBalanceIntent_By_TxId(ctx,
		dbx.StripecoinpaymentsApplyBalanceIntent_TxId(id.String()),
		dbx.StripecoinpaymentsApplyBalanceIntent_Update_Fields{
			State: dbx.StripecoinpaymentsApplyBalanceIntent_State(applyBalanceIntentStateConsumed.Int()),
		},
	)
	return err
}

// ListPending returns paginated list of pending transactions.
func (db *coinpaymentsTransactions) ListPending(ctx context.Context, offset int64, limit int, before time.Time) (stripecoinpayments.TransactionsPage, error) {
	var page stripecoinpayments.TransactionsPage

	dbxTXs, err := db.db.Limited_CoinpaymentsTransaction_By_CreatedAt_LessOrEqual_And_Status_OrderBy_Desc_CreatedAt(
		ctx,
		dbx.CoinpaymentsTransaction_CreatedAt(before.UTC()),
		dbx.CoinpaymentsTransaction_Status(coinpayments.StatusPending.Int()),
		limit+1,
		offset,
	)
	if err != nil {
		return stripecoinpayments.TransactionsPage{}, err
	}

	if len(dbxTXs) == limit+1 {
		page.Next = true
		page.NextOffset = offset + int64(limit) + 1

		dbxTXs = dbxTXs[:len(dbxTXs)-1]
	}

	var txs []stripecoinpayments.Transaction
	for _, dbxTX := range dbxTXs {
		tx, err := fromDBXCoinpaymentsTransaction(dbxTX)
		if err != nil {
			return stripecoinpayments.TransactionsPage{}, err
		}

		txs = append(txs, *tx)
	}

	page.Transactions = txs
	return page, nil
}

// List Unapplied returns TransactionsPage with transactions completed transaction that should be applied to account balance.
func (db *coinpaymentsTransactions) ListUnapplied(ctx context.Context, offset int64, limit int, before time.Time) (_ stripecoinpayments.TransactionsPage, err error) {
	query := db.db.Rebind(`SELECT 
				txs.id,
				txs.user_id,
				txs.address,
				txs.amount,
				txs.received,
				txs.key,
				txs.created_at
			FROM coinpayments_transactions as txs 
			INNER JOIN stripecoinpayments_apply_balance_intents as ints
			ON txs.id = ints.tx_id
			WHERE txs.status = ?
			AND txs.created_at <= ?
			AND ints.state = ?
			ORDER by txs.created_at DESC
			LIMIT ? OFFSET ?`)

	rows, err := db.db.QueryContext(ctx, query, coinpayments.StatusReceived, before, applyBalanceIntentStateUnapplied, limit+1, offset)
	if err != nil {
		return stripecoinpayments.TransactionsPage{}, err
	}

	defer func() {
		err = errs.Combine(err, rows.Close())
	}()

	var page stripecoinpayments.TransactionsPage

	for rows.Next() {
		var id, address string
		var userIDB []byte
		var amountB, receivedB []byte
		var key string
		var createdAt time.Time

		err := rows.Scan(&id, &userIDB, &address, &amountB, &receivedB, &key, &createdAt)
		if err != nil {
			return stripecoinpayments.TransactionsPage{}, err
		}

		userID, err := bytesToUUID(userIDB)
		if err != nil {
			return stripecoinpayments.TransactionsPage{}, errs.Wrap(err)
		}

		var amount, received big.Float
		if err := amount.GobDecode(amountB); err != nil {
			return stripecoinpayments.TransactionsPage{}, errs.Wrap(err)
		}
		if err := received.GobDecode(receivedB); err != nil {
			return stripecoinpayments.TransactionsPage{}, errs.Wrap(err)
		}

		page.Transactions = append(page.Transactions,
			stripecoinpayments.Transaction{
				ID:        coinpayments.TransactionID(id),
				AccountID: userID,
				Address:   address,
				Amount:    amount,
				Received:  received,
				Status:    coinpayments.StatusReceived,
				Key:       key,
				CreatedAt: createdAt,
			},
		)
	}

	if err = rows.Err(); err != nil {
		return stripecoinpayments.TransactionsPage{}, err
	}

	if len(page.Transactions) == limit+1 {
		page.Next = true
		page.NextOffset = offset + int64(limit) + 1
		page.Transactions = page.Transactions[:len(page.Transactions)-1]
	}

	return page, nil
}

// fromDBXCoinpaymentsTransaction converts *dbx.CoinpaymentsTransaction to *stripecoinpayments.Transaction.
func fromDBXCoinpaymentsTransaction(dbxCPTX *dbx.CoinpaymentsTransaction) (*stripecoinpayments.Transaction, error) {
	userID, err := bytesToUUID(dbxCPTX.UserId)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	var amount, received big.Float
	if err := amount.GobDecode(dbxCPTX.Amount); err != nil {
		return nil, errs.Wrap(err)
	}
	if err := received.GobDecode(dbxCPTX.Received); err != nil {
		return nil, errs.Wrap(err)
	}

	return &stripecoinpayments.Transaction{
		ID:        coinpayments.TransactionID(dbxCPTX.Id),
		AccountID: userID,
		Address:   dbxCPTX.Address,
		Amount:    amount,
		Received:  received,
		Status:    coinpayments.Status(dbxCPTX.Status),
		Key:       dbxCPTX.Key,
		CreatedAt: dbxCPTX.CreatedAt,
	}, nil
}