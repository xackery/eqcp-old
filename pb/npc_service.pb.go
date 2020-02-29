// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/npc_service.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type NpcSearchRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int64    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Orderby              string   `protobuf:"bytes,4,opt,name=orderby,proto3" json:"orderby,omitempty"`
	Orderdesc            bool     `protobuf:"varint,5,opt,name=orderdesc,proto3" json:"orderdesc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NpcSearchRequest) Reset()         { *m = NpcSearchRequest{} }
func (m *NpcSearchRequest) String() string { return proto.CompactTextString(m) }
func (*NpcSearchRequest) ProtoMessage()    {}
func (*NpcSearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb39d325093f256a, []int{0}
}

func (m *NpcSearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NpcSearchRequest.Unmarshal(m, b)
}
func (m *NpcSearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NpcSearchRequest.Marshal(b, m, deterministic)
}
func (m *NpcSearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NpcSearchRequest.Merge(m, src)
}
func (m *NpcSearchRequest) XXX_Size() int {
	return xxx_messageInfo_NpcSearchRequest.Size(m)
}
func (m *NpcSearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NpcSearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NpcSearchRequest proto.InternalMessageInfo

func (m *NpcSearchRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NpcSearchRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *NpcSearchRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *NpcSearchRequest) GetOrderby() string {
	if m != nil {
		return m.Orderby
	}
	return ""
}

func (m *NpcSearchRequest) GetOrderdesc() bool {
	if m != nil {
		return m.Orderdesc
	}
	return false
}

type NpcSearchResponse struct {
	Npcs                 []*Npc   `protobuf:"bytes,1,rep,name=Npcs,proto3" json:"Npcs,omitempty"`
	Total                int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NpcSearchResponse) Reset()         { *m = NpcSearchResponse{} }
func (m *NpcSearchResponse) String() string { return proto.CompactTextString(m) }
func (*NpcSearchResponse) ProtoMessage()    {}
func (*NpcSearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb39d325093f256a, []int{1}
}

func (m *NpcSearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NpcSearchResponse.Unmarshal(m, b)
}
func (m *NpcSearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NpcSearchResponse.Marshal(b, m, deterministic)
}
func (m *NpcSearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NpcSearchResponse.Merge(m, src)
}
func (m *NpcSearchResponse) XXX_Size() int {
	return xxx_messageInfo_NpcSearchResponse.Size(m)
}
func (m *NpcSearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NpcSearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NpcSearchResponse proto.InternalMessageInfo

func (m *NpcSearchResponse) GetNpcs() []*Npc {
	if m != nil {
		return m.Npcs
	}
	return nil
}

func (m *NpcSearchResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type NpcCreateRequest struct {
	Values               map[string]string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NpcCreateRequest) Reset()         { *m = NpcCreateRequest{} }
func (m *NpcCreateRequest) String() string { return proto.CompactTextString(m) }
func (*NpcCreateRequest) ProtoMessage()    {}
func (*NpcCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb39d325093f256a, []int{2}
}

func (m *NpcCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NpcCreateRequest.Unmarshal(m, b)
}
func (m *NpcCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NpcCreateRequest.Marshal(b, m, deterministic)
}
func (m *NpcCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NpcCreateRequest.Merge(m, src)
}
func (m *NpcCreateRequest) XXX_Size() int {
	return xxx_messageInfo_NpcCreateRequest.Size(m)
}
func (m *NpcCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NpcCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NpcCreateRequest proto.InternalMessageInfo

func (m *NpcCreateRequest) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

type NpcCreateResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NpcCreateResponse) Reset()         { *m = NpcCreateResponse{} }
func (m *NpcCreateResponse) String() string { return proto.CompactTextString(m) }
func (*NpcCreateResponse) ProtoMessage()    {}
func (*NpcCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb39d325093f256a, []int{3}
}

func (m *NpcCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NpcCreateResponse.Unmarshal(m, b)
}
func (m *NpcCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NpcCreateResponse.Marshal(b, m, deterministic)
}
func (m *NpcCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NpcCreateResponse.Merge(m, src)
}
func (m *NpcCreateResponse) XXX_Size() int {
	return xxx_messageInfo_NpcCreateResponse.Size(m)
}
func (m *NpcCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NpcCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NpcCreateResponse proto.InternalMessageInfo

func (m *NpcCreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type NpcReadRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NpcReadRequest) Reset()         { *m = NpcReadRequest{} }
func (m *NpcReadRequest) String() string { return proto.CompactTextString(m) }
func (*NpcReadRequest) ProtoMessage()    {}
func (*NpcReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb39d325093f256a, []int{4}
}

func (m *NpcReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NpcReadRequest.Unmarshal(m, b)
}
func (m *NpcReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NpcReadRequest.Marshal(b, m, deterministic)
}
func (m *NpcReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NpcReadRequest.Merge(m, src)
}
func (m *NpcReadRequest) XXX_Size() int {
	return xxx_messageInfo_NpcReadRequest.Size(m)
}
func (m *NpcReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NpcReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NpcReadRequest proto.InternalMessageInfo

func (m *NpcReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type NpcReadResponse struct {
	Npc                  *Npc     `protobuf:"bytes,1,opt,name=npc,proto3" json:"npc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NpcReadResponse) Reset()         { *m = NpcReadResponse{} }
func (m *NpcReadResponse) String() string { return proto.CompactTextString(m) }
func (*NpcReadResponse) ProtoMessage()    {}
func (*NpcReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb39d325093f256a, []int{5}
}

func (m *NpcReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NpcReadResponse.Unmarshal(m, b)
}
func (m *NpcReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NpcReadResponse.Marshal(b, m, deterministic)
}
func (m *NpcReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NpcReadResponse.Merge(m, src)
}
func (m *NpcReadResponse) XXX_Size() int {
	return xxx_messageInfo_NpcReadResponse.Size(m)
}
func (m *NpcReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NpcReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NpcReadResponse proto.InternalMessageInfo

func (m *NpcReadResponse) GetNpc() *Npc {
	if m != nil {
		return m.Npc
	}
	return nil
}

type NpcUpdateRequest struct {
	Id                   int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Values               map[string]string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NpcUpdateRequest) Reset()         { *m = NpcUpdateRequest{} }
func (m *NpcUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*NpcUpdateRequest) ProtoMessage()    {}
func (*NpcUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb39d325093f256a, []int{6}
}

func (m *NpcUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NpcUpdateRequest.Unmarshal(m, b)
}
func (m *NpcUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NpcUpdateRequest.Marshal(b, m, deterministic)
}
func (m *NpcUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NpcUpdateRequest.Merge(m, src)
}
func (m *NpcUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_NpcUpdateRequest.Size(m)
}
func (m *NpcUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NpcUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NpcUpdateRequest proto.InternalMessageInfo

func (m *NpcUpdateRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NpcUpdateRequest) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

type NpcUpdateResponse struct {
	Rowsaffected         int64    `protobuf:"varint,1,opt,name=rowsaffected,proto3" json:"rowsaffected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NpcUpdateResponse) Reset()         { *m = NpcUpdateResponse{} }
func (m *NpcUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*NpcUpdateResponse) ProtoMessage()    {}
func (*NpcUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb39d325093f256a, []int{7}
}

func (m *NpcUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NpcUpdateResponse.Unmarshal(m, b)
}
func (m *NpcUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NpcUpdateResponse.Marshal(b, m, deterministic)
}
func (m *NpcUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NpcUpdateResponse.Merge(m, src)
}
func (m *NpcUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_NpcUpdateResponse.Size(m)
}
func (m *NpcUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NpcUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NpcUpdateResponse proto.InternalMessageInfo

func (m *NpcUpdateResponse) GetRowsaffected() int64 {
	if m != nil {
		return m.Rowsaffected
	}
	return 0
}

type NpcDeleteRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NpcDeleteRequest) Reset()         { *m = NpcDeleteRequest{} }
func (m *NpcDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*NpcDeleteRequest) ProtoMessage()    {}
func (*NpcDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb39d325093f256a, []int{8}
}

func (m *NpcDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NpcDeleteRequest.Unmarshal(m, b)
}
func (m *NpcDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NpcDeleteRequest.Marshal(b, m, deterministic)
}
func (m *NpcDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NpcDeleteRequest.Merge(m, src)
}
func (m *NpcDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_NpcDeleteRequest.Size(m)
}
func (m *NpcDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NpcDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NpcDeleteRequest proto.InternalMessageInfo

func (m *NpcDeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type NpcDeleteResponse struct {
	Rowsaffected         int64    `protobuf:"varint,1,opt,name=rowsaffected,proto3" json:"rowsaffected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NpcDeleteResponse) Reset()         { *m = NpcDeleteResponse{} }
func (m *NpcDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*NpcDeleteResponse) ProtoMessage()    {}
func (*NpcDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb39d325093f256a, []int{9}
}

func (m *NpcDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NpcDeleteResponse.Unmarshal(m, b)
}
func (m *NpcDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NpcDeleteResponse.Marshal(b, m, deterministic)
}
func (m *NpcDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NpcDeleteResponse.Merge(m, src)
}
func (m *NpcDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_NpcDeleteResponse.Size(m)
}
func (m *NpcDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NpcDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NpcDeleteResponse proto.InternalMessageInfo

func (m *NpcDeleteResponse) GetRowsaffected() int64 {
	if m != nil {
		return m.Rowsaffected
	}
	return 0
}

type NpcPatchRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NpcPatchRequest) Reset()         { *m = NpcPatchRequest{} }
func (m *NpcPatchRequest) String() string { return proto.CompactTextString(m) }
func (*NpcPatchRequest) ProtoMessage()    {}
func (*NpcPatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb39d325093f256a, []int{10}
}

func (m *NpcPatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NpcPatchRequest.Unmarshal(m, b)
}
func (m *NpcPatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NpcPatchRequest.Marshal(b, m, deterministic)
}
func (m *NpcPatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NpcPatchRequest.Merge(m, src)
}
func (m *NpcPatchRequest) XXX_Size() int {
	return xxx_messageInfo_NpcPatchRequest.Size(m)
}
func (m *NpcPatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NpcPatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NpcPatchRequest proto.InternalMessageInfo

func (m *NpcPatchRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NpcPatchRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *NpcPatchRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type NpcPatchResponse struct {
	Rowsaffected         int64    `protobuf:"varint,1,opt,name=rowsaffected,proto3" json:"rowsaffected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NpcPatchResponse) Reset()         { *m = NpcPatchResponse{} }
func (m *NpcPatchResponse) String() string { return proto.CompactTextString(m) }
func (*NpcPatchResponse) ProtoMessage()    {}
func (*NpcPatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb39d325093f256a, []int{11}
}

func (m *NpcPatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NpcPatchResponse.Unmarshal(m, b)
}
func (m *NpcPatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NpcPatchResponse.Marshal(b, m, deterministic)
}
func (m *NpcPatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NpcPatchResponse.Merge(m, src)
}
func (m *NpcPatchResponse) XXX_Size() int {
	return xxx_messageInfo_NpcPatchResponse.Size(m)
}
func (m *NpcPatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NpcPatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NpcPatchResponse proto.InternalMessageInfo

func (m *NpcPatchResponse) GetRowsaffected() int64 {
	if m != nil {
		return m.Rowsaffected
	}
	return 0
}

func init() {
	proto.RegisterType((*NpcSearchRequest)(nil), "pb.NpcSearchRequest")
	proto.RegisterType((*NpcSearchResponse)(nil), "pb.NpcSearchResponse")
	proto.RegisterType((*NpcCreateRequest)(nil), "pb.NpcCreateRequest")
	proto.RegisterMapType((map[string]string)(nil), "pb.NpcCreateRequest.ValuesEntry")
	proto.RegisterType((*NpcCreateResponse)(nil), "pb.NpcCreateResponse")
	proto.RegisterType((*NpcReadRequest)(nil), "pb.NpcReadRequest")
	proto.RegisterType((*NpcReadResponse)(nil), "pb.NpcReadResponse")
	proto.RegisterType((*NpcUpdateRequest)(nil), "pb.NpcUpdateRequest")
	proto.RegisterMapType((map[string]string)(nil), "pb.NpcUpdateRequest.ValuesEntry")
	proto.RegisterType((*NpcUpdateResponse)(nil), "pb.NpcUpdateResponse")
	proto.RegisterType((*NpcDeleteRequest)(nil), "pb.NpcDeleteRequest")
	proto.RegisterType((*NpcDeleteResponse)(nil), "pb.NpcDeleteResponse")
	proto.RegisterType((*NpcPatchRequest)(nil), "pb.NpcPatchRequest")
	proto.RegisterType((*NpcPatchResponse)(nil), "pb.NpcPatchResponse")
}

func init() { proto.RegisterFile("proto/npc_service.proto", fileDescriptor_bb39d325093f256a) }

var fileDescriptor_bb39d325093f256a = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x95, 0xed, 0x34, 0x69, 0x6e, 0xa3, 0xb4, 0x9d, 0x2f, 0xfd, 0x62, 0x4c, 0x91, 0xac, 0x61,
	0x13, 0x55, 0x28, 0x16, 0x41, 0x82, 0xd2, 0x2d, 0x3f, 0x02, 0x09, 0x45, 0x60, 0x04, 0x2c, 0x91,
	0x63, 0x4f, 0x8a, 0x45, 0x6a, 0x0f, 0x9e, 0x69, 0x50, 0x54, 0x75, 0xc3, 0x86, 0x0d, 0x3b, 0x16,
	0x3c, 0x18, 0xaf, 0xc0, 0x03, 0xf0, 0x08, 0xc8, 0xd7, 0x33, 0x8e, 0x27, 0xa2, 0x12, 0x48, 0xec,
	0xe6, 0x9e, 0xdc, 0x39, 0x73, 0xce, 0xf5, 0xb9, 0x81, 0x21, 0x2f, 0x72, 0x99, 0x07, 0x19, 0x8f,
	0xdf, 0x0a, 0x56, 0x2c, 0xd3, 0x98, 0x8d, 0x11, 0x21, 0x36, 0x9f, 0x79, 0x87, 0xa7, 0x79, 0x7e,
	0xba, 0x60, 0x41, 0xc4, 0xd3, 0x20, 0xca, 0xb2, 0x5c, 0x46, 0x32, 0xcd, 0x33, 0x51, 0x75, 0x78,
	0xbb, 0xf5, 0xd5, 0x0a, 0xa0, 0x5f, 0x2c, 0xd8, 0x9b, 0xf2, 0xf8, 0x25, 0x8b, 0x8a, 0xf8, 0x5d,
	0xc8, 0x3e, 0x9c, 0x33, 0x21, 0x09, 0x81, 0x56, 0x16, 0x9d, 0x31, 0xd7, 0xf2, 0xad, 0x51, 0x37,
	0xc4, 0x33, 0x19, 0xc0, 0xd6, 0x22, 0x3d, 0x4b, 0xa5, 0x6b, 0xfb, 0xd6, 0xc8, 0x09, 0xab, 0x82,
	0xfc, 0x0f, 0xed, 0x7c, 0x3e, 0x17, 0x4c, 0xba, 0x0e, 0xc2, 0xaa, 0x22, 0x2e, 0x74, 0xf2, 0x22,
	0x61, 0xc5, 0x6c, 0xe5, 0xb6, 0x90, 0x44, 0x97, 0xe4, 0x10, 0xba, 0x78, 0x4c, 0x98, 0x88, 0xdd,
	0x2d, 0xdf, 0x1a, 0x6d, 0x87, 0x6b, 0x80, 0x3e, 0x86, 0xfd, 0x86, 0x1a, 0xc1, 0xf3, 0x4c, 0x30,
	0x72, 0x1d, 0x5a, 0x53, 0x1e, 0x0b, 0xd7, 0xf2, 0x9d, 0xd1, 0xce, 0xa4, 0x33, 0xe6, 0xb3, 0xf1,
	0x94, 0xc7, 0x21, 0x82, 0xa5, 0x2e, 0x99, 0xcb, 0x68, 0xa1, 0x75, 0x61, 0x41, 0x3f, 0x57, 0xb6,
	0x1e, 0x14, 0x2c, 0x92, 0x4c, 0xdb, 0x3a, 0x86, 0xf6, 0x32, 0x5a, 0x9c, 0x33, 0xe1, 0xda, 0xc8,
	0xe4, 0x2b, 0x26, 0xa3, 0x6b, 0xfc, 0x1a, 0x5b, 0x1e, 0x65, 0xb2, 0x58, 0x85, 0xaa, 0xdf, 0xbb,
	0x0f, 0x3b, 0x0d, 0x98, 0xec, 0x81, 0xf3, 0x9e, 0xad, 0xd4, 0x78, 0xca, 0x63, 0xa9, 0x02, 0x5b,
	0x51, 0x45, 0x37, 0xac, 0x8a, 0x13, 0xfb, 0xd8, 0xa2, 0x37, 0xd1, 0x91, 0x7e, 0x42, 0x39, 0xea,
	0x83, 0x9d, 0x26, 0x78, 0xdf, 0x09, 0xed, 0x34, 0xa1, 0x3e, 0xf4, 0x4b, 0x47, 0x2c, 0x4a, 0xb4,
	0xd6, 0xcd, 0x8e, 0x5b, 0xb0, 0x5b, 0x77, 0x28, 0x92, 0x6b, 0xe0, 0x64, 0x3c, 0xc6, 0x9e, 0xc6,
	0x54, 0x4a, 0x8c, 0x7e, 0xab, 0xec, 0xbf, 0xe2, 0x49, 0xc3, 0xfe, 0x06, 0xe5, 0x95, 0xe3, 0x30,
	0x6e, 0xfd, 0xeb, 0x71, 0xdc, 0xc3, 0x71, 0xe8, 0x27, 0x94, 0x13, 0x0a, 0xbd, 0x22, 0xff, 0x28,
	0xa2, 0xf9, 0x9c, 0xc5, 0x92, 0x69, 0x8d, 0x06, 0x46, 0x29, 0x3a, 0x7a, 0xc8, 0x16, 0xec, 0x4a,
	0x47, 0x8a, 0x5c, 0xf7, 0xfc, 0x05, 0xf9, 0x53, 0x9c, 0xee, 0xf3, 0x48, 0xae, 0x77, 0x60, 0x73,
	0x5a, 0xca, 0xa4, 0xfd, 0x1b, 0x93, 0x4e, 0xc3, 0x24, 0xbd, 0x8b, 0x3a, 0x15, 0xd5, 0x9f, 0x4b,
	0x98, 0xfc, 0x74, 0x00, 0x30, 0xfa, 0xb8, 0xd0, 0xe4, 0x0d, 0x74, 0xeb, 0x45, 0x20, 0x03, 0xf5,
	0x65, 0x8c, 0x2d, 0xf5, 0x0e, 0x36, 0xd0, 0xea, 0x31, 0x7a, 0xe3, 0xd3, 0xf7, 0x1f, 0x5f, 0xed,
	0x21, 0x39, 0x08, 0x96, 0xb7, 0xcb, 0x45, 0x0f, 0x04, 0xfe, 0x1e, 0x5c, 0x94, 0x6b, 0x7c, 0x49,
	0x9e, 0x21, 0x71, 0x95, 0xc7, 0x9a, 0xd8, 0xd8, 0x80, 0x9a, 0xd8, 0x0c, 0x2d, 0x25, 0x48, 0xdc,
	0xa3, 0x1d, 0x45, 0x7c, 0x62, 0x1d, 0x91, 0x27, 0xd0, 0x51, 0xb1, 0x24, 0x44, 0x27, 0x70, 0x9d,
	0x62, 0xef, 0x3f, 0x03, 0x53, 0x3c, 0x03, 0xe4, 0xe9, 0x93, 0x9e, 0x16, 0x78, 0x91, 0x26, 0x97,
	0xe4, 0x05, 0xea, 0xaa, 0x82, 0x51, 0xeb, 0x32, 0xa2, 0x58, 0xeb, 0x32, 0xd3, 0x43, 0x87, 0xc8,
	0xb7, 0xef, 0x19, 0x7c, 0xa5, 0xb8, 0x29, 0x52, 0x56, 0x71, 0xa8, 0x29, 0x8d, 0x04, 0xd5, 0x94,
	0x66, 0x66, 0xb4, 0xc4, 0x23, 0x53, 0xe2, 0x14, 0xb6, 0xf5, 0xa7, 0x25, 0xda, 0x59, 0x33, 0x33,
	0xde, 0xc0, 0x04, 0x4d, 0x7d, 0x93, 0x4d, 0x7d, 0xb3, 0x36, 0xfe, 0x05, 0xdf, 0xf9, 0x15, 0x00,
	0x00, 0xff, 0xff, 0x41, 0x5e, 0xbf, 0xe9, 0xd0, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NpcServiceClient is the client API for NpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NpcServiceClient interface {
	NpcSearch(ctx context.Context, in *NpcSearchRequest, opts ...grpc.CallOption) (*NpcSearchResponse, error)
	NpcCreate(ctx context.Context, in *NpcCreateRequest, opts ...grpc.CallOption) (*NpcCreateResponse, error)
	NpcRead(ctx context.Context, in *NpcReadRequest, opts ...grpc.CallOption) (*NpcReadResponse, error)
	NpcUpdate(ctx context.Context, in *NpcUpdateRequest, opts ...grpc.CallOption) (*NpcUpdateResponse, error)
	NpcDelete(ctx context.Context, in *NpcDeleteRequest, opts ...grpc.CallOption) (*NpcDeleteResponse, error)
	NpcPatch(ctx context.Context, in *NpcPatchRequest, opts ...grpc.CallOption) (*NpcPatchResponse, error)
}

type npcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNpcServiceClient(cc grpc.ClientConnInterface) NpcServiceClient {
	return &npcServiceClient{cc}
}

func (c *npcServiceClient) NpcSearch(ctx context.Context, in *NpcSearchRequest, opts ...grpc.CallOption) (*NpcSearchResponse, error) {
	out := new(NpcSearchResponse)
	err := c.cc.Invoke(ctx, "/pb.NpcService/NpcSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *npcServiceClient) NpcCreate(ctx context.Context, in *NpcCreateRequest, opts ...grpc.CallOption) (*NpcCreateResponse, error) {
	out := new(NpcCreateResponse)
	err := c.cc.Invoke(ctx, "/pb.NpcService/NpcCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *npcServiceClient) NpcRead(ctx context.Context, in *NpcReadRequest, opts ...grpc.CallOption) (*NpcReadResponse, error) {
	out := new(NpcReadResponse)
	err := c.cc.Invoke(ctx, "/pb.NpcService/NpcRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *npcServiceClient) NpcUpdate(ctx context.Context, in *NpcUpdateRequest, opts ...grpc.CallOption) (*NpcUpdateResponse, error) {
	out := new(NpcUpdateResponse)
	err := c.cc.Invoke(ctx, "/pb.NpcService/NpcUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *npcServiceClient) NpcDelete(ctx context.Context, in *NpcDeleteRequest, opts ...grpc.CallOption) (*NpcDeleteResponse, error) {
	out := new(NpcDeleteResponse)
	err := c.cc.Invoke(ctx, "/pb.NpcService/NpcDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *npcServiceClient) NpcPatch(ctx context.Context, in *NpcPatchRequest, opts ...grpc.CallOption) (*NpcPatchResponse, error) {
	out := new(NpcPatchResponse)
	err := c.cc.Invoke(ctx, "/pb.NpcService/NpcPatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NpcServiceServer is the server API for NpcService service.
type NpcServiceServer interface {
	NpcSearch(context.Context, *NpcSearchRequest) (*NpcSearchResponse, error)
	NpcCreate(context.Context, *NpcCreateRequest) (*NpcCreateResponse, error)
	NpcRead(context.Context, *NpcReadRequest) (*NpcReadResponse, error)
	NpcUpdate(context.Context, *NpcUpdateRequest) (*NpcUpdateResponse, error)
	NpcDelete(context.Context, *NpcDeleteRequest) (*NpcDeleteResponse, error)
	NpcPatch(context.Context, *NpcPatchRequest) (*NpcPatchResponse, error)
}

// UnimplementedNpcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNpcServiceServer struct {
}

func (*UnimplementedNpcServiceServer) NpcSearch(ctx context.Context, req *NpcSearchRequest) (*NpcSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NpcSearch not implemented")
}
func (*UnimplementedNpcServiceServer) NpcCreate(ctx context.Context, req *NpcCreateRequest) (*NpcCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NpcCreate not implemented")
}
func (*UnimplementedNpcServiceServer) NpcRead(ctx context.Context, req *NpcReadRequest) (*NpcReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NpcRead not implemented")
}
func (*UnimplementedNpcServiceServer) NpcUpdate(ctx context.Context, req *NpcUpdateRequest) (*NpcUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NpcUpdate not implemented")
}
func (*UnimplementedNpcServiceServer) NpcDelete(ctx context.Context, req *NpcDeleteRequest) (*NpcDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NpcDelete not implemented")
}
func (*UnimplementedNpcServiceServer) NpcPatch(ctx context.Context, req *NpcPatchRequest) (*NpcPatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NpcPatch not implemented")
}

func RegisterNpcServiceServer(s *grpc.Server, srv NpcServiceServer) {
	s.RegisterService(&_NpcService_serviceDesc, srv)
}

func _NpcService_NpcSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NpcSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NpcServiceServer).NpcSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NpcService/NpcSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NpcServiceServer).NpcSearch(ctx, req.(*NpcSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NpcService_NpcCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NpcCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NpcServiceServer).NpcCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NpcService/NpcCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NpcServiceServer).NpcCreate(ctx, req.(*NpcCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NpcService_NpcRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NpcReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NpcServiceServer).NpcRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NpcService/NpcRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NpcServiceServer).NpcRead(ctx, req.(*NpcReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NpcService_NpcUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NpcUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NpcServiceServer).NpcUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NpcService/NpcUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NpcServiceServer).NpcUpdate(ctx, req.(*NpcUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NpcService_NpcDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NpcDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NpcServiceServer).NpcDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NpcService/NpcDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NpcServiceServer).NpcDelete(ctx, req.(*NpcDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NpcService_NpcPatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NpcPatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NpcServiceServer).NpcPatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.NpcService/NpcPatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NpcServiceServer).NpcPatch(ctx, req.(*NpcPatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NpcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.NpcService",
	HandlerType: (*NpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NpcSearch",
			Handler:    _NpcService_NpcSearch_Handler,
		},
		{
			MethodName: "NpcCreate",
			Handler:    _NpcService_NpcCreate_Handler,
		},
		{
			MethodName: "NpcRead",
			Handler:    _NpcService_NpcRead_Handler,
		},
		{
			MethodName: "NpcUpdate",
			Handler:    _NpcService_NpcUpdate_Handler,
		},
		{
			MethodName: "NpcDelete",
			Handler:    _NpcService_NpcDelete_Handler,
		},
		{
			MethodName: "NpcPatch",
			Handler:    _NpcService_NpcPatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/npc_service.proto",
}
