// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/display_keyword_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A display keyword view.
type DisplayKeywordView struct {
	// Immutable. The resource name of the display keyword view.
	// Display Keyword view resource names have the form:
	//
	// `customers/{customer_id}/displayKeywordViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisplayKeywordView) Reset()         { *m = DisplayKeywordView{} }
func (m *DisplayKeywordView) String() string { return proto.CompactTextString(m) }
func (*DisplayKeywordView) ProtoMessage()    {}
func (*DisplayKeywordView) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbae9f6689a25c70, []int{0}
}

func (m *DisplayKeywordView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisplayKeywordView.Unmarshal(m, b)
}
func (m *DisplayKeywordView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisplayKeywordView.Marshal(b, m, deterministic)
}
func (m *DisplayKeywordView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisplayKeywordView.Merge(m, src)
}
func (m *DisplayKeywordView) XXX_Size() int {
	return xxx_messageInfo_DisplayKeywordView.Size(m)
}
func (m *DisplayKeywordView) XXX_DiscardUnknown() {
	xxx_messageInfo_DisplayKeywordView.DiscardUnknown(m)
}

var xxx_messageInfo_DisplayKeywordView proto.InternalMessageInfo

func (m *DisplayKeywordView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*DisplayKeywordView)(nil), "google.ads.googleads.v2.resources.DisplayKeywordView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/display_keyword_view.proto", fileDescriptor_bbae9f6689a25c70)
}

var fileDescriptor_bbae9f6689a25c70 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4a, 0xeb, 0x40,
	0x18, 0x85, 0x49, 0x2e, 0xf7, 0xc2, 0x0d, 0xf7, 0x6e, 0xb2, 0x51, 0x8b, 0xa0, 0x15, 0x0a, 0x82,
	0x38, 0x03, 0xe9, 0x6e, 0x14, 0x64, 0x8a, 0x50, 0x50, 0x90, 0xd2, 0x45, 0x10, 0x09, 0x84, 0x69,
	0x66, 0x8c, 0x83, 0x49, 0xfe, 0x98, 0x49, 0x53, 0x4a, 0xe9, 0xce, 0x27, 0x71, 0xe9, 0xa3, 0xf8,
	0x14, 0xba, 0xed, 0x23, 0xb8, 0x92, 0x76, 0x3a, 0x69, 0xa1, 0xa2, 0xb8, 0x3b, 0xf0, 0x7f, 0xe7,
	0xcc, 0xc9, 0x21, 0xce, 0x69, 0x0c, 0x10, 0x27, 0x02, 0x33, 0xae, 0xb0, 0x96, 0x73, 0x55, 0x79,
	0xb8, 0x10, 0x0a, 0x86, 0x45, 0x24, 0x14, 0xe6, 0x52, 0xe5, 0x09, 0x1b, 0x87, 0xf7, 0x62, 0x3c,
	0x82, 0x82, 0x87, 0x95, 0x14, 0x23, 0x94, 0x17, 0x50, 0x82, 0xdb, 0xd4, 0x16, 0xc4, 0xb8, 0x42,
	0xb5, 0x1b, 0x55, 0x1e, 0xaa, 0xdd, 0x8d, 0x3d, 0xf3, 0x40, 0x2e, 0xf1, 0xad, 0x14, 0x09, 0x0f,
	0x07, 0xe2, 0x8e, 0x55, 0x12, 0x0a, 0x9d, 0xd1, 0xd8, 0x59, 0x03, 0x8c, 0x6d, 0x79, 0xda, 0x5d,
	0x3b, 0xb1, 0x2c, 0x83, 0x92, 0x95, 0x12, 0x32, 0xa5, 0xaf, 0x07, 0x6f, 0x96, 0xe3, 0x9e, 0xeb,
	0x6e, 0x97, 0xba, 0x9a, 0x2f, 0xc5, 0xc8, 0xbd, 0x76, 0xfe, 0x9b, 0x98, 0x30, 0x63, 0xa9, 0xd8,
	0xb6, 0xf6, 0xad, 0xc3, 0xbf, 0x9d, 0xf6, 0x2b, 0xfd, 0xfd, 0x4e, 0x8f, 0x9d, 0xa3, 0x55, 0xcf,
	0xa5, 0xca, 0xa5, 0x42, 0x11, 0xa4, 0x78, 0x33, 0xab, 0xff, 0xcf, 0x24, 0x5d, 0xb1, 0x54, 0x90,
	0x87, 0x19, 0xcd, 0x7e, 0xe4, 0x77, 0xcf, 0xa2, 0xa1, 0x2a, 0x21, 0x15, 0x85, 0xc2, 0x13, 0x23,
	0xa7, 0x66, 0xd0, 0x35, 0x50, 0xe1, 0xc9, 0x67, 0x2b, 0x4f, 0x3b, 0x8f, 0xb6, 0xd3, 0x8a, 0x20,
	0x45, 0xdf, 0xee, 0xdc, 0xd9, 0xda, 0x7c, 0xbe, 0x37, 0x9f, 0xa9, 0x67, 0xdd, 0x5c, 0x2c, 0xdd,
	0x31, 0x24, 0x2c, 0x8b, 0x11, 0x14, 0x31, 0x8e, 0x45, 0xb6, 0x18, 0x11, 0xaf, 0xbe, 0xe1, 0x8b,
	0x5f, 0xe0, 0xa4, 0x56, 0x4f, 0xf6, 0xaf, 0x2e, 0xa5, 0xcf, 0x76, 0xb3, 0xab, 0x23, 0x29, 0x57,
	0x48, 0xcb, 0xb9, 0xf2, 0x3d, 0xd4, 0x37, 0xe4, 0x8b, 0x61, 0x02, 0xca, 0x55, 0x50, 0x33, 0x81,
	0xef, 0x05, 0x35, 0x33, 0xb3, 0x5b, 0xfa, 0x40, 0x08, 0xe5, 0x8a, 0x90, 0x9a, 0x22, 0xc4, 0xf7,
	0x08, 0xa9, 0xb9, 0xc1, 0x9f, 0x45, 0xd9, 0xf6, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x5f,
	0x1d, 0xc8, 0xae, 0x02, 0x00, 0x00,
}
