// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cpb.proto

package cpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type CGI_Hello struct {
	ID                   *int32   `protobuf:"varint,1,req,name=ID" json:"ID,omitempty"`
	Desc                 *string  `protobuf:"bytes,2,req,name=Desc" json:"Desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CGI_Hello) Reset()         { *m = CGI_Hello{} }
func (m *CGI_Hello) String() string { return proto.CompactTextString(m) }
func (*CGI_Hello) ProtoMessage()    {}
func (*CGI_Hello) Descriptor() ([]byte, []int) {
	return fileDescriptor_030085b7308ea53a, []int{0}
}

func (m *CGI_Hello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CGI_Hello.Unmarshal(m, b)
}
func (m *CGI_Hello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CGI_Hello.Marshal(b, m, deterministic)
}
func (m *CGI_Hello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CGI_Hello.Merge(m, src)
}
func (m *CGI_Hello) XXX_Size() int {
	return xxx_messageInfo_CGI_Hello.Size(m)
}
func (m *CGI_Hello) XXX_DiscardUnknown() {
	xxx_messageInfo_CGI_Hello.DiscardUnknown(m)
}

var xxx_messageInfo_CGI_Hello proto.InternalMessageInfo

func (m *CGI_Hello) GetID() int32 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *CGI_Hello) GetDesc() string {
	if m != nil && m.Desc != nil {
		return *m.Desc
	}
	return ""
}

func init() {
	proto.RegisterType((*CGI_Hello)(nil), "cpb.CGI_Hello")
}

func init() { proto.RegisterFile("cpb.proto", fileDescriptor_030085b7308ea53a) }

var fileDescriptor_030085b7308ea53a = []byte{
	// 79 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2e, 0x48, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2e, 0x48, 0x52, 0x52, 0xe5, 0xe2, 0x74, 0x76,
	0xf7, 0x8c, 0xf7, 0x48, 0xcd, 0xc9, 0xc9, 0x17, 0xe2, 0xe2, 0x62, 0xf2, 0x74, 0x91, 0x60, 0x54,
	0x60, 0xd2, 0x60, 0x15, 0xe2, 0xe1, 0x62, 0x71, 0x49, 0x2d, 0x4e, 0x96, 0x60, 0x52, 0x60, 0xd2,
	0xe0, 0x04, 0x04, 0x00, 0x00, 0xff, 0xff, 0x06, 0xa0, 0x80, 0x00, 0x37, 0x00, 0x00, 0x00,
}
