// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mycel/furnace/burn_amount.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type BurnAmount struct {
	Index                 uint64     `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	BurnStarted           bool       `protobuf:"varint,2,opt,name=burnStarted,proto3" json:"burnStarted,omitempty"`
	TotalEpochs           uint64     `protobuf:"varint,3,opt,name=totalEpochs,proto3" json:"totalEpochs,omitempty"`
	CurrentEpoch          uint64     `protobuf:"varint,4,opt,name=currentEpoch,proto3" json:"currentEpoch,omitempty"`
	TotalBurnAmount       types.Coin `protobuf:"bytes,5,opt,name=totalBurnAmount,proto3" json:"totalBurnAmount"`
	CumulativeBurntAmount types.Coin `protobuf:"bytes,6,opt,name=cumulativeBurntAmount,proto3" json:"cumulativeBurntAmount"`
}

func (m *BurnAmount) Reset()         { *m = BurnAmount{} }
func (m *BurnAmount) String() string { return proto.CompactTextString(m) }
func (*BurnAmount) ProtoMessage()    {}
func (*BurnAmount) Descriptor() ([]byte, []int) {
	return fileDescriptor_c51e14c665b742e5, []int{0}
}
func (m *BurnAmount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BurnAmount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BurnAmount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BurnAmount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BurnAmount.Merge(m, src)
}
func (m *BurnAmount) XXX_Size() int {
	return m.Size()
}
func (m *BurnAmount) XXX_DiscardUnknown() {
	xxx_messageInfo_BurnAmount.DiscardUnknown(m)
}

var xxx_messageInfo_BurnAmount proto.InternalMessageInfo

func (m *BurnAmount) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *BurnAmount) GetBurnStarted() bool {
	if m != nil {
		return m.BurnStarted
	}
	return false
}

func (m *BurnAmount) GetTotalEpochs() uint64 {
	if m != nil {
		return m.TotalEpochs
	}
	return 0
}

func (m *BurnAmount) GetCurrentEpoch() uint64 {
	if m != nil {
		return m.CurrentEpoch
	}
	return 0
}

func (m *BurnAmount) GetTotalBurnAmount() types.Coin {
	if m != nil {
		return m.TotalBurnAmount
	}
	return types.Coin{}
}

func (m *BurnAmount) GetCumulativeBurntAmount() types.Coin {
	if m != nil {
		return m.CumulativeBurntAmount
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*BurnAmount)(nil), "mycel.furnace.BurnAmount")
}

func init() { proto.RegisterFile("mycel/furnace/burn_amount.proto", fileDescriptor_c51e14c665b742e5) }

var fileDescriptor_c51e14c665b742e5 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4e, 0x3a, 0x31,
	0x10, 0xc6, 0xb7, 0xfc, 0x81, 0xfc, 0x53, 0x34, 0x26, 0x1b, 0x4c, 0x56, 0x0e, 0x65, 0xc3, 0x89,
	0x0b, 0x6d, 0xd0, 0x27, 0x10, 0x63, 0x8c, 0x57, 0x8c, 0x17, 0x2f, 0xa6, 0x5b, 0x2a, 0x6c, 0xc2,
	0x76, 0x48, 0x77, 0x4a, 0xe0, 0x2d, 0x7c, 0x10, 0x1f, 0x84, 0x23, 0x47, 0x4f, 0xc6, 0xc0, 0x8b,
	0x98, 0xed, 0x6e, 0x14, 0x8d, 0x07, 0x6f, 0x9d, 0x6f, 0x7e, 0xdf, 0xb4, 0x5f, 0x87, 0x76, 0xb3,
	0xb5, 0xd2, 0x73, 0xf1, 0xe4, 0xac, 0x91, 0x4a, 0x8b, 0xc4, 0x59, 0xf3, 0x28, 0x33, 0x70, 0x06,
	0xf9, 0xc2, 0x02, 0x42, 0x78, 0xec, 0x01, 0x5e, 0x01, 0x9d, 0xf6, 0x14, 0xa6, 0xe0, 0x3b, 0xa2,
	0x38, 0x95, 0x50, 0x87, 0x29, 0xc8, 0x33, 0xc8, 0x45, 0x22, 0x73, 0x2d, 0x96, 0xc3, 0x44, 0xa3,
	0x1c, 0x0a, 0x05, 0xa9, 0x29, 0xfb, 0xbd, 0x97, 0x1a, 0xa5, 0x23, 0x67, 0xcd, 0xa5, 0x9f, 0x1c,
	0xb6, 0x69, 0x23, 0x35, 0x13, 0xbd, 0x8a, 0x48, 0x4c, 0xfa, 0xf5, 0x71, 0x59, 0x84, 0x31, 0x6d,
	0x15, 0xd7, 0xdf, 0xa1, 0xb4, 0xa8, 0x27, 0x51, 0x2d, 0x26, 0xfd, 0xff, 0xe3, 0x43, 0xa9, 0x20,
	0x10, 0x50, 0xce, 0xaf, 0x17, 0xa0, 0x66, 0x79, 0xf4, 0xcf, 0xbb, 0x0f, 0xa5, 0xb0, 0x47, 0x8f,
	0x94, 0xb3, 0x56, 0x1b, 0xf4, 0x42, 0x54, 0xf7, 0xc8, 0x37, 0x2d, 0xbc, 0xa5, 0x27, 0xde, 0xf2,
	0xf5, 0xa0, 0xa8, 0x11, 0x93, 0x7e, 0xeb, 0xfc, 0x8c, 0x97, 0x31, 0x78, 0x11, 0x83, 0x57, 0x31,
	0xf8, 0x15, 0xa4, 0x66, 0x54, 0xdf, 0xbc, 0x75, 0x83, 0xf1, 0x4f, 0x5f, 0x78, 0x4f, 0x4f, 0x95,
	0xcb, 0xdc, 0x5c, 0x62, 0xba, 0xd4, 0x85, 0x8e, 0xd5, 0xc0, 0xe6, 0xdf, 0x06, 0xfe, 0xee, 0x1e,
	0xdd, 0x6c, 0x76, 0x8c, 0x6c, 0x77, 0x8c, 0xbc, 0xef, 0x18, 0x79, 0xde, 0xb3, 0x60, 0xbb, 0x67,
	0xc1, 0xeb, 0x9e, 0x05, 0x0f, 0x83, 0x69, 0x8a, 0x33, 0x97, 0x70, 0x05, 0x99, 0xf0, 0x8b, 0x19,
	0x4c, 0x20, 0x93, 0xa9, 0x29, 0x0b, 0xb1, 0xfa, 0x5c, 0x24, 0xae, 0x17, 0x3a, 0x4f, 0x9a, 0xfe,
	0xfb, 0x2f, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x34, 0xf3, 0x88, 0x10, 0xe6, 0x01, 0x00, 0x00,
}

func (m *BurnAmount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BurnAmount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BurnAmount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CumulativeBurntAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBurnAmount(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.TotalBurnAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBurnAmount(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.CurrentEpoch != 0 {
		i = encodeVarintBurnAmount(dAtA, i, uint64(m.CurrentEpoch))
		i--
		dAtA[i] = 0x20
	}
	if m.TotalEpochs != 0 {
		i = encodeVarintBurnAmount(dAtA, i, uint64(m.TotalEpochs))
		i--
		dAtA[i] = 0x18
	}
	if m.BurnStarted {
		i--
		if m.BurnStarted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Index != 0 {
		i = encodeVarintBurnAmount(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBurnAmount(dAtA []byte, offset int, v uint64) int {
	offset -= sovBurnAmount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BurnAmount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovBurnAmount(uint64(m.Index))
	}
	if m.BurnStarted {
		n += 2
	}
	if m.TotalEpochs != 0 {
		n += 1 + sovBurnAmount(uint64(m.TotalEpochs))
	}
	if m.CurrentEpoch != 0 {
		n += 1 + sovBurnAmount(uint64(m.CurrentEpoch))
	}
	l = m.TotalBurnAmount.Size()
	n += 1 + l + sovBurnAmount(uint64(l))
	l = m.CumulativeBurntAmount.Size()
	n += 1 + l + sovBurnAmount(uint64(l))
	return n
}

func sovBurnAmount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBurnAmount(x uint64) (n int) {
	return sovBurnAmount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BurnAmount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBurnAmount
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
			return fmt.Errorf("proto: BurnAmount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BurnAmount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBurnAmount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnStarted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBurnAmount
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
			m.BurnStarted = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalEpochs", wireType)
			}
			m.TotalEpochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBurnAmount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalEpochs |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpoch", wireType)
			}
			m.CurrentEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBurnAmount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalBurnAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBurnAmount
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
				return ErrInvalidLengthBurnAmount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBurnAmount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalBurnAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CumulativeBurntAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBurnAmount
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
				return ErrInvalidLengthBurnAmount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBurnAmount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CumulativeBurntAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBurnAmount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBurnAmount
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
func skipBurnAmount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBurnAmount
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
					return 0, ErrIntOverflowBurnAmount
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
					return 0, ErrIntOverflowBurnAmount
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
				return 0, ErrInvalidLengthBurnAmount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBurnAmount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBurnAmount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBurnAmount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBurnAmount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBurnAmount = fmt.Errorf("proto: unexpected end of group")
)
