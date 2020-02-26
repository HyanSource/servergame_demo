// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

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

//給玩家的資訊
type PlayerData struct {
	PlayerID             int32    `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	PlayerName           string   `protobuf:"bytes,2,opt,name=PlayerName,proto3" json:"PlayerName,omitempty"`
	PlayerMoney          int32    `protobuf:"varint,3,opt,name=PlayerMoney,proto3" json:"PlayerMoney,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerData) Reset()         { *m = PlayerData{} }
func (m *PlayerData) String() string { return proto.CompactTextString(m) }
func (*PlayerData) ProtoMessage()    {}
func (*PlayerData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *PlayerData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerData.Unmarshal(m, b)
}
func (m *PlayerData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerData.Marshal(b, m, deterministic)
}
func (m *PlayerData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerData.Merge(m, src)
}
func (m *PlayerData) XXX_Size() int {
	return xxx_messageInfo_PlayerData.Size(m)
}
func (m *PlayerData) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerData.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerData proto.InternalMessageInfo

func (m *PlayerData) GetPlayerID() int32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *PlayerData) GetPlayerName() string {
	if m != nil {
		return m.PlayerName
	}
	return ""
}

func (m *PlayerData) GetPlayerMoney() int32 {
	if m != nil {
		return m.PlayerMoney
	}
	return 0
}

//給伺服器下注資訊
type Bet struct {
	PlayerID             int32    `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	Bet                  int32    `protobuf:"varint,2,opt,name=Bet,proto3" json:"Bet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bet) Reset()         { *m = Bet{} }
func (m *Bet) String() string { return proto.CompactTextString(m) }
func (*Bet) ProtoMessage()    {}
func (*Bet) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *Bet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bet.Unmarshal(m, b)
}
func (m *Bet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bet.Marshal(b, m, deterministic)
}
func (m *Bet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bet.Merge(m, src)
}
func (m *Bet) XXX_Size() int {
	return xxx_messageInfo_Bet.Size(m)
}
func (m *Bet) XXX_DiscardUnknown() {
	xxx_messageInfo_Bet.DiscardUnknown(m)
}

var xxx_messageInfo_Bet proto.InternalMessageInfo

func (m *Bet) GetPlayerID() int32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *Bet) GetBet() int32 {
	if m != nil {
		return m.Bet
	}
	return 0
}

func init() {
	proto.RegisterType((*PlayerData)(nil), "pb.PlayerData")
	proto.RegisterType((*Bet)(nil), "pb.Bet")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xca, 0xe2, 0xe2, 0x0a, 0xc8, 0x49,
	0xac, 0x4c, 0x2d, 0x72, 0x49, 0x2c, 0x49, 0x14, 0x92, 0xe2, 0xe2, 0x80, 0xf0, 0x3c, 0x5d, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0xe0, 0x7c, 0x21, 0x39, 0x98, 0x4a, 0xbf, 0xc4, 0xdc, 0x54,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x24, 0x11, 0x21, 0x05, 0x2e, 0x6e, 0x08, 0xcf, 0x37,
	0x3f, 0x2f, 0xb5, 0x52, 0x82, 0x19, 0xac, 0x1d, 0x59, 0x48, 0xc9, 0x98, 0x8b, 0xd9, 0x29, 0xb5,
	0x04, 0xaf, 0x25, 0x02, 0x60, 0x25, 0x60, 0xd3, 0x59, 0x83, 0x40, 0x4c, 0x27, 0xd6, 0x55, 0x4c,
	0x4c, 0x01, 0x49, 0x49, 0x6c, 0x60, 0x27, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x10, 0x66,
	0xac, 0xad, 0xbf, 0x00, 0x00, 0x00,
}
