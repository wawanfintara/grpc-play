// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payment.proto

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

type Payment struct {
	Id                   int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Invoice              string           `protobuf:"bytes,2,opt,name=invoice,proto3" json:"invoice,omitempty"`
	Amount               float32          `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Status               int64            `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	PaidAmount           float32          `protobuf:"fixed32,5,opt,name=paid_amount,json=paidAmount,proto3" json:"paid_amount,omitempty"`
	Method               int64            `protobuf:"varint,6,opt,name=method,proto3" json:"method,omitempty"`
	Details              []*PaymentDetail `protobuf:"bytes,7,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Payment) Reset()         { *m = Payment{} }
func (m *Payment) String() string { return proto.CompactTextString(m) }
func (*Payment) ProtoMessage()    {}
func (*Payment) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8400961b101b93a7, []int{0}
}
func (m *Payment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payment.Unmarshal(m, b)
}
func (m *Payment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payment.Marshal(b, m, deterministic)
}
func (dst *Payment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payment.Merge(dst, src)
}
func (m *Payment) XXX_Size() int {
	return xxx_messageInfo_Payment.Size(m)
}
func (m *Payment) XXX_DiscardUnknown() {
	xxx_messageInfo_Payment.DiscardUnknown(m)
}

var xxx_messageInfo_Payment proto.InternalMessageInfo

func (m *Payment) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Payment) GetInvoice() string {
	if m != nil {
		return m.Invoice
	}
	return ""
}

func (m *Payment) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Payment) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Payment) GetPaidAmount() float32 {
	if m != nil {
		return m.PaidAmount
	}
	return 0
}

func (m *Payment) GetMethod() int64 {
	if m != nil {
		return m.Method
	}
	return 0
}

func (m *Payment) GetDetails() []*PaymentDetail {
	if m != nil {
		return m.Details
	}
	return nil
}

type PaymentDetail struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PaymentId            int64                `protobuf:"varint,2,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	TxId                 int64                `protobuf:"varint,3,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	TxType               int64                `protobuf:"varint,4,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty"`
	TxName               string               `protobuf:"bytes,5,opt,name=tx_name,json=txName,proto3" json:"tx_name,omitempty"`
	TxCreatedTime        *timestamp.Timestamp `protobuf:"bytes,7,opt,name=tx_created_time,json=txCreatedTime,proto3" json:"tx_created_time,omitempty"`
	TxUpdatedTime        *timestamp.Timestamp `protobuf:"bytes,8,opt,name=tx_updated_time,json=txUpdatedTime,proto3" json:"tx_updated_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PaymentDetail) Reset()         { *m = PaymentDetail{} }
func (m *PaymentDetail) String() string { return proto.CompactTextString(m) }
func (*PaymentDetail) ProtoMessage()    {}
func (*PaymentDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8400961b101b93a7, []int{1}
}
func (m *PaymentDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentDetail.Unmarshal(m, b)
}
func (m *PaymentDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentDetail.Marshal(b, m, deterministic)
}
func (dst *PaymentDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentDetail.Merge(dst, src)
}
func (m *PaymentDetail) XXX_Size() int {
	return xxx_messageInfo_PaymentDetail.Size(m)
}
func (m *PaymentDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentDetail.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentDetail proto.InternalMessageInfo

func (m *PaymentDetail) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PaymentDetail) GetPaymentId() int64 {
	if m != nil {
		return m.PaymentId
	}
	return 0
}

func (m *PaymentDetail) GetTxId() int64 {
	if m != nil {
		return m.TxId
	}
	return 0
}

func (m *PaymentDetail) GetTxType() int64 {
	if m != nil {
		return m.TxType
	}
	return 0
}

func (m *PaymentDetail) GetTxName() string {
	if m != nil {
		return m.TxName
	}
	return ""
}

func (m *PaymentDetail) GetTxCreatedTime() *timestamp.Timestamp {
	if m != nil {
		return m.TxCreatedTime
	}
	return nil
}

func (m *PaymentDetail) GetTxUpdatedTime() *timestamp.Timestamp {
	if m != nil {
		return m.TxUpdatedTime
	}
	return nil
}

type PaymentDetailAppend struct {
	TxId                 int64    `protobuf:"varint,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	TxType               int64    `protobuf:"varint,2,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty"`
	TxName               string   `protobuf:"bytes,3,opt,name=tx_name,json=txName,proto3" json:"tx_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentDetailAppend) Reset()         { *m = PaymentDetailAppend{} }
func (m *PaymentDetailAppend) String() string { return proto.CompactTextString(m) }
func (*PaymentDetailAppend) ProtoMessage()    {}
func (*PaymentDetailAppend) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8400961b101b93a7, []int{2}
}
func (m *PaymentDetailAppend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentDetailAppend.Unmarshal(m, b)
}
func (m *PaymentDetailAppend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentDetailAppend.Marshal(b, m, deterministic)
}
func (dst *PaymentDetailAppend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentDetailAppend.Merge(dst, src)
}
func (m *PaymentDetailAppend) XXX_Size() int {
	return xxx_messageInfo_PaymentDetailAppend.Size(m)
}
func (m *PaymentDetailAppend) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentDetailAppend.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentDetailAppend proto.InternalMessageInfo

func (m *PaymentDetailAppend) GetTxId() int64 {
	if m != nil {
		return m.TxId
	}
	return 0
}

func (m *PaymentDetailAppend) GetTxType() int64 {
	if m != nil {
		return m.TxType
	}
	return 0
}

func (m *PaymentDetailAppend) GetTxName() string {
	if m != nil {
		return m.TxName
	}
	return ""
}

type CreatePaymentRequest struct {
	Invoice              string                 `protobuf:"bytes,1,opt,name=invoice,proto3" json:"invoice,omitempty"`
	Amount               float32                `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Status               int64                  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	PaidAmount           float32                `protobuf:"fixed32,4,opt,name=paid_amount,json=paidAmount,proto3" json:"paid_amount,omitempty"`
	Method               int64                  `protobuf:"varint,5,opt,name=method,proto3" json:"method,omitempty"`
	Details              []*PaymentDetailAppend `protobuf:"bytes,6,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CreatePaymentRequest) Reset()         { *m = CreatePaymentRequest{} }
func (m *CreatePaymentRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePaymentRequest) ProtoMessage()    {}
func (*CreatePaymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8400961b101b93a7, []int{3}
}
func (m *CreatePaymentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePaymentRequest.Unmarshal(m, b)
}
func (m *CreatePaymentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePaymentRequest.Marshal(b, m, deterministic)
}
func (dst *CreatePaymentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePaymentRequest.Merge(dst, src)
}
func (m *CreatePaymentRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePaymentRequest.Size(m)
}
func (m *CreatePaymentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePaymentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePaymentRequest proto.InternalMessageInfo

func (m *CreatePaymentRequest) GetInvoice() string {
	if m != nil {
		return m.Invoice
	}
	return ""
}

func (m *CreatePaymentRequest) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *CreatePaymentRequest) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CreatePaymentRequest) GetPaidAmount() float32 {
	if m != nil {
		return m.PaidAmount
	}
	return 0
}

func (m *CreatePaymentRequest) GetMethod() int64 {
	if m != nil {
		return m.Method
	}
	return 0
}

func (m *CreatePaymentRequest) GetDetails() []*PaymentDetailAppend {
	if m != nil {
		return m.Details
	}
	return nil
}

type CreatePaymentResponse struct {
	PaymentId            int64    `protobuf:"varint,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePaymentResponse) Reset()         { *m = CreatePaymentResponse{} }
func (m *CreatePaymentResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePaymentResponse) ProtoMessage()    {}
func (*CreatePaymentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8400961b101b93a7, []int{4}
}
func (m *CreatePaymentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePaymentResponse.Unmarshal(m, b)
}
func (m *CreatePaymentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePaymentResponse.Marshal(b, m, deterministic)
}
func (dst *CreatePaymentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePaymentResponse.Merge(dst, src)
}
func (m *CreatePaymentResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePaymentResponse.Size(m)
}
func (m *CreatePaymentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePaymentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePaymentResponse proto.InternalMessageInfo

func (m *CreatePaymentResponse) GetPaymentId() int64 {
	if m != nil {
		return m.PaymentId
	}
	return 0
}

type CreatePaymentDetailRequest struct {
	PaymentId            int64                `protobuf:"varint,1,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	TxId                 int64                `protobuf:"varint,2,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	TxType               int64                `protobuf:"varint,3,opt,name=tx_type,json=txType,proto3" json:"tx_type,omitempty"`
	TxName               string               `protobuf:"bytes,4,opt,name=tx_name,json=txName,proto3" json:"tx_name,omitempty"`
	TxCreatedTime        *timestamp.Timestamp `protobuf:"bytes,5,opt,name=tx_created_time,json=txCreatedTime,proto3" json:"tx_created_time,omitempty"`
	TxUpdatedTime        *timestamp.Timestamp `protobuf:"bytes,6,opt,name=tx_updated_time,json=txUpdatedTime,proto3" json:"tx_updated_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreatePaymentDetailRequest) Reset()         { *m = CreatePaymentDetailRequest{} }
func (m *CreatePaymentDetailRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePaymentDetailRequest) ProtoMessage()    {}
func (*CreatePaymentDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8400961b101b93a7, []int{5}
}
func (m *CreatePaymentDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePaymentDetailRequest.Unmarshal(m, b)
}
func (m *CreatePaymentDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePaymentDetailRequest.Marshal(b, m, deterministic)
}
func (dst *CreatePaymentDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePaymentDetailRequest.Merge(dst, src)
}
func (m *CreatePaymentDetailRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePaymentDetailRequest.Size(m)
}
func (m *CreatePaymentDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePaymentDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePaymentDetailRequest proto.InternalMessageInfo

func (m *CreatePaymentDetailRequest) GetPaymentId() int64 {
	if m != nil {
		return m.PaymentId
	}
	return 0
}

func (m *CreatePaymentDetailRequest) GetTxId() int64 {
	if m != nil {
		return m.TxId
	}
	return 0
}

func (m *CreatePaymentDetailRequest) GetTxType() int64 {
	if m != nil {
		return m.TxType
	}
	return 0
}

func (m *CreatePaymentDetailRequest) GetTxName() string {
	if m != nil {
		return m.TxName
	}
	return ""
}

func (m *CreatePaymentDetailRequest) GetTxCreatedTime() *timestamp.Timestamp {
	if m != nil {
		return m.TxCreatedTime
	}
	return nil
}

func (m *CreatePaymentDetailRequest) GetTxUpdatedTime() *timestamp.Timestamp {
	if m != nil {
		return m.TxUpdatedTime
	}
	return nil
}

type CreatePaymentDetailResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePaymentDetailResponse) Reset()         { *m = CreatePaymentDetailResponse{} }
func (m *CreatePaymentDetailResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePaymentDetailResponse) ProtoMessage()    {}
func (*CreatePaymentDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8400961b101b93a7, []int{6}
}
func (m *CreatePaymentDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePaymentDetailResponse.Unmarshal(m, b)
}
func (m *CreatePaymentDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePaymentDetailResponse.Marshal(b, m, deterministic)
}
func (dst *CreatePaymentDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePaymentDetailResponse.Merge(dst, src)
}
func (m *CreatePaymentDetailResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePaymentDetailResponse.Size(m)
}
func (m *CreatePaymentDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePaymentDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePaymentDetailResponse proto.InternalMessageInfo

type GetPaymentRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPaymentRequest) Reset()         { *m = GetPaymentRequest{} }
func (m *GetPaymentRequest) String() string { return proto.CompactTextString(m) }
func (*GetPaymentRequest) ProtoMessage()    {}
func (*GetPaymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8400961b101b93a7, []int{7}
}
func (m *GetPaymentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPaymentRequest.Unmarshal(m, b)
}
func (m *GetPaymentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPaymentRequest.Marshal(b, m, deterministic)
}
func (dst *GetPaymentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPaymentRequest.Merge(dst, src)
}
func (m *GetPaymentRequest) XXX_Size() int {
	return xxx_messageInfo_GetPaymentRequest.Size(m)
}
func (m *GetPaymentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPaymentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPaymentRequest proto.InternalMessageInfo

func (m *GetPaymentRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetPaymentResponse struct {
	Payment              *Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPaymentResponse) Reset()         { *m = GetPaymentResponse{} }
func (m *GetPaymentResponse) String() string { return proto.CompactTextString(m) }
func (*GetPaymentResponse) ProtoMessage()    {}
func (*GetPaymentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8400961b101b93a7, []int{8}
}
func (m *GetPaymentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPaymentResponse.Unmarshal(m, b)
}
func (m *GetPaymentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPaymentResponse.Marshal(b, m, deterministic)
}
func (dst *GetPaymentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPaymentResponse.Merge(dst, src)
}
func (m *GetPaymentResponse) XXX_Size() int {
	return xxx_messageInfo_GetPaymentResponse.Size(m)
}
func (m *GetPaymentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPaymentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPaymentResponse proto.InternalMessageInfo

func (m *GetPaymentResponse) GetPayment() *Payment {
	if m != nil {
		return m.Payment
	}
	return nil
}

type GetPaymentDetailRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPaymentDetailRequest) Reset()         { *m = GetPaymentDetailRequest{} }
func (m *GetPaymentDetailRequest) String() string { return proto.CompactTextString(m) }
func (*GetPaymentDetailRequest) ProtoMessage()    {}
func (*GetPaymentDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8400961b101b93a7, []int{9}
}
func (m *GetPaymentDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPaymentDetailRequest.Unmarshal(m, b)
}
func (m *GetPaymentDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPaymentDetailRequest.Marshal(b, m, deterministic)
}
func (dst *GetPaymentDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPaymentDetailRequest.Merge(dst, src)
}
func (m *GetPaymentDetailRequest) XXX_Size() int {
	return xxx_messageInfo_GetPaymentDetailRequest.Size(m)
}
func (m *GetPaymentDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPaymentDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPaymentDetailRequest proto.InternalMessageInfo

func (m *GetPaymentDetailRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetPaymentDetailResponse struct {
	Detail               *PaymentDetail `protobuf:"bytes,1,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetPaymentDetailResponse) Reset()         { *m = GetPaymentDetailResponse{} }
func (m *GetPaymentDetailResponse) String() string { return proto.CompactTextString(m) }
func (*GetPaymentDetailResponse) ProtoMessage()    {}
func (*GetPaymentDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8400961b101b93a7, []int{10}
}
func (m *GetPaymentDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPaymentDetailResponse.Unmarshal(m, b)
}
func (m *GetPaymentDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPaymentDetailResponse.Marshal(b, m, deterministic)
}
func (dst *GetPaymentDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPaymentDetailResponse.Merge(dst, src)
}
func (m *GetPaymentDetailResponse) XXX_Size() int {
	return xxx_messageInfo_GetPaymentDetailResponse.Size(m)
}
func (m *GetPaymentDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPaymentDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPaymentDetailResponse proto.InternalMessageInfo

func (m *GetPaymentDetailResponse) GetDetail() *PaymentDetail {
	if m != nil {
		return m.Detail
	}
	return nil
}

type GetPaymentDetailsResponse struct {
	Details              []*PaymentDetail `protobuf:"bytes,1,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetPaymentDetailsResponse) Reset()         { *m = GetPaymentDetailsResponse{} }
func (m *GetPaymentDetailsResponse) String() string { return proto.CompactTextString(m) }
func (*GetPaymentDetailsResponse) ProtoMessage()    {}
func (*GetPaymentDetailsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8400961b101b93a7, []int{11}
}
func (m *GetPaymentDetailsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPaymentDetailsResponse.Unmarshal(m, b)
}
func (m *GetPaymentDetailsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPaymentDetailsResponse.Marshal(b, m, deterministic)
}
func (dst *GetPaymentDetailsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPaymentDetailsResponse.Merge(dst, src)
}
func (m *GetPaymentDetailsResponse) XXX_Size() int {
	return xxx_messageInfo_GetPaymentDetailsResponse.Size(m)
}
func (m *GetPaymentDetailsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPaymentDetailsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPaymentDetailsResponse proto.InternalMessageInfo

func (m *GetPaymentDetailsResponse) GetDetails() []*PaymentDetail {
	if m != nil {
		return m.Details
	}
	return nil
}

type UpdatePaymentDataRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Invoice              string   `protobuf:"bytes,2,opt,name=invoice,proto3" json:"invoice,omitempty"`
	Amount               float32  `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Status               int64    `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	PaidAmount           float32  `protobuf:"fixed32,5,opt,name=paid_amount,json=paidAmount,proto3" json:"paid_amount,omitempty"`
	Method               int64    `protobuf:"varint,6,opt,name=method,proto3" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePaymentDataRequest) Reset()         { *m = UpdatePaymentDataRequest{} }
func (m *UpdatePaymentDataRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePaymentDataRequest) ProtoMessage()    {}
func (*UpdatePaymentDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8400961b101b93a7, []int{12}
}
func (m *UpdatePaymentDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePaymentDataRequest.Unmarshal(m, b)
}
func (m *UpdatePaymentDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePaymentDataRequest.Marshal(b, m, deterministic)
}
func (dst *UpdatePaymentDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePaymentDataRequest.Merge(dst, src)
}
func (m *UpdatePaymentDataRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePaymentDataRequest.Size(m)
}
func (m *UpdatePaymentDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePaymentDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePaymentDataRequest proto.InternalMessageInfo

func (m *UpdatePaymentDataRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdatePaymentDataRequest) GetInvoice() string {
	if m != nil {
		return m.Invoice
	}
	return ""
}

func (m *UpdatePaymentDataRequest) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *UpdatePaymentDataRequest) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UpdatePaymentDataRequest) GetPaidAmount() float32 {
	if m != nil {
		return m.PaidAmount
	}
	return 0
}

func (m *UpdatePaymentDataRequest) GetMethod() int64 {
	if m != nil {
		return m.Method
	}
	return 0
}

type UpdatePaymentDataResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePaymentDataResponse) Reset()         { *m = UpdatePaymentDataResponse{} }
func (m *UpdatePaymentDataResponse) String() string { return proto.CompactTextString(m) }
func (*UpdatePaymentDataResponse) ProtoMessage()    {}
func (*UpdatePaymentDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8400961b101b93a7, []int{13}
}
func (m *UpdatePaymentDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePaymentDataResponse.Unmarshal(m, b)
}
func (m *UpdatePaymentDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePaymentDataResponse.Marshal(b, m, deterministic)
}
func (dst *UpdatePaymentDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePaymentDataResponse.Merge(dst, src)
}
func (m *UpdatePaymentDataResponse) XXX_Size() int {
	return xxx_messageInfo_UpdatePaymentDataResponse.Size(m)
}
func (m *UpdatePaymentDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePaymentDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePaymentDataResponse proto.InternalMessageInfo

type UpdatePaymentDetailRequest struct {
	Paymentdetail        *PaymentDetail `protobuf:"bytes,1,opt,name=paymentdetail,proto3" json:"paymentdetail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdatePaymentDetailRequest) Reset()         { *m = UpdatePaymentDetailRequest{} }
func (m *UpdatePaymentDetailRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePaymentDetailRequest) ProtoMessage()    {}
func (*UpdatePaymentDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8400961b101b93a7, []int{14}
}
func (m *UpdatePaymentDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePaymentDetailRequest.Unmarshal(m, b)
}
func (m *UpdatePaymentDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePaymentDetailRequest.Marshal(b, m, deterministic)
}
func (dst *UpdatePaymentDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePaymentDetailRequest.Merge(dst, src)
}
func (m *UpdatePaymentDetailRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePaymentDetailRequest.Size(m)
}
func (m *UpdatePaymentDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePaymentDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePaymentDetailRequest proto.InternalMessageInfo

func (m *UpdatePaymentDetailRequest) GetPaymentdetail() *PaymentDetail {
	if m != nil {
		return m.Paymentdetail
	}
	return nil
}

type UpdatePaymentDetailResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePaymentDetailResponse) Reset()         { *m = UpdatePaymentDetailResponse{} }
func (m *UpdatePaymentDetailResponse) String() string { return proto.CompactTextString(m) }
func (*UpdatePaymentDetailResponse) ProtoMessage()    {}
func (*UpdatePaymentDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_8400961b101b93a7, []int{15}
}
func (m *UpdatePaymentDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePaymentDetailResponse.Unmarshal(m, b)
}
func (m *UpdatePaymentDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePaymentDetailResponse.Marshal(b, m, deterministic)
}
func (dst *UpdatePaymentDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePaymentDetailResponse.Merge(dst, src)
}
func (m *UpdatePaymentDetailResponse) XXX_Size() int {
	return xxx_messageInfo_UpdatePaymentDetailResponse.Size(m)
}
func (m *UpdatePaymentDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePaymentDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePaymentDetailResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Payment)(nil), "Payment")
	proto.RegisterType((*PaymentDetail)(nil), "PaymentDetail")
	proto.RegisterType((*PaymentDetailAppend)(nil), "PaymentDetailAppend")
	proto.RegisterType((*CreatePaymentRequest)(nil), "CreatePaymentRequest")
	proto.RegisterType((*CreatePaymentResponse)(nil), "CreatePaymentResponse")
	proto.RegisterType((*CreatePaymentDetailRequest)(nil), "CreatePaymentDetailRequest")
	proto.RegisterType((*CreatePaymentDetailResponse)(nil), "CreatePaymentDetailResponse")
	proto.RegisterType((*GetPaymentRequest)(nil), "GetPaymentRequest")
	proto.RegisterType((*GetPaymentResponse)(nil), "GetPaymentResponse")
	proto.RegisterType((*GetPaymentDetailRequest)(nil), "GetPaymentDetailRequest")
	proto.RegisterType((*GetPaymentDetailResponse)(nil), "GetPaymentDetailResponse")
	proto.RegisterType((*GetPaymentDetailsResponse)(nil), "GetPaymentDetailsResponse")
	proto.RegisterType((*UpdatePaymentDataRequest)(nil), "UpdatePaymentDataRequest")
	proto.RegisterType((*UpdatePaymentDataResponse)(nil), "UpdatePaymentDataResponse")
	proto.RegisterType((*UpdatePaymentDetailRequest)(nil), "UpdatePaymentDetailRequest")
	proto.RegisterType((*UpdatePaymentDetailResponse)(nil), "UpdatePaymentDetailResponse")
}

func init() { proto.RegisterFile("payment.proto", fileDescriptor_payment_8400961b101b93a7) }

var fileDescriptor_payment_8400961b101b93a7 = []byte{
	// 590 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xdd, 0x6a, 0xdb, 0x4c,
	0x10, 0x65, 0x25, 0x5b, 0x8a, 0xc7, 0x38, 0xe1, 0x53, 0xf2, 0x35, 0x8a, 0x43, 0x88, 0xd9, 0x42,
	0x51, 0x6f, 0x14, 0x70, 0x4b, 0xe9, 0x6d, 0xdc, 0x96, 0x92, 0x9b, 0x52, 0x44, 0x7a, 0x53, 0x0a,
	0x66, 0x63, 0x6d, 0x53, 0x81, 0xf5, 0x53, 0xef, 0x38, 0xc8, 0xaf, 0x50, 0xe8, 0xa3, 0xf4, 0x35,
	0x0a, 0x7d, 0xab, 0xa2, 0xdd, 0x95, 0x6d, 0x29, 0xb2, 0xda, 0x42, 0x2f, 0x7a, 0xb9, 0x33, 0x73,
	0x66, 0x76, 0xce, 0x39, 0x5a, 0xc1, 0x20, 0x63, 0xab, 0x98, 0x27, 0xe8, 0x67, 0x8b, 0x14, 0xd3,
	0xe1, 0xf9, 0x6d, 0x9a, 0xde, 0xce, 0xf9, 0x85, 0x3c, 0xdd, 0x2c, 0x3f, 0x5e, 0x60, 0x14, 0x73,
	0x81, 0x2c, 0xce, 0x54, 0x01, 0xfd, 0x4e, 0xc0, 0x7e, 0xab, 0x20, 0xce, 0x3e, 0x18, 0x51, 0xe8,
	0x92, 0x11, 0xf1, 0xcc, 0xc0, 0x88, 0x42, 0xc7, 0x05, 0x3b, 0x4a, 0xee, 0xd2, 0x68, 0xc6, 0x5d,
	0x63, 0x44, 0xbc, 0x5e, 0x50, 0x1e, 0x9d, 0x07, 0x60, 0xb1, 0x38, 0x5d, 0x26, 0xe8, 0x9a, 0x23,
	0xe2, 0x19, 0x81, 0x3e, 0x15, 0x71, 0x81, 0x0c, 0x97, 0xc2, 0xed, 0xc8, 0x2e, 0xfa, 0xe4, 0x9c,
	0x43, 0x3f, 0x63, 0x51, 0x38, 0xd5, 0xa0, 0xae, 0x04, 0x41, 0x11, 0xba, 0x5c, 0x03, 0x63, 0x8e,
	0x9f, 0xd2, 0xd0, 0xb5, 0x14, 0x50, 0x9d, 0x1c, 0x0f, 0xec, 0x90, 0x23, 0x8b, 0xe6, 0xc2, 0xb5,
	0x47, 0xa6, 0xd7, 0x1f, 0xef, 0xfb, 0xfa, 0xb6, 0x2f, 0x65, 0x38, 0x28, 0xd3, 0xf4, 0xab, 0x01,
	0x83, 0x4a, 0xea, 0xde, 0x3a, 0x67, 0x00, 0x9a, 0x9c, 0x69, 0x14, 0xca, 0x8d, 0xcc, 0xa0, 0xa7,
	0x23, 0x57, 0xa1, 0x73, 0x08, 0x5d, 0xcc, 0x8b, 0x8c, 0x29, 0x33, 0x1d, 0xcc, 0xaf, 0x42, 0xe7,
	0x18, 0x6c, 0xcc, 0xa7, 0xb8, 0xca, 0x78, 0xb9, 0x11, 0xe6, 0xd7, 0xab, 0x8c, 0xeb, 0x44, 0xc2,
	0x62, 0x2e, 0xb7, 0xe9, 0x15, 0x89, 0x37, 0x2c, 0xe6, 0xce, 0x04, 0x0e, 0x30, 0x9f, 0xce, 0x16,
	0x9c, 0x21, 0x0f, 0xa7, 0x05, 0xdd, 0xae, 0x3d, 0x22, 0x5e, 0x7f, 0x3c, 0xf4, 0x95, 0x16, 0x7e,
	0xa9, 0x85, 0x7f, 0x5d, 0x6a, 0x11, 0x0c, 0x30, 0x7f, 0xa1, 0x10, 0x45, 0x4c, 0xf7, 0x58, 0x66,
	0xe1, 0xa6, 0xc7, 0xde, 0xef, 0xf4, 0x78, 0xa7, 0x10, 0x45, 0x8c, 0x7e, 0x80, 0xc3, 0x0a, 0x1d,
	0x97, 0x59, 0xc6, 0x93, 0xad, 0x2d, 0x49, 0xf3, 0x96, 0xc6, 0xae, 0x2d, 0xcd, 0xed, 0x2d, 0xe9,
	0x0f, 0x02, 0x47, 0xea, 0xc6, 0x7a, 0x48, 0xc0, 0x3f, 0x2f, 0xb9, 0xc0, 0x6d, 0xcf, 0x90, 0x5d,
	0x9e, 0x31, 0x76, 0x78, 0xc6, 0x6c, 0xf3, 0x4c, 0xa7, 0xc5, 0x33, 0xdd, 0x8a, 0x67, 0xfc, 0x8d,
	0x67, 0x2c, 0xe9, 0x99, 0x23, 0xbf, 0x81, 0x89, 0x8d, 0x73, 0x9e, 0xc1, 0xff, 0xb5, 0x55, 0x44,
	0x96, 0x26, 0x82, 0xd7, 0x0c, 0x43, 0x6a, 0x86, 0xa1, 0x5f, 0x0c, 0x18, 0x56, 0x80, 0xda, 0x92,
	0x9a, 0x89, 0x76, 0xf4, 0x46, 0x08, 0xa3, 0x59, 0x08, 0x73, 0x97, 0x10, 0x9d, 0x5f, 0xd9, 0xad,
	0xfb, 0x17, 0xec, 0x66, 0xfd, 0xa9, 0xdd, 0xce, 0xe0, 0xb4, 0x91, 0x0b, 0x45, 0x25, 0x7d, 0x08,
	0xff, 0xbd, 0xe6, 0x58, 0xf3, 0x4a, 0xed, 0x03, 0xa5, 0xcf, 0xc1, 0xd9, 0x2e, 0xd2, 0x2a, 0x50,
	0xb0, 0x35, 0x6b, 0xb2, 0xb4, 0x3f, 0xde, 0x2b, 0xe5, 0x0c, 0xca, 0x04, 0x7d, 0x0c, 0xc7, 0x1b,
	0x64, 0x55, 0x86, 0xfa, 0x90, 0x09, 0xb8, 0xf7, 0x4b, 0xf5, 0xa8, 0x47, 0x60, 0x29, 0x53, 0xe8,
	0x49, 0xf5, 0xc7, 0x46, 0x67, 0xe9, 0x2b, 0x38, 0xa9, 0xf7, 0x10, 0xeb, 0x26, 0x5b, 0x4f, 0x16,
	0x69, 0x7f, 0xb2, 0xbe, 0x11, 0x70, 0x15, 0x87, 0x65, 0x01, 0x43, 0xb6, 0xe3, 0xde, 0xff, 0xc0,
	0x63, 0x4c, 0x4f, 0xe1, 0xa4, 0xe1, 0xba, 0x5a, 0xe1, 0x00, 0x86, 0xd5, 0x64, 0x45, 0x85, 0xa7,
	0xeb, 0x1f, 0x53, 0x2b, 0xc1, 0xd5, 0xa2, 0xc2, 0x54, 0x8d, 0x3d, 0xd5, 0xc8, 0xc9, 0xc1, 0xfb,
	0xc1, 0x3c, 0x9d, 0xb1, 0xb9, 0xe0, 0x8b, 0xbb, 0x68, 0xc6, 0xc5, 0x8d, 0x25, 0x7d, 0xfa, 0xe4,
	0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xf6, 0xdc, 0x99, 0x05, 0x07, 0x00, 0x00,
}
