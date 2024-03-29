// Code generated by protoc-gen-go. DO NOT EDIT.
// source: categories.proto

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

type Category struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ParentId             int64    `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_categories_da622cb92c5b2ead, []int{0}
}
func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (dst *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(dst, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Category) GetParentId() int64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

type CategoryMap struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CategoryId           int64    `protobuf:"varint,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	VendorCategoryId     int64    `protobuf:"varint,3,opt,name=vendor_category_id,json=vendorCategoryId,proto3" json:"vendor_category_id,omitempty"`
	VendorId             int64    `protobuf:"varint,4,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryMap) Reset()         { *m = CategoryMap{} }
func (m *CategoryMap) String() string { return proto.CompactTextString(m) }
func (*CategoryMap) ProtoMessage()    {}
func (*CategoryMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_categories_da622cb92c5b2ead, []int{1}
}
func (m *CategoryMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryMap.Unmarshal(m, b)
}
func (m *CategoryMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryMap.Marshal(b, m, deterministic)
}
func (dst *CategoryMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryMap.Merge(dst, src)
}
func (m *CategoryMap) XXX_Size() int {
	return xxx_messageInfo_CategoryMap.Size(m)
}
func (m *CategoryMap) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryMap.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryMap proto.InternalMessageInfo

func (m *CategoryMap) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CategoryMap) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *CategoryMap) GetVendorCategoryId() int64 {
	if m != nil {
		return m.VendorCategoryId
	}
	return 0
}

func (m *CategoryMap) GetVendorId() int64 {
	if m != nil {
		return m.VendorId
	}
	return 0
}

type GetCategoriesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCategoriesRequest) Reset()         { *m = GetCategoriesRequest{} }
func (m *GetCategoriesRequest) String() string { return proto.CompactTextString(m) }
func (*GetCategoriesRequest) ProtoMessage()    {}
func (*GetCategoriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_categories_da622cb92c5b2ead, []int{2}
}
func (m *GetCategoriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCategoriesRequest.Unmarshal(m, b)
}
func (m *GetCategoriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCategoriesRequest.Marshal(b, m, deterministic)
}
func (dst *GetCategoriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCategoriesRequest.Merge(dst, src)
}
func (m *GetCategoriesRequest) XXX_Size() int {
	return xxx_messageInfo_GetCategoriesRequest.Size(m)
}
func (m *GetCategoriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCategoriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCategoriesRequest proto.InternalMessageInfo

type Categories struct {
	Category             []*Category `protobuf:"bytes,1,rep,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Categories) Reset()         { *m = Categories{} }
func (m *Categories) String() string { return proto.CompactTextString(m) }
func (*Categories) ProtoMessage()    {}
func (*Categories) Descriptor() ([]byte, []int) {
	return fileDescriptor_categories_da622cb92c5b2ead, []int{3}
}
func (m *Categories) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Categories.Unmarshal(m, b)
}
func (m *Categories) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Categories.Marshal(b, m, deterministic)
}
func (dst *Categories) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Categories.Merge(dst, src)
}
func (m *Categories) XXX_Size() int {
	return xxx_messageInfo_Categories.Size(m)
}
func (m *Categories) XXX_DiscardUnknown() {
	xxx_messageInfo_Categories.DiscardUnknown(m)
}

var xxx_messageInfo_Categories proto.InternalMessageInfo

func (m *Categories) GetCategory() []*Category {
	if m != nil {
		return m.Category
	}
	return nil
}

type CreateCategoryMapRequest struct {
	CatMap               *CategoryMap `protobuf:"bytes,1,opt,name=cat_map,json=catMap,proto3" json:"cat_map,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateCategoryMapRequest) Reset()         { *m = CreateCategoryMapRequest{} }
func (m *CreateCategoryMapRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCategoryMapRequest) ProtoMessage()    {}
func (*CreateCategoryMapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_categories_da622cb92c5b2ead, []int{4}
}
func (m *CreateCategoryMapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCategoryMapRequest.Unmarshal(m, b)
}
func (m *CreateCategoryMapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCategoryMapRequest.Marshal(b, m, deterministic)
}
func (dst *CreateCategoryMapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCategoryMapRequest.Merge(dst, src)
}
func (m *CreateCategoryMapRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCategoryMapRequest.Size(m)
}
func (m *CreateCategoryMapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCategoryMapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCategoryMapRequest proto.InternalMessageInfo

func (m *CreateCategoryMapRequest) GetCatMap() *CategoryMap {
	if m != nil {
		return m.CatMap
	}
	return nil
}

type CreateCategoryMapResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCategoryMapResponse) Reset()         { *m = CreateCategoryMapResponse{} }
func (m *CreateCategoryMapResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCategoryMapResponse) ProtoMessage()    {}
func (*CreateCategoryMapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_categories_da622cb92c5b2ead, []int{5}
}
func (m *CreateCategoryMapResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCategoryMapResponse.Unmarshal(m, b)
}
func (m *CreateCategoryMapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCategoryMapResponse.Marshal(b, m, deterministic)
}
func (dst *CreateCategoryMapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCategoryMapResponse.Merge(dst, src)
}
func (m *CreateCategoryMapResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCategoryMapResponse.Size(m)
}
func (m *CreateCategoryMapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCategoryMapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCategoryMapResponse proto.InternalMessageInfo

type GetCategoryMapsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCategoryMapsRequest) Reset()         { *m = GetCategoryMapsRequest{} }
func (m *GetCategoryMapsRequest) String() string { return proto.CompactTextString(m) }
func (*GetCategoryMapsRequest) ProtoMessage()    {}
func (*GetCategoryMapsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_categories_da622cb92c5b2ead, []int{6}
}
func (m *GetCategoryMapsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCategoryMapsRequest.Unmarshal(m, b)
}
func (m *GetCategoryMapsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCategoryMapsRequest.Marshal(b, m, deterministic)
}
func (dst *GetCategoryMapsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCategoryMapsRequest.Merge(dst, src)
}
func (m *GetCategoryMapsRequest) XXX_Size() int {
	return xxx_messageInfo_GetCategoryMapsRequest.Size(m)
}
func (m *GetCategoryMapsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCategoryMapsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCategoryMapsRequest proto.InternalMessageInfo

type GetCategoryMapsResponse struct {
	CatMap               []*CategoryMap `protobuf:"bytes,1,rep,name=cat_map,json=catMap,proto3" json:"cat_map,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetCategoryMapsResponse) Reset()         { *m = GetCategoryMapsResponse{} }
func (m *GetCategoryMapsResponse) String() string { return proto.CompactTextString(m) }
func (*GetCategoryMapsResponse) ProtoMessage()    {}
func (*GetCategoryMapsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_categories_da622cb92c5b2ead, []int{7}
}
func (m *GetCategoryMapsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCategoryMapsResponse.Unmarshal(m, b)
}
func (m *GetCategoryMapsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCategoryMapsResponse.Marshal(b, m, deterministic)
}
func (dst *GetCategoryMapsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCategoryMapsResponse.Merge(dst, src)
}
func (m *GetCategoryMapsResponse) XXX_Size() int {
	return xxx_messageInfo_GetCategoryMapsResponse.Size(m)
}
func (m *GetCategoryMapsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCategoryMapsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCategoryMapsResponse proto.InternalMessageInfo

func (m *GetCategoryMapsResponse) GetCatMap() []*CategoryMap {
	if m != nil {
		return m.CatMap
	}
	return nil
}

type GetCategoryMapByIdRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCategoryMapByIdRequest) Reset()         { *m = GetCategoryMapByIdRequest{} }
func (m *GetCategoryMapByIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetCategoryMapByIdRequest) ProtoMessage()    {}
func (*GetCategoryMapByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_categories_da622cb92c5b2ead, []int{8}
}
func (m *GetCategoryMapByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCategoryMapByIdRequest.Unmarshal(m, b)
}
func (m *GetCategoryMapByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCategoryMapByIdRequest.Marshal(b, m, deterministic)
}
func (dst *GetCategoryMapByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCategoryMapByIdRequest.Merge(dst, src)
}
func (m *GetCategoryMapByIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetCategoryMapByIdRequest.Size(m)
}
func (m *GetCategoryMapByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCategoryMapByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCategoryMapByIdRequest proto.InternalMessageInfo

func (m *GetCategoryMapByIdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetCategoryMapByIdResponse struct {
	CatMap               *CategoryMap `protobuf:"bytes,1,opt,name=cat_map,json=catMap,proto3" json:"cat_map,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetCategoryMapByIdResponse) Reset()         { *m = GetCategoryMapByIdResponse{} }
func (m *GetCategoryMapByIdResponse) String() string { return proto.CompactTextString(m) }
func (*GetCategoryMapByIdResponse) ProtoMessage()    {}
func (*GetCategoryMapByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_categories_da622cb92c5b2ead, []int{9}
}
func (m *GetCategoryMapByIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCategoryMapByIdResponse.Unmarshal(m, b)
}
func (m *GetCategoryMapByIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCategoryMapByIdResponse.Marshal(b, m, deterministic)
}
func (dst *GetCategoryMapByIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCategoryMapByIdResponse.Merge(dst, src)
}
func (m *GetCategoryMapByIdResponse) XXX_Size() int {
	return xxx_messageInfo_GetCategoryMapByIdResponse.Size(m)
}
func (m *GetCategoryMapByIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCategoryMapByIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCategoryMapByIdResponse proto.InternalMessageInfo

func (m *GetCategoryMapByIdResponse) GetCatMap() *CategoryMap {
	if m != nil {
		return m.CatMap
	}
	return nil
}

type UpdateCategoryMapRequest struct {
	CatMap               *CategoryMap `protobuf:"bytes,1,opt,name=cat_map,json=catMap,proto3" json:"cat_map,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateCategoryMapRequest) Reset()         { *m = UpdateCategoryMapRequest{} }
func (m *UpdateCategoryMapRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCategoryMapRequest) ProtoMessage()    {}
func (*UpdateCategoryMapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_categories_da622cb92c5b2ead, []int{10}
}
func (m *UpdateCategoryMapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCategoryMapRequest.Unmarshal(m, b)
}
func (m *UpdateCategoryMapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCategoryMapRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateCategoryMapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCategoryMapRequest.Merge(dst, src)
}
func (m *UpdateCategoryMapRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCategoryMapRequest.Size(m)
}
func (m *UpdateCategoryMapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCategoryMapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCategoryMapRequest proto.InternalMessageInfo

func (m *UpdateCategoryMapRequest) GetCatMap() *CategoryMap {
	if m != nil {
		return m.CatMap
	}
	return nil
}

type UpdateCategoryMapResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCategoryMapResponse) Reset()         { *m = UpdateCategoryMapResponse{} }
func (m *UpdateCategoryMapResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateCategoryMapResponse) ProtoMessage()    {}
func (*UpdateCategoryMapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_categories_da622cb92c5b2ead, []int{11}
}
func (m *UpdateCategoryMapResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCategoryMapResponse.Unmarshal(m, b)
}
func (m *UpdateCategoryMapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCategoryMapResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateCategoryMapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCategoryMapResponse.Merge(dst, src)
}
func (m *UpdateCategoryMapResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateCategoryMapResponse.Size(m)
}
func (m *UpdateCategoryMapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCategoryMapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCategoryMapResponse proto.InternalMessageInfo

type DeleteCategoryMapRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCategoryMapRequest) Reset()         { *m = DeleteCategoryMapRequest{} }
func (m *DeleteCategoryMapRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCategoryMapRequest) ProtoMessage()    {}
func (*DeleteCategoryMapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_categories_da622cb92c5b2ead, []int{12}
}
func (m *DeleteCategoryMapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCategoryMapRequest.Unmarshal(m, b)
}
func (m *DeleteCategoryMapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCategoryMapRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteCategoryMapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCategoryMapRequest.Merge(dst, src)
}
func (m *DeleteCategoryMapRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCategoryMapRequest.Size(m)
}
func (m *DeleteCategoryMapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCategoryMapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCategoryMapRequest proto.InternalMessageInfo

func (m *DeleteCategoryMapRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteCategoryMapResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCategoryMapResponse) Reset()         { *m = DeleteCategoryMapResponse{} }
func (m *DeleteCategoryMapResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteCategoryMapResponse) ProtoMessage()    {}
func (*DeleteCategoryMapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_categories_da622cb92c5b2ead, []int{13}
}
func (m *DeleteCategoryMapResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCategoryMapResponse.Unmarshal(m, b)
}
func (m *DeleteCategoryMapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCategoryMapResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteCategoryMapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCategoryMapResponse.Merge(dst, src)
}
func (m *DeleteCategoryMapResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteCategoryMapResponse.Size(m)
}
func (m *DeleteCategoryMapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCategoryMapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCategoryMapResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Category)(nil), "Category")
	proto.RegisterType((*CategoryMap)(nil), "CategoryMap")
	proto.RegisterType((*GetCategoriesRequest)(nil), "GetCategoriesRequest")
	proto.RegisterType((*Categories)(nil), "Categories")
	proto.RegisterType((*CreateCategoryMapRequest)(nil), "CreateCategoryMapRequest")
	proto.RegisterType((*CreateCategoryMapResponse)(nil), "CreateCategoryMapResponse")
	proto.RegisterType((*GetCategoryMapsRequest)(nil), "GetCategoryMapsRequest")
	proto.RegisterType((*GetCategoryMapsResponse)(nil), "GetCategoryMapsResponse")
	proto.RegisterType((*GetCategoryMapByIdRequest)(nil), "GetCategoryMapByIdRequest")
	proto.RegisterType((*GetCategoryMapByIdResponse)(nil), "GetCategoryMapByIdResponse")
	proto.RegisterType((*UpdateCategoryMapRequest)(nil), "UpdateCategoryMapRequest")
	proto.RegisterType((*UpdateCategoryMapResponse)(nil), "UpdateCategoryMapResponse")
	proto.RegisterType((*DeleteCategoryMapRequest)(nil), "DeleteCategoryMapRequest")
	proto.RegisterType((*DeleteCategoryMapResponse)(nil), "DeleteCategoryMapResponse")
}

func init() { proto.RegisterFile("categories.proto", fileDescriptor_categories_da622cb92c5b2ead) }

var fileDescriptor_categories_da622cb92c5b2ead = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x5d, 0x4b, 0xf3, 0x30,
	0x18, 0xa5, 0xed, 0xd8, 0xbb, 0x3e, 0x7d, 0xd5, 0x11, 0x64, 0x66, 0xf6, 0xc2, 0x11, 0x18, 0x0c,
	0x95, 0x5d, 0x6c, 0x7f, 0xc0, 0xad, 0x82, 0x14, 0xd9, 0x4d, 0xc1, 0x1b, 0x6f, 0x46, 0x6c, 0x1e,
	0xa4, 0xb0, 0xb5, 0xb1, 0x8d, 0x83, 0xfd, 0x04, 0xff, 0xb5, 0xac, 0x69, 0xf7, 0x99, 0x5d, 0x79,
	0x57, 0xce, 0x39, 0x39, 0x1f, 0x29, 0x81, 0x76, 0xcc, 0x15, 0x7e, 0x66, 0x79, 0x82, 0xc5, 0x50,
	0xe6, 0x99, 0xca, 0xd8, 0x2b, 0xb4, 0x02, 0x8d, 0xad, 0xc9, 0x25, 0xd8, 0x89, 0xa0, 0x56, 0xcf,
	0x1a, 0x38, 0x91, 0x9d, 0x08, 0x42, 0xa0, 0x91, 0xf2, 0x25, 0x52, 0xbb, 0x67, 0x0d, 0xdc, 0xa8,
	0xfc, 0x26, 0x3e, 0xb8, 0x92, 0xe7, 0x98, 0xaa, 0x79, 0x22, 0xa8, 0x53, 0x4a, 0x5b, 0x1a, 0x08,
	0x05, 0xfb, 0xb1, 0xc0, 0xab, 0xdd, 0x66, 0x5c, 0x9e, 0x18, 0xde, 0x81, 0x57, 0x15, 0x58, 0x6f,
	0x8e, 0xdb, 0x25, 0x01, 0x35, 0x14, 0x0a, 0xf2, 0x08, 0x64, 0x85, 0xa9, 0xc8, 0xf2, 0xf9, 0xbe,
	0x4e, 0xc7, 0xb4, 0x35, 0x13, 0xec, 0xd4, 0x3e, 0xb8, 0x95, 0x3a, 0x11, 0xb4, 0xa1, 0xbb, 0x68,
	0x20, 0x14, 0xac, 0x03, 0xd7, 0x2f, 0xa8, 0x82, 0xed, 0xde, 0x08, 0xbf, 0xbe, 0xb1, 0x50, 0x6c,
	0x0c, 0xb0, 0x03, 0x49, 0x1f, 0x5a, 0x75, 0x12, 0xb5, 0x7a, 0xce, 0xc0, 0x1b, 0xb9, 0xc3, 0x3a,
	0x21, 0xda, 0x52, 0x6c, 0x02, 0x34, 0xc8, 0x91, 0x2b, 0xdc, 0x5b, 0x57, 0x19, 0x92, 0x3e, 0xfc,
	0x8b, 0xb9, 0x9a, 0x2f, 0xb9, 0x2c, 0x97, 0x7a, 0xa3, 0xff, 0xc3, 0x7d, 0x55, 0x33, 0xe6, 0x6a,
	0xc6, 0x25, 0xf3, 0xa1, 0x6b, 0xb0, 0x28, 0x64, 0x96, 0x16, 0xc8, 0x28, 0x74, 0x76, 0x65, 0x37,
	0xcc, 0xb6, 0xee, 0x13, 0xdc, 0x9c, 0x30, 0xfa, 0xd0, 0x61, 0xb0, 0x73, 0x36, 0xf8, 0x01, 0xba,
	0x87, 0x0e, 0xd3, 0x75, 0x28, 0xea, 0xf2, 0x47, 0x7f, 0x88, 0x05, 0x70, 0x6b, 0x12, 0x9b, 0x12,
	0xcf, 0x4f, 0x9d, 0x00, 0x7d, 0x93, 0xe2, 0xaf, 0xb7, 0x65, 0xb0, 0xa8, 0x6e, 0xeb, 0x1e, 0xe8,
	0x33, 0x2e, 0xd0, 0xe8, 0x7f, 0x3c, 0xc8, 0x87, 0xae, 0x41, 0xab, 0x8d, 0xa6, 0x57, 0xef, 0x17,
	0x8b, 0x2c, 0xe6, 0x8b, 0x02, 0xf3, 0x55, 0x12, 0x63, 0xf1, 0xd1, 0x2c, 0x1f, 0xc5, 0xf8, 0x37,
	0x00, 0x00, 0xff, 0xff, 0x93, 0xef, 0xe8, 0xf4, 0x28, 0x03, 0x00, 0x00,
}
