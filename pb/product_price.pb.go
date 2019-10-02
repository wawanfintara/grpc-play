// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product_price.proto

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

type Area struct {
	DistrictId           int64    `protobuf:"varint,1,opt,name=district_id,json=districtId,proto3" json:"district_id,omitempty"`
	CityId               int64    `protobuf:"varint,2,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
	ProvinceId           int64    `protobuf:"varint,3,opt,name=province_id,json=provinceId,proto3" json:"province_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Area) Reset()         { *m = Area{} }
func (m *Area) String() string { return proto.CompactTextString(m) }
func (*Area) ProtoMessage()    {}
func (*Area) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_price_0e26de6b21ab960d, []int{0}
}
func (m *Area) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Area.Unmarshal(m, b)
}
func (m *Area) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Area.Marshal(b, m, deterministic)
}
func (dst *Area) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Area.Merge(dst, src)
}
func (m *Area) XXX_Size() int {
	return xxx_messageInfo_Area.Size(m)
}
func (m *Area) XXX_DiscardUnknown() {
	xxx_messageInfo_Area.DiscardUnknown(m)
}

var xxx_messageInfo_Area proto.InternalMessageInfo

func (m *Area) GetDistrictId() int64 {
	if m != nil {
		return m.DistrictId
	}
	return 0
}

func (m *Area) GetCityId() int64 {
	if m != nil {
		return m.CityId
	}
	return 0
}

func (m *Area) GetProvinceId() int64 {
	if m != nil {
		return m.ProvinceId
	}
	return 0
}

type GetProductPriceRequest struct {
	ProductId            int64    `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	SpecSet              string   `protobuf:"bytes,2,opt,name=spec_set,json=specSet,proto3" json:"spec_set,omitempty"`
	Area                 *Area    `protobuf:"bytes,3,opt,name=area,proto3" json:"area,omitempty"`
	Quantity             int64    `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductPriceRequest) Reset()         { *m = GetProductPriceRequest{} }
func (m *GetProductPriceRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductPriceRequest) ProtoMessage()    {}
func (*GetProductPriceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_price_0e26de6b21ab960d, []int{1}
}
func (m *GetProductPriceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductPriceRequest.Unmarshal(m, b)
}
func (m *GetProductPriceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductPriceRequest.Marshal(b, m, deterministic)
}
func (dst *GetProductPriceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductPriceRequest.Merge(dst, src)
}
func (m *GetProductPriceRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductPriceRequest.Size(m)
}
func (m *GetProductPriceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductPriceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductPriceRequest proto.InternalMessageInfo

func (m *GetProductPriceRequest) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *GetProductPriceRequest) GetSpecSet() string {
	if m != nil {
		return m.SpecSet
	}
	return ""
}

func (m *GetProductPriceRequest) GetArea() *Area {
	if m != nil {
		return m.Area
	}
	return nil
}

func (m *GetProductPriceRequest) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

type GetProductPriceResponse struct {
	Prices               []*Price `protobuf:"bytes,1,rep,name=prices,proto3" json:"prices,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductPriceResponse) Reset()         { *m = GetProductPriceResponse{} }
func (m *GetProductPriceResponse) String() string { return proto.CompactTextString(m) }
func (*GetProductPriceResponse) ProtoMessage()    {}
func (*GetProductPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_price_0e26de6b21ab960d, []int{2}
}
func (m *GetProductPriceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductPriceResponse.Unmarshal(m, b)
}
func (m *GetProductPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductPriceResponse.Marshal(b, m, deterministic)
}
func (dst *GetProductPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductPriceResponse.Merge(dst, src)
}
func (m *GetProductPriceResponse) XXX_Size() int {
	return xxx_messageInfo_GetProductPriceResponse.Size(m)
}
func (m *GetProductPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductPriceResponse proto.InternalMessageInfo

func (m *GetProductPriceResponse) GetPrices() []*Price {
	if m != nil {
		return m.Prices
	}
	return nil
}

type Price struct {
	Qty                  int64            `protobuf:"varint,1,opt,name=qty,proto3" json:"qty,omitempty"`
	Unit                 string           `protobuf:"bytes,2,opt,name=unit,proto3" json:"unit,omitempty"`
	Partner              *PartnerLocation `protobuf:"bytes,3,opt,name=partner,proto3" json:"partner,omitempty"`
	Option               []*PriceOption   `protobuf:"bytes,4,rep,name=option,proto3" json:"option,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Price) Reset()         { *m = Price{} }
func (m *Price) String() string { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()    {}
func (*Price) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_price_0e26de6b21ab960d, []int{3}
}
func (m *Price) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Price.Unmarshal(m, b)
}
func (m *Price) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Price.Marshal(b, m, deterministic)
}
func (dst *Price) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Price.Merge(dst, src)
}
func (m *Price) XXX_Size() int {
	return xxx_messageInfo_Price.Size(m)
}
func (m *Price) XXX_DiscardUnknown() {
	xxx_messageInfo_Price.DiscardUnknown(m)
}

var xxx_messageInfo_Price proto.InternalMessageInfo

func (m *Price) GetQty() int64 {
	if m != nil {
		return m.Qty
	}
	return 0
}

func (m *Price) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *Price) GetPartner() *PartnerLocation {
	if m != nil {
		return m.Partner
	}
	return nil
}

func (m *Price) GetOption() []*PriceOption {
	if m != nil {
		return m.Option
	}
	return nil
}

type PartnerLocation struct {
	Lat                  float64  `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Long                 float64  `protobuf:"fixed64,2,opt,name=long,proto3" json:"long,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartnerLocation) Reset()         { *m = PartnerLocation{} }
func (m *PartnerLocation) String() string { return proto.CompactTextString(m) }
func (*PartnerLocation) ProtoMessage()    {}
func (*PartnerLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_price_0e26de6b21ab960d, []int{4}
}
func (m *PartnerLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartnerLocation.Unmarshal(m, b)
}
func (m *PartnerLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartnerLocation.Marshal(b, m, deterministic)
}
func (dst *PartnerLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartnerLocation.Merge(dst, src)
}
func (m *PartnerLocation) XXX_Size() int {
	return xxx_messageInfo_PartnerLocation.Size(m)
}
func (m *PartnerLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_PartnerLocation.DiscardUnknown(m)
}

var xxx_messageInfo_PartnerLocation proto.InternalMessageInfo

func (m *PartnerLocation) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *PartnerLocation) GetLong() float64 {
	if m != nil {
		return m.Long
	}
	return 0
}

type PriceOption struct {
	ProductionFinishDate string   `protobuf:"bytes,1,opt,name=production_finish_date,json=productionFinishDate,proto3" json:"production_finish_date,omitempty"`
	Duration             int64    `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Weight               int64    `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	UnitPrice            int64    `protobuf:"varint,4,opt,name=unit_price,json=unitPrice,proto3" json:"unit_price,omitempty"`
	Price                int64    `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PriceOption) Reset()         { *m = PriceOption{} }
func (m *PriceOption) String() string { return proto.CompactTextString(m) }
func (*PriceOption) ProtoMessage()    {}
func (*PriceOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_price_0e26de6b21ab960d, []int{5}
}
func (m *PriceOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PriceOption.Unmarshal(m, b)
}
func (m *PriceOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PriceOption.Marshal(b, m, deterministic)
}
func (dst *PriceOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceOption.Merge(dst, src)
}
func (m *PriceOption) XXX_Size() int {
	return xxx_messageInfo_PriceOption.Size(m)
}
func (m *PriceOption) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceOption.DiscardUnknown(m)
}

var xxx_messageInfo_PriceOption proto.InternalMessageInfo

func (m *PriceOption) GetProductionFinishDate() string {
	if m != nil {
		return m.ProductionFinishDate
	}
	return ""
}

func (m *PriceOption) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *PriceOption) GetWeight() int64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *PriceOption) GetUnitPrice() int64 {
	if m != nil {
		return m.UnitPrice
	}
	return 0
}

func (m *PriceOption) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto.RegisterType((*Area)(nil), "Area")
	proto.RegisterType((*GetProductPriceRequest)(nil), "GetProductPriceRequest")
	proto.RegisterType((*GetProductPriceResponse)(nil), "GetProductPriceResponse")
	proto.RegisterType((*Price)(nil), "Price")
	proto.RegisterType((*PartnerLocation)(nil), "PartnerLocation")
	proto.RegisterType((*PriceOption)(nil), "PriceOption")
}

func init() { proto.RegisterFile("product_price.proto", fileDescriptor_product_price_0e26de6b21ab960d) }

var fileDescriptor_product_price_0e26de6b21ab960d = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x41, 0x8b, 0x13, 0x31,
	0x14, 0x66, 0xec, 0x74, 0xba, 0x7d, 0x55, 0x76, 0x89, 0x4b, 0xb7, 0x2b, 0xa8, 0xcb, 0xe0, 0x61,
	0xf1, 0xd0, 0xc3, 0x2a, 0x88, 0x47, 0x45, 0x94, 0x82, 0x60, 0x89, 0x37, 0x2f, 0x43, 0x9c, 0xc4,
	0xdd, 0xc0, 0x90, 0x4c, 0x93, 0x37, 0x95, 0x9e, 0x3c, 0xfb, 0x5f, 0xfc, 0x91, 0xf2, 0x5e, 0x32,
	0x55, 0xf4, 0xf6, 0xbe, 0xef, 0x9b, 0xbc, 0xef, 0x9b, 0x2f, 0x81, 0x87, 0x7d, 0xf0, 0x7a, 0x68,
	0xb1, 0xe9, 0x83, 0x6d, 0xcd, 0xba, 0x0f, 0x1e, 0x7d, 0xad, 0xa0, 0x7c, 0x13, 0x8c, 0x12, 0x4f,
	0x61, 0xa1, 0x6d, 0xc4, 0x60, 0x5b, 0x6c, 0xac, 0x5e, 0x15, 0x57, 0xc5, 0xf5, 0x44, 0xc2, 0x48,
	0x6d, 0xb4, 0xb8, 0x80, 0x59, 0x6b, 0xf1, 0x40, 0xe2, 0x3d, 0x16, 0x2b, 0x82, 0x1b, 0x4d, 0x27,
	0xfb, 0xe0, 0xf7, 0xd6, 0xb5, 0x86, 0xc4, 0x49, 0x3a, 0x39, 0x52, 0x1b, 0x5d, 0xff, 0x2c, 0x60,
	0xf9, 0xc1, 0xe0, 0x36, 0xb9, 0x6f, 0xc9, 0x5c, 0x9a, 0xdd, 0x60, 0x22, 0x8a, 0xc7, 0x00, 0x63,
	0xa8, 0xa3, 0xe9, 0x3c, 0x33, 0x1b, 0x2d, 0x2e, 0xe1, 0x24, 0xf6, 0xa6, 0x6d, 0xa2, 0x41, 0x36,
	0x9d, 0xcb, 0x19, 0xe1, 0xcf, 0x06, 0xc5, 0x25, 0x94, 0x2a, 0x18, 0xc5, 0x76, 0x8b, 0x9b, 0xe9,
	0x9a, 0x7e, 0x42, 0x32, 0x25, 0x1e, 0xc1, 0xc9, 0x6e, 0x50, 0x0e, 0x2d, 0x1e, 0x56, 0x25, 0xaf,
	0x3c, 0xe2, 0xfa, 0x35, 0x5c, 0xfc, 0x17, 0x25, 0xf6, 0xde, 0x45, 0x23, 0x9e, 0x40, 0xc5, 0xc5,
	0xc4, 0x55, 0x71, 0x35, 0xb9, 0x5e, 0xdc, 0x54, 0xeb, 0xa4, 0x67, 0xb6, 0xfe, 0x01, 0x53, 0x26,
	0xc4, 0x19, 0x4c, 0x76, 0x78, 0xc8, 0x69, 0x69, 0x14, 0x02, 0xca, 0xc1, 0xd9, 0x31, 0x23, 0xcf,
	0xe2, 0x39, 0xcc, 0x7a, 0x15, 0xd0, 0x99, 0x90, 0x33, 0x9e, 0xad, 0xb7, 0x09, 0x7f, 0xf4, 0xad,
	0x42, 0xeb, 0x9d, 0x1c, 0x3f, 0x10, 0xcf, 0xa0, 0xf2, 0x3d, 0x51, 0xab, 0x92, 0xad, 0xef, 0x27,
	0xeb, 0x4f, 0xcc, 0xc9, 0xac, 0xd5, 0xaf, 0xe0, 0xf4, 0x9f, 0x0d, 0x14, 0xa5, 0x53, 0xc8, 0x51,
	0x0a, 0x49, 0x23, 0x45, 0xe9, 0xbc, 0xbb, 0xe5, 0x28, 0x85, 0xe4, 0xb9, 0xfe, 0x55, 0xc0, 0xe2,
	0xaf, 0x85, 0xe2, 0x25, 0x2c, 0x73, 0xc7, 0xd6, 0xbb, 0xe6, 0x9b, 0x75, 0x36, 0xde, 0x35, 0x5a,
	0xa1, 0xe1, 0x45, 0x73, 0x79, 0xfe, 0x47, 0x7d, 0xcf, 0xe2, 0x3b, 0x85, 0x86, 0x6a, 0xd5, 0x43,
	0x60, 0xdf, 0xfc, 0x02, 0x8e, 0x58, 0x2c, 0xa1, 0xfa, 0x6e, 0xec, 0xed, 0x1d, 0xe6, 0xeb, 0xcf,
	0x88, 0xee, 0x97, 0xca, 0x48, 0x2f, 0x2e, 0x5f, 0xc6, 0x9c, 0x98, 0xd4, 0xe4, 0x39, 0x4c, 0x93,
	0x32, 0x65, 0x25, 0x81, 0xb7, 0xa7, 0x5f, 0x1e, 0x74, 0xbe, 0x55, 0x5d, 0x34, 0x61, 0x4f, 0xcd,
	0x7f, 0xad, 0xf8, 0xa9, 0xbe, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x33, 0x30, 0xb7, 0xc8, 0xc1,
	0x02, 0x00, 0x00,
}
