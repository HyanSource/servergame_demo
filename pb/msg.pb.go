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

//單線的賠率
type LineOdds struct {
	GetID                int32    `protobuf:"varint,1,opt,name=GetID,proto3" json:"GetID,omitempty"`
	GetCount             int32    `protobuf:"varint,2,opt,name=GetCount,proto3" json:"GetCount,omitempty"`
	GetOdds              int32    `protobuf:"varint,3,opt,name=GetOdds,proto3" json:"GetOdds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LineOdds) Reset()         { *m = LineOdds{} }
func (m *LineOdds) String() string { return proto.CompactTextString(m) }
func (*LineOdds) ProtoMessage()    {}
func (*LineOdds) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}

func (m *LineOdds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LineOdds.Unmarshal(m, b)
}
func (m *LineOdds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LineOdds.Marshal(b, m, deterministic)
}
func (m *LineOdds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LineOdds.Merge(m, src)
}
func (m *LineOdds) XXX_Size() int {
	return xxx_messageInfo_LineOdds.Size(m)
}
func (m *LineOdds) XXX_DiscardUnknown() {
	xxx_messageInfo_LineOdds.DiscardUnknown(m)
}

var xxx_messageInfo_LineOdds proto.InternalMessageInfo

func (m *LineOdds) GetGetID() int32 {
	if m != nil {
		return m.GetID
	}
	return 0
}

func (m *LineOdds) GetGetCount() int32 {
	if m != nil {
		return m.GetCount
	}
	return 0
}

func (m *LineOdds) GetGetOdds() int32 {
	if m != nil {
		return m.GetOdds
	}
	return 0
}

//遊戲結果
type GameResult struct {
	Result               string      `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	LinesOdds            []*LineOdds `protobuf:"bytes,2,rep,name=LinesOdds,proto3" json:"LinesOdds,omitempty"`
	ScatterCount         int32       `protobuf:"varint,3,opt,name=ScatterCount,proto3" json:"ScatterCount,omitempty"`
	AllOdds              int32       `protobuf:"varint,4,opt,name=AllOdds,proto3" json:"AllOdds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GameResult) Reset()         { *m = GameResult{} }
func (m *GameResult) String() string { return proto.CompactTextString(m) }
func (*GameResult) ProtoMessage()    {}
func (*GameResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{3}
}

func (m *GameResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameResult.Unmarshal(m, b)
}
func (m *GameResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameResult.Marshal(b, m, deterministic)
}
func (m *GameResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameResult.Merge(m, src)
}
func (m *GameResult) XXX_Size() int {
	return xxx_messageInfo_GameResult.Size(m)
}
func (m *GameResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GameResult.DiscardUnknown(m)
}

var xxx_messageInfo_GameResult proto.InternalMessageInfo

func (m *GameResult) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *GameResult) GetLinesOdds() []*LineOdds {
	if m != nil {
		return m.LinesOdds
	}
	return nil
}

func (m *GameResult) GetScatterCount() int32 {
	if m != nil {
		return m.ScatterCount
	}
	return 0
}

func (m *GameResult) GetAllOdds() int32 {
	if m != nil {
		return m.AllOdds
	}
	return 0
}

//金額和遊戲結果
type ReturnGameResult struct {
	AllMoney             int32       `protobuf:"varint,1,opt,name=AllMoney,proto3" json:"AllMoney,omitempty"`
	GetMoney             int32       `protobuf:"varint,2,opt,name=GetMoney,proto3" json:"GetMoney,omitempty"`
	FreeCount            int32       `protobuf:"varint,3,opt,name=FreeCount,proto3" json:"FreeCount,omitempty"`
	FreeRound            int32       `protobuf:"varint,4,opt,name=FreeRound,proto3" json:"FreeRound,omitempty"`
	GameResult           *GameResult `protobuf:"bytes,5,opt,name=GameResult,proto3" json:"GameResult,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReturnGameResult) Reset()         { *m = ReturnGameResult{} }
func (m *ReturnGameResult) String() string { return proto.CompactTextString(m) }
func (*ReturnGameResult) ProtoMessage()    {}
func (*ReturnGameResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{4}
}

func (m *ReturnGameResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReturnGameResult.Unmarshal(m, b)
}
func (m *ReturnGameResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReturnGameResult.Marshal(b, m, deterministic)
}
func (m *ReturnGameResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReturnGameResult.Merge(m, src)
}
func (m *ReturnGameResult) XXX_Size() int {
	return xxx_messageInfo_ReturnGameResult.Size(m)
}
func (m *ReturnGameResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ReturnGameResult.DiscardUnknown(m)
}

var xxx_messageInfo_ReturnGameResult proto.InternalMessageInfo

func (m *ReturnGameResult) GetAllMoney() int32 {
	if m != nil {
		return m.AllMoney
	}
	return 0
}

func (m *ReturnGameResult) GetGetMoney() int32 {
	if m != nil {
		return m.GetMoney
	}
	return 0
}

func (m *ReturnGameResult) GetFreeCount() int32 {
	if m != nil {
		return m.FreeCount
	}
	return 0
}

func (m *ReturnGameResult) GetFreeRound() int32 {
	if m != nil {
		return m.FreeRound
	}
	return 0
}

func (m *ReturnGameResult) GetGameResult() *GameResult {
	if m != nil {
		return m.GameResult
	}
	return nil
}

//聊天訊息
type TalkMsg struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TalkMsg) Reset()         { *m = TalkMsg{} }
func (m *TalkMsg) String() string { return proto.CompactTextString(m) }
func (*TalkMsg) ProtoMessage()    {}
func (*TalkMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{5}
}

func (m *TalkMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TalkMsg.Unmarshal(m, b)
}
func (m *TalkMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TalkMsg.Marshal(b, m, deterministic)
}
func (m *TalkMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TalkMsg.Merge(m, src)
}
func (m *TalkMsg) XXX_Size() int {
	return xxx_messageInfo_TalkMsg.Size(m)
}
func (m *TalkMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_TalkMsg.DiscardUnknown(m)
}

var xxx_messageInfo_TalkMsg proto.InternalMessageInfo

func (m *TalkMsg) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

//廣播
type BroadCast struct {
	PlayerID             int32    `protobuf:"varint,1,opt,name=PlayerID,proto3" json:"PlayerID,omitempty"`
	TalkMsg              *TalkMsg `protobuf:"bytes,2,opt,name=TalkMsg,proto3" json:"TalkMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BroadCast) Reset()         { *m = BroadCast{} }
func (m *BroadCast) String() string { return proto.CompactTextString(m) }
func (*BroadCast) ProtoMessage()    {}
func (*BroadCast) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{6}
}

func (m *BroadCast) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadCast.Unmarshal(m, b)
}
func (m *BroadCast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadCast.Marshal(b, m, deterministic)
}
func (m *BroadCast) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadCast.Merge(m, src)
}
func (m *BroadCast) XXX_Size() int {
	return xxx_messageInfo_BroadCast.Size(m)
}
func (m *BroadCast) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadCast.DiscardUnknown(m)
}

var xxx_messageInfo_BroadCast proto.InternalMessageInfo

func (m *BroadCast) GetPlayerID() int32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *BroadCast) GetTalkMsg() *TalkMsg {
	if m != nil {
		return m.TalkMsg
	}
	return nil
}

func init() {
	proto.RegisterType((*PlayerData)(nil), "pb.PlayerData")
	proto.RegisterType((*Bet)(nil), "pb.Bet")
	proto.RegisterType((*LineOdds)(nil), "pb.LineOdds")
	proto.RegisterType((*GameResult)(nil), "pb.GameResult")
	proto.RegisterType((*ReturnGameResult)(nil), "pb.ReturnGameResult")
	proto.RegisterType((*TalkMsg)(nil), "pb.TalkMsg")
	proto.RegisterType((*BroadCast)(nil), "pb.BroadCast")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x6a, 0xe3, 0x30,
	0x10, 0xc6, 0xce, 0x3a, 0x89, 0xc7, 0x61, 0x09, 0x62, 0x59, 0x4c, 0x58, 0x16, 0xa3, 0x65, 0xc1,
	0xf4, 0xe0, 0x43, 0xf2, 0x04, 0xf9, 0xa1, 0xa6, 0xd0, 0xa4, 0x41, 0x2d, 0xbd, 0xcb, 0xf1, 0x10,
	0xda, 0x3a, 0x76, 0xb0, 0xe5, 0x43, 0x5e, 0xa2, 0x0f, 0xd2, 0x07, 0xe8, 0xf3, 0x15, 0x49, 0x96,
	0xed, 0x5e, 0x72, 0xd3, 0xf7, 0xcd, 0xcc, 0x37, 0xdf, 0x37, 0x08, 0xdc, 0x53, 0x75, 0x8c, 0xce,
	0x65, 0x21, 0x0a, 0x62, 0x9f, 0x13, 0xfa, 0x0a, 0xb0, 0xcf, 0xf8, 0x05, 0xcb, 0x0d, 0x17, 0x9c,
	0xcc, 0x60, 0xac, 0xd1, 0xdd, 0xc6, 0xb7, 0x02, 0x2b, 0x74, 0x58, 0x8b, 0xc9, 0x5f, 0xd3, 0xb9,
	0xe3, 0x27, 0xf4, 0xed, 0xc0, 0x0a, 0x5d, 0xd6, 0x63, 0x48, 0x00, 0x9e, 0x46, 0xdb, 0x22, 0xc7,
	0x8b, 0x3f, 0x50, 0xe3, 0x7d, 0x8a, 0x2e, 0x60, 0xb0, 0x42, 0x71, 0x75, 0xc9, 0x54, 0xb5, 0x28,
	0x75, 0x87, 0xc9, 0x27, 0x7d, 0x86, 0xf1, 0xfd, 0x4b, 0x8e, 0x0f, 0x69, 0x5a, 0x91, 0x5f, 0xe0,
	0xc4, 0x28, 0xda, 0x31, 0x0d, 0xa4, 0x5e, 0x8c, 0x62, 0x5d, 0xd4, 0xb9, 0x19, 0x6c, 0x31, 0xf1,
	0x61, 0x14, 0xa3, 0x90, 0xc3, 0x8d, 0x21, 0x03, 0xe9, 0xbb, 0x05, 0x10, 0xf3, 0x13, 0x32, 0xac,
	0xea, 0x4c, 0x90, 0xdf, 0x30, 0xd4, 0x2f, 0xa5, 0xed, 0xb2, 0x06, 0x91, 0x1b, 0x70, 0xe5, 0xfa,
	0x4a, 0x49, 0xd8, 0xc1, 0x20, 0xf4, 0xe6, 0x93, 0xe8, 0x9c, 0x44, 0xc6, 0x13, 0xeb, 0xca, 0x84,
	0xc2, 0xe4, 0xf1, 0xc0, 0x85, 0xc0, 0x52, 0x9b, 0xd1, 0x1b, 0xbf, 0x71, 0xd2, 0xd0, 0x32, 0xcb,
	0x94, 0xda, 0x0f, 0x6d, 0xa8, 0x81, 0xf4, 0xd3, 0x82, 0x29, 0x43, 0x51, 0x97, 0x79, 0xcf, 0xd6,
	0x0c, 0xc6, 0xcb, 0x2c, 0xd3, 0x17, 0x6d, 0x6e, 0x65, 0x70, 0x93, 0x5b, 0xd7, 0xba, 0xdc, 0xba,
	0xf6, 0x07, 0xdc, 0xdb, 0x12, 0xb1, 0xef, 0xa3, 0x23, 0x4c, 0x95, 0x15, 0x75, 0x9e, 0x36, 0x36,
	0x3a, 0x82, 0x44, 0xfd, 0xc3, 0xf8, 0x4e, 0x60, 0x85, 0xde, 0xfc, 0xa7, 0xcc, 0xdc, 0xb1, 0xac,
	0xd7, 0x41, 0xff, 0xc1, 0xe8, 0x89, 0x67, 0x6f, 0xdb, 0xea, 0x28, 0xd3, 0x1d, 0x8a, 0x5c, 0x60,
	0x6e, 0xce, 0x68, 0x20, 0xdd, 0x81, 0xbb, 0x2a, 0x0b, 0x9e, 0xae, 0x79, 0x75, 0xfd, 0x07, 0xfc,
	0x6f, 0xd5, 0x54, 0x28, 0x6f, 0xee, 0xc9, 0xd5, 0x0d, 0xc5, 0x4c, 0x6d, 0xe5, 0x7c, 0xd8, 0xf6,
	0x3e, 0x49, 0x86, 0xea, 0x27, 0x2f, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x82, 0x3e, 0xa6, 0x0e,
	0xd6, 0x02, 0x00, 0x00,
}
