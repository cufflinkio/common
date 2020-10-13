// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: delegated_repair.proto

package pb

import (
	context "context"
	fmt "fmt"
	math "math"
	time "time"

	proto "github.com/gogo/protobuf/proto"

	drpc "storj.io/drpc"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RepairJobRequest struct {
	// When not the first request, this will include the result of the last job
	LastJobResult        *RepairJobResult `protobuf:"bytes,1,opt,name=last_job_result,json=lastJobResult,proto3" json:"last_job_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RepairJobRequest) Reset()         { *m = RepairJobRequest{} }
func (m *RepairJobRequest) String() string { return proto.CompactTextString(m) }
func (*RepairJobRequest) ProtoMessage()    {}
func (*RepairJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d00d18c724d5a7, []int{0}
}
func (m *RepairJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepairJobRequest.Unmarshal(m, b)
}
func (m *RepairJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepairJobRequest.Marshal(b, m, deterministic)
}
func (m *RepairJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepairJobRequest.Merge(m, src)
}
func (m *RepairJobRequest) XXX_Size() int {
	return xxx_messageInfo_RepairJobRequest.Size(m)
}
func (m *RepairJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RepairJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RepairJobRequest proto.InternalMessageInfo

func (m *RepairJobRequest) GetLastJobResult() *RepairJobResult {
	if m != nil {
		return m.LastJobResult
	}
	return nil
}

type RepairJobResponse struct {
	// When a job is available, this will be filled in
	NewJob *RepairJobDefinition `protobuf:"bytes,1,opt,name=new_job,json=newJob,proto3" json:"new_job,omitempty"`
	// Otherwise, client should wait this many milliseconds and then try again
	ComeBackInMillis     int32    `protobuf:"varint,2,opt,name=come_back_in_millis,json=comeBackInMillis,proto3" json:"come_back_in_millis,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepairJobResponse) Reset()         { *m = RepairJobResponse{} }
func (m *RepairJobResponse) String() string { return proto.CompactTextString(m) }
func (*RepairJobResponse) ProtoMessage()    {}
func (*RepairJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d00d18c724d5a7, []int{1}
}
func (m *RepairJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepairJobResponse.Unmarshal(m, b)
}
func (m *RepairJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepairJobResponse.Marshal(b, m, deterministic)
}
func (m *RepairJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepairJobResponse.Merge(m, src)
}
func (m *RepairJobResponse) XXX_Size() int {
	return xxx_messageInfo_RepairJobResponse.Size(m)
}
func (m *RepairJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RepairJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RepairJobResponse proto.InternalMessageInfo

func (m *RepairJobResponse) GetNewJob() *RepairJobDefinition {
	if m != nil {
		return m.NewJob
	}
	return nil
}

func (m *RepairJobResponse) GetComeBackInMillis() int32 {
	if m != nil {
		return m.ComeBackInMillis
	}
	return 0
}

type RepairJobDefinition struct {
	// Identifier for this job
	JobId []byte `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// Signed GET orders for all believed-healthy pieces to be downloaded
	GetOrders []*AddressedOrderLimit `protobuf:"bytes,2,rep,name=get_orders,json=getOrders,proto3" json:"get_orders,omitempty"`
	// Private piece key to use for fetching
	PrivateKeyForGet []byte `protobuf:"bytes,3,opt,name=private_key_for_get,json=privateKeyForGet,proto3" json:"private_key_for_get,omitempty"`
	// Signed PUT orders for all possible pieces to be uploaded (not including
	// piece numbers in get_orders)
	PutOrders []*AddressedOrderLimit `protobuf:"bytes,4,rep,name=put_orders,json=putOrders,proto3" json:"put_orders,omitempty"`
	// Private piece key to use for storing
	PrivateKeyForPut []byte `protobuf:"bytes,5,opt,name=private_key_for_put,json=privateKeyForPut,proto3" json:"private_key_for_put,omitempty"`
	// Redundancy scheme used by the segment to be repaired
	Redundancy *RedundancyScheme `protobuf:"bytes,6,opt,name=redundancy,proto3" json:"redundancy,omitempty"`
	// Size of the segment to be repaired
	SegmentSize int64 `protobuf:"varint,7,opt,name=segment_size,json=segmentSize,proto3" json:"segment_size,omitempty"`
	// Target piece count (worker should try to upload enough pieces so that
	// this count is achieved)
	DesiredPieceCount int32 `protobuf:"varint,8,opt,name=desired_piece_count,json=desiredPieceCount,proto3" json:"desired_piece_count,omitempty"`
	// Job expiration time
	ExpirationTime       time.Time `protobuf:"bytes,9,opt,name=expiration_time,json=expirationTime,proto3,stdtime" json:"expiration_time"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RepairJobDefinition) Reset()         { *m = RepairJobDefinition{} }
func (m *RepairJobDefinition) String() string { return proto.CompactTextString(m) }
func (*RepairJobDefinition) ProtoMessage()    {}
func (*RepairJobDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d00d18c724d5a7, []int{2}
}
func (m *RepairJobDefinition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepairJobDefinition.Unmarshal(m, b)
}
func (m *RepairJobDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepairJobDefinition.Marshal(b, m, deterministic)
}
func (m *RepairJobDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepairJobDefinition.Merge(m, src)
}
func (m *RepairJobDefinition) XXX_Size() int {
	return xxx_messageInfo_RepairJobDefinition.Size(m)
}
func (m *RepairJobDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_RepairJobDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_RepairJobDefinition proto.InternalMessageInfo

func (m *RepairJobDefinition) GetJobId() []byte {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *RepairJobDefinition) GetGetOrders() []*AddressedOrderLimit {
	if m != nil {
		return m.GetOrders
	}
	return nil
}

func (m *RepairJobDefinition) GetPrivateKeyForGet() []byte {
	if m != nil {
		return m.PrivateKeyForGet
	}
	return nil
}

func (m *RepairJobDefinition) GetPutOrders() []*AddressedOrderLimit {
	if m != nil {
		return m.PutOrders
	}
	return nil
}

func (m *RepairJobDefinition) GetPrivateKeyForPut() []byte {
	if m != nil {
		return m.PrivateKeyForPut
	}
	return nil
}

func (m *RepairJobDefinition) GetRedundancy() *RedundancyScheme {
	if m != nil {
		return m.Redundancy
	}
	return nil
}

func (m *RepairJobDefinition) GetSegmentSize() int64 {
	if m != nil {
		return m.SegmentSize
	}
	return 0
}

func (m *RepairJobDefinition) GetDesiredPieceCount() int32 {
	if m != nil {
		return m.DesiredPieceCount
	}
	return 0
}

func (m *RepairJobDefinition) GetExpirationTime() time.Time {
	if m != nil {
		return m.ExpirationTime
	}
	return time.Time{}
}

type RepairJobResult struct {
	// Identifier for this job, as given in RepairJobResponse
	JobId []byte `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// Set nonzero only if the segment could not be reconstructed because of
	// too few pieces available.
	IrreparablePiecesRetrieved int32 `protobuf:"varint,2,opt,name=irreparable_pieces_retrieved,json=irreparablePiecesRetrieved,proto3" json:"irreparable_pieces_retrieved,omitempty"`
	// Set only if the segment could not be reconstructed.
	ReconstructError string `protobuf:"bytes,3,opt,name=reconstruct_error,json=reconstructError,proto3" json:"reconstruct_error,omitempty"`
	// Set only if new pieces could not be stored to any new nodes.
	StoreError string `protobuf:"bytes,4,opt,name=store_error,json=storeError,proto3" json:"store_error,omitempty"`
	// PieceHashes signed by storage nodes which were used to accomplish repair
	NewPiecesStored []*PieceHash `protobuf:"bytes,5,rep,name=new_pieces_stored,json=newPiecesStored,proto3" json:"new_pieces_stored,omitempty"`
	// A copy of the put_orders list as provided in the corresponding
	// RepairJobDefinition
	PutOrders []*AddressedOrderLimit `protobuf:"bytes,6,rep,name=put_orders,json=putOrders,proto3" json:"put_orders,omitempty"`
	// Pieces which should be _removed_ from the pointer. This will include
	// pieces for which the expected owning storage node returned a "not found"
	// error, as well as pieces which were downloaded but failed their
	// validation check.
	DeletePieceNums      []int32  `protobuf:"varint,7,rep,packed,name=delete_piece_nums,json=deletePieceNums,proto3" json:"delete_piece_nums,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepairJobResult) Reset()         { *m = RepairJobResult{} }
func (m *RepairJobResult) String() string { return proto.CompactTextString(m) }
func (*RepairJobResult) ProtoMessage()    {}
func (*RepairJobResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d00d18c724d5a7, []int{3}
}
func (m *RepairJobResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepairJobResult.Unmarshal(m, b)
}
func (m *RepairJobResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepairJobResult.Marshal(b, m, deterministic)
}
func (m *RepairJobResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepairJobResult.Merge(m, src)
}
func (m *RepairJobResult) XXX_Size() int {
	return xxx_messageInfo_RepairJobResult.Size(m)
}
func (m *RepairJobResult) XXX_DiscardUnknown() {
	xxx_messageInfo_RepairJobResult.DiscardUnknown(m)
}

var xxx_messageInfo_RepairJobResult proto.InternalMessageInfo

func (m *RepairJobResult) GetJobId() []byte {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *RepairJobResult) GetIrreparablePiecesRetrieved() int32 {
	if m != nil {
		return m.IrreparablePiecesRetrieved
	}
	return 0
}

func (m *RepairJobResult) GetReconstructError() string {
	if m != nil {
		return m.ReconstructError
	}
	return ""
}

func (m *RepairJobResult) GetStoreError() string {
	if m != nil {
		return m.StoreError
	}
	return ""
}

func (m *RepairJobResult) GetNewPiecesStored() []*PieceHash {
	if m != nil {
		return m.NewPiecesStored
	}
	return nil
}

func (m *RepairJobResult) GetPutOrders() []*AddressedOrderLimit {
	if m != nil {
		return m.PutOrders
	}
	return nil
}

func (m *RepairJobResult) GetDeletePieceNums() []int32 {
	if m != nil {
		return m.DeletePieceNums
	}
	return nil
}

func init() {
	proto.RegisterType((*RepairJobRequest)(nil), "delegated_repair.RepairJobRequest")
	proto.RegisterType((*RepairJobResponse)(nil), "delegated_repair.RepairJobResponse")
	proto.RegisterType((*RepairJobDefinition)(nil), "delegated_repair.RepairJobDefinition")
	proto.RegisterType((*RepairJobResult)(nil), "delegated_repair.RepairJobResult")
}

func init() { proto.RegisterFile("delegated_repair.proto", fileDescriptor_04d00d18c724d5a7) }

var fileDescriptor_04d00d18c724d5a7 = []byte{
	// 691 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdb, 0x6e, 0x13, 0x3b,
	0x14, 0x6d, 0x9a, 0x26, 0x6d, 0x9c, 0x9e, 0x26, 0x71, 0xcf, 0x39, 0x1a, 0x05, 0x50, 0xd2, 0x20,
	0xa4, 0x08, 0xc4, 0x44, 0x2a, 0x8f, 0x5c, 0x04, 0x2d, 0xb7, 0x16, 0x0a, 0x95, 0x8b, 0x78, 0x40,
	0x42, 0xd6, 0xcc, 0x78, 0x67, 0xea, 0x26, 0x63, 0x0f, 0xb6, 0xa7, 0xa5, 0x7d, 0xe4, 0x0b, 0xf8,
	0x2c, 0xbe, 0x02, 0xc4, 0x03, 0xff, 0x81, 0xec, 0x99, 0xa4, 0x51, 0x29, 0x11, 0x3c, 0x7a, 0xaf,
	0x65, 0xaf, 0xe5, 0xed, 0xe5, 0x8d, 0xfe, 0x67, 0x30, 0x86, 0x38, 0x30, 0xc0, 0xa8, 0x82, 0x34,
	0xe0, 0xca, 0x4f, 0x95, 0x34, 0x12, 0x37, 0x2f, 0xd6, 0xdb, 0x28, 0x96, 0xb1, 0xcc, 0xd1, 0x76,
	0x27, 0x96, 0x32, 0x1e, 0xc3, 0xc0, 0xad, 0xc2, 0x6c, 0x38, 0x30, 0x3c, 0x01, 0x6d, 0x82, 0x24,
	0x2d, 0x08, 0x6b, 0x09, 0x98, 0x80, 0x8b, 0xe1, 0x64, 0xc3, 0xaa, 0x54, 0x0c, 0x94, 0x2e, 0x56,
	0x8d, 0x54, 0x72, 0x61, 0x40, 0xb1, 0x30, 0x2f, 0xf4, 0xde, 0xa3, 0x26, 0x71, 0x2a, 0xbb, 0x32,
	0x24, 0xf0, 0x21, 0x03, 0x6d, 0xf0, 0x0e, 0x6a, 0x8c, 0x03, 0x6d, 0xe8, 0x91, 0x0c, 0xa9, 0x02,
	0x9d, 0x8d, 0x8d, 0x57, 0xea, 0x96, 0xfa, 0xf5, 0xcd, 0x0d, 0xff, 0x17, 0xcf, 0x33, 0x9b, 0x2d,
	0x91, 0xfc, 0x63, 0x77, 0x4e, 0x97, 0xbd, 0x4f, 0x25, 0xd4, 0x9a, 0xa5, 0xa4, 0x52, 0x68, 0xc0,
	0x0f, 0xd0, 0xb2, 0x80, 0x13, 0x7b, 0x7e, 0x71, 0xf0, 0x8d, 0x39, 0x07, 0x3f, 0x86, 0x21, 0x17,
	0xdc, 0x70, 0x29, 0x48, 0x55, 0xc0, 0xc9, 0xae, 0x0c, 0xf1, 0x6d, 0xb4, 0x1e, 0xc9, 0x04, 0x68,
	0x18, 0x44, 0x23, 0xca, 0x05, 0x4d, 0xf8, 0x78, 0xcc, 0xb5, 0xb7, 0xd8, 0x2d, 0xf5, 0x2b, 0xa4,
	0x69, 0xa1, 0xad, 0x20, 0x1a, 0xed, 0x88, 0x3d, 0x57, 0xef, 0xfd, 0x28, 0xa3, 0xf5, 0x4b, 0x8e,
	0xc3, 0xff, 0xa1, 0xaa, 0xbd, 0x22, 0x67, 0xce, 0xc5, 0x2a, 0xa9, 0x1c, 0xc9, 0x70, 0x87, 0xe1,
	0x7b, 0x08, 0xc5, 0x60, 0x68, 0xde, 0x37, 0x6f, 0xb1, 0x5b, 0xee, 0xd7, 0x37, 0xaf, 0xf9, 0xd3,
	0xb6, 0x3e, 0x62, 0x4c, 0x81, 0xd6, 0xc0, 0x5e, 0x5b, 0xc2, 0x4b, 0x9e, 0x70, 0x43, 0x6a, 0x31,
	0x18, 0xb7, 0xd4, 0xd6, 0x5b, 0xaa, 0xf8, 0x71, 0x60, 0x80, 0x8e, 0xe0, 0x94, 0x0e, 0xa5, 0xa2,
	0x31, 0x18, 0xaf, 0xec, 0x14, 0x9a, 0x05, 0xf4, 0x02, 0x4e, 0x9f, 0x4a, 0xf5, 0x0c, 0x8c, 0x15,
	0x4b, 0xb3, 0xa9, 0xd8, 0xd2, 0x1f, 0x89, 0xa5, 0xd9, 0x1c, 0xb1, 0x34, 0x33, 0x5e, 0xe5, 0x12,
	0xb1, 0xfd, 0xcc, 0xe0, 0xbb, 0x08, 0x29, 0x60, 0x99, 0x60, 0x81, 0x88, 0x4e, 0xbd, 0xaa, 0x6b,
	0xfd, 0x15, 0xff, 0x3c, 0x12, 0x64, 0x0a, 0x1e, 0x44, 0x87, 0x90, 0x00, 0x99, 0xa1, 0xe3, 0x0d,
	0xb4, 0xaa, 0x21, 0x4e, 0x40, 0x18, 0xaa, 0xf9, 0x19, 0x78, 0xcb, 0xdd, 0x52, 0xbf, 0x4c, 0xea,
	0x45, 0xed, 0x80, 0x9f, 0x01, 0xf6, 0xd1, 0x3a, 0x03, 0xcd, 0x15, 0x30, 0x9a, 0x72, 0x88, 0x80,
	0x46, 0x32, 0x13, 0xc6, 0x5b, 0x71, 0xef, 0xd2, 0x2a, 0xa0, 0x7d, 0x8b, 0x6c, 0x5b, 0x00, 0xef,
	0xa1, 0x06, 0x7c, 0x4c, 0xb9, 0x0a, 0xec, 0x73, 0x50, 0x9b, 0x64, 0xaf, 0xe6, 0x4c, 0xb5, 0xfd,
	0x3c, 0xe6, 0xfe, 0x24, 0xe6, 0xfe, 0x9b, 0x49, 0xcc, 0xb7, 0x56, 0xbe, 0x7c, 0xed, 0x2c, 0x7c,
	0xfe, 0xd6, 0x29, 0x91, 0xb5, 0xf3, 0xcd, 0x16, 0xee, 0x7d, 0x5f, 0x44, 0x8d, 0x0b, 0x79, 0xfc,
	0xdd, 0x1b, 0x3f, 0x44, 0x57, 0xb9, 0xb2, 0x49, 0x53, 0x41, 0x38, 0x86, 0xdc, 0xad, 0xa6, 0x0a,
	0x8c, 0xe2, 0x70, 0x0c, 0xac, 0x88, 0x52, 0x7b, 0x86, 0xe3, 0x6c, 0x6b, 0x32, 0x61, 0xe0, 0x5b,
	0xa8, 0xa5, 0x20, 0x92, 0x42, 0x1b, 0x95, 0x45, 0x86, 0x82, 0x52, 0x52, 0xb9, 0x57, 0xae, 0x91,
	0xe6, 0x0c, 0xf0, 0xc4, 0xd6, 0x71, 0x07, 0xd5, 0xb5, 0x91, 0x0a, 0x0a, 0xda, 0x92, 0xa3, 0x21,
	0x57, 0xca, 0x09, 0xf7, 0x51, 0xcb, 0xfe, 0x88, 0xc2, 0x87, 0x03, 0x98, 0x57, 0x71, 0x69, 0x68,
	0xf9, 0xc5, 0x0f, 0x76, 0x0e, 0x9e, 0x07, 0xfa, 0x90, 0x34, 0x04, 0x9c, 0xe4, 0x7e, 0x0e, 0x1c,
	0xf3, 0x42, 0x8a, 0xaa, 0x7f, 0x99, 0xa2, 0x9b, 0xa8, 0x65, 0xbf, 0x9f, 0x29, 0xfa, 0x40, 0x45,
	0x96, 0x68, 0x6f, 0xb9, 0x5b, 0xee, 0x57, 0x48, 0x23, 0x07, 0x9c, 0xd8, 0xab, 0x2c, 0xd1, 0x9b,
	0xa3, 0xc9, 0x7f, 0xde, 0x96, 0x52, 0x31, 0x2e, 0x02, 0x23, 0x15, 0x7e, 0x8b, 0x6a, 0xd3, 0xbe,
	0xe3, 0xde, 0xdc, 0x21, 0xe1, 0x26, 0x4c, 0xfb, 0xfa, 0xfc, 0x41, 0xe2, 0xa6, 0x44, 0x6f, 0x61,
	0xeb, 0xdf, 0x77, 0xd8, 0xb6, 0xe2, 0xc8, 0xe7, 0x72, 0x10, 0xc9, 0x24, 0x91, 0x62, 0x90, 0x86,
	0x61, 0xd5, 0x85, 0xe2, 0xce, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0x72, 0xa0, 0x8e, 0x41,
	0x05, 0x00, 0x00,
}

// --- DRPC BEGIN ---

type DRPCRepairCoordinatorClient interface {
	DRPCConn() drpc.Conn

	RepairJob(ctx context.Context, in *RepairJobRequest) (*RepairJobResponse, error)
}

type drpcRepairCoordinatorClient struct {
	cc drpc.Conn
}

func NewDRPCRepairCoordinatorClient(cc drpc.Conn) DRPCRepairCoordinatorClient {
	return &drpcRepairCoordinatorClient{cc}
}

func (c *drpcRepairCoordinatorClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcRepairCoordinatorClient) RepairJob(ctx context.Context, in *RepairJobRequest) (*RepairJobResponse, error) {
	out := new(RepairJobResponse)
	err := c.cc.Invoke(ctx, "/delegated_repair.RepairCoordinator/RepairJob", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCRepairCoordinatorServer interface {
	RepairJob(context.Context, *RepairJobRequest) (*RepairJobResponse, error)
}

type DRPCRepairCoordinatorDescription struct{}

func (DRPCRepairCoordinatorDescription) NumMethods() int { return 1 }

func (DRPCRepairCoordinatorDescription) Method(n int) (string, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/delegated_repair.RepairCoordinator/RepairJob",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCRepairCoordinatorServer).
					RepairJob(
						ctx,
						in1.(*RepairJobRequest),
					)
			}, DRPCRepairCoordinatorServer.RepairJob, true
	default:
		return "", nil, nil, false
	}
}

func DRPCRegisterRepairCoordinator(mux drpc.Mux, impl DRPCRepairCoordinatorServer) error {
	return mux.Register(impl, DRPCRepairCoordinatorDescription{})
}

type DRPCRepairCoordinator_RepairJobStream interface {
	drpc.Stream
	SendAndClose(*RepairJobResponse) error
}

type drpcRepairCoordinatorRepairJobStream struct {
	drpc.Stream
}

func (x *drpcRepairCoordinatorRepairJobStream) SendAndClose(m *RepairJobResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

// --- DRPC END ---