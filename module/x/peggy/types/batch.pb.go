// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: module/peggy/v1beta1/batch.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
	time "time"
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

type BatchStatus int32

const (
	BATCH_STATUS_UNKNOWN   BatchStatus = 0
	BATCH_STATUS_PENDING   BatchStatus = 1
	BATCH_STATUS_SUBMITTED BatchStatus = 2
	BATCH_STATUS_PROCESSED BatchStatus = 3
	BATCH_STATUS_CANCELED  BatchStatus = 4
)

var BatchStatus_name = map[int32]string{
	0: "BATCH_STATUS_UNKNOWN",
	1: "BATCH_STATUS_PENDING",
	2: "BATCH_STATUS_SUBMITTED",
	3: "BATCH_STATUS_PROCESSED",
	4: "BATCH_STATUS_CANCELED",
}

var BatchStatus_value = map[string]int32{
	"BATCH_STATUS_UNKNOWN":   0,
	"BATCH_STATUS_PENDING":   1,
	"BATCH_STATUS_SUBMITTED": 2,
	"BATCH_STATUS_PROCESSED": 3,
	"BATCH_STATUS_CANCELED":  4,
}

func (x BatchStatus) String() string {
	return proto.EnumName(BatchStatus_name, int32(x))
}

func (BatchStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9735c82168ab3388, []int{0}
}

type OutgoingTxBatch struct {
	Nonce              uint64                `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Elements           []*OutgoingTransferTx `protobuf:"bytes,2,rep,name=elements,proto3" json:"elements,omitempty"`
	CreatedAt          time.Time             `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
	TotalFee           *ERC20Token           `protobuf:"bytes,4,opt,name=total_fee,json=totalFee,proto3" json:"total_fee,omitempty"`
	BridgedDenominator *BridgedDenominator   `protobuf:"bytes,5,opt,name=bridged_denominator,json=bridgedDenominator,proto3" json:"bridged_denominator,omitempty"`
	BatchStatus        BatchStatus           `protobuf:"varint,6,opt,name=batch_status,json=batchStatus,proto3,enum=module.peggy.v1beta1.BatchStatus" json:"batch_status,omitempty"`
	Valset             *Valset               `protobuf:"bytes,7,opt,name=valset,proto3" json:"valset,omitempty"`
	TokenContract      []byte                `protobuf:"bytes,8,opt,name=token_contract,json=tokenContract,proto3" json:"token_contract,omitempty"`
}

func (m *OutgoingTxBatch) Reset()         { *m = OutgoingTxBatch{} }
func (m *OutgoingTxBatch) String() string { return proto.CompactTextString(m) }
func (*OutgoingTxBatch) ProtoMessage()    {}
func (*OutgoingTxBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_9735c82168ab3388, []int{0}
}
func (m *OutgoingTxBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutgoingTxBatch.Unmarshal(m, b)
}
func (m *OutgoingTxBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutgoingTxBatch.Marshal(b, m, deterministic)
}
func (m *OutgoingTxBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutgoingTxBatch.Merge(m, src)
}
func (m *OutgoingTxBatch) XXX_Size() int {
	return xxx_messageInfo_OutgoingTxBatch.Size(m)
}
func (m *OutgoingTxBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_OutgoingTxBatch.DiscardUnknown(m)
}

var xxx_messageInfo_OutgoingTxBatch proto.InternalMessageInfo

func (m *OutgoingTxBatch) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *OutgoingTxBatch) GetElements() []*OutgoingTransferTx {
	if m != nil {
		return m.Elements
	}
	return nil
}

func (m *OutgoingTxBatch) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *OutgoingTxBatch) GetTotalFee() *ERC20Token {
	if m != nil {
		return m.TotalFee
	}
	return nil
}

func (m *OutgoingTxBatch) GetBridgedDenominator() *BridgedDenominator {
	if m != nil {
		return m.BridgedDenominator
	}
	return nil
}

func (m *OutgoingTxBatch) GetBatchStatus() BatchStatus {
	if m != nil {
		return m.BatchStatus
	}
	return BATCH_STATUS_UNKNOWN
}

func (m *OutgoingTxBatch) GetValset() *Valset {
	if m != nil {
		return m.Valset
	}
	return nil
}

func (m *OutgoingTxBatch) GetTokenContract() []byte {
	if m != nil {
		return m.TokenContract
	}
	return nil
}

type OutgoingTransferTx struct {
	Id          uint64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender      string      `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	DestAddress []byte      `protobuf:"bytes,3,opt,name=dest_address,json=destAddress,proto3" json:"dest_address,omitempty"`
	Amount      *ERC20Token `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	BridgeFee   *ERC20Token `protobuf:"bytes,5,opt,name=bridge_fee,json=bridgeFee,proto3" json:"bridge_fee,omitempty"`
}

func (m *OutgoingTransferTx) Reset()         { *m = OutgoingTransferTx{} }
func (m *OutgoingTransferTx) String() string { return proto.CompactTextString(m) }
func (*OutgoingTransferTx) ProtoMessage()    {}
func (*OutgoingTransferTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_9735c82168ab3388, []int{1}
}
func (m *OutgoingTransferTx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutgoingTransferTx.Unmarshal(m, b)
}
func (m *OutgoingTransferTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutgoingTransferTx.Marshal(b, m, deterministic)
}
func (m *OutgoingTransferTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutgoingTransferTx.Merge(m, src)
}
func (m *OutgoingTransferTx) XXX_Size() int {
	return xxx_messageInfo_OutgoingTransferTx.Size(m)
}
func (m *OutgoingTransferTx) XXX_DiscardUnknown() {
	xxx_messageInfo_OutgoingTransferTx.DiscardUnknown(m)
}

var xxx_messageInfo_OutgoingTransferTx proto.InternalMessageInfo

func (m *OutgoingTransferTx) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OutgoingTransferTx) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *OutgoingTransferTx) GetDestAddress() []byte {
	if m != nil {
		return m.DestAddress
	}
	return nil
}

func (m *OutgoingTransferTx) GetAmount() *ERC20Token {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *OutgoingTransferTx) GetBridgeFee() *ERC20Token {
	if m != nil {
		return m.BridgeFee
	}
	return nil
}

func init() {
	proto.RegisterEnum("module.peggy.v1beta1.BatchStatus", BatchStatus_name, BatchStatus_value)
	proto.RegisterType((*OutgoingTxBatch)(nil), "module.peggy.v1beta1.OutgoingTxBatch")
	proto.RegisterType((*OutgoingTransferTx)(nil), "module.peggy.v1beta1.OutgoingTransferTx")
}

func init() { proto.RegisterFile("module/peggy/v1beta1/batch.proto", fileDescriptor_9735c82168ab3388) }

var fileDescriptor_9735c82168ab3388 = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6e, 0xd3, 0x5a,
	0x10, 0xc6, 0xed, 0x24, 0xcd, 0x4d, 0x4e, 0x72, 0x7b, 0xa3, 0x73, 0x73, 0x2b, 0xdf, 0x08, 0x39,
	0x6e, 0x25, 0x90, 0x85, 0x84, 0xd3, 0x06, 0x16, 0x08, 0x09, 0xa1, 0xc4, 0x76, 0xa1, 0xfc, 0x49,
	0x2b, 0xc7, 0x05, 0xc1, 0xc6, 0x3a, 0x8e, 0xa7, 0xae, 0x45, 0xec, 0x13, 0xd9, 0x27, 0x55, 0xfb,
	0x06, 0x5d, 0x76, 0xcd, 0x8a, 0x8a, 0x97, 0xe9, 0xb2, 0x4b, 0x56, 0x80, 0xda, 0x0d, 0x2f, 0x81,
	0x84, 0x7c, 0xec, 0xb6, 0x54, 0x49, 0xa5, 0xee, 0x3c, 0x33, 0xbf, 0x6f, 0x34, 0xdf, 0x19, 0x0f,
	0x52, 0x42, 0xea, 0x4d, 0xc7, 0xd0, 0x99, 0x80, 0xef, 0x1f, 0x74, 0xf6, 0xd6, 0x5c, 0x60, 0x64,
	0xad, 0xe3, 0x12, 0x36, 0xda, 0xd5, 0x26, 0x31, 0x65, 0x14, 0x37, 0x33, 0x42, 0xe3, 0x84, 0x96,
	0x13, 0xad, 0xa6, 0x4f, 0x7d, 0xca, 0x81, 0x4e, 0xfa, 0x95, 0xb1, 0xad, 0x7b, 0x73, 0xbb, 0x11,
	0xc6, 0x20, 0x61, 0x84, 0x05, 0x34, 0xca, 0xb9, 0xb6, 0x4f, 0xa9, 0x9f, 0x72, 0x69, 0xe4, 0x4e,
	0x77, 0x3a, 0x2c, 0x08, 0x53, 0x24, 0x9c, 0x5c, 0x00, 0x73, 0x1b, 0x4d, 0x28, 0x1d, 0xe7, 0xc0,
	0xfc, 0xb9, 0xd9, 0xc1, 0x04, 0x92, 0x8c, 0x58, 0xf9, 0x55, 0x44, 0xff, 0x6c, 0x4e, 0x99, 0x4f,
	0x83, 0xc8, 0xb7, 0xf7, 0xfb, 0xa9, 0x23, 0xdc, 0x44, 0x0b, 0x11, 0x8d, 0x46, 0x20, 0x89, 0x8a,
	0xa8, 0x96, 0xac, 0x2c, 0xc0, 0x06, 0xaa, 0xc0, 0x18, 0x42, 0x88, 0x58, 0x22, 0x15, 0x94, 0xa2,
	0x5a, 0xeb, 0xaa, 0xda, 0x3c, 0xd3, 0xda, 0x65, 0xbb, 0x98, 0x44, 0xc9, 0x0e, 0xc4, 0xf6, 0xbe,
	0x75, 0xa9, 0xc4, 0x3a, 0x42, 0xa3, 0x18, 0x08, 0x03, 0xcf, 0x21, 0x4c, 0x2a, 0x2a, 0xa2, 0x5a,
	0xeb, 0xb6, 0xb4, 0xcc, 0xa8, 0x76, 0x61, 0x54, 0xb3, 0x2f, 0x8c, 0xf6, 0x2b, 0x27, 0xdf, 0xda,
	0xc2, 0xd1, 0xf7, 0xb6, 0x68, 0x55, 0x73, 0x5d, 0x8f, 0xe1, 0xa7, 0xa8, 0xca, 0x28, 0x23, 0x63,
	0x67, 0x07, 0x40, 0x2a, 0xf1, 0x1e, 0xca, 0xfc, 0x59, 0x4c, 0x4b, 0xef, 0xae, 0xda, 0xf4, 0x23,
	0x44, 0x56, 0x85, 0x4b, 0xd6, 0x01, 0xf0, 0x7b, 0xf4, 0xaf, 0x1b, 0x07, 0x9e, 0x0f, 0x9e, 0xe3,
	0x41, 0x44, 0xc3, 0x20, 0x22, 0x8c, 0xc6, 0xd2, 0x02, 0x6f, 0x74, 0x83, 0xa9, 0x7e, 0x26, 0x30,
	0xae, 0x78, 0x0b, 0xbb, 0x33, 0x39, 0x6c, 0xa0, 0x3a, 0xff, 0x2b, 0x9c, 0x74, 0x93, 0xd3, 0x44,
	0x2a, 0x2b, 0xa2, 0xba, 0xd8, 0x5d, 0xbe, 0xa1, 0x67, 0x4a, 0x0e, 0x39, 0x68, 0xd5, 0xdc, 0xab,
	0x00, 0x3f, 0x42, 0xe5, 0x3d, 0x32, 0x4e, 0x80, 0x49, 0x7f, 0xf1, 0x99, 0xee, 0xcc, 0xd7, 0xbf,
	0xe5, 0x8c, 0x95, 0xb3, 0xf8, 0x2e, 0x5a, 0x64, 0xa9, 0x53, 0x67, 0x44, 0x23, 0x16, 0x93, 0x11,
	0x93, 0x2a, 0x8a, 0xa8, 0xd6, 0xad, 0xbf, 0x79, 0x56, 0xcf, 0x93, 0x4f, 0xea, 0x87, 0xc7, 0x6d,
	0xe1, 0xe8, 0xb8, 0x2d, 0x7c, 0x3e, 0x6e, 0x0b, 0x2b, 0x3f, 0x45, 0x84, 0x67, 0x17, 0x86, 0x17,
	0x51, 0x21, 0xf0, 0xf2, 0xfd, 0x17, 0x02, 0x0f, 0x2f, 0xa1, 0x72, 0x02, 0x91, 0x07, 0xb1, 0x54,
	0x50, 0x44, 0xb5, 0x6a, 0xe5, 0x11, 0x5e, 0x46, 0x75, 0x0f, 0x12, 0xe6, 0x10, 0xcf, 0x8b, 0x21,
	0x49, 0xf8, 0x42, 0xeb, 0x56, 0x2d, 0xcd, 0xf5, 0xb2, 0x14, 0x7e, 0x8c, 0xca, 0x24, 0xa4, 0xd3,
	0x88, 0xdd, 0x7a, 0x53, 0x39, 0x8f, 0x9f, 0x21, 0x94, 0x3d, 0x31, 0xdf, 0xf3, 0xc2, 0x2d, 0xd5,
	0xd5, 0x4c, 0xb3, 0x0e, 0x70, 0xdd, 0xea, 0xfd, 0x4f, 0x22, 0xaa, 0xfd, 0xf1, 0xe4, 0x58, 0x42,
	0xcd, 0x7e, 0xcf, 0xd6, 0x5f, 0x38, 0x43, 0xbb, 0x67, 0x6f, 0x0f, 0x9d, 0xed, 0xc1, 0xab, 0xc1,
	0xe6, 0xbb, 0x41, 0x43, 0x98, 0xa9, 0x6c, 0x99, 0x03, 0x63, 0x63, 0xf0, 0xbc, 0x21, 0xe2, 0x16,
	0x5a, 0xba, 0x56, 0x19, 0x6e, 0xf7, 0xdf, 0x6c, 0xd8, 0xb6, 0x69, 0x34, 0x0a, 0x33, 0xb5, 0x2d,
	0x6b, 0x53, 0x37, 0x87, 0x43, 0xd3, 0x68, 0x14, 0xf1, 0xff, 0xe8, 0xbf, 0x6b, 0x35, 0xbd, 0x37,
	0xd0, 0xcd, 0xd7, 0xa6, 0xd1, 0x28, 0xb5, 0x4a, 0x87, 0x5f, 0x64, 0xa1, 0xff, 0xf2, 0xe4, 0x4c,
	0x16, 0x4f, 0xcf, 0x64, 0xf1, 0xc7, 0x99, 0x2c, 0x1e, 0x9d, 0xcb, 0xc2, 0xe9, 0xb9, 0x2c, 0x7c,
	0x3d, 0x97, 0x85, 0x0f, 0xab, 0x7e, 0xc0, 0x76, 0xa7, 0xae, 0x36, 0xa2, 0x61, 0x87, 0x8c, 0xd9,
	0x2e, 0x90, 0x07, 0x11, 0xb0, 0xfc, 0xa4, 0xf3, 0xfb, 0xde, 0xcf, 0x43, 0x7e, 0xd9, 0x6e, 0x99,
	0xdf, 0xd1, 0xc3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xee, 0xcc, 0x51, 0x48, 0xb6, 0x04, 0x00,
	0x00,
}
