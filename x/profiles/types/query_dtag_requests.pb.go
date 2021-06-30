// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/query_dtag_requests.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryIncomingDTagTransferRequestsRequest is the request type for the
// Query/IncomingDTagTransferRequests RPC endpoint
type QueryIncomingDTagTransferRequestsRequest struct {
	// Receiver represents the address of the user to which query the incoming
	// requests for
	Receiver string `protobuf:"bytes,1,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// Pagination defines an optional pagination for the request
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryIncomingDTagTransferRequestsRequest) Reset() {
	*m = QueryIncomingDTagTransferRequestsRequest{}
}
func (m *QueryIncomingDTagTransferRequestsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIncomingDTagTransferRequestsRequest) ProtoMessage()    {}
func (*QueryIncomingDTagTransferRequestsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2be2e173228b210d, []int{0}
}
func (m *QueryIncomingDTagTransferRequestsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIncomingDTagTransferRequestsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIncomingDTagTransferRequestsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIncomingDTagTransferRequestsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIncomingDTagTransferRequestsRequest.Merge(m, src)
}
func (m *QueryIncomingDTagTransferRequestsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIncomingDTagTransferRequestsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIncomingDTagTransferRequestsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIncomingDTagTransferRequestsRequest proto.InternalMessageInfo

// QueryIncomingDTagTransferRequestsResponse is the response type for the
// Query/IncomingDTagTransferRequests RPC method.
type QueryIncomingDTagTransferRequestsResponse struct {
	// Requests represent the list of all the DTag transfer requests made towards
	// the user
	Requests []DTagTransferRequest `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests"`
	// Pagination defines the pagination response
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryIncomingDTagTransferRequestsResponse) Reset() {
	*m = QueryIncomingDTagTransferRequestsResponse{}
}
func (m *QueryIncomingDTagTransferRequestsResponse) String() string {
	return proto.CompactTextString(m)
}
func (*QueryIncomingDTagTransferRequestsResponse) ProtoMessage() {}
func (*QueryIncomingDTagTransferRequestsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2be2e173228b210d, []int{1}
}
func (m *QueryIncomingDTagTransferRequestsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIncomingDTagTransferRequestsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIncomingDTagTransferRequestsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIncomingDTagTransferRequestsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIncomingDTagTransferRequestsResponse.Merge(m, src)
}
func (m *QueryIncomingDTagTransferRequestsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryIncomingDTagTransferRequestsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIncomingDTagTransferRequestsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIncomingDTagTransferRequestsResponse proto.InternalMessageInfo

func (m *QueryIncomingDTagTransferRequestsResponse) GetRequests() []DTagTransferRequest {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *QueryIncomingDTagTransferRequestsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryIncomingDTagTransferRequestsRequest)(nil), "desmos.profiles.v1beta1.QueryIncomingDTagTransferRequestsRequest")
	proto.RegisterType((*QueryIncomingDTagTransferRequestsResponse)(nil), "desmos.profiles.v1beta1.QueryIncomingDTagTransferRequestsResponse")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/query_dtag_requests.proto", fileDescriptor_2be2e173228b210d)
}

var fileDescriptor_2be2e173228b210d = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3f, 0x8f, 0xd3, 0x30,
	0x18, 0xc6, 0x63, 0x40, 0xa8, 0xa4, 0x5b, 0x84, 0x44, 0x1b, 0xa1, 0xb4, 0xea, 0x00, 0x01, 0x81,
	0xad, 0x96, 0x8d, 0xb1, 0x42, 0xfc, 0x59, 0x10, 0x44, 0x9d, 0x58, 0x2a, 0x27, 0x7d, 0x6b, 0x22,
	0x25, 0x76, 0x6a, 0x3b, 0x15, 0xfd, 0x06, 0x8c, 0x37, 0xde, 0xd8, 0x0f, 0x73, 0x43, 0xc7, 0x8e,
	0x37, 0x9d, 0x4e, 0xed, 0x72, 0x1f, 0xe3, 0x94, 0xd8, 0xfd, 0x33, 0xb4, 0xba, 0x9b, 0xe2, 0x37,
	0xef, 0xfb, 0x3c, 0xf9, 0x3d, 0x79, 0xed, 0xf6, 0x27, 0xa0, 0x72, 0xa1, 0x48, 0x21, 0xc5, 0x34,
	0xcd, 0x40, 0x91, 0x79, 0x3f, 0x06, 0x4d, 0xfb, 0x64, 0x56, 0x82, 0x5c, 0x8c, 0x27, 0x9a, 0xb2,
	0xb1, 0x84, 0x59, 0x09, 0x4a, 0x2b, 0x5c, 0x48, 0xa1, 0x85, 0xf7, 0xca, 0x48, 0xf0, 0x4e, 0x82,
	0xad, 0xc4, 0x7f, 0xc9, 0x04, 0x13, 0xf5, 0x0c, 0xa9, 0x4e, 0x66, 0xdc, 0x7f, 0xcd, 0x84, 0x60,
	0x19, 0x10, 0x5a, 0xa4, 0x84, 0x72, 0x2e, 0x34, 0xd5, 0xa9, 0xe0, 0xd6, 0xcc, 0x6f, 0xdb, 0x6e,
	0x5d, 0xc5, 0xe5, 0x94, 0x50, 0xbe, 0xb0, 0xad, 0xc1, 0x39, 0xb4, 0x5c, 0x4c, 0x20, 0x53, 0xa7,
	0xd8, 0xfc, 0x76, 0x22, 0x2a, 0xcd, 0xd8, 0x50, 0x98, 0xc2, 0xb6, 0xde, 0x9b, 0x8a, 0xc4, 0x54,
	0x81, 0x49, 0xb7, 0x37, 0x2c, 0x28, 0x4b, 0x79, 0x8d, 0x65, 0x66, 0x7b, 0x97, 0xc8, 0x0d, 0x7f,
	0x57, 0x23, 0x3f, 0x78, 0x22, 0xf2, 0x94, 0xb3, 0x2f, 0x23, 0xca, 0x46, 0x92, 0x72, 0x35, 0x05,
	0x19, 0xd9, 0x4f, 0xda, 0xa7, 0xe7, 0xbb, 0x0d, 0x09, 0x09, 0xa4, 0x73, 0x90, 0x2d, 0xd4, 0x45,
	0xe1, 0x8b, 0x68, 0x5f, 0x7b, 0x5f, 0x5d, 0xf7, 0x60, 0xde, 0x7a, 0xd2, 0x45, 0x61, 0x73, 0xf0,
	0x06, 0x5b, 0xae, 0x8a, 0x04, 0xd7, 0x24, 0xbb, 0x5f, 0x88, 0x7f, 0x51, 0x06, 0xd6, 0x37, 0x3a,
	0x52, 0x7e, 0x6e, 0xfc, 0x5f, 0x76, 0x9c, 0xbb, 0x65, 0xc7, 0xe9, 0x5d, 0x21, 0xf7, 0xdd, 0x23,
	0xd0, 0x54, 0x21, 0xb8, 0x02, 0xef, 0x67, 0xc5, 0x66, 0xde, 0xb5, 0x50, 0xf7, 0x69, 0xd8, 0x1c,
	0x7c, 0xc0, 0x67, 0xd6, 0x87, 0x4f, 0x18, 0x0d, 0x9f, 0xad, 0x6e, 0x3a, 0x4e, 0xb4, 0xf7, 0xf0,
	0xbe, 0x9d, 0xc8, 0xf3, 0xf6, 0xc1, 0x3c, 0x06, 0xe6, 0x38, 0xd0, 0xf0, 0xfb, 0x6a, 0x13, 0xa0,
	0xf5, 0x26, 0x40, 0xb7, 0x9b, 0x00, 0x5d, 0x6c, 0x03, 0x67, 0xbd, 0x0d, 0x9c, 0xeb, 0x6d, 0xe0,
	0xfc, 0xc1, 0x2c, 0xd5, 0x7f, 0xcb, 0x18, 0x27, 0x22, 0x27, 0x06, 0xf5, 0x63, 0x46, 0x63, 0x65,
	0xcf, 0xe4, 0xdf, 0xe1, 0x3e, 0xe8, 0x45, 0x01, 0x2a, 0x7e, 0x5e, 0xaf, 0xec, 0xd3, 0x7d, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x6f, 0x26, 0x2a, 0xd8, 0xca, 0x02, 0x00, 0x00,
}

func (m *QueryIncomingDTagTransferRequestsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIncomingDTagTransferRequestsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIncomingDTagTransferRequestsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryDtagRequests(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintQueryDtagRequests(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryIncomingDTagTransferRequestsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIncomingDTagTransferRequestsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIncomingDTagTransferRequestsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryDtagRequests(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Requests) > 0 {
		for iNdEx := len(m.Requests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Requests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryDtagRequests(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryDtagRequests(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryDtagRequests(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryIncomingDTagTransferRequestsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovQueryDtagRequests(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQueryDtagRequests(uint64(l))
	}
	return n
}

func (m *QueryIncomingDTagTransferRequestsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Requests) > 0 {
		for _, e := range m.Requests {
			l = e.Size()
			n += 1 + l + sovQueryDtagRequests(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQueryDtagRequests(uint64(l))
	}
	return n
}

func sovQueryDtagRequests(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryDtagRequests(x uint64) (n int) {
	return sovQueryDtagRequests(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryIncomingDTagTransferRequestsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryDtagRequests
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
			return fmt.Errorf("proto: QueryIncomingDTagTransferRequestsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIncomingDTagTransferRequestsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDtagRequests
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
				return ErrInvalidLengthQueryDtagRequests
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDtagRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDtagRequests
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
				return ErrInvalidLengthQueryDtagRequests
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDtagRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryDtagRequests(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryDtagRequests
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
func (m *QueryIncomingDTagTransferRequestsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryDtagRequests
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
			return fmt.Errorf("proto: QueryIncomingDTagTransferRequestsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIncomingDTagTransferRequestsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDtagRequests
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
				return ErrInvalidLengthQueryDtagRequests
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDtagRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requests = append(m.Requests, DTagTransferRequest{})
			if err := m.Requests[len(m.Requests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryDtagRequests
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
				return ErrInvalidLengthQueryDtagRequests
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryDtagRequests
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryDtagRequests(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryDtagRequests
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
func skipQueryDtagRequests(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryDtagRequests
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
					return 0, ErrIntOverflowQueryDtagRequests
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
					return 0, ErrIntOverflowQueryDtagRequests
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
				return 0, ErrInvalidLengthQueryDtagRequests
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryDtagRequests
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryDtagRequests
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryDtagRequests        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryDtagRequests          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryDtagRequests = fmt.Errorf("proto: unexpected end of group")
)
