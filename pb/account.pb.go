// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/account.proto

package pb

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

type Account struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Charname             string   `protobuf:"bytes,3,opt,name=charname,proto3" json:"charname,omitempty"`
	Sharedplat           int64    `protobuf:"varint,4,opt,name=sharedplat,proto3" json:"sharedplat,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	Status               int64    `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Lsid                 string   `protobuf:"bytes,7,opt,name=lsid,proto3" json:"lsid,omitempty"`
	Lsaccountid          int64    `protobuf:"varint,8,opt,name=lsaccountid,proto3" json:"lsaccountid,omitempty"`
	Gmspeed              int64    `protobuf:"varint,9,opt,name=gmspeed,proto3" json:"gmspeed,omitempty"`
	Revoked              int64    `protobuf:"varint,10,opt,name=revoked,proto3" json:"revoked,omitempty"`
	Karma                int64    `protobuf:"varint,11,opt,name=karma,proto3" json:"karma,omitempty"`
	Miniloginip          string   `protobuf:"bytes,12,opt,name=miniloginip,proto3" json:"miniloginip,omitempty"`
	Hideme               int64    `protobuf:"varint,13,opt,name=hideme,proto3" json:"hideme,omitempty"`
	Rulesflag            int64    `protobuf:"varint,14,opt,name=rulesflag,proto3" json:"rulesflag,omitempty"`
	Suspendeduntil       int64    `protobuf:"varint,15,opt,name=suspendeduntil,proto3" json:"suspendeduntil,omitempty"`
	Timecreation         int64    `protobuf:"varint,16,opt,name=timecreation,proto3" json:"timecreation,omitempty"`
	Expansion            int64    `protobuf:"varint,17,opt,name=expansion,proto3" json:"expansion,omitempty"`
	Banreason            string   `protobuf:"bytes,18,opt,name=banreason,proto3" json:"banreason,omitempty"`
	Suspendreason        string   `protobuf:"bytes,19,opt,name=suspendreason,proto3" json:"suspendreason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_477cbf5ae5b46edf, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Account) GetCharname() string {
	if m != nil {
		return m.Charname
	}
	return ""
}

func (m *Account) GetSharedplat() int64 {
	if m != nil {
		return m.Sharedplat
	}
	return 0
}

func (m *Account) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Account) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Account) GetLsid() string {
	if m != nil {
		return m.Lsid
	}
	return ""
}

func (m *Account) GetLsaccountid() int64 {
	if m != nil {
		return m.Lsaccountid
	}
	return 0
}

func (m *Account) GetGmspeed() int64 {
	if m != nil {
		return m.Gmspeed
	}
	return 0
}

func (m *Account) GetRevoked() int64 {
	if m != nil {
		return m.Revoked
	}
	return 0
}

func (m *Account) GetKarma() int64 {
	if m != nil {
		return m.Karma
	}
	return 0
}

func (m *Account) GetMiniloginip() string {
	if m != nil {
		return m.Miniloginip
	}
	return ""
}

func (m *Account) GetHideme() int64 {
	if m != nil {
		return m.Hideme
	}
	return 0
}

func (m *Account) GetRulesflag() int64 {
	if m != nil {
		return m.Rulesflag
	}
	return 0
}

func (m *Account) GetSuspendeduntil() int64 {
	if m != nil {
		return m.Suspendeduntil
	}
	return 0
}

func (m *Account) GetTimecreation() int64 {
	if m != nil {
		return m.Timecreation
	}
	return 0
}

func (m *Account) GetExpansion() int64 {
	if m != nil {
		return m.Expansion
	}
	return 0
}

func (m *Account) GetBanreason() string {
	if m != nil {
		return m.Banreason
	}
	return ""
}

func (m *Account) GetSuspendreason() string {
	if m != nil {
		return m.Suspendreason
	}
	return ""
}

func init() {
	proto.RegisterType((*Account)(nil), "pb.Account")
}

func init() { proto.RegisterFile("proto/account.proto", fileDescriptor_477cbf5ae5b46edf) }

var fileDescriptor_477cbf5ae5b46edf = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x6f, 0x8e, 0xd3, 0x30,
	0x10, 0xc5, 0xd5, 0x6c, 0xb7, 0xdd, 0xce, 0xee, 0x16, 0xf0, 0x22, 0x34, 0x5a, 0xad, 0x50, 0x55,
	0x21, 0xd4, 0x4f, 0xf4, 0x03, 0x27, 0xe0, 0x0a, 0xbd, 0xc1, 0x34, 0x1e, 0x52, 0xab, 0x8e, 0x6d,
	0xd9, 0x0e, 0x70, 0x1f, 0x2e, 0x8a, 0x3c, 0x4e, 0xff, 0xf1, 0x6d, 0xde, 0x6f, 0xde, 0xcb, 0xbc,
	0x28, 0x81, 0x97, 0x10, 0x7d, 0xf6, 0x5b, 0x6a, 0x5b, 0x3f, 0xb8, 0xfc, 0x4d, 0x94, 0x6a, 0xc2,
	0xfe, 0xf5, 0xad, 0xf3, 0xbe, 0xb3, 0xbc, 0xa5, 0x60, 0xb6, 0xe4, 0x9c, 0xcf, 0x94, 0x8d, 0x77,
	0xa9, 0x3a, 0xd6, 0x7f, 0xa7, 0x30, 0xff, 0x51, 0x33, 0x6a, 0x09, 0x8d, 0xd1, 0x38, 0x59, 0x4d,
	0x36, 0x77, 0xbb, 0xc6, 0x68, 0xa5, 0x60, 0xea, 0xa8, 0x67, 0x6c, 0x56, 0x93, 0xcd, 0x62, 0x27,
	0xb3, 0x7a, 0x85, 0x87, 0xf6, 0x40, 0x51, 0xf8, 0x9d, 0xf0, 0xb3, 0x56, 0x9f, 0x01, 0xd2, 0x81,
	0x22, 0xeb, 0x60, 0x29, 0xe3, 0x54, 0x9e, 0x73, 0x45, 0x4a, 0x36, 0x50, 0x4a, 0xbf, 0x7d, 0xd4,
	0x78, 0x5f, 0xb3, 0x27, 0xad, 0x3e, 0xc1, 0x2c, 0x65, 0xca, 0x43, 0xc2, 0x99, 0xe4, 0x46, 0x55,
	0x3a, 0xd8, 0x64, 0x34, 0xce, 0x6b, 0x87, 0x32, 0xab, 0x15, 0x3c, 0xda, 0x34, 0xbe, 0xa8, 0xd1,
	0xf8, 0x20, 0x81, 0x6b, 0xa4, 0x10, 0xe6, 0x5d, 0x9f, 0x02, 0xb3, 0xc6, 0x85, 0x6c, 0x4f, 0xb2,
	0x6c, 0x22, 0xff, 0xf2, 0x47, 0xd6, 0x08, 0x75, 0x33, 0x4a, 0xf5, 0x11, 0xee, 0x8f, 0x14, 0x7b,
	0xc2, 0x47, 0xe1, 0x55, 0x94, 0x5b, 0xbd, 0x71, 0xc6, 0xfa, 0xce, 0x38, 0x13, 0xf0, 0x49, 0x6a,
	0x5c, 0xa3, 0xd2, 0xfc, 0x60, 0x34, 0xf7, 0x8c, 0xcf, 0xb5, 0x79, 0x55, 0xea, 0x0d, 0x16, 0x71,
	0xb0, 0x9c, 0x7e, 0x5a, 0xea, 0x70, 0x29, 0xab, 0x0b, 0x50, 0x5f, 0x61, 0x99, 0x86, 0x14, 0xd8,
	0x69, 0xd6, 0xa5, 0xb3, 0xc5, 0x77, 0x62, 0xf9, 0x8f, 0xaa, 0x35, 0x3c, 0x65, 0xd3, 0x73, 0x1b,
	0x59, 0x3e, 0x1b, 0xbe, 0x17, 0xd7, 0x0d, 0x2b, 0x97, 0xf8, 0x4f, 0x20, 0x97, 0x8a, 0xe1, 0x43,
	0xbd, 0x74, 0x06, 0x65, 0xbb, 0x27, 0x17, 0x99, 0x92, 0x77, 0xa8, 0xa4, 0xff, 0x05, 0xa8, 0x2f,
	0xf0, 0x3c, 0x5e, 0x1c, 0x1d, 0x2f, 0xe2, 0xb8, 0x85, 0xfb, 0x99, 0xfc, 0x2c, 0xdf, 0xff, 0x05,
	0x00, 0x00, 0xff, 0xff, 0xf2, 0x1e, 0xe7, 0x80, 0x65, 0x02, 0x00, 0x00,
}
