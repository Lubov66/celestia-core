// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cometbft/consensus/v1/wal.proto

package v1

import (
	fmt "fmt"
	v1 "github.com/cometbft/cometbft/api/cometbft/types/v1"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/cosmos/gogoproto/types"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "github.com/golang/protobuf/ptypes/duration"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgInfo are msgs from the reactor which may update the state
type MsgInfo struct {
	Msg    Message `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
	PeerID string  `protobuf:"bytes,2,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
}

func (m *MsgInfo) Reset()         { *m = MsgInfo{} }
func (m *MsgInfo) String() string { return proto.CompactTextString(m) }
func (*MsgInfo) ProtoMessage()    {}
func (*MsgInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f013b385bb75d435, []int{0}
}
func (m *MsgInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInfo.Merge(m, src)
}
func (m *MsgInfo) XXX_Size() int {
	return m.Size()
}
func (m *MsgInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInfo proto.InternalMessageInfo

func (m *MsgInfo) GetMsg() Message {
	if m != nil {
		return m.Msg
	}
	return Message{}
}

func (m *MsgInfo) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

// TimeoutInfo internally generated messages which may update the state
type TimeoutInfo struct {
	Duration time.Duration `protobuf:"bytes,1,opt,name=duration,proto3,stdduration" json:"duration"`
	Height   int64         `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Round    int32         `protobuf:"varint,3,opt,name=round,proto3" json:"round,omitempty"`
	Step     uint32        `protobuf:"varint,4,opt,name=step,proto3" json:"step,omitempty"`
}

func (m *TimeoutInfo) Reset()         { *m = TimeoutInfo{} }
func (m *TimeoutInfo) String() string { return proto.CompactTextString(m) }
func (*TimeoutInfo) ProtoMessage()    {}
func (*TimeoutInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f013b385bb75d435, []int{1}
}
func (m *TimeoutInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimeoutInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimeoutInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimeoutInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeoutInfo.Merge(m, src)
}
func (m *TimeoutInfo) XXX_Size() int {
	return m.Size()
}
func (m *TimeoutInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeoutInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TimeoutInfo proto.InternalMessageInfo

func (m *TimeoutInfo) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *TimeoutInfo) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *TimeoutInfo) GetRound() int32 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *TimeoutInfo) GetStep() uint32 {
	if m != nil {
		return m.Step
	}
	return 0
}

// EndHeight marks the end of the given height inside WAL.
// @internal used by scripts/wal2json util.
type EndHeight struct {
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *EndHeight) Reset()         { *m = EndHeight{} }
func (m *EndHeight) String() string { return proto.CompactTextString(m) }
func (*EndHeight) ProtoMessage()    {}
func (*EndHeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_f013b385bb75d435, []int{2}
}
func (m *EndHeight) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EndHeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EndHeight.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EndHeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndHeight.Merge(m, src)
}
func (m *EndHeight) XXX_Size() int {
	return m.Size()
}
func (m *EndHeight) XXX_DiscardUnknown() {
	xxx_messageInfo_EndHeight.DiscardUnknown(m)
}

var xxx_messageInfo_EndHeight proto.InternalMessageInfo

func (m *EndHeight) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type WALMessage struct {
	// Types that are valid to be assigned to Sum:
	//	*WALMessage_EventDataRoundState
	//	*WALMessage_MsgInfo
	//	*WALMessage_TimeoutInfo
	//	*WALMessage_EndHeight
	Sum isWALMessage_Sum `protobuf_oneof:"sum"`
}

func (m *WALMessage) Reset()         { *m = WALMessage{} }
func (m *WALMessage) String() string { return proto.CompactTextString(m) }
func (*WALMessage) ProtoMessage()    {}
func (*WALMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f013b385bb75d435, []int{3}
}
func (m *WALMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WALMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WALMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WALMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WALMessage.Merge(m, src)
}
func (m *WALMessage) XXX_Size() int {
	return m.Size()
}
func (m *WALMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_WALMessage.DiscardUnknown(m)
}

var xxx_messageInfo_WALMessage proto.InternalMessageInfo

type isWALMessage_Sum interface {
	isWALMessage_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type WALMessage_EventDataRoundState struct {
	EventDataRoundState *v1.EventDataRoundState `protobuf:"bytes,1,opt,name=event_data_round_state,json=eventDataRoundState,proto3,oneof" json:"event_data_round_state,omitempty"`
}
type WALMessage_MsgInfo struct {
	MsgInfo *MsgInfo `protobuf:"bytes,2,opt,name=msg_info,json=msgInfo,proto3,oneof" json:"msg_info,omitempty"`
}
type WALMessage_TimeoutInfo struct {
	TimeoutInfo *TimeoutInfo `protobuf:"bytes,3,opt,name=timeout_info,json=timeoutInfo,proto3,oneof" json:"timeout_info,omitempty"`
}
type WALMessage_EndHeight struct {
	EndHeight *EndHeight `protobuf:"bytes,4,opt,name=end_height,json=endHeight,proto3,oneof" json:"end_height,omitempty"`
}

func (*WALMessage_EventDataRoundState) isWALMessage_Sum() {}
func (*WALMessage_MsgInfo) isWALMessage_Sum()             {}
func (*WALMessage_TimeoutInfo) isWALMessage_Sum()         {}
func (*WALMessage_EndHeight) isWALMessage_Sum()           {}

func (m *WALMessage) GetSum() isWALMessage_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *WALMessage) GetEventDataRoundState() *v1.EventDataRoundState {
	if x, ok := m.GetSum().(*WALMessage_EventDataRoundState); ok {
		return x.EventDataRoundState
	}
	return nil
}

func (m *WALMessage) GetMsgInfo() *MsgInfo {
	if x, ok := m.GetSum().(*WALMessage_MsgInfo); ok {
		return x.MsgInfo
	}
	return nil
}

func (m *WALMessage) GetTimeoutInfo() *TimeoutInfo {
	if x, ok := m.GetSum().(*WALMessage_TimeoutInfo); ok {
		return x.TimeoutInfo
	}
	return nil
}

func (m *WALMessage) GetEndHeight() *EndHeight {
	if x, ok := m.GetSum().(*WALMessage_EndHeight); ok {
		return x.EndHeight
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*WALMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*WALMessage_EventDataRoundState)(nil),
		(*WALMessage_MsgInfo)(nil),
		(*WALMessage_TimeoutInfo)(nil),
		(*WALMessage_EndHeight)(nil),
	}
}

// TimedWALMessage wraps WALMessage and adds Time for debugging purposes.
type TimedWALMessage struct {
	Time time.Time   `protobuf:"bytes,1,opt,name=time,proto3,stdtime" json:"time"`
	Msg  *WALMessage `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *TimedWALMessage) Reset()         { *m = TimedWALMessage{} }
func (m *TimedWALMessage) String() string { return proto.CompactTextString(m) }
func (*TimedWALMessage) ProtoMessage()    {}
func (*TimedWALMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f013b385bb75d435, []int{4}
}
func (m *TimedWALMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimedWALMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimedWALMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimedWALMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimedWALMessage.Merge(m, src)
}
func (m *TimedWALMessage) XXX_Size() int {
	return m.Size()
}
func (m *TimedWALMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TimedWALMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TimedWALMessage proto.InternalMessageInfo

func (m *TimedWALMessage) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *TimedWALMessage) GetMsg() *WALMessage {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgInfo)(nil), "cometbft.consensus.v1.MsgInfo")
	proto.RegisterType((*TimeoutInfo)(nil), "cometbft.consensus.v1.TimeoutInfo")
	proto.RegisterType((*EndHeight)(nil), "cometbft.consensus.v1.EndHeight")
	proto.RegisterType((*WALMessage)(nil), "cometbft.consensus.v1.WALMessage")
	proto.RegisterType((*TimedWALMessage)(nil), "cometbft.consensus.v1.TimedWALMessage")
}

func init() { proto.RegisterFile("cometbft/consensus/v1/wal.proto", fileDescriptor_f013b385bb75d435) }

var fileDescriptor_f013b385bb75d435 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0xd6, 0xc6, 0xff, 0xe3, 0x96, 0x82, 0x9a, 0x06, 0xd7, 0x07, 0xd9, 0x51, 0xa0, 0xf8, 0x24,
	0xe1, 0x06, 0x42, 0xa1, 0x87, 0x12, 0xe3, 0x50, 0x1b, 0x1a, 0x08, 0xdb, 0x40, 0xa1, 0x50, 0x84,
	0x1c, 0xad, 0xd7, 0x82, 0x48, 0x2b, 0xbc, 0x2b, 0x97, 0xde, 0xfa, 0x08, 0x3e, 0xf6, 0x5d, 0xfa,
	0x02, 0x39, 0xe6, 0xd8, 0x53, 0x5a, 0xec, 0x17, 0x29, 0xfb, 0xe3, 0x1f, 0x1a, 0xbb, 0xb7, 0x9d,
	0x9d, 0x6f, 0xbe, 0x99, 0xf9, 0xbe, 0x5d, 0x68, 0xdd, 0xb0, 0x84, 0x88, 0xd1, 0x58, 0xf8, 0x37,
	0x2c, 0xe5, 0x24, 0xe5, 0x39, 0xf7, 0x67, 0x5d, 0xff, 0x6b, 0x78, 0xeb, 0x65, 0x53, 0x26, 0x98,
	0xfd, 0x62, 0x05, 0xf0, 0xd6, 0x00, 0x6f, 0xd6, 0x6d, 0x1e, 0x52, 0x46, 0x99, 0x42, 0xf8, 0xf2,
	0xa4, 0xc1, 0xcd, 0xe3, 0xdd, 0x6c, 0xe2, 0x5b, 0x46, 0xb8, 0x81, 0x38, 0x6b, 0x88, 0xba, 0x95,
	0x69, 0x32, 0x23, 0xa9, 0x58, 0xe7, 0x29, 0x63, 0xf4, 0x96, 0xf8, 0x2a, 0x1a, 0xe5, 0x63, 0x3f,
	0xca, 0xa7, 0xa1, 0x88, 0x59, 0x6a, 0xf2, 0xad, 0x7f, 0xf3, 0x22, 0x4e, 0x08, 0x17, 0x61, 0x92,
	0x69, 0x80, 0x3b, 0x86, 0xca, 0x25, 0xa7, 0xc3, 0x74, 0xcc, 0xec, 0x33, 0x28, 0x24, 0x9c, 0x36,
	0x50, 0x1b, 0x75, 0xea, 0xaf, 0x1d, 0x6f, 0xe7, 0x26, 0xde, 0x25, 0xe1, 0x3c, 0xa4, 0xa4, 0x57,
	0xbc, 0x7b, 0x68, 0x59, 0x58, 0x16, 0xd8, 0x27, 0x50, 0xc9, 0x08, 0x99, 0x06, 0x71, 0xd4, 0x38,
	0x68, 0xa3, 0x4e, 0xad, 0x07, 0x8b, 0x87, 0x56, 0xf9, 0x8a, 0x90, 0xe9, 0xb0, 0x8f, 0xcb, 0x32,
	0x35, 0x8c, 0xdc, 0x39, 0x82, 0xfa, 0x75, 0x9c, 0x10, 0x96, 0x0b, 0xd5, 0xec, 0x1d, 0x54, 0x57,
	0xa3, 0x9a, 0x8e, 0x2f, 0x3d, 0x3d, 0xab, 0xb7, 0x9a, 0xd5, 0xeb, 0x1b, 0x40, 0xaf, 0x2a, 0x9b,
	0xfd, 0xf8, 0xdd, 0x42, 0x78, 0x5d, 0x64, 0x1f, 0x41, 0x79, 0x42, 0x62, 0x3a, 0x11, 0xaa, 0x69,
	0x01, 0x9b, 0xc8, 0x3e, 0x84, 0xd2, 0x94, 0xe5, 0x69, 0xd4, 0x28, 0xb4, 0x51, 0xa7, 0x84, 0x75,
	0x60, 0xdb, 0x50, 0xe4, 0x82, 0x64, 0x8d, 0x62, 0x1b, 0x75, 0x9e, 0x62, 0x75, 0x76, 0x4f, 0xa0,
	0x76, 0x91, 0x46, 0x03, 0x5d, 0xb6, 0xa1, 0x43, 0xdb, 0x74, 0xee, 0xcf, 0x03, 0x80, 0x4f, 0xe7,
	0x1f, 0xcc, 0xda, 0xf6, 0x17, 0x38, 0x52, 0xfa, 0x07, 0x51, 0x28, 0xc2, 0x40, 0x71, 0x07, 0x5c,
	0x84, 0x82, 0x98, 0x25, 0x5e, 0x6d, 0x64, 0xd3, 0x36, 0xce, 0xba, 0xde, 0x85, 0x2c, 0xe8, 0x87,
	0x22, 0xc4, 0x12, 0xfe, 0x51, 0xa2, 0x07, 0x16, 0x7e, 0x4e, 0x1e, 0x5f, 0xdb, 0x6f, 0xa1, 0x9a,
	0x70, 0x1a, 0xc4, 0xe9, 0x98, 0xa9, 0xb5, 0xfe, 0xe3, 0x83, 0x36, 0x6d, 0x60, 0xe1, 0x4a, 0x62,
	0xfc, 0x7b, 0x0f, 0x4f, 0x84, 0x56, 0x58, 0x13, 0x14, 0x14, 0x81, 0xbb, 0x87, 0x60, 0xcb, 0x8c,
	0x81, 0x85, 0xeb, 0x62, 0xcb, 0x9b, 0x73, 0x00, 0x92, 0x46, 0x81, 0xd1, 0xa3, 0xa8, 0x68, 0xda,
	0x7b, 0x68, 0xd6, 0x0a, 0x0e, 0x2c, 0x5c, 0x23, 0xab, 0xa0, 0x57, 0x82, 0x02, 0xcf, 0x13, 0xf7,
	0x3b, 0x82, 0x67, 0xb2, 0x51, 0xb4, 0x25, 0xe1, 0x1b, 0x28, 0xca, 0x66, 0x46, 0xb0, 0xe6, 0x23,
	0xd7, 0xaf, 0x57, 0x2f, 0x54, 0xdb, 0x3e, 0x97, 0xb6, 0xab, 0x0a, 0xfb, 0x54, 0x3f, 0x50, 0x2d,
	0xcc, 0xf1, 0x9e, 0x81, 0x36, 0x9d, 0xd4, 0xeb, 0xec, 0x5d, 0xdd, 0x2d, 0x1c, 0x74, 0xbf, 0x70,
	0xd0, 0x9f, 0x85, 0x83, 0xe6, 0x4b, 0xc7, 0xba, 0x5f, 0x3a, 0xd6, 0xaf, 0xa5, 0x63, 0x7d, 0x3e,
	0xa3, 0xb1, 0x98, 0xe4, 0x23, 0xc9, 0xe3, 0x6f, 0xfd, 0x44, 0x73, 0x08, 0xb3, 0xd8, 0xdf, 0xf9,
	0x3f, 0x47, 0x65, 0x35, 0xea, 0xe9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xb9, 0xcf, 0xe3,
	0x0d, 0x04, 0x00, 0x00,
}

func (m *MsgInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PeerID) > 0 {
		i -= len(m.PeerID)
		copy(dAtA[i:], m.PeerID)
		i = encodeVarintWal(dAtA, i, uint64(len(m.PeerID)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Msg.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintWal(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *TimeoutInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimeoutInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimeoutInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Step != 0 {
		i = encodeVarintWal(dAtA, i, uint64(m.Step))
		i--
		dAtA[i] = 0x20
	}
	if m.Round != 0 {
		i = encodeVarintWal(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x18
	}
	if m.Height != 0 {
		i = encodeVarintWal(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintWal(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *EndHeight) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EndHeight) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EndHeight) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintWal(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *WALMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WALMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WALMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *WALMessage_EventDataRoundState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WALMessage_EventDataRoundState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.EventDataRoundState != nil {
		{
			size, err := m.EventDataRoundState.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *WALMessage_MsgInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WALMessage_MsgInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MsgInfo != nil {
		{
			size, err := m.MsgInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *WALMessage_TimeoutInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WALMessage_TimeoutInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.TimeoutInfo != nil {
		{
			size, err := m.TimeoutInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *WALMessage_EndHeight) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WALMessage_EndHeight) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.EndHeight != nil {
		{
			size, err := m.EndHeight.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *TimedWALMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimedWALMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimedWALMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Msg != nil {
		{
			size, err := m.Msg.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	n8, err8 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time):])
	if err8 != nil {
		return 0, err8
	}
	i -= n8
	i = encodeVarintWal(dAtA, i, uint64(n8))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintWal(dAtA []byte, offset int, v uint64) int {
	offset -= sovWal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Msg.Size()
	n += 1 + l + sovWal(uint64(l))
	l = len(m.PeerID)
	if l > 0 {
		n += 1 + l + sovWal(uint64(l))
	}
	return n
}

func (m *TimeoutInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovWal(uint64(l))
	if m.Height != 0 {
		n += 1 + sovWal(uint64(m.Height))
	}
	if m.Round != 0 {
		n += 1 + sovWal(uint64(m.Round))
	}
	if m.Step != 0 {
		n += 1 + sovWal(uint64(m.Step))
	}
	return n
}

func (m *EndHeight) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovWal(uint64(m.Height))
	}
	return n
}

func (m *WALMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *WALMessage_EventDataRoundState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EventDataRoundState != nil {
		l = m.EventDataRoundState.Size()
		n += 1 + l + sovWal(uint64(l))
	}
	return n
}
func (m *WALMessage_MsgInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MsgInfo != nil {
		l = m.MsgInfo.Size()
		n += 1 + l + sovWal(uint64(l))
	}
	return n
}
func (m *WALMessage_TimeoutInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TimeoutInfo != nil {
		l = m.TimeoutInfo.Size()
		n += 1 + l + sovWal(uint64(l))
	}
	return n
}
func (m *WALMessage_EndHeight) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EndHeight != nil {
		l = m.EndHeight.Size()
		n += 1 + l + sovWal(uint64(l))
	}
	return n
}
func (m *TimedWALMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovWal(uint64(l))
	if m.Msg != nil {
		l = m.Msg.Size()
		n += 1 + l + sovWal(uint64(l))
	}
	return n
}

func sovWal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWal(x uint64) (n int) {
	return sovWal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Msg.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TimeoutInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TimeoutInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeoutInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Step", wireType)
			}
			m.Step = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Step |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EndHeight) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EndHeight: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EndHeight: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WALMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WALMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WALMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventDataRoundState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &v1.EventDataRoundState{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &WALMessage_EventDataRoundState{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &MsgInfo{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &WALMessage_MsgInfo{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TimeoutInfo{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &WALMessage_TimeoutInfo{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &EndHeight{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &WALMessage_EndHeight{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TimedWALMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWal
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TimedWALMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimedWALMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Msg == nil {
				m.Msg = &WALMessage{}
			}
			if err := m.Msg.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWal
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthWal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWal = fmt.Errorf("proto: unexpected end of group")
)