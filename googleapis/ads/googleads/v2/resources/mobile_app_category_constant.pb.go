// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/mobile_app_category_constant.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A mobile application category constant.
type MobileAppCategoryConstant struct {
	// Immutable. The resource name of the mobile app category constant.
	// Mobile app category constant resource names have the form:
	//
	// `mobileAppCategoryConstants/{mobile_app_category_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the mobile app category constant.
	Id *wrappers.Int32Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. Mobile app category name.
	Name                 *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MobileAppCategoryConstant) Reset()         { *m = MobileAppCategoryConstant{} }
func (m *MobileAppCategoryConstant) String() string { return proto.CompactTextString(m) }
func (*MobileAppCategoryConstant) ProtoMessage()    {}
func (*MobileAppCategoryConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fad794294b415b1, []int{0}
}

func (m *MobileAppCategoryConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileAppCategoryConstant.Unmarshal(m, b)
}
func (m *MobileAppCategoryConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileAppCategoryConstant.Marshal(b, m, deterministic)
}
func (m *MobileAppCategoryConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileAppCategoryConstant.Merge(m, src)
}
func (m *MobileAppCategoryConstant) XXX_Size() int {
	return xxx_messageInfo_MobileAppCategoryConstant.Size(m)
}
func (m *MobileAppCategoryConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileAppCategoryConstant.DiscardUnknown(m)
}

var xxx_messageInfo_MobileAppCategoryConstant proto.InternalMessageInfo

func (m *MobileAppCategoryConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *MobileAppCategoryConstant) GetId() *wrappers.Int32Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MobileAppCategoryConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func init() {
	proto.RegisterType((*MobileAppCategoryConstant)(nil), "google.ads.googleads.v2.resources.MobileAppCategoryConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/mobile_app_category_constant.proto", fileDescriptor_8fad794294b415b1)
}

var fileDescriptor_8fad794294b415b1 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcf, 0x6a, 0x14, 0x31,
	0x18, 0x67, 0xb3, 0x2a, 0x38, 0xea, 0x65, 0x4e, 0x6d, 0x2d, 0xb5, 0x15, 0x0a, 0x3d, 0x25, 0x92,
	0xea, 0xc1, 0x78, 0xca, 0x56, 0x28, 0x0a, 0x4a, 0x59, 0x61, 0x0f, 0xb2, 0x30, 0x64, 0x26, 0x69,
	0x0c, 0xcc, 0xe4, 0x0b, 0x49, 0x76, 0x45, 0xc4, 0x27, 0xf0, 0x2d, 0x3c, 0xfa, 0x28, 0xbe, 0x81,
	0xb7, 0x9e, 0xfb, 0x08, 0x9e, 0x64, 0x67, 0x32, 0xd3, 0x05, 0xd9, 0x15, 0x7a, 0xfb, 0x85, 0xdf,
	0x9f, 0xef, 0xfb, 0xf1, 0x25, 0x7b, 0xad, 0x01, 0x74, 0xad, 0x88, 0x90, 0x81, 0x74, 0x70, 0x85,
	0x96, 0x94, 0x78, 0x15, 0x60, 0xe1, 0x2b, 0x15, 0x48, 0x03, 0xa5, 0xa9, 0x55, 0x21, 0x9c, 0x2b,
	0x2a, 0x11, 0x95, 0x06, 0xff, 0xa5, 0xa8, 0xc0, 0x86, 0x28, 0x6c, 0xc4, 0xce, 0x43, 0x84, 0xfc,
	0xa8, 0xb3, 0x62, 0x21, 0x03, 0x1e, 0x52, 0xf0, 0x92, 0xe2, 0x21, 0x65, 0xef, 0x49, 0x3f, 0xc8,
	0x19, 0x72, 0x69, 0x54, 0x2d, 0x8b, 0x52, 0x7d, 0x12, 0x4b, 0x03, 0xbe, 0xcb, 0xd8, 0xdb, 0x5d,
	0x13, 0xf4, 0xb6, 0x44, 0x1d, 0x24, 0xaa, 0x7d, 0x95, 0x8b, 0x4b, 0xf2, 0xd9, 0x0b, 0xe7, 0x94,
	0x0f, 0x89, 0xdf, 0x5f, 0xb3, 0x0a, 0x6b, 0x21, 0x8a, 0x68, 0xc0, 0x26, 0xf6, 0xe9, 0x6f, 0x94,
	0xed, 0xbe, 0x6b, 0x3b, 0x70, 0xe7, 0xce, 0x52, 0x83, 0xb3, 0x54, 0x20, 0x2f, 0xb2, 0x47, 0xfd,
	0xb4, 0xc2, 0x8a, 0x46, 0xed, 0x8c, 0x0e, 0x47, 0x27, 0xf7, 0x27, 0xec, 0x8a, 0xdf, 0xfd, 0xc3,
	0x9f, 0x67, 0xf4, 0xa6, 0x4e, 0x42, 0xce, 0x04, 0x5c, 0x41, 0x43, 0x36, 0x46, 0x4e, 0x1f, 0xf6,
	0x81, 0xef, 0x45, 0xa3, 0xf2, 0x67, 0x19, 0x32, 0x72, 0x07, 0x1d, 0x8e, 0x4e, 0x1e, 0xd0, 0xc7,
	0x29, 0x04, 0xf7, 0x4d, 0xf0, 0x1b, 0x1b, 0x4f, 0xe9, 0x4c, 0xd4, 0x0b, 0x35, 0x19, 0x5f, 0xf1,
	0xf1, 0x14, 0x19, 0x99, 0xbf, 0xc8, 0xee, 0xb4, 0x9b, 0x8c, 0x5b, 0xcf, 0xfe, 0x3f, 0x9e, 0x0f,
	0xd1, 0x1b, 0xab, 0xd7, 0x4c, 0xad, 0x9c, 0xf9, 0x6b, 0x0e, 0xb7, 0xd9, 0x37, 0x7f, 0xd9, 0x6c,
	0xa2, 0x02, 0xf9, 0xba, 0xed, 0xfa, 0xdf, 0x26, 0xdf, 0x51, 0x76, 0x5c, 0x41, 0x83, 0xff, 0x7b,
	0xff, 0xc9, 0xc1, 0xc6, 0xf9, 0x17, 0xab, 0x5e, 0x17, 0xa3, 0x8f, 0x6f, 0x53, 0x88, 0x86, 0x5a,
	0x58, 0x8d, 0xc1, 0x6b, 0xa2, 0x95, 0x6d, 0x5b, 0x93, 0x9b, 0x2e, 0x5b, 0x7e, 0xea, 0xab, 0x01,
	0xfd, 0x40, 0xe3, 0x73, 0xce, 0x7f, 0xa2, 0xa3, 0xf3, 0x2e, 0x92, 0xcb, 0x80, 0x3b, 0xb8, 0x42,
	0x33, 0x8a, 0xa7, 0xbd, 0xf2, 0x57, 0xaf, 0x99, 0x73, 0x19, 0xe6, 0x83, 0x66, 0x3e, 0xa3, 0xf3,
	0x41, 0x73, 0x8d, 0x8e, 0x3b, 0x82, 0x31, 0x2e, 0x03, 0x63, 0x83, 0x8a, 0xb1, 0x19, 0x65, 0x6c,
	0xd0, 0x95, 0xf7, 0xda, 0x65, 0x4f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xb7, 0xc8, 0x8a,
	0x55, 0x03, 0x00, 0x00,
}
