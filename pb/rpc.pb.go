// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc.proto

package pb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RPC struct {
	Subscriptions        []*RPC_SubOpts `protobuf:"bytes,1,rep,name=subscriptions" json:"subscriptions,omitempty"`
	Rpcs                 []*TopicRpc    `protobuf:"bytes,2,rep,name=rpcs" json:"rpcs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RPC) Reset()         { *m = RPC{} }
func (m *RPC) String() string { return proto.CompactTextString(m) }
func (*RPC) ProtoMessage()    {}
func (*RPC) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}
func (m *RPC) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RPC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RPC.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RPC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPC.Merge(m, src)
}
func (m *RPC) XXX_Size() int {
	return m.Size()
}
func (m *RPC) XXX_DiscardUnknown() {
	xxx_messageInfo_RPC.DiscardUnknown(m)
}

var xxx_messageInfo_RPC proto.InternalMessageInfo

func (m *RPC) GetSubscriptions() []*RPC_SubOpts {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

func (m *RPC) GetRpcs() []*TopicRpc {
	if m != nil {
		return m.Rpcs
	}
	return nil
}

type RPC_SubOpts struct {
	Subscribe            *bool    `protobuf:"varint,1,opt,name=subscribe" json:"subscribe,omitempty"`
	Topicid              *string  `protobuf:"bytes,2,opt,name=topicid" json:"topicid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPC_SubOpts) Reset()         { *m = RPC_SubOpts{} }
func (m *RPC_SubOpts) String() string { return proto.CompactTextString(m) }
func (*RPC_SubOpts) ProtoMessage()    {}
func (*RPC_SubOpts) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0, 0}
}
func (m *RPC_SubOpts) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RPC_SubOpts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RPC_SubOpts.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RPC_SubOpts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPC_SubOpts.Merge(m, src)
}
func (m *RPC_SubOpts) XXX_Size() int {
	return m.Size()
}
func (m *RPC_SubOpts) XXX_DiscardUnknown() {
	xxx_messageInfo_RPC_SubOpts.DiscardUnknown(m)
}

var xxx_messageInfo_RPC_SubOpts proto.InternalMessageInfo

func (m *RPC_SubOpts) GetSubscribe() bool {
	if m != nil && m.Subscribe != nil {
		return *m.Subscribe
	}
	return false
}

func (m *RPC_SubOpts) GetTopicid() string {
	if m != nil && m.Topicid != nil {
		return *m.Topicid
	}
	return ""
}

type TopicRpc struct {
	Topicid *string `protobuf:"bytes,1,opt,name=topicid" json:"topicid,omitempty"`
	// Router-specific types
	Cat                  *CatRpc      `protobuf:"bytes,2,opt,name=cat" json:"cat,omitempty"`
	Floodsub             *FloodsubRpc `protobuf:"bytes,3,opt,name=floodsub" json:"floodsub,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TopicRpc) Reset()         { *m = TopicRpc{} }
func (m *TopicRpc) String() string { return proto.CompactTextString(m) }
func (*TopicRpc) ProtoMessage()    {}
func (*TopicRpc) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}
func (m *TopicRpc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TopicRpc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TopicRpc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TopicRpc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicRpc.Merge(m, src)
}
func (m *TopicRpc) XXX_Size() int {
	return m.Size()
}
func (m *TopicRpc) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicRpc.DiscardUnknown(m)
}

var xxx_messageInfo_TopicRpc proto.InternalMessageInfo

func (m *TopicRpc) GetTopicid() string {
	if m != nil && m.Topicid != nil {
		return *m.Topicid
	}
	return ""
}

func (m *TopicRpc) GetCat() *CatRpc {
	if m != nil {
		return m.Cat
	}
	return nil
}

func (m *TopicRpc) GetFloodsub() *FloodsubRpc {
	if m != nil {
		return m.Floodsub
	}
	return nil
}

type CatRpc struct {
	Chunks               []*CatRpc_Chunk `protobuf:"bytes,1,rep,name=chunks" json:"chunks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CatRpc) Reset()         { *m = CatRpc{} }
func (m *CatRpc) String() string { return proto.CompactTextString(m) }
func (*CatRpc) ProtoMessage()    {}
func (*CatRpc) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}
func (m *CatRpc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CatRpc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CatRpc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CatRpc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatRpc.Merge(m, src)
}
func (m *CatRpc) XXX_Size() int {
	return m.Size()
}
func (m *CatRpc) XXX_DiscardUnknown() {
	xxx_messageInfo_CatRpc.DiscardUnknown(m)
}

var xxx_messageInfo_CatRpc proto.InternalMessageInfo

func (m *CatRpc) GetChunks() []*CatRpc_Chunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

type CatRpc_Chunk struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Coefficients         [][]byte `protobuf:"bytes,2,rep,name=coefficients" json:"coefficients,omitempty"`
	Extra                []byte   `protobuf:"bytes,3,opt,name=extra" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CatRpc_Chunk) Reset()         { *m = CatRpc_Chunk{} }
func (m *CatRpc_Chunk) String() string { return proto.CompactTextString(m) }
func (*CatRpc_Chunk) ProtoMessage()    {}
func (*CatRpc_Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2, 0}
}
func (m *CatRpc_Chunk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CatRpc_Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CatRpc_Chunk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CatRpc_Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatRpc_Chunk.Merge(m, src)
}
func (m *CatRpc_Chunk) XXX_Size() int {
	return m.Size()
}
func (m *CatRpc_Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_CatRpc_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_CatRpc_Chunk proto.InternalMessageInfo

func (m *CatRpc_Chunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CatRpc_Chunk) GetCoefficients() [][]byte {
	if m != nil {
		return m.Coefficients
	}
	return nil
}

func (m *CatRpc_Chunk) GetExtra() []byte {
	if m != nil {
		return m.Extra
	}
	return nil
}

type FloodsubRpc struct {
	Messages             [][]byte `protobuf:"bytes,1,rep,name=messages" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FloodsubRpc) Reset()         { *m = FloodsubRpc{} }
func (m *FloodsubRpc) String() string { return proto.CompactTextString(m) }
func (*FloodsubRpc) ProtoMessage()    {}
func (*FloodsubRpc) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{3}
}
func (m *FloodsubRpc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FloodsubRpc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FloodsubRpc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FloodsubRpc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FloodsubRpc.Merge(m, src)
}
func (m *FloodsubRpc) XXX_Size() int {
	return m.Size()
}
func (m *FloodsubRpc) XXX_DiscardUnknown() {
	xxx_messageInfo_FloodsubRpc.DiscardUnknown(m)
}

var xxx_messageInfo_FloodsubRpc proto.InternalMessageInfo

func (m *FloodsubRpc) GetMessages() [][]byte {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*RPC)(nil), "pb.RPC")
	proto.RegisterType((*RPC_SubOpts)(nil), "pb.RPC.SubOpts")
	proto.RegisterType((*TopicRpc)(nil), "pb.TopicRpc")
	proto.RegisterType((*CatRpc)(nil), "pb.CatRpc")
	proto.RegisterType((*CatRpc_Chunk)(nil), "pb.CatRpc.Chunk")
	proto.RegisterType((*FloodsubRpc)(nil), "pb.FloodsubRpc")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x5d, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x99, 0xa6, 0x3f, 0xe9, 0xed, 0x88, 0x72, 0xf1, 0x21, 0x94, 0x52, 0x42, 0x9e, 0x22,
	0x42, 0x1e, 0x0a, 0x2e, 0x40, 0x03, 0xbe, 0x5a, 0xae, 0xba, 0x80, 0x64, 0x9a, 0xea, 0xa0, 0x76,
	0x86, 0xcc, 0x04, 0xdc, 0x82, 0xcb, 0x70, 0x37, 0x3e, 0xba, 0x04, 0xe9, 0x4a, 0x64, 0x26, 0xe9,
	0xdf, 0xdb, 0xdc, 0x73, 0xbe, 0x73, 0x39, 0x73, 0x61, 0x5c, 0x6b, 0x91, 0xe9, 0x5a, 0x59, 0x85,
	0x3d, 0x5d, 0x26, 0xdf, 0x0c, 0x02, 0x5a, 0xe6, 0x78, 0x03, 0x67, 0xa6, 0x29, 0x8d, 0xa8, 0xa5,
	0xb6, 0x52, 0x6d, 0x4c, 0xc4, 0xe2, 0x20, 0x9d, 0x2c, 0xce, 0x33, 0x5d, 0x66, 0xb4, 0xcc, 0xb3,
	0xc7, 0xa6, 0x7c, 0xd0, 0xd6, 0xd0, 0x29, 0x85, 0x31, 0xf4, 0x6b, 0x2d, 0x4c, 0xd4, 0xf3, 0x34,
	0x77, 0xf4, 0x93, 0xd2, 0x52, 0x90, 0x16, 0xe4, 0x9d, 0xe9, 0x2d, 0x8c, 0xba, 0x2c, 0xce, 0x60,
	0xdc, 0xa5, 0xcb, 0x2a, 0x62, 0x31, 0x4b, 0x43, 0x3a, 0x08, 0x18, 0xc1, 0xc8, 0xba, 0xa8, 0x5c,
	0x45, 0xbd, 0x98, 0xa5, 0x63, 0xda, 0x8d, 0x89, 0x82, 0x70, 0xb7, 0xf4, 0x98, 0x62, 0x27, 0x14,
	0xce, 0x20, 0x10, 0x85, 0xf5, 0xd9, 0xc9, 0x02, 0x5c, 0x93, 0xbc, 0xb0, 0xae, 0x87, 0x93, 0xf1,
	0x1a, 0xc2, 0xf5, 0xbb, 0x52, 0x2b, 0xd3, 0x94, 0x51, 0xe0, 0x11, 0xff, 0xb5, 0xfb, 0x4e, 0x73,
	0xdc, 0x1e, 0x48, 0xbe, 0x18, 0x0c, 0xdb, 0x30, 0xa6, 0x30, 0x14, 0xaf, 0xcd, 0xe6, 0x6d, 0x77,
	0x90, 0x8b, 0xc3, 0xe2, 0x2c, 0x77, 0x06, 0x75, 0xfe, 0xf4, 0x19, 0x06, 0x5e, 0x40, 0x84, 0xfe,
	0xaa, 0xb0, 0x85, 0xef, 0xc7, 0xc9, 0xbf, 0x31, 0x01, 0x2e, 0x54, 0xb5, 0x5e, 0x4b, 0x21, 0xab,
	0x8d, 0x6d, 0xef, 0xc5, 0xe9, 0x44, 0xc3, 0x4b, 0x18, 0x54, 0x9f, 0xb6, 0x2e, 0x7c, 0x3f, 0x4e,
	0xed, 0x90, 0x5c, 0xc1, 0xe4, 0xa8, 0x24, 0x4e, 0x21, 0xfc, 0xa8, 0x8c, 0x29, 0x5e, 0xaa, 0xb6,
	0x11, 0xa7, 0xfd, 0x7c, 0xc7, 0x7f, 0xb6, 0x73, 0xf6, 0xbb, 0x9d, 0xb3, 0xbf, 0xed, 0x9c, 0xfd,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xda, 0x7d, 0x24, 0xe9, 0x01, 0x00, 0x00,
}

func (m *RPC) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RPC) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RPC) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Rpcs) > 0 {
		for iNdEx := len(m.Rpcs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rpcs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Subscriptions) > 0 {
		for iNdEx := len(m.Subscriptions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Subscriptions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RPC_SubOpts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RPC_SubOpts) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RPC_SubOpts) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Topicid != nil {
		i -= len(*m.Topicid)
		copy(dAtA[i:], *m.Topicid)
		i = encodeVarintRpc(dAtA, i, uint64(len(*m.Topicid)))
		i--
		dAtA[i] = 0x12
	}
	if m.Subscribe != nil {
		i--
		if *m.Subscribe {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TopicRpc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TopicRpc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TopicRpc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Floodsub != nil {
		{
			size, err := m.Floodsub.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Cat != nil {
		{
			size, err := m.Cat.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Topicid != nil {
		i -= len(*m.Topicid)
		copy(dAtA[i:], *m.Topicid)
		i = encodeVarintRpc(dAtA, i, uint64(len(*m.Topicid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CatRpc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CatRpc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CatRpc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Chunks) > 0 {
		for iNdEx := len(m.Chunks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Chunks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CatRpc_Chunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CatRpc_Chunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CatRpc_Chunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Extra != nil {
		i -= len(m.Extra)
		copy(dAtA[i:], m.Extra)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Extra)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Coefficients) > 0 {
		for iNdEx := len(m.Coefficients) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Coefficients[iNdEx])
			copy(dAtA[i:], m.Coefficients[iNdEx])
			i = encodeVarintRpc(dAtA, i, uint64(len(m.Coefficients[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Data != nil {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintRpc(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FloodsubRpc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FloodsubRpc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FloodsubRpc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Messages) > 0 {
		for iNdEx := len(m.Messages) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Messages[iNdEx])
			copy(dAtA[i:], m.Messages[iNdEx])
			i = encodeVarintRpc(dAtA, i, uint64(len(m.Messages[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintRpc(dAtA []byte, offset int, v uint64) int {
	offset -= sovRpc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RPC) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Subscriptions) > 0 {
		for _, e := range m.Subscriptions {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if len(m.Rpcs) > 0 {
		for _, e := range m.Rpcs {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RPC_SubOpts) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Subscribe != nil {
		n += 2
	}
	if m.Topicid != nil {
		l = len(*m.Topicid)
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TopicRpc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Topicid != nil {
		l = len(*m.Topicid)
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Cat != nil {
		l = m.Cat.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.Floodsub != nil {
		l = m.Floodsub.Size()
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CatRpc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Chunks) > 0 {
		for _, e := range m.Chunks {
			l = e.Size()
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CatRpc_Chunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovRpc(uint64(l))
	}
	if len(m.Coefficients) > 0 {
		for _, b := range m.Coefficients {
			l = len(b)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.Extra != nil {
		l = len(m.Extra)
		n += 1 + l + sovRpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FloodsubRpc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for _, b := range m.Messages {
			l = len(b)
			n += 1 + l + sovRpc(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRpc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRpc(x uint64) (n int) {
	return sovRpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RPC) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: RPC: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RPC: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscriptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subscriptions = append(m.Subscriptions, &RPC_SubOpts{})
			if err := m.Subscriptions[len(m.Subscriptions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rpcs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rpcs = append(m.Rpcs, &TopicRpc{})
			if err := m.Rpcs[len(m.Rpcs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RPC_SubOpts) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: SubOpts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubOpts: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscribe", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Subscribe = &b
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topicid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Topicid = &s
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TopicRpc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: TopicRpc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TopicRpc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topicid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Topicid = &s
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cat", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Cat == nil {
				m.Cat = &CatRpc{}
			}
			if err := m.Cat.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Floodsub", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Floodsub == nil {
				m.Floodsub = &FloodsubRpc{}
			}
			if err := m.Floodsub.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CatRpc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: CatRpc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CatRpc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
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
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chunks = append(m.Chunks, &CatRpc_Chunk{})
			if err := m.Chunks[len(m.Chunks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CatRpc_Chunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: Chunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Chunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coefficients", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coefficients = append(m.Coefficients, make([]byte, postIndex-iNdEx))
			copy(m.Coefficients[len(m.Coefficients)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extra", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Extra = append(m.Extra[:0], dAtA[iNdEx:postIndex]...)
			if m.Extra == nil {
				m.Extra = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FloodsubRpc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpc
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
			return fmt.Errorf("proto: FloodsubRpc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FloodsubRpc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRpc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Messages = append(m.Messages, make([]byte, postIndex-iNdEx))
			copy(m.Messages[len(m.Messages)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRpc
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
					return 0, ErrIntOverflowRpc
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
					return 0, ErrIntOverflowRpc
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
				return 0, ErrInvalidLengthRpc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRpc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRpc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRpc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRpc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRpc = fmt.Errorf("proto: unexpected end of group")
)
