// Code generated by protoc-gen-go. DO NOT EDIT.
// source: object_type.proto

package runm

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

// Indicates the scope of a type of object. The object type's scope indicates
// the level at which an object's name is guaranteed to be unique. Objects that
// have an object type with a PROJECT object type scope must be created with a
// specific project identifier. Objects with a PARTITION object type scope must
// be created with a partition UUID.
type ObjectTypeScope int32

const (
	ObjectTypeScope_PARTITION ObjectTypeScope = 0
	ObjectTypeScope_PROJECT   ObjectTypeScope = 1
)

var ObjectTypeScope_name = map[int32]string{
	0: "PARTITION",
	1: "PROJECT",
}
var ObjectTypeScope_value = map[string]int32{
	"PARTITION": 0,
	"PROJECT":   1,
}

func (x ObjectTypeScope) String() string {
	return proto.EnumName(ObjectTypeScope_name, int32(x))
}
func (ObjectTypeScope) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_object_type_c82d8226e6f15b4e, []int{0}
}

// An object type is a simple classification for various types of things known
// to the runm system
type ObjectType struct {
	Code        string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// Indicates the scope that names of objects of this type must guarantee
	// uniqueness for
	Scope                ObjectTypeScope `protobuf:"varint,3,opt,name=scope,enum=runm.ObjectTypeScope" json:"scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ObjectType) Reset()         { *m = ObjectType{} }
func (m *ObjectType) String() string { return proto.CompactTextString(m) }
func (*ObjectType) ProtoMessage()    {}
func (*ObjectType) Descriptor() ([]byte, []int) {
	return fileDescriptor_object_type_c82d8226e6f15b4e, []int{0}
}
func (m *ObjectType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectType.Unmarshal(m, b)
}
func (m *ObjectType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectType.Marshal(b, m, deterministic)
}
func (dst *ObjectType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectType.Merge(dst, src)
}
func (m *ObjectType) XXX_Size() int {
	return xxx_messageInfo_ObjectType.Size(m)
}
func (m *ObjectType) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectType.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectType proto.InternalMessageInfo

func (m *ObjectType) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ObjectType) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ObjectType) GetScope() ObjectTypeScope {
	if m != nil {
		return m.Scope
	}
	return ObjectTypeScope_PARTITION
}

// Used in matching object type records
type ObjectTypeFilter struct {
	// A search term on the object type's string code
	Search string `protobuf:"bytes,1,opt,name=search" json:"search,omitempty"`
	// Indicates the search should be a prefix expression
	UsePrefix            bool     `protobuf:"varint,2,opt,name=use_prefix,json=usePrefix" json:"use_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectTypeFilter) Reset()         { *m = ObjectTypeFilter{} }
func (m *ObjectTypeFilter) String() string { return proto.CompactTextString(m) }
func (*ObjectTypeFilter) ProtoMessage()    {}
func (*ObjectTypeFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_object_type_c82d8226e6f15b4e, []int{1}
}
func (m *ObjectTypeFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectTypeFilter.Unmarshal(m, b)
}
func (m *ObjectTypeFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectTypeFilter.Marshal(b, m, deterministic)
}
func (dst *ObjectTypeFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectTypeFilter.Merge(dst, src)
}
func (m *ObjectTypeFilter) XXX_Size() int {
	return xxx_messageInfo_ObjectTypeFilter.Size(m)
}
func (m *ObjectTypeFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectTypeFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectTypeFilter proto.InternalMessageInfo

func (m *ObjectTypeFilter) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

func (m *ObjectTypeFilter) GetUsePrefix() bool {
	if m != nil {
		return m.UsePrefix
	}
	return false
}

func init() {
	proto.RegisterType((*ObjectType)(nil), "runm.ObjectType")
	proto.RegisterType((*ObjectTypeFilter)(nil), "runm.ObjectTypeFilter")
	proto.RegisterEnum("runm.ObjectTypeScope", ObjectTypeScope_name, ObjectTypeScope_value)
}

func init() { proto.RegisterFile("object_type.proto", fileDescriptor_object_type_c82d8226e6f15b4e) }

var fileDescriptor_object_type_c82d8226e6f15b4e = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x41, 0x4b, 0x86, 0x40,
	0x10, 0x86, 0xdb, 0xfa, 0xfa, 0xca, 0x91, 0xca, 0x06, 0x0a, 0x2f, 0x81, 0x78, 0x92, 0x22, 0x0f,
	0xf5, 0x0b, 0x22, 0x0a, 0xec, 0x90, 0xb2, 0xed, 0x5d, 0x72, 0x9d, 0x68, 0xa3, 0xdc, 0x65, 0x77,
	0x85, 0xfc, 0xf7, 0xe1, 0x26, 0x18, 0xdd, 0x66, 0xde, 0x67, 0x78, 0x1f, 0x06, 0x4e, 0x75, 0xf7,
	0x41, 0xd2, 0xb7, 0x7e, 0x32, 0x54, 0x1a, 0xab, 0xbd, 0xc6, 0x8d, 0x1d, 0x87, 0xaf, 0x5c, 0x03,
	0xd4, 0x01, 0x89, 0xc9, 0x10, 0x22, 0x6c, 0xa4, 0xee, 0x29, 0x65, 0x19, 0x2b, 0x22, 0x1e, 0x66,
	0xcc, 0x20, 0xee, 0xc9, 0x49, 0xab, 0x8c, 0x57, 0x7a, 0x48, 0x77, 0x03, 0xfa, 0x1b, 0xe1, 0x15,
	0xec, 0x3b, 0xa9, 0x0d, 0xa5, 0x7b, 0x19, 0x2b, 0x8e, 0x6f, 0xce, 0xca, 0xb9, 0xb9, 0x5c, 0x6b,
	0x5f, 0x66, 0xc8, 0x7f, 0x6f, 0xf2, 0x0a, 0x92, 0x95, 0x3c, 0xaa, 0x4f, 0x4f, 0x16, 0xcf, 0x61,
	0xeb, 0xe8, 0xd5, 0xca, 0xf7, 0x45, 0xbc, 0x6c, 0x78, 0x01, 0x30, 0x3a, 0x6a, 0x8d, 0xa5, 0x37,
	0xf5, 0x1d, 0xcc, 0x87, 0x3c, 0x1a, 0x1d, 0x35, 0x21, 0xb8, 0xbc, 0x86, 0x93, 0x7f, 0x12, 0x3c,
	0x82, 0xa8, 0xb9, 0xe3, 0xa2, 0x12, 0x55, 0xfd, 0x9c, 0xec, 0x60, 0x0c, 0x07, 0x0d, 0xaf, 0x9f,
	0x1e, 0xee, 0x45, 0xc2, 0xba, 0x6d, 0xf8, 0xfb, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x82, 0x90,
	0x51, 0x0e, 0x0c, 0x01, 0x00, 0x00,
}