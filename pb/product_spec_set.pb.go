// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product_spec_set.proto

package localservices

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProductSpecSet struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductSpecSet) Reset()         { *m = ProductSpecSet{} }
func (m *ProductSpecSet) String() string { return proto.CompactTextString(m) }
func (*ProductSpecSet) ProtoMessage()    {}
func (*ProductSpecSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{0}
}
func (m *ProductSpecSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductSpecSet.Unmarshal(m, b)
}
func (m *ProductSpecSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductSpecSet.Marshal(b, m, deterministic)
}
func (dst *ProductSpecSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductSpecSet.Merge(dst, src)
}
func (m *ProductSpecSet) XXX_Size() int {
	return xxx_messageInfo_ProductSpecSet.Size(m)
}
func (m *ProductSpecSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductSpecSet.DiscardUnknown(m)
}

var xxx_messageInfo_ProductSpecSet proto.InternalMessageInfo

func (m *ProductSpecSet) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductSpecSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateProductSpecSetRequest struct {
	SpecSet              *ProductSpecSet `protobuf:"bytes,1,opt,name=spec_set,json=specSet,proto3" json:"spec_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateProductSpecSetRequest) Reset()         { *m = CreateProductSpecSetRequest{} }
func (m *CreateProductSpecSetRequest) String() string { return proto.CompactTextString(m) }
func (*CreateProductSpecSetRequest) ProtoMessage()    {}
func (*CreateProductSpecSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{1}
}
func (m *CreateProductSpecSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProductSpecSetRequest.Unmarshal(m, b)
}
func (m *CreateProductSpecSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProductSpecSetRequest.Marshal(b, m, deterministic)
}
func (dst *CreateProductSpecSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProductSpecSetRequest.Merge(dst, src)
}
func (m *CreateProductSpecSetRequest) XXX_Size() int {
	return xxx_messageInfo_CreateProductSpecSetRequest.Size(m)
}
func (m *CreateProductSpecSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProductSpecSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProductSpecSetRequest proto.InternalMessageInfo

func (m *CreateProductSpecSetRequest) GetSpecSet() *ProductSpecSet {
	if m != nil {
		return m.SpecSet
	}
	return nil
}

type CreateProductSpecSetResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProductSpecSetResponse) Reset()         { *m = CreateProductSpecSetResponse{} }
func (m *CreateProductSpecSetResponse) String() string { return proto.CompactTextString(m) }
func (*CreateProductSpecSetResponse) ProtoMessage()    {}
func (*CreateProductSpecSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{2}
}
func (m *CreateProductSpecSetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProductSpecSetResponse.Unmarshal(m, b)
}
func (m *CreateProductSpecSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProductSpecSetResponse.Marshal(b, m, deterministic)
}
func (dst *CreateProductSpecSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProductSpecSetResponse.Merge(dst, src)
}
func (m *CreateProductSpecSetResponse) XXX_Size() int {
	return xxx_messageInfo_CreateProductSpecSetResponse.Size(m)
}
func (m *CreateProductSpecSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProductSpecSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProductSpecSetResponse proto.InternalMessageInfo

type GetProductSpecSetsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductSpecSetsRequest) Reset()         { *m = GetProductSpecSetsRequest{} }
func (m *GetProductSpecSetsRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductSpecSetsRequest) ProtoMessage()    {}
func (*GetProductSpecSetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{3}
}
func (m *GetProductSpecSetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductSpecSetsRequest.Unmarshal(m, b)
}
func (m *GetProductSpecSetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductSpecSetsRequest.Marshal(b, m, deterministic)
}
func (dst *GetProductSpecSetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductSpecSetsRequest.Merge(dst, src)
}
func (m *GetProductSpecSetsRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductSpecSetsRequest.Size(m)
}
func (m *GetProductSpecSetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductSpecSetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductSpecSetsRequest proto.InternalMessageInfo

type GetProductSpecSetsResponse struct {
	SpecSet              []*ProductSpecSet `protobuf:"bytes,1,rep,name=spec_set,json=specSet,proto3" json:"spec_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetProductSpecSetsResponse) Reset()         { *m = GetProductSpecSetsResponse{} }
func (m *GetProductSpecSetsResponse) String() string { return proto.CompactTextString(m) }
func (*GetProductSpecSetsResponse) ProtoMessage()    {}
func (*GetProductSpecSetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{4}
}
func (m *GetProductSpecSetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductSpecSetsResponse.Unmarshal(m, b)
}
func (m *GetProductSpecSetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductSpecSetsResponse.Marshal(b, m, deterministic)
}
func (dst *GetProductSpecSetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductSpecSetsResponse.Merge(dst, src)
}
func (m *GetProductSpecSetsResponse) XXX_Size() int {
	return xxx_messageInfo_GetProductSpecSetsResponse.Size(m)
}
func (m *GetProductSpecSetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductSpecSetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductSpecSetsResponse proto.InternalMessageInfo

func (m *GetProductSpecSetsResponse) GetSpecSet() []*ProductSpecSet {
	if m != nil {
		return m.SpecSet
	}
	return nil
}

type GetProductSpecSetByIdRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductSpecSetByIdRequest) Reset()         { *m = GetProductSpecSetByIdRequest{} }
func (m *GetProductSpecSetByIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductSpecSetByIdRequest) ProtoMessage()    {}
func (*GetProductSpecSetByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{5}
}
func (m *GetProductSpecSetByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductSpecSetByIdRequest.Unmarshal(m, b)
}
func (m *GetProductSpecSetByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductSpecSetByIdRequest.Marshal(b, m, deterministic)
}
func (dst *GetProductSpecSetByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductSpecSetByIdRequest.Merge(dst, src)
}
func (m *GetProductSpecSetByIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductSpecSetByIdRequest.Size(m)
}
func (m *GetProductSpecSetByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductSpecSetByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductSpecSetByIdRequest proto.InternalMessageInfo

func (m *GetProductSpecSetByIdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetProductSpecSetByIdResponse struct {
	SpecSet              *ProductSpecSet `protobuf:"bytes,1,opt,name=spec_set,json=specSet,proto3" json:"spec_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetProductSpecSetByIdResponse) Reset()         { *m = GetProductSpecSetByIdResponse{} }
func (m *GetProductSpecSetByIdResponse) String() string { return proto.CompactTextString(m) }
func (*GetProductSpecSetByIdResponse) ProtoMessage()    {}
func (*GetProductSpecSetByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{6}
}
func (m *GetProductSpecSetByIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductSpecSetByIdResponse.Unmarshal(m, b)
}
func (m *GetProductSpecSetByIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductSpecSetByIdResponse.Marshal(b, m, deterministic)
}
func (dst *GetProductSpecSetByIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductSpecSetByIdResponse.Merge(dst, src)
}
func (m *GetProductSpecSetByIdResponse) XXX_Size() int {
	return xxx_messageInfo_GetProductSpecSetByIdResponse.Size(m)
}
func (m *GetProductSpecSetByIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductSpecSetByIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductSpecSetByIdResponse proto.InternalMessageInfo

func (m *GetProductSpecSetByIdResponse) GetSpecSet() *ProductSpecSet {
	if m != nil {
		return m.SpecSet
	}
	return nil
}

type UpdateProductSpecSetRequest struct {
	SpecSet              *ProductSpecSet `protobuf:"bytes,1,opt,name=spec_set,json=specSet,proto3" json:"spec_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdateProductSpecSetRequest) Reset()         { *m = UpdateProductSpecSetRequest{} }
func (m *UpdateProductSpecSetRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateProductSpecSetRequest) ProtoMessage()    {}
func (*UpdateProductSpecSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{7}
}
func (m *UpdateProductSpecSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProductSpecSetRequest.Unmarshal(m, b)
}
func (m *UpdateProductSpecSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProductSpecSetRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateProductSpecSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProductSpecSetRequest.Merge(dst, src)
}
func (m *UpdateProductSpecSetRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateProductSpecSetRequest.Size(m)
}
func (m *UpdateProductSpecSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProductSpecSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProductSpecSetRequest proto.InternalMessageInfo

func (m *UpdateProductSpecSetRequest) GetSpecSet() *ProductSpecSet {
	if m != nil {
		return m.SpecSet
	}
	return nil
}

type UpdateProductSpecSetResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateProductSpecSetResponse) Reset()         { *m = UpdateProductSpecSetResponse{} }
func (m *UpdateProductSpecSetResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateProductSpecSetResponse) ProtoMessage()    {}
func (*UpdateProductSpecSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{8}
}
func (m *UpdateProductSpecSetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProductSpecSetResponse.Unmarshal(m, b)
}
func (m *UpdateProductSpecSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProductSpecSetResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateProductSpecSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProductSpecSetResponse.Merge(dst, src)
}
func (m *UpdateProductSpecSetResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateProductSpecSetResponse.Size(m)
}
func (m *UpdateProductSpecSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProductSpecSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProductSpecSetResponse proto.InternalMessageInfo

type DeleteProductSpecSetRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProductSpecSetRequest) Reset()         { *m = DeleteProductSpecSetRequest{} }
func (m *DeleteProductSpecSetRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteProductSpecSetRequest) ProtoMessage()    {}
func (*DeleteProductSpecSetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{9}
}
func (m *DeleteProductSpecSetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProductSpecSetRequest.Unmarshal(m, b)
}
func (m *DeleteProductSpecSetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProductSpecSetRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteProductSpecSetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProductSpecSetRequest.Merge(dst, src)
}
func (m *DeleteProductSpecSetRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteProductSpecSetRequest.Size(m)
}
func (m *DeleteProductSpecSetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProductSpecSetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProductSpecSetRequest proto.InternalMessageInfo

func (m *DeleteProductSpecSetRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteProductSpecSetResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProductSpecSetResponse) Reset()         { *m = DeleteProductSpecSetResponse{} }
func (m *DeleteProductSpecSetResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteProductSpecSetResponse) ProtoMessage()    {}
func (*DeleteProductSpecSetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{10}
}
func (m *DeleteProductSpecSetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProductSpecSetResponse.Unmarshal(m, b)
}
func (m *DeleteProductSpecSetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProductSpecSetResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteProductSpecSetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProductSpecSetResponse.Merge(dst, src)
}
func (m *DeleteProductSpecSetResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteProductSpecSetResponse.Size(m)
}
func (m *DeleteProductSpecSetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProductSpecSetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProductSpecSetResponse proto.InternalMessageInfo

type ProductSpecSetMap struct {
	Id                     int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	VendorId               int64    `protobuf:"varint,2,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty"`
	ProductSpecSetId       int64    `protobuf:"varint,3,opt,name=product_spec_set_id,json=productSpecSetId,proto3" json:"product_spec_set_id,omitempty"`
	ProductVendorSpecSetId int64    `protobuf:"varint,4,opt,name=product_vendor_spec_set_id,json=productVendorSpecSetId,proto3" json:"product_vendor_spec_set_id,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *ProductSpecSetMap) Reset()         { *m = ProductSpecSetMap{} }
func (m *ProductSpecSetMap) String() string { return proto.CompactTextString(m) }
func (*ProductSpecSetMap) ProtoMessage()    {}
func (*ProductSpecSetMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{11}
}
func (m *ProductSpecSetMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductSpecSetMap.Unmarshal(m, b)
}
func (m *ProductSpecSetMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductSpecSetMap.Marshal(b, m, deterministic)
}
func (dst *ProductSpecSetMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductSpecSetMap.Merge(dst, src)
}
func (m *ProductSpecSetMap) XXX_Size() int {
	return xxx_messageInfo_ProductSpecSetMap.Size(m)
}
func (m *ProductSpecSetMap) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductSpecSetMap.DiscardUnknown(m)
}

var xxx_messageInfo_ProductSpecSetMap proto.InternalMessageInfo

func (m *ProductSpecSetMap) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductSpecSetMap) GetVendorId() int64 {
	if m != nil {
		return m.VendorId
	}
	return 0
}

func (m *ProductSpecSetMap) GetProductSpecSetId() int64 {
	if m != nil {
		return m.ProductSpecSetId
	}
	return 0
}

func (m *ProductSpecSetMap) GetProductVendorSpecSetId() int64 {
	if m != nil {
		return m.ProductVendorSpecSetId
	}
	return 0
}

type CreateProductSpecSetMapRequest struct {
	SpecSetMap           *ProductSpecSetMap `protobuf:"bytes,1,opt,name=spec_set_map,json=specSetMap,proto3" json:"spec_set_map,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateProductSpecSetMapRequest) Reset()         { *m = CreateProductSpecSetMapRequest{} }
func (m *CreateProductSpecSetMapRequest) String() string { return proto.CompactTextString(m) }
func (*CreateProductSpecSetMapRequest) ProtoMessage()    {}
func (*CreateProductSpecSetMapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{12}
}
func (m *CreateProductSpecSetMapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProductSpecSetMapRequest.Unmarshal(m, b)
}
func (m *CreateProductSpecSetMapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProductSpecSetMapRequest.Marshal(b, m, deterministic)
}
func (dst *CreateProductSpecSetMapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProductSpecSetMapRequest.Merge(dst, src)
}
func (m *CreateProductSpecSetMapRequest) XXX_Size() int {
	return xxx_messageInfo_CreateProductSpecSetMapRequest.Size(m)
}
func (m *CreateProductSpecSetMapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProductSpecSetMapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProductSpecSetMapRequest proto.InternalMessageInfo

func (m *CreateProductSpecSetMapRequest) GetSpecSetMap() *ProductSpecSetMap {
	if m != nil {
		return m.SpecSetMap
	}
	return nil
}

type CreateProductSpecSetMapResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProductSpecSetMapResponse) Reset()         { *m = CreateProductSpecSetMapResponse{} }
func (m *CreateProductSpecSetMapResponse) String() string { return proto.CompactTextString(m) }
func (*CreateProductSpecSetMapResponse) ProtoMessage()    {}
func (*CreateProductSpecSetMapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{13}
}
func (m *CreateProductSpecSetMapResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProductSpecSetMapResponse.Unmarshal(m, b)
}
func (m *CreateProductSpecSetMapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProductSpecSetMapResponse.Marshal(b, m, deterministic)
}
func (dst *CreateProductSpecSetMapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProductSpecSetMapResponse.Merge(dst, src)
}
func (m *CreateProductSpecSetMapResponse) XXX_Size() int {
	return xxx_messageInfo_CreateProductSpecSetMapResponse.Size(m)
}
func (m *CreateProductSpecSetMapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProductSpecSetMapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProductSpecSetMapResponse proto.InternalMessageInfo

type GetProductSpecSetMapsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductSpecSetMapsRequest) Reset()         { *m = GetProductSpecSetMapsRequest{} }
func (m *GetProductSpecSetMapsRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductSpecSetMapsRequest) ProtoMessage()    {}
func (*GetProductSpecSetMapsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{14}
}
func (m *GetProductSpecSetMapsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductSpecSetMapsRequest.Unmarshal(m, b)
}
func (m *GetProductSpecSetMapsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductSpecSetMapsRequest.Marshal(b, m, deterministic)
}
func (dst *GetProductSpecSetMapsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductSpecSetMapsRequest.Merge(dst, src)
}
func (m *GetProductSpecSetMapsRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductSpecSetMapsRequest.Size(m)
}
func (m *GetProductSpecSetMapsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductSpecSetMapsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductSpecSetMapsRequest proto.InternalMessageInfo

type GetProductSpecSetMapsResponse struct {
	SpecSetMap           []*ProductSpecSetMap `protobuf:"bytes,1,rep,name=spec_set_map,json=specSetMap,proto3" json:"spec_set_map,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetProductSpecSetMapsResponse) Reset()         { *m = GetProductSpecSetMapsResponse{} }
func (m *GetProductSpecSetMapsResponse) String() string { return proto.CompactTextString(m) }
func (*GetProductSpecSetMapsResponse) ProtoMessage()    {}
func (*GetProductSpecSetMapsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{15}
}
func (m *GetProductSpecSetMapsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductSpecSetMapsResponse.Unmarshal(m, b)
}
func (m *GetProductSpecSetMapsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductSpecSetMapsResponse.Marshal(b, m, deterministic)
}
func (dst *GetProductSpecSetMapsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductSpecSetMapsResponse.Merge(dst, src)
}
func (m *GetProductSpecSetMapsResponse) XXX_Size() int {
	return xxx_messageInfo_GetProductSpecSetMapsResponse.Size(m)
}
func (m *GetProductSpecSetMapsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductSpecSetMapsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductSpecSetMapsResponse proto.InternalMessageInfo

func (m *GetProductSpecSetMapsResponse) GetSpecSetMap() []*ProductSpecSetMap {
	if m != nil {
		return m.SpecSetMap
	}
	return nil
}

type GetProductSpecSetMapByIdRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductSpecSetMapByIdRequest) Reset()         { *m = GetProductSpecSetMapByIdRequest{} }
func (m *GetProductSpecSetMapByIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductSpecSetMapByIdRequest) ProtoMessage()    {}
func (*GetProductSpecSetMapByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{16}
}
func (m *GetProductSpecSetMapByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductSpecSetMapByIdRequest.Unmarshal(m, b)
}
func (m *GetProductSpecSetMapByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductSpecSetMapByIdRequest.Marshal(b, m, deterministic)
}
func (dst *GetProductSpecSetMapByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductSpecSetMapByIdRequest.Merge(dst, src)
}
func (m *GetProductSpecSetMapByIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductSpecSetMapByIdRequest.Size(m)
}
func (m *GetProductSpecSetMapByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductSpecSetMapByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductSpecSetMapByIdRequest proto.InternalMessageInfo

func (m *GetProductSpecSetMapByIdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetProductSpecSetMapByIdResponse struct {
	SpecSetMap           *ProductSpecSetMap `protobuf:"bytes,1,opt,name=spec_set_map,json=specSetMap,proto3" json:"spec_set_map,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetProductSpecSetMapByIdResponse) Reset()         { *m = GetProductSpecSetMapByIdResponse{} }
func (m *GetProductSpecSetMapByIdResponse) String() string { return proto.CompactTextString(m) }
func (*GetProductSpecSetMapByIdResponse) ProtoMessage()    {}
func (*GetProductSpecSetMapByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{17}
}
func (m *GetProductSpecSetMapByIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductSpecSetMapByIdResponse.Unmarshal(m, b)
}
func (m *GetProductSpecSetMapByIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductSpecSetMapByIdResponse.Marshal(b, m, deterministic)
}
func (dst *GetProductSpecSetMapByIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductSpecSetMapByIdResponse.Merge(dst, src)
}
func (m *GetProductSpecSetMapByIdResponse) XXX_Size() int {
	return xxx_messageInfo_GetProductSpecSetMapByIdResponse.Size(m)
}
func (m *GetProductSpecSetMapByIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductSpecSetMapByIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductSpecSetMapByIdResponse proto.InternalMessageInfo

func (m *GetProductSpecSetMapByIdResponse) GetSpecSetMap() *ProductSpecSetMap {
	if m != nil {
		return m.SpecSetMap
	}
	return nil
}

type UpdateProductSpecSetMapRequest struct {
	SpecSetMap           *ProductSpecSetMap `protobuf:"bytes,1,opt,name=spec_set_map,json=specSetMap,proto3" json:"spec_set_map,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UpdateProductSpecSetMapRequest) Reset()         { *m = UpdateProductSpecSetMapRequest{} }
func (m *UpdateProductSpecSetMapRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateProductSpecSetMapRequest) ProtoMessage()    {}
func (*UpdateProductSpecSetMapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{18}
}
func (m *UpdateProductSpecSetMapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProductSpecSetMapRequest.Unmarshal(m, b)
}
func (m *UpdateProductSpecSetMapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProductSpecSetMapRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateProductSpecSetMapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProductSpecSetMapRequest.Merge(dst, src)
}
func (m *UpdateProductSpecSetMapRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateProductSpecSetMapRequest.Size(m)
}
func (m *UpdateProductSpecSetMapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProductSpecSetMapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProductSpecSetMapRequest proto.InternalMessageInfo

func (m *UpdateProductSpecSetMapRequest) GetSpecSetMap() *ProductSpecSetMap {
	if m != nil {
		return m.SpecSetMap
	}
	return nil
}

type UpdateProductSpecSetMapResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateProductSpecSetMapResponse) Reset()         { *m = UpdateProductSpecSetMapResponse{} }
func (m *UpdateProductSpecSetMapResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateProductSpecSetMapResponse) ProtoMessage()    {}
func (*UpdateProductSpecSetMapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{19}
}
func (m *UpdateProductSpecSetMapResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProductSpecSetMapResponse.Unmarshal(m, b)
}
func (m *UpdateProductSpecSetMapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProductSpecSetMapResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateProductSpecSetMapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProductSpecSetMapResponse.Merge(dst, src)
}
func (m *UpdateProductSpecSetMapResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateProductSpecSetMapResponse.Size(m)
}
func (m *UpdateProductSpecSetMapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProductSpecSetMapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProductSpecSetMapResponse proto.InternalMessageInfo

type DeleteProductSpecSetMapRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProductSpecSetMapRequest) Reset()         { *m = DeleteProductSpecSetMapRequest{} }
func (m *DeleteProductSpecSetMapRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteProductSpecSetMapRequest) ProtoMessage()    {}
func (*DeleteProductSpecSetMapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{20}
}
func (m *DeleteProductSpecSetMapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProductSpecSetMapRequest.Unmarshal(m, b)
}
func (m *DeleteProductSpecSetMapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProductSpecSetMapRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteProductSpecSetMapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProductSpecSetMapRequest.Merge(dst, src)
}
func (m *DeleteProductSpecSetMapRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteProductSpecSetMapRequest.Size(m)
}
func (m *DeleteProductSpecSetMapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProductSpecSetMapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProductSpecSetMapRequest proto.InternalMessageInfo

func (m *DeleteProductSpecSetMapRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteProductSpecSetMapResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProductSpecSetMapResponse) Reset()         { *m = DeleteProductSpecSetMapResponse{} }
func (m *DeleteProductSpecSetMapResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteProductSpecSetMapResponse) ProtoMessage()    {}
func (*DeleteProductSpecSetMapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_spec_set_0574eca9f2d9dd4e, []int{21}
}
func (m *DeleteProductSpecSetMapResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProductSpecSetMapResponse.Unmarshal(m, b)
}
func (m *DeleteProductSpecSetMapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProductSpecSetMapResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteProductSpecSetMapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProductSpecSetMapResponse.Merge(dst, src)
}
func (m *DeleteProductSpecSetMapResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteProductSpecSetMapResponse.Size(m)
}
func (m *DeleteProductSpecSetMapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProductSpecSetMapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProductSpecSetMapResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProductSpecSet)(nil), "ProductSpecSet")
	proto.RegisterType((*CreateProductSpecSetRequest)(nil), "CreateProductSpecSetRequest")
	proto.RegisterType((*CreateProductSpecSetResponse)(nil), "CreateProductSpecSetResponse")
	proto.RegisterType((*GetProductSpecSetsRequest)(nil), "GetProductSpecSetsRequest")
	proto.RegisterType((*GetProductSpecSetsResponse)(nil), "GetProductSpecSetsResponse")
	proto.RegisterType((*GetProductSpecSetByIdRequest)(nil), "GetProductSpecSetByIdRequest")
	proto.RegisterType((*GetProductSpecSetByIdResponse)(nil), "GetProductSpecSetByIdResponse")
	proto.RegisterType((*UpdateProductSpecSetRequest)(nil), "UpdateProductSpecSetRequest")
	proto.RegisterType((*UpdateProductSpecSetResponse)(nil), "UpdateProductSpecSetResponse")
	proto.RegisterType((*DeleteProductSpecSetRequest)(nil), "DeleteProductSpecSetRequest")
	proto.RegisterType((*DeleteProductSpecSetResponse)(nil), "DeleteProductSpecSetResponse")
	proto.RegisterType((*ProductSpecSetMap)(nil), "ProductSpecSetMap")
	proto.RegisterType((*CreateProductSpecSetMapRequest)(nil), "CreateProductSpecSetMapRequest")
	proto.RegisterType((*CreateProductSpecSetMapResponse)(nil), "CreateProductSpecSetMapResponse")
	proto.RegisterType((*GetProductSpecSetMapsRequest)(nil), "GetProductSpecSetMapsRequest")
	proto.RegisterType((*GetProductSpecSetMapsResponse)(nil), "GetProductSpecSetMapsResponse")
	proto.RegisterType((*GetProductSpecSetMapByIdRequest)(nil), "GetProductSpecSetMapByIdRequest")
	proto.RegisterType((*GetProductSpecSetMapByIdResponse)(nil), "GetProductSpecSetMapByIdResponse")
	proto.RegisterType((*UpdateProductSpecSetMapRequest)(nil), "UpdateProductSpecSetMapRequest")
	proto.RegisterType((*UpdateProductSpecSetMapResponse)(nil), "UpdateProductSpecSetMapResponse")
	proto.RegisterType((*DeleteProductSpecSetMapRequest)(nil), "DeleteProductSpecSetMapRequest")
	proto.RegisterType((*DeleteProductSpecSetMapResponse)(nil), "DeleteProductSpecSetMapResponse")
}

func init() {
	proto.RegisterFile("product_spec_set.proto", fileDescriptor_product_spec_set_0574eca9f2d9dd4e)
}

var fileDescriptor_product_spec_set_0574eca9f2d9dd4e = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x4f, 0x6b, 0xea, 0x40,
	0x14, 0xc5, 0x89, 0x91, 0xf7, 0xf4, 0xbe, 0xf7, 0xf4, 0x35, 0x85, 0x62, 0x8d, 0x8d, 0x3a, 0x2b,
	0x29, 0x18, 0xfa, 0xc7, 0x55, 0x97, 0xb6, 0xd0, 0x86, 0x22, 0x94, 0x88, 0x52, 0xba, 0x91, 0x34,
	0x73, 0x17, 0x01, 0x35, 0xd3, 0x4c, 0x14, 0xfa, 0x99, 0xfa, 0x25, 0x4b, 0x93, 0x49, 0x30, 0xc9,
	0x44, 0x5c, 0xb8, 0x1b, 0x72, 0xef, 0x39, 0xf7, 0xde, 0xc3, 0x8f, 0xc0, 0x19, 0x0b, 0x7c, 0xba,
	0x71, 0xc3, 0x05, 0x67, 0xe8, 0x2e, 0x38, 0x86, 0x26, 0x0b, 0xfc, 0xd0, 0x27, 0x23, 0x68, 0xbc,
	0xc4, 0x95, 0x29, 0x43, 0x77, 0x8a, 0xa1, 0xd6, 0x80, 0x8a, 0x47, 0x5b, 0x4a, 0x4f, 0x19, 0xa8,
	0x76, 0xc5, 0xa3, 0x9a, 0x06, 0xd5, 0xb5, 0xb3, 0xc2, 0x56, 0xa5, 0xa7, 0x0c, 0xea, 0x76, 0xf4,
	0x26, 0x16, 0xe8, 0xf7, 0x01, 0x3a, 0x21, 0x66, 0xb5, 0x36, 0x7e, 0x6c, 0x90, 0x87, 0xda, 0x25,
	0xd4, 0x92, 0x31, 0x91, 0xd1, 0x9f, 0x9b, 0xa6, 0x99, 0xeb, 0xfc, 0xcd, 0xe3, 0x07, 0x31, 0xa0,
	0x23, 0xb7, 0xe2, 0xcc, 0x5f, 0x73, 0x24, 0x3a, 0x9c, 0x3f, 0x62, 0x98, 0x2d, 0x72, 0x31, 0x88,
	0x3c, 0x41, 0x5b, 0x56, 0x8c, 0xa5, 0xb9, 0x35, 0xd4, 0xbd, 0x6b, 0x98, 0xd0, 0x29, 0x38, 0x8d,
	0x3f, 0x2d, 0x9a, 0x9c, 0x94, 0x4b, 0x85, 0x3c, 0xc3, 0x45, 0x49, 0xbf, 0x74, 0xf8, 0xfe, 0x0c,
	0x2c, 0xd0, 0x67, 0x8c, 0x1e, 0x2b, 0x4e, 0xb9, 0x95, 0x88, 0x73, 0x08, 0xfa, 0x03, 0x2e, 0xb1,
	0x6c, 0x54, 0xfe, 0x4c, 0x03, 0x3a, 0xf2, 0x76, 0x61, 0xf7, 0xa5, 0xc0, 0x49, 0xb6, 0x34, 0x71,
	0x58, 0x01, 0x21, 0x1d, 0xea, 0x5b, 0x5c, 0x53, 0x3f, 0x58, 0x78, 0x34, 0xe2, 0x48, 0xb5, 0x6b,
	0xf1, 0x07, 0x8b, 0x6a, 0x43, 0x38, 0xcd, 0xb3, 0xf9, 0xd3, 0xa6, 0x46, 0x6d, 0xff, 0x59, 0xc6,
	0xdc, 0xa2, 0xda, 0x1d, 0xb4, 0x93, 0x76, 0xe1, 0xb9, 0xab, 0xaa, 0x46, 0xaa, 0x04, 0xf6, 0x79,
	0xd4, 0x90, 0x6a, 0xc9, 0x1c, 0x0c, 0x19, 0x6b, 0x13, 0x87, 0x25, 0xf7, 0x8f, 0xe0, 0x6f, 0x6a,
	0xb7, 0x72, 0x98, 0x88, 0x5b, 0x33, 0x8b, 0x02, 0xe0, 0xe9, 0x9b, 0xf4, 0xa1, 0x5b, 0xea, 0x2b,
	0x82, 0x32, 0x24, 0x7c, 0x4d, 0x1c, 0x96, 0x92, 0x3c, 0x93, 0xf0, 0x14, 0xd7, 0x05, 0x4f, 0xc5,
	0xcd, 0xd4, 0x03, 0x36, 0xbb, 0x86, 0xae, 0xcc, 0x76, 0x1f, 0xd9, 0xaf, 0xd0, 0x2b, 0x97, 0x94,
	0x2e, 0x73, 0x48, 0x4c, 0x73, 0x30, 0x64, 0x6c, 0x1e, 0x23, 0xfe, 0x52, 0x5f, 0x11, 0xff, 0x15,
	0x18, 0x32, 0x8e, 0x77, 0x46, 0xe7, 0x63, 0xe8, 0x43, 0xb7, 0x54, 0x11, 0x9b, 0x8e, 0x9b, 0x6f,
	0xff, 0x96, 0xbe, 0xeb, 0x2c, 0x39, 0x06, 0x5b, 0xcf, 0x45, 0xfe, 0xfe, 0x2b, 0xfa, 0xa7, 0xde,
	0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x21, 0x63, 0xe9, 0xae, 0x6d, 0x05, 0x00, 0x00,
}
