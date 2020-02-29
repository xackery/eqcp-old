// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/bug.proto

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

type Bug struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Zone                 string   `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	Clientversionid      int64    `protobuf:"varint,3,opt,name=clientversionid,proto3" json:"clientversionid,omitempty"`
	Clientversionname    string   `protobuf:"bytes,4,opt,name=clientversionname,proto3" json:"clientversionname,omitempty"`
	Accountid            int64    `protobuf:"varint,5,opt,name=accountid,proto3" json:"accountid,omitempty"`
	Characterid          int64    `protobuf:"varint,6,opt,name=characterid,proto3" json:"characterid,omitempty"`
	Charactername        string   `protobuf:"bytes,7,opt,name=charactername,proto3" json:"charactername,omitempty"`
	Reporterspoof        int64    `protobuf:"varint,8,opt,name=reporterspoof,proto3" json:"reporterspoof,omitempty"`
	Categoryid           int64    `protobuf:"varint,9,opt,name=categoryid,proto3" json:"categoryid,omitempty"`
	Categoryname         string   `protobuf:"bytes,10,opt,name=categoryname,proto3" json:"categoryname,omitempty"`
	Reportername         string   `protobuf:"bytes,11,opt,name=reportername,proto3" json:"reportername,omitempty"`
	Uipath               string   `protobuf:"bytes,12,opt,name=uipath,proto3" json:"uipath,omitempty"`
	Posx                 int64    `protobuf:"varint,13,opt,name=posx,proto3" json:"posx,omitempty"`
	Posy                 int64    `protobuf:"varint,14,opt,name=posy,proto3" json:"posy,omitempty"`
	Posz                 int64    `protobuf:"varint,15,opt,name=posz,proto3" json:"posz,omitempty"`
	Heading              int64    `protobuf:"varint,16,opt,name=heading,proto3" json:"heading,omitempty"`
	Timeplayed           int64    `protobuf:"varint,17,opt,name=timeplayed,proto3" json:"timeplayed,omitempty"`
	Targetid             int64    `protobuf:"varint,18,opt,name=targetid,proto3" json:"targetid,omitempty"`
	Targetname           string   `protobuf:"bytes,19,opt,name=targetname,proto3" json:"targetname,omitempty"`
	Optionalinfomask     int64    `protobuf:"varint,20,opt,name=optionalinfomask,proto3" json:"optionalinfomask,omitempty"`
	Canduplicate         int64    `protobuf:"varint,21,opt,name=canduplicate,proto3" json:"canduplicate,omitempty"`
	Crashbug             int64    `protobuf:"varint,22,opt,name=crashbug,proto3" json:"crashbug,omitempty"`
	Targetinfo           int64    `protobuf:"varint,23,opt,name=targetinfo,proto3" json:"targetinfo,omitempty"`
	Characterflags       int64    `protobuf:"varint,24,opt,name=characterflags,proto3" json:"characterflags,omitempty"`
	Unknownvalue         int64    `protobuf:"varint,25,opt,name=unknownvalue,proto3" json:"unknownvalue,omitempty"`
	Bugreport            string   `protobuf:"bytes,26,opt,name=bugreport,proto3" json:"bugreport,omitempty"`
	Systeminfo           string   `protobuf:"bytes,27,opt,name=systeminfo,proto3" json:"systeminfo,omitempty"`
	Reportdatetime       int64    `protobuf:"varint,28,opt,name=reportdatetime,proto3" json:"reportdatetime,omitempty"`
	Bugstatus            int64    `protobuf:"varint,29,opt,name=bugstatus,proto3" json:"bugstatus,omitempty"`
	Lastreview           int64    `protobuf:"varint,30,opt,name=lastreview,proto3" json:"lastreview,omitempty"`
	Lastreviewer         string   `protobuf:"bytes,31,opt,name=lastreviewer,proto3" json:"lastreviewer,omitempty"`
	Reviewernotes        string   `protobuf:"bytes,32,opt,name=reviewernotes,proto3" json:"reviewernotes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bug) Reset()         { *m = Bug{} }
func (m *Bug) String() string { return proto.CompactTextString(m) }
func (*Bug) ProtoMessage()    {}
func (*Bug) Descriptor() ([]byte, []int) {
	return fileDescriptor_e891962b823e3f16, []int{0}
}

func (m *Bug) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bug.Unmarshal(m, b)
}
func (m *Bug) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bug.Marshal(b, m, deterministic)
}
func (m *Bug) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bug.Merge(m, src)
}
func (m *Bug) XXX_Size() int {
	return xxx_messageInfo_Bug.Size(m)
}
func (m *Bug) XXX_DiscardUnknown() {
	xxx_messageInfo_Bug.DiscardUnknown(m)
}

var xxx_messageInfo_Bug proto.InternalMessageInfo

func (m *Bug) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Bug) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *Bug) GetClientversionid() int64 {
	if m != nil {
		return m.Clientversionid
	}
	return 0
}

func (m *Bug) GetClientversionname() string {
	if m != nil {
		return m.Clientversionname
	}
	return ""
}

func (m *Bug) GetAccountid() int64 {
	if m != nil {
		return m.Accountid
	}
	return 0
}

func (m *Bug) GetCharacterid() int64 {
	if m != nil {
		return m.Characterid
	}
	return 0
}

func (m *Bug) GetCharactername() string {
	if m != nil {
		return m.Charactername
	}
	return ""
}

func (m *Bug) GetReporterspoof() int64 {
	if m != nil {
		return m.Reporterspoof
	}
	return 0
}

func (m *Bug) GetCategoryid() int64 {
	if m != nil {
		return m.Categoryid
	}
	return 0
}

func (m *Bug) GetCategoryname() string {
	if m != nil {
		return m.Categoryname
	}
	return ""
}

func (m *Bug) GetReportername() string {
	if m != nil {
		return m.Reportername
	}
	return ""
}

func (m *Bug) GetUipath() string {
	if m != nil {
		return m.Uipath
	}
	return ""
}

func (m *Bug) GetPosx() int64 {
	if m != nil {
		return m.Posx
	}
	return 0
}

func (m *Bug) GetPosy() int64 {
	if m != nil {
		return m.Posy
	}
	return 0
}

func (m *Bug) GetPosz() int64 {
	if m != nil {
		return m.Posz
	}
	return 0
}

func (m *Bug) GetHeading() int64 {
	if m != nil {
		return m.Heading
	}
	return 0
}

func (m *Bug) GetTimeplayed() int64 {
	if m != nil {
		return m.Timeplayed
	}
	return 0
}

func (m *Bug) GetTargetid() int64 {
	if m != nil {
		return m.Targetid
	}
	return 0
}

func (m *Bug) GetTargetname() string {
	if m != nil {
		return m.Targetname
	}
	return ""
}

func (m *Bug) GetOptionalinfomask() int64 {
	if m != nil {
		return m.Optionalinfomask
	}
	return 0
}

func (m *Bug) GetCanduplicate() int64 {
	if m != nil {
		return m.Canduplicate
	}
	return 0
}

func (m *Bug) GetCrashbug() int64 {
	if m != nil {
		return m.Crashbug
	}
	return 0
}

func (m *Bug) GetTargetinfo() int64 {
	if m != nil {
		return m.Targetinfo
	}
	return 0
}

func (m *Bug) GetCharacterflags() int64 {
	if m != nil {
		return m.Characterflags
	}
	return 0
}

func (m *Bug) GetUnknownvalue() int64 {
	if m != nil {
		return m.Unknownvalue
	}
	return 0
}

func (m *Bug) GetBugreport() string {
	if m != nil {
		return m.Bugreport
	}
	return ""
}

func (m *Bug) GetSysteminfo() string {
	if m != nil {
		return m.Systeminfo
	}
	return ""
}

func (m *Bug) GetReportdatetime() int64 {
	if m != nil {
		return m.Reportdatetime
	}
	return 0
}

func (m *Bug) GetBugstatus() int64 {
	if m != nil {
		return m.Bugstatus
	}
	return 0
}

func (m *Bug) GetLastreview() int64 {
	if m != nil {
		return m.Lastreview
	}
	return 0
}

func (m *Bug) GetLastreviewer() string {
	if m != nil {
		return m.Lastreviewer
	}
	return ""
}

func (m *Bug) GetReviewernotes() string {
	if m != nil {
		return m.Reviewernotes
	}
	return ""
}

func init() {
	proto.RegisterType((*Bug)(nil), "pb.Bug")
}

func init() { proto.RegisterFile("proto/bug.proto", fileDescriptor_e891962b823e3f16) }

var fileDescriptor_e891962b823e3f16 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x94, 0xdd, 0x6e, 0x1a, 0x3d,
	0x10, 0x86, 0x05, 0xe4, 0x23, 0xe0, 0x24, 0x90, 0xf8, 0x6b, 0xd3, 0x29, 0xa5, 0x29, 0x8a, 0xaa,
	0x0a, 0x55, 0x55, 0x39, 0xe8, 0x1d, 0xf4, 0x12, 0x72, 0x07, 0xc3, 0xee, 0xb0, 0x58, 0x59, 0xec,
	0x95, 0x7f, 0x48, 0xe1, 0xbc, 0xf7, 0x5d, 0x79, 0x0c, 0xfb, 0x93, 0x9c, 0x79, 0x1e, 0xbf, 0x3b,
	0xf3, 0xce, 0x78, 0x40, 0x4c, 0x2b, 0x6b, 0xbc, 0x59, 0xad, 0x43, 0xf1, 0x93, 0x4f, 0xb2, 0x5f,
	0xad, 0x67, 0xf3, 0xc2, 0x98, 0xa2, 0xa4, 0x15, 0x56, 0x6a, 0x85, 0x5a, 0x1b, 0x8f, 0x5e, 0x19,
	0xed, 0x92, 0xe2, 0xf1, 0xef, 0x48, 0x0c, 0x7e, 0x87, 0x42, 0x4e, 0x44, 0x5f, 0xe5, 0xd0, 0x5b,
	0xf4, 0x96, 0x83, 0xa7, 0xbe, 0xca, 0xa5, 0x14, 0x17, 0x47, 0xa3, 0x09, 0xfa, 0x8b, 0xde, 0x72,
	0xfc, 0xc4, 0x67, 0xb9, 0x14, 0xd3, 0xac, 0x54, 0xa4, 0xfd, 0x9e, 0xac, 0x53, 0x46, 0xab, 0x1c,
	0x06, 0xfc, 0xc1, 0x6b, 0x2c, 0x7f, 0x88, 0xbb, 0x0e, 0xd2, 0xb8, 0x23, 0xb8, 0xe0, 0x54, 0x6f,
	0x2f, 0xe4, 0x5c, 0x8c, 0x31, 0xcb, 0x4c, 0xd0, 0x5e, 0xe5, 0xf0, 0x1f, 0x67, 0x6c, 0x80, 0x5c,
	0x88, 0xab, 0x6c, 0x8b, 0x16, 0x33, 0x4f, 0x56, 0xe5, 0x30, 0xe4, 0xfb, 0x36, 0x92, 0x5f, 0xc5,
	0x4d, 0x1d, 0x72, 0xa5, 0x4b, 0xae, 0xd4, 0x85, 0x51, 0x65, 0xa9, 0x32, 0xd6, 0x93, 0x75, 0x95,
	0x31, 0x1b, 0x18, 0x71, 0xa6, 0x2e, 0x94, 0x0f, 0x42, 0x64, 0xe8, 0xa9, 0x30, 0xf6, 0xa0, 0x72,
	0x18, 0xb3, 0xa4, 0x45, 0xe4, 0xa3, 0xb8, 0x3e, 0x47, 0x5c, 0x4a, 0x70, 0xa9, 0x0e, 0x8b, 0x9a,
	0x73, 0x52, 0xd6, 0x5c, 0x25, 0x4d, 0x9b, 0xc9, 0x7b, 0x31, 0x0c, 0xaa, 0x42, 0xbf, 0x85, 0x6b,
	0xbe, 0x3d, 0x45, 0x71, 0xee, 0x95, 0x71, 0x7f, 0xe0, 0x86, 0x2b, 0xf3, 0xf9, 0xc4, 0x0e, 0x30,
	0xa9, 0xd9, 0xe1, 0xc4, 0x8e, 0x30, 0xad, 0xd9, 0x51, 0x82, 0xb8, 0xdc, 0x12, 0xe6, 0x4a, 0x17,
	0x70, 0xcb, 0xf8, 0x1c, 0xc6, 0xae, 0xbc, 0xda, 0x51, 0x55, 0xe2, 0x81, 0x72, 0xb8, 0x4b, 0x5d,
	0x35, 0x44, 0xce, 0xc4, 0xc8, 0xa3, 0x2d, 0x28, 0x3e, 0x80, 0xe4, 0xdb, 0x3a, 0xe6, 0x6f, 0xf9,
	0xcc, 0xbd, 0xfc, 0xcf, 0x6e, 0x5b, 0x44, 0x7e, 0x17, 0xb7, 0xa6, 0x8a, 0x2b, 0x85, 0xa5, 0xd2,
	0x1b, 0xb3, 0x43, 0xf7, 0x0c, 0xef, 0x38, 0xc7, 0x1b, 0x9e, 0xa6, 0xa7, 0xf3, 0x50, 0x95, 0x2a,
	0x4e, 0x0c, 0xde, 0xb3, 0xae, 0xc3, 0xa2, 0x97, 0xcc, 0xa2, 0xdb, 0xae, 0x43, 0x01, 0xf7, 0xc9,
	0xcb, 0x39, 0x6e, 0xbc, 0xc4, 0x8c, 0xf0, 0xe1, 0xd4, 0x47, 0x4d, 0xe4, 0x37, 0x31, 0xa9, 0x1f,
	0x7d, 0x53, 0x62, 0xe1, 0x00, 0x58, 0xf3, 0x8a, 0x46, 0x1f, 0x41, 0x3f, 0x6b, 0xf3, 0xa2, 0xf7,
	0x58, 0x06, 0x82, 0x8f, 0xc9, 0x47, 0x9b, 0xc5, 0xad, 0x5c, 0x87, 0x22, 0x3d, 0x1a, 0xcc, 0xb8,
	0xed, 0x06, 0x44, 0x27, 0xee, 0xe0, 0x3c, 0xed, 0xd8, 0xc9, 0xa7, 0x34, 0x95, 0x86, 0x44, 0x27,
	0x49, 0x99, 0xa3, 0xa7, 0x38, 0x69, 0x98, 0x27, 0x27, 0x5d, 0x7a, 0xaa, 0xe2, 0x3c, 0xfa, 0xe0,
	0xe0, 0x73, 0xda, 0xfd, 0x1a, 0xc4, 0x2a, 0x25, 0x3a, 0x6f, 0x69, 0xaf, 0xe8, 0x05, 0x1e, 0x52,
	0xbf, 0x0d, 0x89, 0x7d, 0x34, 0x11, 0x59, 0xf8, 0x92, 0x36, 0xad, 0xcd, 0xd2, 0xde, 0xa7, 0xb3,
	0x36, 0x9e, 0x1c, 0x2c, 0xd2, 0xaf, 0xa3, 0x03, 0xd7, 0x43, 0xfe, 0x3b, 0xf8, 0xf5, 0x2f, 0x00,
	0x00, 0xff, 0xff, 0x96, 0x19, 0xb6, 0xe9, 0x43, 0x04, 0x00, 0x00,
}
