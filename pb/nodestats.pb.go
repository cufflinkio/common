// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nodestats.proto

package pb

import (
	fmt "fmt"
	math "math"
	time "time"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ReputationStats struct {
	TotalCount             int64    `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	SuccessCount           int64    `protobuf:"varint,2,opt,name=success_count,json=successCount,proto3" json:"success_count,omitempty"`
	ReputationAlpha        float64  `protobuf:"fixed64,3,opt,name=reputation_alpha,json=reputationAlpha,proto3" json:"reputation_alpha,omitempty"`
	ReputationBeta         float64  `protobuf:"fixed64,4,opt,name=reputation_beta,json=reputationBeta,proto3" json:"reputation_beta,omitempty"`
	ReputationScore        float64  `protobuf:"fixed64,5,opt,name=reputation_score,json=reputationScore,proto3" json:"reputation_score,omitempty"`
	UnknownReputationAlpha float64  `protobuf:"fixed64,6,opt,name=unknown_reputation_alpha,json=unknownReputationAlpha,proto3" json:"unknown_reputation_alpha,omitempty"`
	UnknownReputationBeta  float64  `protobuf:"fixed64,7,opt,name=unknown_reputation_beta,json=unknownReputationBeta,proto3" json:"unknown_reputation_beta,omitempty"`
	UnknownReputationScore float64  `protobuf:"fixed64,8,opt,name=unknown_reputation_score,json=unknownReputationScore,proto3" json:"unknown_reputation_score,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *ReputationStats) Reset()         { *m = ReputationStats{} }
func (m *ReputationStats) String() string { return proto.CompactTextString(m) }
func (*ReputationStats) ProtoMessage()    {}
func (*ReputationStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{0}
}
func (m *ReputationStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReputationStats.Unmarshal(m, b)
}
func (m *ReputationStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReputationStats.Marshal(b, m, deterministic)
}
func (m *ReputationStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReputationStats.Merge(m, src)
}
func (m *ReputationStats) XXX_Size() int {
	return xxx_messageInfo_ReputationStats.Size(m)
}
func (m *ReputationStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ReputationStats.DiscardUnknown(m)
}

var xxx_messageInfo_ReputationStats proto.InternalMessageInfo

func (m *ReputationStats) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ReputationStats) GetSuccessCount() int64 {
	if m != nil {
		return m.SuccessCount
	}
	return 0
}

func (m *ReputationStats) GetReputationAlpha() float64 {
	if m != nil {
		return m.ReputationAlpha
	}
	return 0
}

func (m *ReputationStats) GetReputationBeta() float64 {
	if m != nil {
		return m.ReputationBeta
	}
	return 0
}

func (m *ReputationStats) GetReputationScore() float64 {
	if m != nil {
		return m.ReputationScore
	}
	return 0
}

func (m *ReputationStats) GetUnknownReputationAlpha() float64 {
	if m != nil {
		return m.UnknownReputationAlpha
	}
	return 0
}

func (m *ReputationStats) GetUnknownReputationBeta() float64 {
	if m != nil {
		return m.UnknownReputationBeta
	}
	return 0
}

func (m *ReputationStats) GetUnknownReputationScore() float64 {
	if m != nil {
		return m.UnknownReputationScore
	}
	return 0
}

type GetStatsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatsRequest) Reset()         { *m = GetStatsRequest{} }
func (m *GetStatsRequest) String() string { return proto.CompactTextString(m) }
func (*GetStatsRequest) ProtoMessage()    {}
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{1}
}
func (m *GetStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsRequest.Unmarshal(m, b)
}
func (m *GetStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsRequest.Marshal(b, m, deterministic)
}
func (m *GetStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsRequest.Merge(m, src)
}
func (m *GetStatsRequest) XXX_Size() int {
	return xxx_messageInfo_GetStatsRequest.Size(m)
}
func (m *GetStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsRequest proto.InternalMessageInfo

type GetStatsResponse struct {
	UptimeCheck          *ReputationStats `protobuf:"bytes,1,opt,name=uptime_check,json=uptimeCheck,proto3" json:"uptime_check,omitempty"`
	AuditCheck           *ReputationStats `protobuf:"bytes,2,opt,name=audit_check,json=auditCheck,proto3" json:"audit_check,omitempty"`
	Disqualified         *time.Time       `protobuf:"bytes,3,opt,name=disqualified,proto3,stdtime" json:"disqualified,omitempty"`
	Suspended            *time.Time       `protobuf:"bytes,4,opt,name=suspended,proto3,stdtime" json:"suspended,omitempty"`
	JoinedAt             time.Time        `protobuf:"bytes,5,opt,name=joined_at,json=joinedAt,proto3,stdtime" json:"joined_at"`
	OfflineSuspended     *time.Time       `protobuf:"bytes,6,opt,name=offline_suspended,json=offlineSuspended,proto3,stdtime" json:"offline_suspended,omitempty"`
	OnlineScore          float64          `protobuf:"fixed64,7,opt,name=online_score,json=onlineScore,proto3" json:"online_score,omitempty"`
	OfflineUnderReview   *time.Time       `protobuf:"bytes,8,opt,name=offline_under_review,json=offlineUnderReview,proto3,stdtime" json:"offline_under_review,omitempty"`
	AuditHistory         *AuditHistory    `protobuf:"bytes,9,opt,name=audit_history,json=auditHistory,proto3" json:"audit_history,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetStatsResponse) Reset()         { *m = GetStatsResponse{} }
func (m *GetStatsResponse) String() string { return proto.CompactTextString(m) }
func (*GetStatsResponse) ProtoMessage()    {}
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{2}
}
func (m *GetStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsResponse.Unmarshal(m, b)
}
func (m *GetStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsResponse.Marshal(b, m, deterministic)
}
func (m *GetStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsResponse.Merge(m, src)
}
func (m *GetStatsResponse) XXX_Size() int {
	return xxx_messageInfo_GetStatsResponse.Size(m)
}
func (m *GetStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsResponse proto.InternalMessageInfo

func (m *GetStatsResponse) GetUptimeCheck() *ReputationStats {
	if m != nil {
		return m.UptimeCheck
	}
	return nil
}

func (m *GetStatsResponse) GetAuditCheck() *ReputationStats {
	if m != nil {
		return m.AuditCheck
	}
	return nil
}

func (m *GetStatsResponse) GetDisqualified() *time.Time {
	if m != nil {
		return m.Disqualified
	}
	return nil
}

func (m *GetStatsResponse) GetSuspended() *time.Time {
	if m != nil {
		return m.Suspended
	}
	return nil
}

func (m *GetStatsResponse) GetJoinedAt() time.Time {
	if m != nil {
		return m.JoinedAt
	}
	return time.Time{}
}

func (m *GetStatsResponse) GetOfflineSuspended() *time.Time {
	if m != nil {
		return m.OfflineSuspended
	}
	return nil
}

func (m *GetStatsResponse) GetOnlineScore() float64 {
	if m != nil {
		return m.OnlineScore
	}
	return 0
}

func (m *GetStatsResponse) GetOfflineUnderReview() *time.Time {
	if m != nil {
		return m.OfflineUnderReview
	}
	return nil
}

func (m *GetStatsResponse) GetAuditHistory() *AuditHistory {
	if m != nil {
		return m.AuditHistory
	}
	return nil
}

type DailyStorageUsageRequest struct {
	From                 time.Time `protobuf:"bytes,1,opt,name=from,proto3,stdtime" json:"from"`
	To                   time.Time `protobuf:"bytes,2,opt,name=to,proto3,stdtime" json:"to"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DailyStorageUsageRequest) Reset()         { *m = DailyStorageUsageRequest{} }
func (m *DailyStorageUsageRequest) String() string { return proto.CompactTextString(m) }
func (*DailyStorageUsageRequest) ProtoMessage()    {}
func (*DailyStorageUsageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{3}
}
func (m *DailyStorageUsageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DailyStorageUsageRequest.Unmarshal(m, b)
}
func (m *DailyStorageUsageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DailyStorageUsageRequest.Marshal(b, m, deterministic)
}
func (m *DailyStorageUsageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DailyStorageUsageRequest.Merge(m, src)
}
func (m *DailyStorageUsageRequest) XXX_Size() int {
	return xxx_messageInfo_DailyStorageUsageRequest.Size(m)
}
func (m *DailyStorageUsageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DailyStorageUsageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DailyStorageUsageRequest proto.InternalMessageInfo

func (m *DailyStorageUsageRequest) GetFrom() time.Time {
	if m != nil {
		return m.From
	}
	return time.Time{}
}

func (m *DailyStorageUsageRequest) GetTo() time.Time {
	if m != nil {
		return m.To
	}
	return time.Time{}
}

type DailyStorageUsageResponse struct {
	NodeId               NodeID                                    `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3,customtype=NodeID" json:"node_id"`
	DailyStorageUsage    []*DailyStorageUsageResponse_StorageUsage `protobuf:"bytes,2,rep,name=daily_storage_usage,json=dailyStorageUsage,proto3" json:"daily_storage_usage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *DailyStorageUsageResponse) Reset()         { *m = DailyStorageUsageResponse{} }
func (m *DailyStorageUsageResponse) String() string { return proto.CompactTextString(m) }
func (*DailyStorageUsageResponse) ProtoMessage()    {}
func (*DailyStorageUsageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{4}
}
func (m *DailyStorageUsageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DailyStorageUsageResponse.Unmarshal(m, b)
}
func (m *DailyStorageUsageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DailyStorageUsageResponse.Marshal(b, m, deterministic)
}
func (m *DailyStorageUsageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DailyStorageUsageResponse.Merge(m, src)
}
func (m *DailyStorageUsageResponse) XXX_Size() int {
	return xxx_messageInfo_DailyStorageUsageResponse.Size(m)
}
func (m *DailyStorageUsageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DailyStorageUsageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DailyStorageUsageResponse proto.InternalMessageInfo

func (m *DailyStorageUsageResponse) GetDailyStorageUsage() []*DailyStorageUsageResponse_StorageUsage {
	if m != nil {
		return m.DailyStorageUsage
	}
	return nil
}

type DailyStorageUsageResponse_StorageUsage struct {
	AtRestTotal          float64   `protobuf:"fixed64,1,opt,name=at_rest_total,json=atRestTotal,proto3" json:"at_rest_total,omitempty"`
	Timestamp            time.Time `protobuf:"bytes,2,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DailyStorageUsageResponse_StorageUsage) Reset() {
	*m = DailyStorageUsageResponse_StorageUsage{}
}
func (m *DailyStorageUsageResponse_StorageUsage) String() string { return proto.CompactTextString(m) }
func (*DailyStorageUsageResponse_StorageUsage) ProtoMessage()    {}
func (*DailyStorageUsageResponse_StorageUsage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{4, 0}
}
func (m *DailyStorageUsageResponse_StorageUsage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DailyStorageUsageResponse_StorageUsage.Unmarshal(m, b)
}
func (m *DailyStorageUsageResponse_StorageUsage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DailyStorageUsageResponse_StorageUsage.Marshal(b, m, deterministic)
}
func (m *DailyStorageUsageResponse_StorageUsage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DailyStorageUsageResponse_StorageUsage.Merge(m, src)
}
func (m *DailyStorageUsageResponse_StorageUsage) XXX_Size() int {
	return xxx_messageInfo_DailyStorageUsageResponse_StorageUsage.Size(m)
}
func (m *DailyStorageUsageResponse_StorageUsage) XXX_DiscardUnknown() {
	xxx_messageInfo_DailyStorageUsageResponse_StorageUsage.DiscardUnknown(m)
}

var xxx_messageInfo_DailyStorageUsageResponse_StorageUsage proto.InternalMessageInfo

func (m *DailyStorageUsageResponse_StorageUsage) GetAtRestTotal() float64 {
	if m != nil {
		return m.AtRestTotal
	}
	return 0
}

func (m *DailyStorageUsageResponse_StorageUsage) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

type PricingModelRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PricingModelRequest) Reset()         { *m = PricingModelRequest{} }
func (m *PricingModelRequest) String() string { return proto.CompactTextString(m) }
func (*PricingModelRequest) ProtoMessage()    {}
func (*PricingModelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{5}
}
func (m *PricingModelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PricingModelRequest.Unmarshal(m, b)
}
func (m *PricingModelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PricingModelRequest.Marshal(b, m, deterministic)
}
func (m *PricingModelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PricingModelRequest.Merge(m, src)
}
func (m *PricingModelRequest) XXX_Size() int {
	return xxx_messageInfo_PricingModelRequest.Size(m)
}
func (m *PricingModelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PricingModelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PricingModelRequest proto.InternalMessageInfo

type PricingModelResponse struct {
	EgressBandwidthPrice int64    `protobuf:"varint,1,opt,name=egress_bandwidth_price,json=egressBandwidthPrice,proto3" json:"egress_bandwidth_price,omitempty"`
	RepairBandwidthPrice int64    `protobuf:"varint,2,opt,name=repair_bandwidth_price,json=repairBandwidthPrice,proto3" json:"repair_bandwidth_price,omitempty"`
	DiskSpacePrice       int64    `protobuf:"varint,3,opt,name=disk_space_price,json=diskSpacePrice,proto3" json:"disk_space_price,omitempty"`
	AuditBandwidthPrice  int64    `protobuf:"varint,4,opt,name=audit_bandwidth_price,json=auditBandwidthPrice,proto3" json:"audit_bandwidth_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PricingModelResponse) Reset()         { *m = PricingModelResponse{} }
func (m *PricingModelResponse) String() string { return proto.CompactTextString(m) }
func (*PricingModelResponse) ProtoMessage()    {}
func (*PricingModelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{6}
}
func (m *PricingModelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PricingModelResponse.Unmarshal(m, b)
}
func (m *PricingModelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PricingModelResponse.Marshal(b, m, deterministic)
}
func (m *PricingModelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PricingModelResponse.Merge(m, src)
}
func (m *PricingModelResponse) XXX_Size() int {
	return xxx_messageInfo_PricingModelResponse.Size(m)
}
func (m *PricingModelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PricingModelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PricingModelResponse proto.InternalMessageInfo

func (m *PricingModelResponse) GetEgressBandwidthPrice() int64 {
	if m != nil {
		return m.EgressBandwidthPrice
	}
	return 0
}

func (m *PricingModelResponse) GetRepairBandwidthPrice() int64 {
	if m != nil {
		return m.RepairBandwidthPrice
	}
	return 0
}

func (m *PricingModelResponse) GetDiskSpacePrice() int64 {
	if m != nil {
		return m.DiskSpacePrice
	}
	return 0
}

func (m *PricingModelResponse) GetAuditBandwidthPrice() int64 {
	if m != nil {
		return m.AuditBandwidthPrice
	}
	return 0
}

type AuditHistory struct {
	Windows              []*AuditWindow `protobuf:"bytes,1,rep,name=windows,proto3" json:"windows,omitempty"`
	Score                float64        `protobuf:"fixed64,2,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AuditHistory) Reset()         { *m = AuditHistory{} }
func (m *AuditHistory) String() string { return proto.CompactTextString(m) }
func (*AuditHistory) ProtoMessage()    {}
func (*AuditHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{7}
}
func (m *AuditHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditHistory.Unmarshal(m, b)
}
func (m *AuditHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditHistory.Marshal(b, m, deterministic)
}
func (m *AuditHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditHistory.Merge(m, src)
}
func (m *AuditHistory) XXX_Size() int {
	return xxx_messageInfo_AuditHistory.Size(m)
}
func (m *AuditHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditHistory.DiscardUnknown(m)
}

var xxx_messageInfo_AuditHistory proto.InternalMessageInfo

func (m *AuditHistory) GetWindows() []*AuditWindow {
	if m != nil {
		return m.Windows
	}
	return nil
}

func (m *AuditHistory) GetScore() float64 {
	if m != nil {
		return m.Score
	}
	return 0
}

type AuditWindow struct {
	WindowStart          time.Time `protobuf:"bytes,1,opt,name=window_start,json=windowStart,proto3,stdtime" json:"window_start"`
	OnlineCount          int32     `protobuf:"varint,2,opt,name=online_count,json=onlineCount,proto3" json:"online_count,omitempty"`
	TotalCount           int32     `protobuf:"varint,3,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AuditWindow) Reset()         { *m = AuditWindow{} }
func (m *AuditWindow) String() string { return proto.CompactTextString(m) }
func (*AuditWindow) ProtoMessage()    {}
func (*AuditWindow) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{8}
}
func (m *AuditWindow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditWindow.Unmarshal(m, b)
}
func (m *AuditWindow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditWindow.Marshal(b, m, deterministic)
}
func (m *AuditWindow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditWindow.Merge(m, src)
}
func (m *AuditWindow) XXX_Size() int {
	return xxx_messageInfo_AuditWindow.Size(m)
}
func (m *AuditWindow) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditWindow.DiscardUnknown(m)
}

var xxx_messageInfo_AuditWindow proto.InternalMessageInfo

func (m *AuditWindow) GetWindowStart() time.Time {
	if m != nil {
		return m.WindowStart
	}
	return time.Time{}
}

func (m *AuditWindow) GetOnlineCount() int32 {
	if m != nil {
		return m.OnlineCount
	}
	return 0
}

func (m *AuditWindow) GetTotalCount() int32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func init() {
	proto.RegisterType((*ReputationStats)(nil), "nodestats.ReputationStats")
	proto.RegisterType((*GetStatsRequest)(nil), "nodestats.GetStatsRequest")
	proto.RegisterType((*GetStatsResponse)(nil), "nodestats.GetStatsResponse")
	proto.RegisterType((*DailyStorageUsageRequest)(nil), "nodestats.DailyStorageUsageRequest")
	proto.RegisterType((*DailyStorageUsageResponse)(nil), "nodestats.DailyStorageUsageResponse")
	proto.RegisterType((*DailyStorageUsageResponse_StorageUsage)(nil), "nodestats.DailyStorageUsageResponse.StorageUsage")
	proto.RegisterType((*PricingModelRequest)(nil), "nodestats.PricingModelRequest")
	proto.RegisterType((*PricingModelResponse)(nil), "nodestats.PricingModelResponse")
	proto.RegisterType((*AuditHistory)(nil), "nodestats.AuditHistory")
	proto.RegisterType((*AuditWindow)(nil), "nodestats.AuditWindow")
}

func init() { proto.RegisterFile("nodestats.proto", fileDescriptor_e0b184ee117142aa) }

var fileDescriptor_e0b184ee117142aa = []byte{
	// 906 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x4d, 0x6f, 0xdb, 0x46,
	0x10, 0x2d, 0x25, 0x5b, 0xb6, 0x86, 0xf4, 0xd7, 0x5a, 0x76, 0x58, 0x15, 0xa8, 0x5c, 0xa6, 0x40,
	0xdc, 0x8b, 0xdc, 0xba, 0x41, 0x11, 0xa0, 0xed, 0xc1, 0x72, 0x80, 0x24, 0x87, 0x7e, 0x51, 0x49,
	0x0a, 0xf4, 0x50, 0x62, 0xc5, 0x5d, 0xc9, 0x1b, 0x4b, 0x5c, 0x86, 0xbb, 0x8c, 0x90, 0x63, 0x7b,
	0xec, 0xa9, 0xe7, 0x5e, 0xfa, 0x77, 0xfa, 0x1b, 0x7a, 0x48, 0x7b, 0xef, 0x9f, 0x28, 0x76, 0x96,
	0x92, 0x68, 0x5a, 0x4e, 0xe4, 0xa3, 0xde, 0xbc, 0xf7, 0x76, 0xb4, 0xf3, 0x38, 0x0b, 0x3b, 0x89,
	0x64, 0x5c, 0x69, 0xaa, 0x55, 0x37, 0xcd, 0xa4, 0x96, 0xa4, 0x39, 0x07, 0xda, 0x30, 0x92, 0x23,
	0x69, 0xe1, 0x76, 0x67, 0x24, 0xe5, 0x68, 0xcc, 0x4f, 0xf0, 0xd7, 0x20, 0x1f, 0x9e, 0x68, 0x31,
	0x31, 0xb4, 0x49, 0x6a, 0x09, 0xc1, 0xaf, 0x75, 0xd8, 0x09, 0x79, 0x9a, 0x6b, 0xaa, 0x85, 0x4c,
	0xfa, 0xc6, 0x80, 0x74, 0xc0, 0xd5, 0x52, 0xd3, 0x71, 0x14, 0xcb, 0x3c, 0xd1, 0xbe, 0x73, 0xe4,
	0x1c, 0xd7, 0x43, 0x40, 0xe8, 0xdc, 0x20, 0xe4, 0x2e, 0x6c, 0xa9, 0x3c, 0x8e, 0xb9, 0x52, 0x05,
	0xa5, 0x86, 0x14, 0xaf, 0x00, 0x2d, 0xe9, 0x13, 0xd8, 0xcd, 0xe6, 0xc6, 0x11, 0x1d, 0xa7, 0x17,
	0xd4, 0xaf, 0x1f, 0x39, 0xc7, 0x4e, 0xb8, 0xb3, 0xc0, 0xcf, 0x0c, 0x4c, 0xee, 0x41, 0x09, 0x8a,
	0x06, 0x5c, 0x53, 0x7f, 0x0d, 0x99, 0xdb, 0x0b, 0xb8, 0xc7, 0x35, 0xad, 0x78, 0xaa, 0x58, 0x66,
	0xdc, 0x5f, 0xaf, 0x7a, 0xf6, 0x0d, 0x4c, 0x1e, 0x80, 0x9f, 0x27, 0x97, 0x89, 0x9c, 0x26, 0xd1,
	0xb5, 0x36, 0x1a, 0x28, 0x39, 0x2c, 0xea, 0x61, 0xa5, 0x9b, 0x2f, 0xe0, 0xce, 0x12, 0x25, 0x76,
	0xb5, 0x81, 0xc2, 0x83, 0x6b, 0x42, 0x6c, 0x6e, 0xf9, 0x89, 0xb6, 0xc9, 0xcd, 0x1b, 0x4e, 0xc4,
	0x5e, 0x83, 0x3d, 0xd8, 0x79, 0xc4, 0x35, 0x5e, 0x7e, 0xc8, 0x5f, 0xe6, 0x5c, 0xe9, 0xe0, 0xbf,
	0x35, 0xd8, 0x5d, 0x60, 0x2a, 0x95, 0x89, 0xe2, 0xe4, 0x6b, 0xf0, 0xf2, 0xd4, 0x4c, 0x30, 0x8a,
	0x2f, 0x78, 0x7c, 0x89, 0x93, 0x71, 0x4f, 0xdb, 0xdd, 0x45, 0x18, 0x2a, 0xa3, 0x0c, 0x5d, 0xcb,
	0x3f, 0x37, 0x74, 0xf2, 0x25, 0xb8, 0x34, 0x67, 0x42, 0x17, 0xea, 0xda, 0x3b, 0xd5, 0x80, 0x74,
	0x2b, 0x7e, 0x0c, 0x1e, 0x13, 0xea, 0x65, 0x4e, 0xc7, 0x62, 0x28, 0x38, 0xc3, 0x51, 0x1a, 0xb5,
	0x0d, 0x58, 0x77, 0x16, 0xb0, 0xee, 0xd3, 0x59, 0xc0, 0x7a, 0x9b, 0x7f, 0xbd, 0xe9, 0x38, 0xbf,
	0xff, 0xd3, 0x71, 0xc2, 0x2b, 0x4a, 0xd2, 0x83, 0xa6, 0xca, 0x55, 0xca, 0x13, 0xc6, 0x19, 0xce,
	0x79, 0x55, 0x9b, 0x85, 0x8c, 0x9c, 0x41, 0xf3, 0x85, 0x14, 0x09, 0x67, 0x11, 0xd5, 0x98, 0x80,
	0x77, 0x7b, 0xbc, 0x87, 0x1e, 0x9b, 0x56, 0x76, 0xa6, 0xc9, 0x0f, 0xb0, 0x27, 0x87, 0xc3, 0xb1,
	0x48, 0x78, 0xb4, 0x68, 0xa7, 0x71, 0x8b, 0x76, 0x76, 0x0b, 0x79, 0x7f, 0xde, 0xd5, 0x47, 0xe0,
	0xc9, 0xc4, 0x3a, 0xe2, 0xd4, 0x6d, 0x5c, 0x5c, 0x8b, 0xd9, 0x58, 0x3e, 0x87, 0xd6, 0xec, 0xd4,
	0x3c, 0x61, 0x3c, 0x8b, 0x32, 0xfe, 0x4a, 0xf0, 0x29, 0x06, 0x64, 0xd5, 0x83, 0x49, 0xe1, 0xf0,
	0xcc, 0x18, 0x84, 0xa8, 0x27, 0x5f, 0xc1, 0x96, 0x9d, 0xed, 0x85, 0x50, 0x5a, 0x66, 0xaf, 0xfd,
	0x26, 0x1a, 0xde, 0x29, 0x4d, 0xf7, 0xcc, 0xd4, 0x1f, 0xdb, 0x72, 0xe8, 0xd1, 0xd2, 0xaf, 0xe0,
	0x37, 0x07, 0xfc, 0x87, 0x54, 0x8c, 0x5f, 0xf7, 0xb5, 0xcc, 0xe8, 0x88, 0x3f, 0x53, 0x74, 0xc4,
	0x8b, 0x28, 0x92, 0x07, 0xb0, 0x36, 0xcc, 0xe4, 0x64, 0x9e, 0xb6, 0x55, 0xae, 0x19, 0x15, 0xe4,
	0x3e, 0xd4, 0xb4, 0x9c, 0xe7, 0x6c, 0x15, 0x5d, 0x4d, 0xcb, 0xe0, 0xcf, 0x1a, 0xbc, 0xbf, 0xa4,
	0x99, 0xe2, 0x1b, 0xb8, 0x07, 0x1b, 0xe6, 0x2f, 0x45, 0x82, 0x61, 0x43, 0x5e, 0x6f, 0xdb, 0x88,
	0xff, 0x7e, 0xd3, 0x69, 0x7c, 0x2b, 0x19, 0x7f, 0xf2, 0x30, 0x6c, 0x98, 0xf2, 0x13, 0x46, 0x28,
	0xec, 0x33, 0xe3, 0x12, 0x29, 0x6b, 0x13, 0xe5, 0xc6, 0xc7, 0xaf, 0x1d, 0xd5, 0x8f, 0xdd, 0xd3,
	0xcf, 0x4a, 0xf7, 0x72, 0xe3, 0x59, 0xdd, 0x2b, 0xe0, 0x1e, 0xab, 0xf2, 0xda, 0xaf, 0xc0, 0x2b,
	0xff, 0x26, 0x01, 0x6c, 0x51, 0x1d, 0x65, 0x5c, 0xe9, 0x08, 0xb7, 0x25, 0x76, 0xe8, 0x84, 0x2e,
	0xd5, 0x21, 0x57, 0xfa, 0xa9, 0x81, 0x4c, 0xfa, 0xe7, 0x3b, 0xf8, 0x56, 0x57, 0xb3, 0x90, 0x05,
	0x07, 0xb0, 0xff, 0x7d, 0x26, 0x62, 0x91, 0x8c, 0xbe, 0x91, 0x8c, 0x8f, 0x67, 0x3b, 0xe3, 0x5f,
	0x07, 0x5a, 0x57, 0xf1, 0xe2, 0xce, 0xee, 0xc3, 0x21, 0x1f, 0x65, 0x66, 0x5d, 0x0f, 0x68, 0xc2,
	0xa6, 0x82, 0xe9, 0x8b, 0x28, 0xcd, 0x44, 0xcc, 0x8b, 0xdd, 0xde, 0xb2, 0xd5, 0xde, 0xac, 0x68,
	0x4c, 0x50, 0x95, 0xf1, 0x94, 0x8a, 0xec, 0x9a, 0xca, 0xae, 0xfb, 0x96, 0xad, 0x56, 0x54, 0xc7,
	0xb0, 0xcb, 0x84, 0xba, 0x8c, 0x54, 0x4a, 0x63, 0x5e, 0xf0, 0xeb, 0xc8, 0xdf, 0x36, 0x78, 0xdf,
	0xc0, 0x96, 0x79, 0x0a, 0x07, 0x36, 0xb2, 0x55, 0xfb, 0x35, 0xa4, 0xef, 0x63, 0xf1, 0xaa, 0x7b,
	0xf0, 0x1c, 0xbc, 0x72, 0x8c, 0xc9, 0xa7, 0xb0, 0x31, 0x15, 0x09, 0x93, 0x53, 0xe5, 0x3b, 0x38,
	0xd8, 0xc3, 0x6a, 0xe0, 0x7f, 0xc4, 0x72, 0x38, 0xa3, 0x91, 0x16, 0xac, 0xdb, 0x8f, 0xb3, 0x86,
	0xb3, 0xb1, 0x3f, 0x82, 0x3f, 0x1c, 0x70, 0x4b, 0x74, 0xf2, 0x08, 0x3c, 0x2b, 0x88, 0x94, 0xa6,
	0x99, 0xbe, 0x55, 0xf6, 0x5d, 0xab, 0xec, 0x1b, 0x61, 0x69, 0x25, 0x2c, 0x5e, 0xca, 0xf5, 0xd9,
	0x4a, 0xb0, 0x0f, 0x65, 0xe5, 0xb9, 0xad, 0x23, 0xa3, 0xf4, 0xdc, 0x9e, 0xfe, 0x52, 0x83, 0xa6,
	0x09, 0xb7, 0x7d, 0x9d, 0xcf, 0x61, 0x73, 0xf6, 0x30, 0x90, 0xf2, 0xf2, 0xae, 0xbc, 0x20, 0xed,
	0x0f, 0x96, 0xd6, 0x8a, 0x44, 0xfc, 0x0c, 0x7b, 0xd7, 0x62, 0x4f, 0xee, 0xbe, 0xfd, 0xa3, 0xb0,
	0xb6, 0x1f, 0xaf, 0xf2, 0xe5, 0x90, 0xef, 0xc0, 0x2b, 0x27, 0x91, 0x7c, 0x58, 0x52, 0x2d, 0x89,
	0x6e, 0xbb, 0x73, 0x63, 0xdd, 0x1a, 0xf6, 0x5a, 0x3f, 0x11, 0x33, 0xf1, 0x17, 0x5d, 0x21, 0x4f,
	0x62, 0x39, 0x99, 0xc8, 0xe4, 0x24, 0x1d, 0x0c, 0x1a, 0x38, 0x88, 0xcf, 0xff, 0x0f, 0x00, 0x00,
	0xff, 0xff, 0x0f, 0x49, 0xa8, 0xd9, 0x0f, 0x09, 0x00, 0x00,
}
