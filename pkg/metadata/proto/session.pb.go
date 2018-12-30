// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session.proto

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

// A session is a context for a series of requests made to a service endpoint
// or executor
type Session struct {
	// A user takes some action against the system
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	// A project is a tenant/account the user belongs to. It associates a user
	// with billing and quotas.
	Project string `protobuf:"bytes,2,opt,name=project" json:"project,omitempty"`
	// The partition that the user is "targeting".
	Partition            string   `protobuf:"bytes,3,opt,name=partition" json:"partition,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_session_163fb7a2098114d7, []int{0}
}
func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (dst *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(dst, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Session) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Session) GetPartition() string {
	if m != nil {
		return m.Partition
	}
	return ""
}

func init() {
	proto.RegisterType((*Session)(nil), "runm.Session")
}

func init() { proto.RegisterFile("session.proto", fileDescriptor_session_163fb7a2098114d7) }

var fileDescriptor_session_163fb7a2098114d7 = []byte{
	// 112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2e,
	0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x2a, 0xcd, 0xcb, 0x95,
	0xe2, 0x2f, 0x48, 0x2c, 0x2a, 0xc9, 0x2c, 0x81, 0x0b, 0x2b, 0x85, 0x72, 0xb1, 0x07, 0x43, 0xd4,
	0x09, 0x09, 0x71, 0xb1, 0x94, 0x16, 0xa7, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81,
	0xd9, 0x42, 0x12, 0x5c, 0xec, 0x05, 0x45, 0xf9, 0x59, 0xa9, 0xc9, 0x25, 0x12, 0x4c, 0x60, 0x61,
	0x18, 0x57, 0x48, 0x86, 0x8b, 0x13, 0x6e, 0x96, 0x04, 0x33, 0x58, 0x0e, 0x21, 0x90, 0xc4, 0x06,
	0x36, 0xdd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x29, 0xe2, 0xcc, 0x3a, 0x85, 0x00, 0x00, 0x00,
}