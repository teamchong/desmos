// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/msgs_permissioned_contract.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
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

// MsgSavePermissionedContractReference represents the message used to save
// a permissioned contract reference.
type MsgSavePermissionedContractReference struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"nickname"`
	Admin   string `protobuf:"bytes,2,opt,name=admin,proto3" json:"admin,omitempty" yaml:"admin"`
	Message []byte `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty" yaml:"message"`
}

func (m *MsgSavePermissionedContractReference) Reset()         { *m = MsgSavePermissionedContractReference{} }
func (m *MsgSavePermissionedContractReference) String() string { return proto.CompactTextString(m) }
func (*MsgSavePermissionedContractReference) ProtoMessage()    {}
func (*MsgSavePermissionedContractReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_faef539d4afed14c, []int{0}
}
func (m *MsgSavePermissionedContractReference) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSavePermissionedContractReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSavePermissionedContractReference.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSavePermissionedContractReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSavePermissionedContractReference.Merge(m, src)
}
func (m *MsgSavePermissionedContractReference) XXX_Size() int {
	return m.Size()
}
func (m *MsgSavePermissionedContractReference) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSavePermissionedContractReference.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSavePermissionedContractReference proto.InternalMessageInfo

// MsgSavePermissionedContractReferenceResponse defines the Msg/MsgSavePermissionedContractReference response
// type.
type MsgSavePermissionedContractReferenceResponse struct {
}

func (m *MsgSavePermissionedContractReferenceResponse) Reset() {
	*m = MsgSavePermissionedContractReferenceResponse{}
}
func (m *MsgSavePermissionedContractReferenceResponse) String() string {
	return proto.CompactTextString(m)
}
func (*MsgSavePermissionedContractReferenceResponse) ProtoMessage() {}
func (*MsgSavePermissionedContractReferenceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_faef539d4afed14c, []int{1}
}
func (m *MsgSavePermissionedContractReferenceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSavePermissionedContractReferenceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSavePermissionedContractReferenceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSavePermissionedContractReferenceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSavePermissionedContractReferenceResponse.Merge(m, src)
}
func (m *MsgSavePermissionedContractReferenceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSavePermissionedContractReferenceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSavePermissionedContractReferenceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSavePermissionedContractReferenceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSavePermissionedContractReference)(nil), "desmos.profiles.v1beta1.MsgSavePermissionedContractReference")
	proto.RegisterType((*MsgSavePermissionedContractReferenceResponse)(nil), "desmos.profiles.v1beta1.MsgSavePermissionedContractReferenceResponse")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/msgs_permissioned_contract.proto", fileDescriptor_faef539d4afed14c)
}

var fileDescriptor_faef539d4afed14c = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0xae, 0xd3, 0x30,
	0x14, 0x86, 0x13, 0x10, 0x14, 0xa2, 0x0a, 0x50, 0x40, 0xa2, 0x74, 0x48, 0xaa, 0x08, 0xa1, 0x0e,
	0x6d, 0xac, 0x96, 0x05, 0x75, 0x2c, 0x23, 0x42, 0x42, 0x61, 0x63, 0x89, 0x9c, 0xe4, 0xd4, 0xb5,
	0x88, 0xed, 0xe0, 0xe3, 0x56, 0xf4, 0x0d, 0x18, 0x79, 0x84, 0xbe, 0x04, 0xef, 0xc0, 0xd8, 0x91,
	0xa9, 0x42, 0xed, 0x72, 0xe7, 0x3e, 0xc1, 0x55, 0xe3, 0x44, 0xf7, 0x2e, 0x57, 0xb7, 0x9b, 0xcf,
	0xf9, 0xbf, 0xcf, 0x3a, 0xfa, 0xbd, 0x0f, 0x05, 0xa0, 0x50, 0x48, 0x2a, 0xad, 0x16, 0xbc, 0x04,
	0x24, 0xeb, 0x49, 0x06, 0x86, 0x4e, 0x88, 0x40, 0x86, 0x69, 0x05, 0x5a, 0x70, 0x44, 0xae, 0x24,
	0x14, 0x69, 0xae, 0xa4, 0xd1, 0x34, 0x37, 0x71, 0xa5, 0x95, 0x51, 0xfe, 0x6b, 0x6b, 0xc6, 0xad,
	0x19, 0x37, 0x66, 0xff, 0x15, 0x53, 0x4c, 0xd5, 0x0c, 0x39, 0xbf, 0x2c, 0xde, 0x7f, 0xc3, 0x94,
	0x62, 0x25, 0x90, 0x7a, 0xca, 0x56, 0x0b, 0x42, 0xe5, 0xa6, 0x8d, 0x72, 0x75, 0xfe, 0x29, 0xb5,
	0x8e, 0x1d, 0x9a, 0x68, 0x74, 0xe7, 0x79, 0xaa, 0x80, 0xb2, 0x56, 0xce, 0xfb, 0x86, 0x9e, 0xde,
	0x43, 0x6b, 0x28, 0xa9, 0xe1, 0x4a, 0xe2, 0x92, 0x57, 0x78, 0xa1, 0x53, 0x18, 0xca, 0x52, 0x0d,
	0x3f, 0x56, 0x80, 0xa6, 0x71, 0xa2, 0x3f, 0xae, 0xf7, 0xf6, 0x33, 0xb2, 0xaf, 0x74, 0x0d, 0x5f,
	0x6e, 0x35, 0xf4, 0xb1, 0x29, 0x28, 0x81, 0x05, 0x68, 0x90, 0x39, 0xf8, 0x63, 0xaf, 0x43, 0x8b,
	0x42, 0x03, 0x62, 0xcf, 0x1d, 0xb8, 0xc3, 0xa7, 0xf3, 0x97, 0xa7, 0x7d, 0xf8, 0x7c, 0x43, 0x45,
	0x39, 0x8b, 0x24, 0xcf, 0xbf, 0x4b, 0x2a, 0x20, 0x4a, 0x5a, 0xc6, 0x7f, 0xe7, 0x3d, 0xa2, 0x85,
	0xe0, 0xb2, 0xf7, 0xa0, 0x86, 0x5f, 0x9c, 0xf6, 0x61, 0xd7, 0xc2, 0xf5, 0x3a, 0x4a, 0x6c, 0xec,
	0x8f, 0xbc, 0x8e, 0x00, 0x44, 0xca, 0xa0, 0xf7, 0x70, 0xe0, 0x0e, 0xbb, 0x73, 0xff, 0xb4, 0x0f,
	0x9f, 0x59, 0xb2, 0x09, 0xa2, 0xa4, 0x45, 0x66, 0x4f, 0x7e, 0x6d, 0x43, 0xe7, 0x6a, 0x1b, 0x3a,
	0x51, 0xec, 0x8d, 0x2e, 0x39, 0x3b, 0x01, 0xac, 0x94, 0x44, 0x98, 0x7f, 0xfa, 0x7b, 0x08, 0xdc,
	0xdd, 0x21, 0x70, 0xff, 0x1f, 0x02, 0xf7, 0xf7, 0x31, 0x70, 0x76, 0xc7, 0xc0, 0xf9, 0x77, 0x0c,
	0x9c, 0x6f, 0x13, 0xc6, 0xcd, 0x72, 0x95, 0xc5, 0xb9, 0x12, 0xc4, 0x16, 0x38, 0x2e, 0x69, 0x86,
	0xcd, 0x9b, 0xac, 0xa7, 0xe4, 0xe7, 0x4d, 0xa3, 0x66, 0x53, 0x01, 0x66, 0x8f, 0xeb, 0xee, 0xde,
	0x5f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xb5, 0x53, 0x1f, 0x72, 0x02, 0x00, 0x00,
}

func (m *MsgSavePermissionedContractReference) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSavePermissionedContractReference) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSavePermissionedContractReference) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintMsgsPermissionedContract(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Admin) > 0 {
		i -= len(m.Admin)
		copy(dAtA[i:], m.Admin)
		i = encodeVarintMsgsPermissionedContract(dAtA, i, uint64(len(m.Admin)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMsgsPermissionedContract(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSavePermissionedContractReferenceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSavePermissionedContractReferenceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSavePermissionedContractReferenceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsgsPermissionedContract(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgsPermissionedContract(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSavePermissionedContractReference) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMsgsPermissionedContract(uint64(l))
	}
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovMsgsPermissionedContract(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovMsgsPermissionedContract(uint64(l))
	}
	return n
}

func (m *MsgSavePermissionedContractReferenceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsgsPermissionedContract(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgsPermissionedContract(x uint64) (n int) {
	return sovMsgsPermissionedContract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSavePermissionedContractReference) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsPermissionedContract
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
			return fmt.Errorf("proto: MsgSavePermissionedContractReference: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSavePermissionedContractReference: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsPermissionedContract
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
				return ErrInvalidLengthMsgsPermissionedContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsPermissionedContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsPermissionedContract
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
				return ErrInvalidLengthMsgsPermissionedContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsPermissionedContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsPermissionedContract
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
				return ErrInvalidLengthMsgsPermissionedContract
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsPermissionedContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = append(m.Message[:0], dAtA[iNdEx:postIndex]...)
			if m.Message == nil {
				m.Message = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsPermissionedContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsPermissionedContract
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
func (m *MsgSavePermissionedContractReferenceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsPermissionedContract
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
			return fmt.Errorf("proto: MsgSavePermissionedContractReferenceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSavePermissionedContractReferenceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsPermissionedContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsPermissionedContract
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
func skipMsgsPermissionedContract(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgsPermissionedContract
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
					return 0, ErrIntOverflowMsgsPermissionedContract
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
					return 0, ErrIntOverflowMsgsPermissionedContract
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
				return 0, ErrInvalidLengthMsgsPermissionedContract
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgsPermissionedContract
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgsPermissionedContract
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgsPermissionedContract        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgsPermissionedContract          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgsPermissionedContract = fmt.Errorf("proto: unexpected end of group")
)
