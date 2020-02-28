// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/race.proto

package pb

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

type Race struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Icon                 string   `protobuf:"bytes,3,opt,name=icon,proto3" json:"icon,omitempty"`
	Bit                  int64    `protobuf:"varint,4,opt,name=bit,proto3" json:"bit,omitempty"`
	Male                 string   `protobuf:"bytes,5,opt,name=male,proto3" json:"male,omitempty"`
	Female               string   `protobuf:"bytes,6,opt,name=female,proto3" json:"female,omitempty"`
	Neutral              string   `protobuf:"bytes,7,opt,name=neutral,proto3" json:"neutral,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Race) Reset()         { *m = Race{} }
func (m *Race) String() string { return proto.CompactTextString(m) }
func (*Race) ProtoMessage()    {}
func (*Race) Descriptor() ([]byte, []int) {
	return fileDescriptor_a92efc8fd8e97813, []int{0}
}

func (m *Race) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Race.Unmarshal(m, b)
}
func (m *Race) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Race.Marshal(b, m, deterministic)
}
func (m *Race) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Race.Merge(m, src)
}
func (m *Race) XXX_Size() int {
	return xxx_messageInfo_Race.Size(m)
}
func (m *Race) XXX_DiscardUnknown() {
	xxx_messageInfo_Race.DiscardUnknown(m)
}

var xxx_messageInfo_Race proto.InternalMessageInfo

func (m *Race) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Race) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Race) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Race) GetBit() int64 {
	if m != nil {
		return m.Bit
	}
	return 0
}

func (m *Race) GetMale() string {
	if m != nil {
		return m.Male
	}
	return ""
}

func (m *Race) GetFemale() string {
	if m != nil {
		return m.Female
	}
	return ""
}

func (m *Race) GetNeutral() string {
	if m != nil {
		return m.Neutral
	}
	return ""
}

func init() {
	proto.RegisterType((*Race)(nil), "pb.Race")
}

func init() { proto.RegisterFile("proto/race.proto", fileDescriptor_a92efc8fd8e97813) }

var fileDescriptor_a92efc8fd8e97813 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0x41, 0x0a, 0xc2, 0x30,
	0x10, 0x45, 0x49, 0x5a, 0x53, 0x9c, 0x85, 0x94, 0x59, 0xc8, 0x2c, 0x8b, 0xab, 0xae, 0x74, 0xe1,
	0x4d, 0x72, 0x83, 0x24, 0x1d, 0x21, 0xd0, 0x26, 0x25, 0xc4, 0xab, 0x78, 0x5e, 0xc9, 0x54, 0x77,
	0xef, 0xbf, 0xf9, 0x03, 0x1f, 0xc6, 0xbd, 0xe4, 0x9a, 0x1f, 0xc5, 0x05, 0xbe, 0x0b, 0xa2, 0xde,
	0xfd, 0xed, 0xa3, 0xa0, 0xb7, 0x2e, 0x30, 0x5e, 0x40, 0xc7, 0x85, 0xd4, 0xa4, 0xe6, 0xce, 0xea,
	0xb8, 0x20, 0x42, 0x9f, 0xdc, 0xc6, 0xa4, 0x27, 0x35, 0x9f, 0xad, 0x70, 0x73, 0x31, 0xe4, 0x44,
	0xdd, 0xe1, 0x1a, 0xe3, 0x08, 0x9d, 0x8f, 0x95, 0x7a, 0x79, 0x6c, 0xd8, 0x5a, 0x9b, 0x5b, 0x99,
	0x4e, 0x47, 0xab, 0x31, 0x5e, 0xc1, 0xbc, 0x58, 0xac, 0x11, 0xfb, 0x4b, 0x48, 0x30, 0x24, 0x7e,
	0xd7, 0xe2, 0x56, 0x1a, 0xe4, 0xf0, 0x8f, 0xde, 0xc8, 0xc6, 0xe7, 0x37, 0x00, 0x00, 0xff, 0xff,
	0x4a, 0xcb, 0x85, 0x08, 0xb7, 0x00, 0x00, 0x00,
}
