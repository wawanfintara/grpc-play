// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction.proto

package localservices

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TxFulfillment struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AggregatorId         int64                `protobuf:"varint,2,opt,name=aggregator_id,json=aggregatorId,proto3" json:"aggregator_id,omitempty"`
	TxAggregatorId       int64                `protobuf:"varint,3,opt,name=tx_aggregator_id,json=txAggregatorId,proto3" json:"tx_aggregator_id,omitempty"`
	TxDetStatus          int64                `protobuf:"varint,4,opt,name=tx_det_status,json=txDetStatus,proto3" json:"tx_det_status,omitempty"`
	CreatedTime          *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	UpdatedTime          *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TxFulfillment) Reset()         { *m = TxFulfillment{} }
func (m *TxFulfillment) String() string { return proto.CompactTextString(m) }
func (*TxFulfillment) ProtoMessage()    {}
func (*TxFulfillment) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_328fca2998d6afeb, []int{0}
}
func (m *TxFulfillment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxFulfillment.Unmarshal(m, b)
}
func (m *TxFulfillment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxFulfillment.Marshal(b, m, deterministic)
}
func (dst *TxFulfillment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxFulfillment.Merge(dst, src)
}
func (m *TxFulfillment) XXX_Size() int {
	return xxx_messageInfo_TxFulfillment.Size(m)
}
func (m *TxFulfillment) XXX_DiscardUnknown() {
	xxx_messageInfo_TxFulfillment.DiscardUnknown(m)
}

var xxx_messageInfo_TxFulfillment proto.InternalMessageInfo

func (m *TxFulfillment) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TxFulfillment) GetAggregatorId() int64 {
	if m != nil {
		return m.AggregatorId
	}
	return 0
}

func (m *TxFulfillment) GetTxAggregatorId() int64 {
	if m != nil {
		return m.TxAggregatorId
	}
	return 0
}

func (m *TxFulfillment) GetTxDetStatus() int64 {
	if m != nil {
		return m.TxDetStatus
	}
	return 0
}

func (m *TxFulfillment) GetCreatedTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

func (m *TxFulfillment) GetUpdatedTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedTime
	}
	return nil
}

type CreateTxFulfillmentRequest struct {
	TxFulfillment        *TxFulfillment `protobuf:"bytes,1,opt,name=tx_fulfillment,json=txFulfillment,proto3" json:"tx_fulfillment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateTxFulfillmentRequest) Reset()         { *m = CreateTxFulfillmentRequest{} }
func (m *CreateTxFulfillmentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTxFulfillmentRequest) ProtoMessage()    {}
func (*CreateTxFulfillmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_328fca2998d6afeb, []int{1}
}
func (m *CreateTxFulfillmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTxFulfillmentRequest.Unmarshal(m, b)
}
func (m *CreateTxFulfillmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTxFulfillmentRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTxFulfillmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTxFulfillmentRequest.Merge(dst, src)
}
func (m *CreateTxFulfillmentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTxFulfillmentRequest.Size(m)
}
func (m *CreateTxFulfillmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTxFulfillmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTxFulfillmentRequest proto.InternalMessageInfo

func (m *CreateTxFulfillmentRequest) GetTxFulfillment() *TxFulfillment {
	if m != nil {
		return m.TxFulfillment
	}
	return nil
}

type CreateTxFulfillmentResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTxFulfillmentResponse) Reset()         { *m = CreateTxFulfillmentResponse{} }
func (m *CreateTxFulfillmentResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTxFulfillmentResponse) ProtoMessage()    {}
func (*CreateTxFulfillmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_328fca2998d6afeb, []int{2}
}
func (m *CreateTxFulfillmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTxFulfillmentResponse.Unmarshal(m, b)
}
func (m *CreateTxFulfillmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTxFulfillmentResponse.Marshal(b, m, deterministic)
}
func (dst *CreateTxFulfillmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTxFulfillmentResponse.Merge(dst, src)
}
func (m *CreateTxFulfillmentResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTxFulfillmentResponse.Size(m)
}
func (m *CreateTxFulfillmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTxFulfillmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTxFulfillmentResponse proto.InternalMessageInfo

type GetTxFulfillmentsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTxFulfillmentsRequest) Reset()         { *m = GetTxFulfillmentsRequest{} }
func (m *GetTxFulfillmentsRequest) String() string { return proto.CompactTextString(m) }
func (*GetTxFulfillmentsRequest) ProtoMessage()    {}
func (*GetTxFulfillmentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_328fca2998d6afeb, []int{3}
}
func (m *GetTxFulfillmentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTxFulfillmentsRequest.Unmarshal(m, b)
}
func (m *GetTxFulfillmentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTxFulfillmentsRequest.Marshal(b, m, deterministic)
}
func (dst *GetTxFulfillmentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTxFulfillmentsRequest.Merge(dst, src)
}
func (m *GetTxFulfillmentsRequest) XXX_Size() int {
	return xxx_messageInfo_GetTxFulfillmentsRequest.Size(m)
}
func (m *GetTxFulfillmentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTxFulfillmentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTxFulfillmentsRequest proto.InternalMessageInfo

type GetTxFulfillmentsResponse struct {
	TxFulfillment        []*TxFulfillment `protobuf:"bytes,1,rep,name=tx_fulfillment,json=txFulfillment,proto3" json:"tx_fulfillment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetTxFulfillmentsResponse) Reset()         { *m = GetTxFulfillmentsResponse{} }
func (m *GetTxFulfillmentsResponse) String() string { return proto.CompactTextString(m) }
func (*GetTxFulfillmentsResponse) ProtoMessage()    {}
func (*GetTxFulfillmentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_328fca2998d6afeb, []int{4}
}
func (m *GetTxFulfillmentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTxFulfillmentsResponse.Unmarshal(m, b)
}
func (m *GetTxFulfillmentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTxFulfillmentsResponse.Marshal(b, m, deterministic)
}
func (dst *GetTxFulfillmentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTxFulfillmentsResponse.Merge(dst, src)
}
func (m *GetTxFulfillmentsResponse) XXX_Size() int {
	return xxx_messageInfo_GetTxFulfillmentsResponse.Size(m)
}
func (m *GetTxFulfillmentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTxFulfillmentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTxFulfillmentsResponse proto.InternalMessageInfo

func (m *GetTxFulfillmentsResponse) GetTxFulfillment() []*TxFulfillment {
	if m != nil {
		return m.TxFulfillment
	}
	return nil
}

type GetTxFulfillmentByIdRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTxFulfillmentByIdRequest) Reset()         { *m = GetTxFulfillmentByIdRequest{} }
func (m *GetTxFulfillmentByIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetTxFulfillmentByIdRequest) ProtoMessage()    {}
func (*GetTxFulfillmentByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_328fca2998d6afeb, []int{5}
}
func (m *GetTxFulfillmentByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTxFulfillmentByIdRequest.Unmarshal(m, b)
}
func (m *GetTxFulfillmentByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTxFulfillmentByIdRequest.Marshal(b, m, deterministic)
}
func (dst *GetTxFulfillmentByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTxFulfillmentByIdRequest.Merge(dst, src)
}
func (m *GetTxFulfillmentByIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetTxFulfillmentByIdRequest.Size(m)
}
func (m *GetTxFulfillmentByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTxFulfillmentByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTxFulfillmentByIdRequest proto.InternalMessageInfo

func (m *GetTxFulfillmentByIdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetTxFulfillmentByIdResponse struct {
	TxFulfillment        *TxFulfillment `protobuf:"bytes,1,opt,name=tx_fulfillment,json=txFulfillment,proto3" json:"tx_fulfillment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetTxFulfillmentByIdResponse) Reset()         { *m = GetTxFulfillmentByIdResponse{} }
func (m *GetTxFulfillmentByIdResponse) String() string { return proto.CompactTextString(m) }
func (*GetTxFulfillmentByIdResponse) ProtoMessage()    {}
func (*GetTxFulfillmentByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_328fca2998d6afeb, []int{6}
}
func (m *GetTxFulfillmentByIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTxFulfillmentByIdResponse.Unmarshal(m, b)
}
func (m *GetTxFulfillmentByIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTxFulfillmentByIdResponse.Marshal(b, m, deterministic)
}
func (dst *GetTxFulfillmentByIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTxFulfillmentByIdResponse.Merge(dst, src)
}
func (m *GetTxFulfillmentByIdResponse) XXX_Size() int {
	return xxx_messageInfo_GetTxFulfillmentByIdResponse.Size(m)
}
func (m *GetTxFulfillmentByIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTxFulfillmentByIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTxFulfillmentByIdResponse proto.InternalMessageInfo

func (m *GetTxFulfillmentByIdResponse) GetTxFulfillment() *TxFulfillment {
	if m != nil {
		return m.TxFulfillment
	}
	return nil
}

type UpdateTxFulfillmentRequest struct {
	TxFulfillment        *TxFulfillment `protobuf:"bytes,1,opt,name=tx_fulfillment,json=txFulfillment,proto3" json:"tx_fulfillment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateTxFulfillmentRequest) Reset()         { *m = UpdateTxFulfillmentRequest{} }
func (m *UpdateTxFulfillmentRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTxFulfillmentRequest) ProtoMessage()    {}
func (*UpdateTxFulfillmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_328fca2998d6afeb, []int{7}
}
func (m *UpdateTxFulfillmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTxFulfillmentRequest.Unmarshal(m, b)
}
func (m *UpdateTxFulfillmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTxFulfillmentRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateTxFulfillmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTxFulfillmentRequest.Merge(dst, src)
}
func (m *UpdateTxFulfillmentRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTxFulfillmentRequest.Size(m)
}
func (m *UpdateTxFulfillmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTxFulfillmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTxFulfillmentRequest proto.InternalMessageInfo

func (m *UpdateTxFulfillmentRequest) GetTxFulfillment() *TxFulfillment {
	if m != nil {
		return m.TxFulfillment
	}
	return nil
}

type UpdateTxFulfillmentResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTxFulfillmentResponse) Reset()         { *m = UpdateTxFulfillmentResponse{} }
func (m *UpdateTxFulfillmentResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTxFulfillmentResponse) ProtoMessage()    {}
func (*UpdateTxFulfillmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_328fca2998d6afeb, []int{8}
}
func (m *UpdateTxFulfillmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTxFulfillmentResponse.Unmarshal(m, b)
}
func (m *UpdateTxFulfillmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTxFulfillmentResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateTxFulfillmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTxFulfillmentResponse.Merge(dst, src)
}
func (m *UpdateTxFulfillmentResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTxFulfillmentResponse.Size(m)
}
func (m *UpdateTxFulfillmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTxFulfillmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTxFulfillmentResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TxFulfillment)(nil), "TxFulfillment")
	proto.RegisterType((*CreateTxFulfillmentRequest)(nil), "CreateTxFulfillmentRequest")
	proto.RegisterType((*CreateTxFulfillmentResponse)(nil), "CreateTxFulfillmentResponse")
	proto.RegisterType((*GetTxFulfillmentsRequest)(nil), "GetTxFulfillmentsRequest")
	proto.RegisterType((*GetTxFulfillmentsResponse)(nil), "GetTxFulfillmentsResponse")
	proto.RegisterType((*GetTxFulfillmentByIdRequest)(nil), "GetTxFulfillmentByIdRequest")
	proto.RegisterType((*GetTxFulfillmentByIdResponse)(nil), "GetTxFulfillmentByIdResponse")
	proto.RegisterType((*UpdateTxFulfillmentRequest)(nil), "UpdateTxFulfillmentRequest")
	proto.RegisterType((*UpdateTxFulfillmentResponse)(nil), "UpdateTxFulfillmentResponse")
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor_transaction_328fca2998d6afeb) }

var fileDescriptor_transaction_328fca2998d6afeb = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x4f, 0xab, 0x40,
	0x10, 0xc6, 0x53, 0xfa, 0x5e, 0x0f, 0x43, 0xe1, 0x3d, 0x39, 0x21, 0xd5, 0xd8, 0xe0, 0xa5, 0x17,
	0x69, 0x52, 0xe3, 0xd1, 0x83, 0xd5, 0x68, 0x7a, 0xa5, 0xed, 0xc5, 0x0b, 0xd9, 0xb2, 0x53, 0x42,
	0x42, 0x59, 0x64, 0x07, 0x83, 0xff, 0x85, 0x7f, 0xb2, 0x61, 0x69, 0xd3, 0xd2, 0x10, 0xeb, 0xc1,
	0xe3, 0xce, 0xfc, 0x66, 0xf6, 0xfb, 0xbe, 0x0c, 0x9c, 0x51, 0xce, 0x52, 0xc9, 0x42, 0x8a, 0x45,
	0xea, 0x65, 0xb9, 0x20, 0xe1, 0x5c, 0x45, 0x42, 0x44, 0x09, 0x8e, 0xd5, 0x6b, 0x55, 0xac, 0xc7,
	0x14, 0x6f, 0x50, 0x12, 0xdb, 0x64, 0x35, 0xe0, 0x7e, 0x6a, 0x60, 0x2c, 0xca, 0xe7, 0x22, 0x59,
	0xc7, 0x49, 0xb2, 0xc1, 0x94, 0x2c, 0x13, 0xb4, 0x98, 0xdb, 0x9d, 0x61, 0x67, 0xd4, 0xf5, 0xb5,
	0x98, 0x5b, 0xd7, 0x60, 0xb0, 0x28, 0xca, 0x31, 0x62, 0x24, 0xf2, 0x20, 0xe6, 0xb6, 0xa6, 0x5a,
	0xfd, 0x7d, 0x71, 0xc6, 0xad, 0x11, 0xfc, 0xa7, 0x32, 0x68, 0x72, 0x5d, 0xc5, 0x99, 0x54, 0x3e,
	0x1c, 0x92, 0x2e, 0x18, 0x54, 0x06, 0x1c, 0x29, 0x90, 0xc4, 0xa8, 0x90, 0xf6, 0x1f, 0x85, 0xe9,
	0x54, 0x3e, 0x21, 0xcd, 0x55, 0xc9, 0xba, 0x87, 0x7e, 0x98, 0x23, 0x23, 0xe4, 0x41, 0xa5, 0xd7,
	0xfe, 0x3b, 0xec, 0x8c, 0xf4, 0x89, 0xe3, 0xd5, 0x66, 0xbc, 0x9d, 0x19, 0x6f, 0xb1, 0x33, 0xe3,
	0xeb, 0x5b, 0xbe, 0xaa, 0x54, 0xe3, 0x45, 0xc6, 0xf7, 0xe3, 0xbd, 0xd3, 0xe3, 0x5b, 0xbe, 0xaa,
	0xb8, 0x73, 0x70, 0x1e, 0xd5, 0xb6, 0x46, 0x2e, 0x3e, 0xbe, 0x15, 0x28, 0xc9, 0xba, 0x03, 0x93,
	0xca, 0x60, 0xbd, 0x6f, 0xa8, 0xa8, 0xf4, 0x89, 0xe9, 0x35, 0x71, 0x83, 0x0e, 0x9f, 0xee, 0x25,
	0x0c, 0x5a, 0x97, 0xca, 0x4c, 0xa4, 0x12, 0x5d, 0x07, 0xec, 0x17, 0xa4, 0x46, 0x4f, 0x6e, 0x7f,
	0x74, 0x7d, 0x38, 0x6f, 0xe9, 0xd5, 0x83, 0xad, 0x72, 0xba, 0xa7, 0xe5, 0xdc, 0xc0, 0xe0, 0x78,
	0xe7, 0xf4, 0x63, 0xc6, 0x77, 0x26, 0x8f, 0x6e, 0xc0, 0x5d, 0xc2, 0x45, 0x3b, 0xfe, 0x8d, 0x8a,
	0x1f, 0x84, 0x32, 0x07, 0x67, 0xa9, 0x82, 0xff, 0xe5, 0xa4, 0x5b, 0x97, 0xd6, 0x52, 0xa7, 0xff,
	0x5e, 0x8d, 0x44, 0x84, 0x2c, 0x91, 0x98, 0xbf, 0xc7, 0x21, 0xca, 0x55, 0x4f, 0xdd, 0xc3, 0xed,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0x18, 0x44, 0xc3, 0x3e, 0x03, 0x00, 0x00,
}
