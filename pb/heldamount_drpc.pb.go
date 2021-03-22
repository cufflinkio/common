// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.18
// source: heldamount.proto

package pb

import (
	bytes "bytes"
	context "context"
	errors "errors"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"

	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_heldamount_proto struct{}

func (drpcEncoding_File_heldamount_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_heldamount_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_heldamount_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	var buf bytes.Buffer
	err := new(jsonpb.Marshaler).Marshal(&buf, msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (drpcEncoding_File_heldamount_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(buf), msg.(proto.Message))
}

type DRPCHeldAmountClient interface {
	DRPCConn() drpc.Conn

	GetPayStub(ctx context.Context, in *GetHeldAmountRequest) (*GetHeldAmountResponse, error)
	GetAllPaystubs(ctx context.Context, in *GetAllPaystubsRequest) (*GetAllPaystubsResponse, error)
	GetPayment(ctx context.Context, in *GetPaymentRequest) (*GetPaymentResponse, error)
	GetAllPayments(ctx context.Context, in *GetAllPaymentsRequest) (*GetAllPaymentsResponse, error)
}

type drpcHeldAmountClient struct {
	cc drpc.Conn
}

func NewDRPCHeldAmountClient(cc drpc.Conn) DRPCHeldAmountClient {
	return &drpcHeldAmountClient{cc}
}

func (c *drpcHeldAmountClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcHeldAmountClient) GetPayStub(ctx context.Context, in *GetHeldAmountRequest) (*GetHeldAmountResponse, error) {
	out := new(GetHeldAmountResponse)
	err := c.cc.Invoke(ctx, "/heldamount.HeldAmount/GetPayStub", drpcEncoding_File_heldamount_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcHeldAmountClient) GetAllPaystubs(ctx context.Context, in *GetAllPaystubsRequest) (*GetAllPaystubsResponse, error) {
	out := new(GetAllPaystubsResponse)
	err := c.cc.Invoke(ctx, "/heldamount.HeldAmount/GetAllPaystubs", drpcEncoding_File_heldamount_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcHeldAmountClient) GetPayment(ctx context.Context, in *GetPaymentRequest) (*GetPaymentResponse, error) {
	out := new(GetPaymentResponse)
	err := c.cc.Invoke(ctx, "/heldamount.HeldAmount/GetPayment", drpcEncoding_File_heldamount_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcHeldAmountClient) GetAllPayments(ctx context.Context, in *GetAllPaymentsRequest) (*GetAllPaymentsResponse, error) {
	out := new(GetAllPaymentsResponse)
	err := c.cc.Invoke(ctx, "/heldamount.HeldAmount/GetAllPayments", drpcEncoding_File_heldamount_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCHeldAmountServer interface {
	GetPayStub(context.Context, *GetHeldAmountRequest) (*GetHeldAmountResponse, error)
	GetAllPaystubs(context.Context, *GetAllPaystubsRequest) (*GetAllPaystubsResponse, error)
	GetPayment(context.Context, *GetPaymentRequest) (*GetPaymentResponse, error)
	GetAllPayments(context.Context, *GetAllPaymentsRequest) (*GetAllPaymentsResponse, error)
}

type DRPCHeldAmountUnimplementedServer struct{}

func (s *DRPCHeldAmountUnimplementedServer) GetPayStub(context.Context, *GetHeldAmountRequest) (*GetHeldAmountResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), 12)
}

func (s *DRPCHeldAmountUnimplementedServer) GetAllPaystubs(context.Context, *GetAllPaystubsRequest) (*GetAllPaystubsResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), 12)
}

func (s *DRPCHeldAmountUnimplementedServer) GetPayment(context.Context, *GetPaymentRequest) (*GetPaymentResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), 12)
}

func (s *DRPCHeldAmountUnimplementedServer) GetAllPayments(context.Context, *GetAllPaymentsRequest) (*GetAllPaymentsResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), 12)
}

type DRPCHeldAmountDescription struct{}

func (DRPCHeldAmountDescription) NumMethods() int { return 4 }

func (DRPCHeldAmountDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/heldamount.HeldAmount/GetPayStub", drpcEncoding_File_heldamount_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCHeldAmountServer).
					GetPayStub(
						ctx,
						in1.(*GetHeldAmountRequest),
					)
			}, DRPCHeldAmountServer.GetPayStub, true
	case 1:
		return "/heldamount.HeldAmount/GetAllPaystubs", drpcEncoding_File_heldamount_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCHeldAmountServer).
					GetAllPaystubs(
						ctx,
						in1.(*GetAllPaystubsRequest),
					)
			}, DRPCHeldAmountServer.GetAllPaystubs, true
	case 2:
		return "/heldamount.HeldAmount/GetPayment", drpcEncoding_File_heldamount_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCHeldAmountServer).
					GetPayment(
						ctx,
						in1.(*GetPaymentRequest),
					)
			}, DRPCHeldAmountServer.GetPayment, true
	case 3:
		return "/heldamount.HeldAmount/GetAllPayments", drpcEncoding_File_heldamount_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCHeldAmountServer).
					GetAllPayments(
						ctx,
						in1.(*GetAllPaymentsRequest),
					)
			}, DRPCHeldAmountServer.GetAllPayments, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterHeldAmount(mux drpc.Mux, impl DRPCHeldAmountServer) error {
	return mux.Register(impl, DRPCHeldAmountDescription{})
}

type DRPCHeldAmount_GetPayStubStream interface {
	drpc.Stream
	SendAndClose(*GetHeldAmountResponse) error
}

type drpcHeldAmount_GetPayStubStream struct {
	drpc.Stream
}

func (x *drpcHeldAmount_GetPayStubStream) SendAndClose(m *GetHeldAmountResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_heldamount_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCHeldAmount_GetAllPaystubsStream interface {
	drpc.Stream
	SendAndClose(*GetAllPaystubsResponse) error
}

type drpcHeldAmount_GetAllPaystubsStream struct {
	drpc.Stream
}

func (x *drpcHeldAmount_GetAllPaystubsStream) SendAndClose(m *GetAllPaystubsResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_heldamount_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCHeldAmount_GetPaymentStream interface {
	drpc.Stream
	SendAndClose(*GetPaymentResponse) error
}

type drpcHeldAmount_GetPaymentStream struct {
	drpc.Stream
}

func (x *drpcHeldAmount_GetPaymentStream) SendAndClose(m *GetPaymentResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_heldamount_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCHeldAmount_GetAllPaymentsStream interface {
	drpc.Stream
	SendAndClose(*GetAllPaymentsResponse) error
}

type drpcHeldAmount_GetAllPaymentsStream struct {
	drpc.Stream
}

func (x *drpcHeldAmount_GetAllPaymentsStream) SendAndClose(m *GetAllPaymentsResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_heldamount_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}
