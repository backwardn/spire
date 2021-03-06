// Code generated by protoc-gen-go. DO NOT EDIT.
// source: private/server/journal/journal.proto

package journal

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

type X509CAEntry struct {
	// Which X509 CA slot this entry occupied.
	SlotId string `protobuf:"bytes,1,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
	// When the CA was issued (unix epoch in seconds)
	IssuedAt int64 `protobuf:"varint,2,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	// DER encoded CA certificate
	Certificate []byte `protobuf:"bytes,3,opt,name=certificate,proto3" json:"certificate,omitempty"`
	// DER encoded upstream CA chain. See the X509CA struct for details.
	UpstreamChain        [][]byte `protobuf:"bytes,4,rep,name=upstream_chain,json=upstreamChain,proto3" json:"upstream_chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *X509CAEntry) Reset()         { *m = X509CAEntry{} }
func (m *X509CAEntry) String() string { return proto.CompactTextString(m) }
func (*X509CAEntry) ProtoMessage()    {}
func (*X509CAEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_63c6786ba201045d, []int{0}
}

func (m *X509CAEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_X509CAEntry.Unmarshal(m, b)
}
func (m *X509CAEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_X509CAEntry.Marshal(b, m, deterministic)
}
func (m *X509CAEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_X509CAEntry.Merge(m, src)
}
func (m *X509CAEntry) XXX_Size() int {
	return xxx_messageInfo_X509CAEntry.Size(m)
}
func (m *X509CAEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_X509CAEntry.DiscardUnknown(m)
}

var xxx_messageInfo_X509CAEntry proto.InternalMessageInfo

func (m *X509CAEntry) GetSlotId() string {
	if m != nil {
		return m.SlotId
	}
	return ""
}

func (m *X509CAEntry) GetIssuedAt() int64 {
	if m != nil {
		return m.IssuedAt
	}
	return 0
}

func (m *X509CAEntry) GetCertificate() []byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *X509CAEntry) GetUpstreamChain() [][]byte {
	if m != nil {
		return m.UpstreamChain
	}
	return nil
}

type JWTKeyEntry struct {
	// Which JWT Key slot this entry occupied.
	SlotId string `protobuf:"bytes,1,opt,name=slot_id,json=slotId,proto3" json:"slot_id,omitempty"`
	// When the key was issued (unix epoch in seconds)
	IssuedAt int64 `protobuf:"varint,2,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	// When the key expires unix epoch in seconds)
	NotAfter int64 `protobuf:"varint,3,opt,name=not_after,json=notAfter,proto3" json:"not_after,omitempty"`
	// JWT key id (i.e. "kid" claim)
	Kid string `protobuf:"bytes,4,opt,name=kid,proto3" json:"kid,omitempty"`
	// PKIX encoded public key
	PublicKey            []byte   `protobuf:"bytes,5,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWTKeyEntry) Reset()         { *m = JWTKeyEntry{} }
func (m *JWTKeyEntry) String() string { return proto.CompactTextString(m) }
func (*JWTKeyEntry) ProtoMessage()    {}
func (*JWTKeyEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_63c6786ba201045d, []int{1}
}

func (m *JWTKeyEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JWTKeyEntry.Unmarshal(m, b)
}
func (m *JWTKeyEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JWTKeyEntry.Marshal(b, m, deterministic)
}
func (m *JWTKeyEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWTKeyEntry.Merge(m, src)
}
func (m *JWTKeyEntry) XXX_Size() int {
	return xxx_messageInfo_JWTKeyEntry.Size(m)
}
func (m *JWTKeyEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_JWTKeyEntry.DiscardUnknown(m)
}

var xxx_messageInfo_JWTKeyEntry proto.InternalMessageInfo

func (m *JWTKeyEntry) GetSlotId() string {
	if m != nil {
		return m.SlotId
	}
	return ""
}

func (m *JWTKeyEntry) GetIssuedAt() int64 {
	if m != nil {
		return m.IssuedAt
	}
	return 0
}

func (m *JWTKeyEntry) GetNotAfter() int64 {
	if m != nil {
		return m.NotAfter
	}
	return 0
}

func (m *JWTKeyEntry) GetKid() string {
	if m != nil {
		return m.Kid
	}
	return ""
}

func (m *JWTKeyEntry) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type Entries struct {
	X509CAs              []*X509CAEntry `protobuf:"bytes,1,rep,name=x509CAs,proto3" json:"x509CAs,omitempty"`
	JwtKeys              []*JWTKeyEntry `protobuf:"bytes,2,rep,name=jwtKeys,proto3" json:"jwtKeys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Entries) Reset()         { *m = Entries{} }
func (m *Entries) String() string { return proto.CompactTextString(m) }
func (*Entries) ProtoMessage()    {}
func (*Entries) Descriptor() ([]byte, []int) {
	return fileDescriptor_63c6786ba201045d, []int{2}
}

func (m *Entries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entries.Unmarshal(m, b)
}
func (m *Entries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entries.Marshal(b, m, deterministic)
}
func (m *Entries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entries.Merge(m, src)
}
func (m *Entries) XXX_Size() int {
	return xxx_messageInfo_Entries.Size(m)
}
func (m *Entries) XXX_DiscardUnknown() {
	xxx_messageInfo_Entries.DiscardUnknown(m)
}

var xxx_messageInfo_Entries proto.InternalMessageInfo

func (m *Entries) GetX509CAs() []*X509CAEntry {
	if m != nil {
		return m.X509CAs
	}
	return nil
}

func (m *Entries) GetJwtKeys() []*JWTKeyEntry {
	if m != nil {
		return m.JwtKeys
	}
	return nil
}

func init() {
	proto.RegisterType((*X509CAEntry)(nil), "X509CAEntry")
	proto.RegisterType((*JWTKeyEntry)(nil), "JWTKeyEntry")
	proto.RegisterType((*Entries)(nil), "Entries")
}

func init() {
	proto.RegisterFile("private/server/journal/journal.proto", fileDescriptor_63c6786ba201045d)
}

var fileDescriptor_63c6786ba201045d = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xcb, 0x4b, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0x53, 0xdb, 0x66, 0x53, 0x45, 0xf6, 0xe2, 0x42, 0x11, 0x42, 0x51, 0xc9, 0xa9,
	0x11, 0x5f, 0xe0, 0xb1, 0x16, 0x0f, 0xda, 0xdb, 0x22, 0xf8, 0xb8, 0x84, 0x6d, 0x32, 0xb1, 0xdb,
	0x47, 0x36, 0xec, 0x4e, 0xaa, 0xf9, 0x1f, 0xbc, 0xf8, 0x1f, 0xcb, 0xa6, 0x06, 0x7a, 0xf0, 0xe6,
	0x69, 0x76, 0x7e, 0xdf, 0xb0, 0x7c, 0x33, 0x1f, 0x39, 0x29, 0xb4, 0xdc, 0x08, 0x84, 0xc8, 0x80,
	0xde, 0x80, 0x8e, 0x16, 0xaa, 0xd4, 0xb9, 0x58, 0x35, 0x75, 0x54, 0x68, 0x85, 0x6a, 0xf8, 0xe5,
	0x10, 0xff, 0xe5, 0xfa, 0xfc, 0x76, 0x32, 0xbe, 0xcf, 0x51, 0x57, 0xf4, 0x88, 0x74, 0xcd, 0x4a,
	0x61, 0x2c, 0x53, 0xe6, 0x04, 0x4e, 0xe8, 0xf1, 0x8e, 0x6d, 0x1f, 0x52, 0x3a, 0x20, 0x9e, 0x34,
	0xa6, 0x84, 0x34, 0x16, 0xc8, 0x5a, 0x81, 0x13, 0xba, 0xbc, 0xb7, 0x05, 0x63, 0xa4, 0x01, 0xf1,
	0x13, 0xd0, 0x28, 0x33, 0x99, 0x08, 0x04, 0xe6, 0x06, 0x4e, 0xd8, 0xe7, 0xbb, 0x88, 0x9e, 0x92,
	0x83, 0xb2, 0x30, 0xa8, 0x41, 0xac, 0xe3, 0x64, 0x2e, 0x64, 0xce, 0xda, 0x81, 0x1b, 0xf6, 0xf9,
	0x7e, 0x43, 0x27, 0x16, 0x0e, 0xbf, 0x1d, 0xe2, 0x3f, 0x3e, 0x3f, 0x4d, 0xa1, 0xfa, 0x8f, 0x9d,
	0x01, 0xf1, 0x72, 0x85, 0xb1, 0xc8, 0x10, 0x74, 0x6d, 0xc6, 0xe5, 0xbd, 0x5c, 0xe1, 0xd8, 0xf6,
	0xf4, 0x90, 0xb8, 0x4b, 0x99, 0xb2, 0x76, 0xfd, 0x9d, 0x7d, 0xd2, 0x63, 0x42, 0x8a, 0x72, 0xb6,
	0x92, 0x49, 0xbc, 0x84, 0x8a, 0xed, 0xd5, 0xe6, 0xbd, 0x2d, 0x99, 0x42, 0x35, 0x7c, 0x25, 0x5d,
	0x6b, 0x46, 0x82, 0xa1, 0x67, 0xa4, 0xfb, 0x59, 0x1f, 0xcb, 0x30, 0x27, 0x70, 0x43, 0xff, 0xa2,
	0x3f, 0xda, 0x39, 0x1e, 0x6f, 0x44, 0x3b, 0xb7, 0xf8, 0xc0, 0x29, 0x54, 0x86, 0xb5, 0x7e, 0xe7,
	0x76, 0xb6, 0xe2, 0x8d, 0x78, 0x77, 0xf3, 0x76, 0xf5, 0x2e, 0x71, 0x5e, 0xce, 0x46, 0x89, 0x5a,
	0x47, 0xa6, 0x90, 0x59, 0x06, 0xb6, 0x68, 0x88, 0xea, 0x78, 0xa2, 0xbf, 0x33, 0x9c, 0x75, 0x6a,
	0xf5, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x9c, 0xff, 0x3d, 0xe4, 0x01, 0x00, 0x00,
}
