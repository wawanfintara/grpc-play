// Code generated by protoc-gen-go. DO NOT EDIT.
// source: products.proto

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

type GetProductDetailRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductDetailRequest) Reset()         { *m = GetProductDetailRequest{} }
func (m *GetProductDetailRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductDetailRequest) ProtoMessage()    {}
func (*GetProductDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_products_f02f369244215de5, []int{0}
}
func (m *GetProductDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductDetailRequest.Unmarshal(m, b)
}
func (m *GetProductDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductDetailRequest.Marshal(b, m, deterministic)
}
func (dst *GetProductDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductDetailRequest.Merge(dst, src)
}
func (m *GetProductDetailRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductDetailRequest.Size(m)
}
func (m *GetProductDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductDetailRequest proto.InternalMessageInfo

func (m *GetProductDetailRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetProductDetailResponse struct {
	ProductDetail        *ProductDetail `protobuf:"bytes,1,opt,name=product_detail,json=productDetail,proto3" json:"product_detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetProductDetailResponse) Reset()         { *m = GetProductDetailResponse{} }
func (m *GetProductDetailResponse) String() string { return proto.CompactTextString(m) }
func (*GetProductDetailResponse) ProtoMessage()    {}
func (*GetProductDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_products_f02f369244215de5, []int{1}
}
func (m *GetProductDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductDetailResponse.Unmarshal(m, b)
}
func (m *GetProductDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductDetailResponse.Marshal(b, m, deterministic)
}
func (dst *GetProductDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductDetailResponse.Merge(dst, src)
}
func (m *GetProductDetailResponse) XXX_Size() int {
	return xxx_messageInfo_GetProductDetailResponse.Size(m)
}
func (m *GetProductDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductDetailResponse proto.InternalMessageInfo

func (m *GetProductDetailResponse) GetProductDetail() *ProductDetail {
	if m != nil {
		return m.ProductDetail
	}
	return nil
}

type GetProductsRequest struct {
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductsRequest) Reset()         { *m = GetProductsRequest{} }
func (m *GetProductsRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductsRequest) ProtoMessage()    {}
func (*GetProductsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_products_f02f369244215de5, []int{2}
}
func (m *GetProductsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductsRequest.Unmarshal(m, b)
}
func (m *GetProductsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductsRequest.Marshal(b, m, deterministic)
}
func (dst *GetProductsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductsRequest.Merge(dst, src)
}
func (m *GetProductsRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductsRequest.Size(m)
}
func (m *GetProductsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductsRequest proto.InternalMessageInfo

func (m *GetProductsRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetProductsResponse struct {
	Products             *Products `protobuf:"bytes,1,opt,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetProductsResponse) Reset()         { *m = GetProductsResponse{} }
func (m *GetProductsResponse) String() string { return proto.CompactTextString(m) }
func (*GetProductsResponse) ProtoMessage()    {}
func (*GetProductsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_products_f02f369244215de5, []int{3}
}
func (m *GetProductsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductsResponse.Unmarshal(m, b)
}
func (m *GetProductsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductsResponse.Marshal(b, m, deterministic)
}
func (dst *GetProductsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductsResponse.Merge(dst, src)
}
func (m *GetProductsResponse) XXX_Size() int {
	return xxx_messageInfo_GetProductsResponse.Size(m)
}
func (m *GetProductsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductsResponse proto.InternalMessageInfo

func (m *GetProductsResponse) GetProducts() *Products {
	if m != nil {
		return m.Products
	}
	return nil
}

type Product struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	VendorId             int64                `protobuf:"varint,2,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty"`
	CategoryId           int64                `protobuf:"varint,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	ProductName          string               `protobuf:"bytes,4,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	ProductDescription   string               `protobuf:"bytes,5,opt,name=product_description,json=productDescription,proto3" json:"product_description,omitempty"`
	ProductStatus        int32                `protobuf:"varint,6,opt,name=product_status,json=productStatus,proto3" json:"product_status,omitempty"`
	ProductCreatedTime   *timestamp.Timestamp `protobuf:"bytes,7,opt,name=product_created_time,json=productCreatedTime,proto3" json:"product_created_time,omitempty"`
	ProductUpdatedTime   *timestamp.Timestamp `protobuf:"bytes,8,opt,name=product_updated_time,json=productUpdatedTime,proto3" json:"product_updated_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_products_f02f369244215de5, []int{4}
}
func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (dst *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(dst, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Product) GetVendorId() int64 {
	if m != nil {
		return m.VendorId
	}
	return 0
}

func (m *Product) GetCategoryId() int64 {
	if m != nil {
		return m.CategoryId
	}
	return 0
}

func (m *Product) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *Product) GetProductDescription() string {
	if m != nil {
		return m.ProductDescription
	}
	return ""
}

func (m *Product) GetProductStatus() int32 {
	if m != nil {
		return m.ProductStatus
	}
	return 0
}

func (m *Product) GetProductCreatedTime() *timestamp.Timestamp {
	if m != nil {
		return m.ProductCreatedTime
	}
	return nil
}

func (m *Product) GetProductUpdatedTime() *timestamp.Timestamp {
	if m != nil {
		return m.ProductUpdatedTime
	}
	return nil
}

type Products struct {
	Product              []*Product `protobuf:"bytes,1,rep,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Products) Reset()         { *m = Products{} }
func (m *Products) String() string { return proto.CompactTextString(m) }
func (*Products) ProtoMessage()    {}
func (*Products) Descriptor() ([]byte, []int) {
	return fileDescriptor_products_f02f369244215de5, []int{5}
}
func (m *Products) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Products.Unmarshal(m, b)
}
func (m *Products) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Products.Marshal(b, m, deterministic)
}
func (dst *Products) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Products.Merge(dst, src)
}
func (m *Products) XXX_Size() int {
	return xxx_messageInfo_Products.Size(m)
}
func (m *Products) XXX_DiscardUnknown() {
	xxx_messageInfo_Products.DiscardUnknown(m)
}

var xxx_messageInfo_Products proto.InternalMessageInfo

func (m *Products) GetProduct() []*Product {
	if m != nil {
		return m.Product
	}
	return nil
}

type ProductDetail struct {
	Id                   int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StartPrice           float64        `protobuf:"fixed64,3,opt,name=start_price,json=startPrice,proto3" json:"start_price,omitempty"`
	Images               []string       `protobuf:"bytes,4,rep,name=images,proto3" json:"images,omitempty"`
	Description          string         `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	CategoryName         string         `protobuf:"bytes,6,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`
	ProductSpec          []*ProductSpec `protobuf:"bytes,7,rep,name=product_spec,json=productSpec,proto3" json:"product_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ProductDetail) Reset()         { *m = ProductDetail{} }
func (m *ProductDetail) String() string { return proto.CompactTextString(m) }
func (*ProductDetail) ProtoMessage()    {}
func (*ProductDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_products_f02f369244215de5, []int{6}
}
func (m *ProductDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductDetail.Unmarshal(m, b)
}
func (m *ProductDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductDetail.Marshal(b, m, deterministic)
}
func (dst *ProductDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductDetail.Merge(dst, src)
}
func (m *ProductDetail) XXX_Size() int {
	return xxx_messageInfo_ProductDetail.Size(m)
}
func (m *ProductDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductDetail.DiscardUnknown(m)
}

var xxx_messageInfo_ProductDetail proto.InternalMessageInfo

func (m *ProductDetail) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProductDetail) GetStartPrice() float64 {
	if m != nil {
		return m.StartPrice
	}
	return 0
}

func (m *ProductDetail) GetImages() []string {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *ProductDetail) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProductDetail) GetCategoryName() string {
	if m != nil {
		return m.CategoryName
	}
	return ""
}

func (m *ProductDetail) GetProductSpec() []*ProductSpec {
	if m != nil {
		return m.ProductSpec
	}
	return nil
}

type ProductSpec struct {
	Id                   int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SpecName             string        `protobuf:"bytes,2,opt,name=spec_name,json=specName,proto3" json:"spec_name,omitempty"`
	Option               []*SpecOption `protobuf:"bytes,3,rep,name=option,proto3" json:"option,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ProductSpec) Reset()         { *m = ProductSpec{} }
func (m *ProductSpec) String() string { return proto.CompactTextString(m) }
func (*ProductSpec) ProtoMessage()    {}
func (*ProductSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_products_f02f369244215de5, []int{7}
}
func (m *ProductSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductSpec.Unmarshal(m, b)
}
func (m *ProductSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductSpec.Marshal(b, m, deterministic)
}
func (dst *ProductSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductSpec.Merge(dst, src)
}
func (m *ProductSpec) XXX_Size() int {
	return xxx_messageInfo_ProductSpec.Size(m)
}
func (m *ProductSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ProductSpec proto.InternalMessageInfo

func (m *ProductSpec) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductSpec) GetSpecName() string {
	if m != nil {
		return m.SpecName
	}
	return ""
}

func (m *ProductSpec) GetOption() []*SpecOption {
	if m != nil {
		return m.Option
	}
	return nil
}

type SpecOption struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpecOption) Reset()         { *m = SpecOption{} }
func (m *SpecOption) String() string { return proto.CompactTextString(m) }
func (*SpecOption) ProtoMessage()    {}
func (*SpecOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_products_f02f369244215de5, []int{8}
}
func (m *SpecOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecOption.Unmarshal(m, b)
}
func (m *SpecOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecOption.Marshal(b, m, deterministic)
}
func (dst *SpecOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecOption.Merge(dst, src)
}
func (m *SpecOption) XXX_Size() int {
	return xxx_messageInfo_SpecOption.Size(m)
}
func (m *SpecOption) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecOption.DiscardUnknown(m)
}

var xxx_messageInfo_SpecOption proto.InternalMessageInfo

func (m *SpecOption) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SpecOption) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*GetProductDetailRequest)(nil), "GetProductDetailRequest")
	proto.RegisterType((*GetProductDetailResponse)(nil), "GetProductDetailResponse")
	proto.RegisterType((*GetProductsRequest)(nil), "GetProductsRequest")
	proto.RegisterType((*GetProductsResponse)(nil), "GetProductsResponse")
	proto.RegisterType((*Product)(nil), "Product")
	proto.RegisterType((*Products)(nil), "Products")
	proto.RegisterType((*ProductDetail)(nil), "ProductDetail")
	proto.RegisterType((*ProductSpec)(nil), "ProductSpec")
	proto.RegisterType((*SpecOption)(nil), "SpecOption")
}

func init() { proto.RegisterFile("products.proto", fileDescriptor_products_f02f369244215de5) }

var fileDescriptor_products_f02f369244215de5 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0xfa, 0x99, 0x4e, 0xda, 0x22, 0x79, 0x57, 0x10, 0x2d, 0x87, 0x0d, 0x5e, 0xad, 0x54,
	0x38, 0xa4, 0x68, 0x11, 0x37, 0x4e, 0x80, 0x84, 0x56, 0x42, 0xb0, 0x64, 0xe1, 0xc2, 0x25, 0xf2,
	0xc6, 0xa6, 0xb2, 0xd4, 0xd4, 0x26, 0x76, 0x57, 0xe2, 0xff, 0x72, 0xe6, 0x37, 0x20, 0x4f, 0x62,
	0xb7, 0xa5, 0x1c, 0xf6, 0x96, 0x79, 0xef, 0xcd, 0xbc, 0xcc, 0x3c, 0xc3, 0x5c, 0x37, 0x8a, 0x6f,
	0x2b, 0x6b, 0x72, 0xdd, 0x28, 0xab, 0xce, 0xce, 0x57, 0x4a, 0xad, 0xd6, 0x62, 0x89, 0xd5, 0xdd,
	0xf6, 0xc7, 0xd2, 0xca, 0x5a, 0x18, 0xcb, 0x6a, 0xdd, 0x0a, 0xe8, 0x73, 0x78, 0xf2, 0x41, 0xd8,
	0x9b, 0xb6, 0xeb, 0xbd, 0xb0, 0x4c, 0xae, 0x0b, 0xf1, 0x73, 0x2b, 0x8c, 0x25, 0x73, 0xe8, 0x49,
	0x9e, 0x46, 0x59, 0xb4, 0xe8, 0x17, 0x3d, 0xc9, 0xe9, 0x17, 0x48, 0x8f, 0xa5, 0x46, 0xab, 0x8d,
	0x11, 0xe4, 0x75, 0x70, 0x2e, 0x39, 0x32, 0xd8, 0x97, 0x5c, 0xcd, 0xf3, 0x43, 0xfd, 0x4c, 0xef,
	0x97, 0xf4, 0x05, 0x90, 0xdd, 0x48, 0xe3, 0x8d, 0x4f, 0x61, 0xb8, 0x96, 0xb5, 0xb4, 0x9d, 0x77,
	0x5b, 0xd0, 0x37, 0x70, 0x72, 0xa0, 0xed, 0x9c, 0x2f, 0x21, 0xf6, 0x3b, 0x77, 0x9e, 0x93, 0x3c,
	0x88, 0x02, 0x45, 0xff, 0xf4, 0x60, 0xdc, 0xc1, 0xff, 0x2e, 0x46, 0x9e, 0xc2, 0xe4, 0x5e, 0x6c,
	0xb8, 0x6a, 0x4a, 0xc9, 0xd3, 0x1e, 0xc2, 0x71, 0x0b, 0x5c, 0x73, 0x72, 0x0e, 0x49, 0xc5, 0xac,
	0x58, 0xa9, 0xe6, 0x97, 0xa3, 0xfb, 0x48, 0x83, 0x87, 0xae, 0x39, 0x79, 0x06, 0x53, 0xbf, 0xfa,
	0x86, 0xd5, 0x22, 0x1d, 0x64, 0xd1, 0x62, 0x52, 0x24, 0x1d, 0xf6, 0x89, 0xd5, 0x82, 0x2c, 0xe1,
	0x64, 0x77, 0x1d, 0x53, 0x35, 0x52, 0x5b, 0xa9, 0x36, 0xe9, 0x10, 0x95, 0x24, 0x9c, 0x24, 0x30,
	0xe4, 0x72, 0x77, 0x4e, 0x63, 0x99, 0xdd, 0x9a, 0x74, 0x94, 0x45, 0x8b, 0x61, 0x38, 0xdf, 0x2d,
	0x82, 0xe4, 0x23, 0x9c, 0x7a, 0x59, 0xd5, 0x08, 0x66, 0x05, 0x2f, 0x5d, 0xbe, 0xe9, 0x18, 0xef,
	0x70, 0x96, 0xb7, 0xe1, 0xe7, 0x3e, 0xfc, 0xfc, 0xab, 0x0f, 0x3f, 0x98, 0xbe, 0x6b, 0xdb, 0x1c,
	0xb1, 0x3f, 0x6d, 0xab, 0xf9, 0x6e, 0x5a, 0xfc, 0xe0, 0x69, 0xdf, 0xda, 0x36, 0x47, 0xd0, 0x1c,
	0x62, 0x1f, 0x03, 0xa1, 0x30, 0xee, 0x14, 0x69, 0x94, 0xf5, 0x17, 0xc9, 0x55, 0xec, 0x23, 0x2a,
	0x3c, 0x41, 0x7f, 0x47, 0x30, 0x3b, 0x78, 0x2b, 0x47, 0x31, 0x11, 0x18, 0xe0, 0x81, 0x7b, 0x78,
	0x36, 0xfc, 0x76, 0xe9, 0x18, 0xcb, 0x1a, 0x5b, 0xea, 0x46, 0x56, 0x02, 0xd3, 0x89, 0x0a, 0x40,
	0xe8, 0xc6, 0x21, 0xe4, 0x31, 0x8c, 0x64, 0xcd, 0x56, 0xc2, 0xa4, 0x83, 0xac, 0xbf, 0x98, 0x14,
	0x5d, 0x45, 0x32, 0x48, 0x8e, 0xa3, 0xd8, 0x87, 0xc8, 0x05, 0xcc, 0x42, 0xf0, 0xe8, 0x3b, 0x42,
	0xcd, 0xd4, 0x83, 0x5d, 0xb2, 0x21, 0x7c, 0xa3, 0x45, 0x95, 0x8e, 0x71, 0xbd, 0xa9, 0x5f, 0xef,
	0x56, 0x8b, 0x2a, 0x3c, 0x05, 0x57, 0xd0, 0x12, 0x92, 0x3d, 0xee, 0x7f, 0x4f, 0xd1, 0xcd, 0x29,
	0xf7, 0x16, 0x8d, 0x1d, 0x80, 0x66, 0x17, 0x30, 0x52, 0xed, 0xef, 0xf6, 0xd1, 0x26, 0xc9, 0xdd,
	0x8c, 0xcf, 0x08, 0x15, 0x1d, 0x45, 0x5f, 0x02, 0xec, 0xd0, 0x87, 0xdc, 0xf0, 0xed, 0xa3, 0xef,
	0xb3, 0xb5, 0xaa, 0xd8, 0xda, 0x88, 0xe6, 0x5e, 0x56, 0xc2, 0xdc, 0x8d, 0x30, 0xe2, 0x57, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xca, 0x93, 0x5d, 0x4d, 0x04, 0x00, 0x00,
}
