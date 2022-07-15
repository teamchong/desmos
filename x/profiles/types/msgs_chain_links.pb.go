// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v3/msgs_chain_links.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// MsgLinkChainAccount represents a message to link an account to a profile.
type MsgLinkChainAccount struct {
	// ChainAddress contains the details of the external chain address to be
	// linked
	ChainAddress *types.Any `protobuf:"bytes,1,opt,name=chain_address,json=chainAddress,proto3" json:"chain_address,omitempty" yaml:"source_address"`
	// Proof contains the proof of ownership of the external chain address
	Proof Proof `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof" yaml:"source_proof"`
	// ChainConfig contains the configuration of the external chain
	ChainConfig ChainConfig `protobuf:"bytes,3,opt,name=chain_config,json=chainConfig,proto3" json:"chain_config" yaml:"source_chain_config"`
	// Signer represents the Desmos address associated with the
	// profile to which link the external account
	Signer string `protobuf:"bytes,4,opt,name=signer,proto3" json:"signer,omitempty" yaml:"signer"`
}

func (m *MsgLinkChainAccount) Reset()         { *m = MsgLinkChainAccount{} }
func (m *MsgLinkChainAccount) String() string { return proto.CompactTextString(m) }
func (*MsgLinkChainAccount) ProtoMessage()    {}
func (*MsgLinkChainAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_52cd1f5b825f2f4e, []int{0}
}
func (m *MsgLinkChainAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLinkChainAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLinkChainAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLinkChainAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLinkChainAccount.Merge(m, src)
}
func (m *MsgLinkChainAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgLinkChainAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLinkChainAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLinkChainAccount proto.InternalMessageInfo

func (m *MsgLinkChainAccount) GetChainAddress() *types.Any {
	if m != nil {
		return m.ChainAddress
	}
	return nil
}

func (m *MsgLinkChainAccount) GetProof() Proof {
	if m != nil {
		return m.Proof
	}
	return Proof{}
}

func (m *MsgLinkChainAccount) GetChainConfig() ChainConfig {
	if m != nil {
		return m.ChainConfig
	}
	return ChainConfig{}
}

func (m *MsgLinkChainAccount) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

// MsgLinkChainAccountResponse defines the Msg/LinkAccount response type.
type MsgLinkChainAccountResponse struct {
}

func (m *MsgLinkChainAccountResponse) Reset()         { *m = MsgLinkChainAccountResponse{} }
func (m *MsgLinkChainAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgLinkChainAccountResponse) ProtoMessage()    {}
func (*MsgLinkChainAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_52cd1f5b825f2f4e, []int{1}
}
func (m *MsgLinkChainAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLinkChainAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLinkChainAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLinkChainAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLinkChainAccountResponse.Merge(m, src)
}
func (m *MsgLinkChainAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgLinkChainAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLinkChainAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLinkChainAccountResponse proto.InternalMessageInfo

// MsgUnlinkChainAccount represents a message to unlink an account from a
// profile.
type MsgUnlinkChainAccount struct {
	// Owner represents the Desmos profile from which to remove the link
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	// ChainName represents the name of the chain to which the link to remove is
	// associated
	ChainName string `protobuf:"bytes,2,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty" yaml:"chain_name"`
	// Target represents the external address to be removed
	Target string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty" yaml:"target"`
}

func (m *MsgUnlinkChainAccount) Reset()         { *m = MsgUnlinkChainAccount{} }
func (m *MsgUnlinkChainAccount) String() string { return proto.CompactTextString(m) }
func (*MsgUnlinkChainAccount) ProtoMessage()    {}
func (*MsgUnlinkChainAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_52cd1f5b825f2f4e, []int{2}
}
func (m *MsgUnlinkChainAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnlinkChainAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnlinkChainAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnlinkChainAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnlinkChainAccount.Merge(m, src)
}
func (m *MsgUnlinkChainAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnlinkChainAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnlinkChainAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnlinkChainAccount proto.InternalMessageInfo

func (m *MsgUnlinkChainAccount) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgUnlinkChainAccount) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *MsgUnlinkChainAccount) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

// MsgUnlinkChainAccountResponse defines the Msg/UnlinkAccount response type.
type MsgUnlinkChainAccountResponse struct {
}

func (m *MsgUnlinkChainAccountResponse) Reset()         { *m = MsgUnlinkChainAccountResponse{} }
func (m *MsgUnlinkChainAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUnlinkChainAccountResponse) ProtoMessage()    {}
func (*MsgUnlinkChainAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_52cd1f5b825f2f4e, []int{3}
}
func (m *MsgUnlinkChainAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnlinkChainAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnlinkChainAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnlinkChainAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnlinkChainAccountResponse.Merge(m, src)
}
func (m *MsgUnlinkChainAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnlinkChainAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnlinkChainAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnlinkChainAccountResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgLinkChainAccount)(nil), "desmos.profiles.v3.MsgLinkChainAccount")
	proto.RegisterType((*MsgLinkChainAccountResponse)(nil), "desmos.profiles.v3.MsgLinkChainAccountResponse")
	proto.RegisterType((*MsgUnlinkChainAccount)(nil), "desmos.profiles.v3.MsgUnlinkChainAccount")
	proto.RegisterType((*MsgUnlinkChainAccountResponse)(nil), "desmos.profiles.v3.MsgUnlinkChainAccountResponse")
}

func init() {
	proto.RegisterFile("desmos/profiles/v3/msgs_chain_links.proto", fileDescriptor_52cd1f5b825f2f4e)
}

var fileDescriptor_52cd1f5b825f2f4e = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0xb5, 0xd2, 0x26, 0xe0, 0x75, 0x02, 0x8d, 0x12, 0x83, 0xe3, 0x10, 0x29, 0xec, 0xa1, 0x38,
	0x94, 0x68, 0x69, 0x9d, 0x53, 0x6f, 0x56, 0x7a, 0x6b, 0x52, 0x8a, 0xa0, 0x97, 0x5e, 0xcc, 0x7a,
	0xbd, 0xde, 0x88, 0x48, 0x3b, 0x42, 0x2b, 0xbb, 0xf5, 0x5f, 0xf4, 0x13, 0xfa, 0x11, 0xfd, 0x88,
	0xd0, 0x53, 0x8e, 0x85, 0x82, 0x28, 0xf6, 0x1f, 0xe8, 0x0b, 0x8a, 0x76, 0x25, 0x5c, 0x37, 0xba,
	0xcd, 0xcc, 0x7b, 0xf3, 0xe6, 0xcd, 0xec, 0xa2, 0x8b, 0x29, 0x57, 0x31, 0x28, 0x92, 0xa4, 0x30,
	0x0b, 0x23, 0xae, 0xc8, 0x62, 0x48, 0x62, 0x25, 0xd4, 0x98, 0xdd, 0xd1, 0x50, 0x8e, 0xa3, 0x50,
	0xde, 0x2b, 0x2f, 0x49, 0x21, 0x03, 0xdb, 0x36, 0x54, 0xaf, 0xa6, 0x7a, 0x8b, 0x61, 0xff, 0x58,
	0x80, 0x00, 0x0d, 0x93, 0x32, 0x32, 0xcc, 0xfe, 0x89, 0x00, 0x10, 0x11, 0x27, 0x3a, 0x9b, 0xcc,
	0x67, 0x84, 0xca, 0x65, 0x0d, 0x31, 0x28, 0x45, 0xc6, 0xa6, 0xc7, 0x24, 0x15, 0xf4, 0xaa, 0xc9,
	0x0a, 0x4c, 0x79, 0xd4, 0x60, 0x06, 0xff, 0xde, 0x41, 0x47, 0xb7, 0x4a, 0xdc, 0x84, 0xf2, 0xfe,
	0xba, 0x04, 0x47, 0x8c, 0xc1, 0x5c, 0x66, 0x36, 0x43, 0x07, 0x86, 0x4c, 0xa7, 0xd3, 0x94, 0x2b,
	0xd5, 0xb3, 0xce, 0xad, 0x41, 0xe7, 0xcd, 0xb1, 0x67, 0x2c, 0x79, 0xb5, 0x25, 0x6f, 0x24, 0x97,
	0xfe, 0xa0, 0xc8, 0xdd, 0xee, 0x92, 0xc6, 0xd1, 0x5b, 0xac, 0x60, 0x9e, 0x32, 0x5e, 0x77, 0xe1,
	0x9f, 0x3f, 0x2e, 0x3b, 0x23, 0x13, 0xbf, 0xa3, 0x19, 0x0d, 0xf6, 0xb5, 0x68, 0x55, 0xb1, 0x6f,
	0xd0, 0x6e, 0x92, 0x02, 0xcc, 0x7a, 0x3b, 0x5a, 0xfc, 0xc4, 0x7b, 0x7a, 0x19, 0xef, 0x63, 0x49,
	0xf0, 0x4f, 0x1f, 0x72, 0xb7, 0x55, 0xe4, 0xee, 0xd1, 0xd6, 0x14, 0xdd, 0x8c, 0x03, 0x23, 0x62,
	0xcf, 0x90, 0x51, 0x1f, 0x33, 0x90, 0xb3, 0x50, 0xf4, 0x9e, 0x69, 0x51, 0xb7, 0x49, 0x54, 0xaf,
	0x7a, 0xad, 0x69, 0x3e, 0xae, 0xa4, 0xfb, 0x5b, 0xd2, 0xff, 0x2a, 0xe1, 0xa0, 0xc3, 0x36, 0x0d,
	0xf6, 0x05, 0xda, 0x53, 0xa1, 0x90, 0x3c, 0xed, 0x3d, 0x3f, 0xb7, 0x06, 0x6d, 0xff, 0xb0, 0xc8,
	0xdd, 0x83, 0xaa, 0x59, 0xd7, 0x71, 0x50, 0x11, 0xf0, 0x19, 0x3a, 0x6d, 0x38, 0x6e, 0xc0, 0x55,
	0x02, 0x52, 0x71, 0xfc, 0xdd, 0x42, 0xdd, 0x5b, 0x25, 0x3e, 0xc9, 0xe8, 0xff, 0xf3, 0xbf, 0x44,
	0xbb, 0xf0, 0xa5, 0x1c, 0x61, 0xe9, 0x11, 0x2f, 0x8a, 0xdc, 0xdd, 0x37, 0x23, 0x74, 0x19, 0x07,
	0x06, 0xb6, 0xaf, 0x10, 0x32, 0x4e, 0x25, 0x8d, 0xb9, 0x3e, 0x63, 0xdb, 0xef, 0x16, 0xb9, 0x7b,
	0x68, 0xc8, 0x1b, 0x0c, 0x07, 0x6d, 0x9d, 0x7c, 0xa0, 0x31, 0x2f, 0x37, 0xc8, 0x68, 0x2a, 0x78,
	0xa6, 0x6f, 0xb4, 0xb5, 0x81, 0xa9, 0xe3, 0xa0, 0x22, 0x60, 0x17, 0x9d, 0x35, 0x3a, 0xac, 0x77,
	0xf0, 0xdf, 0x3f, 0xac, 0x1c, 0xeb, 0x71, 0xe5, 0x58, 0x7f, 0x56, 0x8e, 0xf5, 0x6d, 0xed, 0xb4,
	0x1e, 0xd7, 0x4e, 0xeb, 0xd7, 0xda, 0x69, 0x7d, 0x7e, 0x2d, 0xc2, 0xec, 0x6e, 0x3e, 0xf1, 0x18,
	0xc4, 0xc4, 0xbc, 0xc1, 0x65, 0x44, 0x27, 0xaa, 0x8a, 0xc9, 0xe2, 0x8a, 0x7c, 0xdd, 0xfc, 0xd1,
	0x6c, 0x99, 0x70, 0x35, 0xd9, 0xd3, 0xdf, 0x6a, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0x95, 0xb7,
	0xb7, 0x15, 0x4e, 0x03, 0x00, 0x00,
}

func (m *MsgLinkChainAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLinkChainAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLinkChainAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.ChainConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.ChainAddress != nil {
		{
			size, err := m.ChainAddress.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMsgsChainLinks(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgLinkChainAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLinkChainAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLinkChainAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUnlinkChainAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnlinkChainAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnlinkChainAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Target) > 0 {
		i -= len(m.Target)
		copy(dAtA[i:], m.Target)
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(len(m.Target)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainName) > 0 {
		i -= len(m.ChainName)
		copy(dAtA[i:], m.ChainName)
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(len(m.ChainName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUnlinkChainAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnlinkChainAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnlinkChainAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsgsChainLinks(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgsChainLinks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgLinkChainAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainAddress != nil {
		l = m.ChainAddress.Size()
		n += 1 + l + sovMsgsChainLinks(uint64(l))
	}
	l = m.Proof.Size()
	n += 1 + l + sovMsgsChainLinks(uint64(l))
	l = m.ChainConfig.Size()
	n += 1 + l + sovMsgsChainLinks(uint64(l))
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgsChainLinks(uint64(l))
	}
	return n
}

func (m *MsgLinkChainAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUnlinkChainAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovMsgsChainLinks(uint64(l))
	}
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovMsgsChainLinks(uint64(l))
	}
	l = len(m.Target)
	if l > 0 {
		n += 1 + l + sovMsgsChainLinks(uint64(l))
	}
	return n
}

func (m *MsgUnlinkChainAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsgsChainLinks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgsChainLinks(x uint64) (n int) {
	return sovMsgsChainLinks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgLinkChainAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsChainLinks
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
			return fmt.Errorf("proto: MsgLinkChainAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLinkChainAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
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
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ChainAddress == nil {
				m.ChainAddress = &types.Any{}
			}
			if err := m.ChainAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
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
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
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
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChainConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
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
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsChainLinks
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
func (m *MsgLinkChainAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsChainLinks
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
			return fmt.Errorf("proto: MsgLinkChainAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLinkChainAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsChainLinks
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
func (m *MsgUnlinkChainAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsChainLinks
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
			return fmt.Errorf("proto: MsgUnlinkChainAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnlinkChainAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
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
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
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
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
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
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Target = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsChainLinks
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
func (m *MsgUnlinkChainAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsChainLinks
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
			return fmt.Errorf("proto: MsgUnlinkChainAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnlinkChainAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsChainLinks
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
func skipMsgsChainLinks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgsChainLinks
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
					return 0, ErrIntOverflowMsgsChainLinks
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
					return 0, ErrIntOverflowMsgsChainLinks
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
				return 0, ErrInvalidLengthMsgsChainLinks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgsChainLinks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgsChainLinks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgsChainLinks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgsChainLinks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgsChainLinks = fmt.Errorf("proto: unexpected end of group")
)
