// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mycel/furnace/v1beta1/burn_amount.proto

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
	BurnStarted           bool       `protobuf:"varint,2,opt,name=burn_started,json=burnStarted,proto3" json:"burn_started,omitempty"`
	TotalEpochs           uint64     `protobuf:"varint,3,opt,name=total_epochs,json=totalEpochs,proto3" json:"total_epochs,omitempty"`
	CurrentEpoch          uint64     `protobuf:"varint,4,opt,name=current_epoch,json=currentEpoch,proto3" json:"current_epoch,omitempty"`
	TotalBurnAmount       types.Coin `protobuf:"bytes,5,opt,name=total_burn_amount,json=totalBurnAmount,proto3" json:"total_burn_amount"`
	CumulativeBurntAmount types.Coin `protobuf:"bytes,6,opt,name=cumulative_burnt_amount,json=cumulativeBurntAmount,proto3" json:"cumulative_burnt_amount"`
}

func (m *BurnAmount) Reset()         { *m = BurnAmount{} }
func (m *BurnAmount) String() string { return proto.CompactTextString(m) }
func (*BurnAmount) ProtoMessage()    {}
func (*BurnAmount) Descriptor() ([]byte, []int) {
	return fileDescriptor_b99a364454e8a693, []int{0}
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
	proto.RegisterType((*BurnAmount)(nil), "mycel.furnace.v1beta1.BurnAmount")
}

func init() {
	proto.RegisterFile("mycel/furnace/v1beta1/burn_amount.proto", fileDescriptor_b99a364454e8a693)
}

var fileDescriptor_b99a364454e8a693 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4e, 0x32, 0x31,
	0x14, 0x85, 0xa7, 0xfc, 0x40, 0xfe, 0x14, 0x8c, 0x71, 0x02, 0x71, 0x64, 0x51, 0x51, 0x17, 0xb2,
	0xa1, 0x0d, 0xfa, 0x04, 0x62, 0x8c, 0x0b, 0x77, 0xb8, 0x30, 0x71, 0x43, 0x3a, 0xa5, 0xc2, 0x24,
	0x4c, 0x4b, 0x3a, 0x2d, 0x81, 0xb7, 0xf0, 0x5d, 0x7c, 0x09, 0x96, 0x2c, 0x5d, 0x19, 0x03, 0x2f,
	0x62, 0xe6, 0x76, 0x10, 0x96, 0xee, 0xee, 0x3d, 0xf9, 0xce, 0x69, 0x4f, 0x2e, 0xbe, 0x4e, 0x97,
	0x42, 0x4e, 0xd9, 0x9b, 0x33, 0x8a, 0x0b, 0xc9, 0xe6, 0xbd, 0x58, 0x5a, 0xde, 0x63, 0xb1, 0x33,
	0x6a, 0xc8, 0x53, 0xed, 0x94, 0xa5, 0x33, 0xa3, 0xad, 0x0e, 0x9b, 0x00, 0xd2, 0x02, 0xa4, 0x05,
	0xd8, 0x22, 0x42, 0x67, 0xa9, 0xce, 0x58, 0xcc, 0xb3, 0xbd, 0x5b, 0xe8, 0x44, 0x79, 0x5b, 0xab,
	0x31, 0xd6, 0x63, 0x0d, 0x23, 0xcb, 0x27, 0xaf, 0x5e, 0x7e, 0x94, 0x30, 0xee, 0x3b, 0xa3, 0xee,
	0xe0, 0x85, 0xb0, 0x81, 0x2b, 0x89, 0x1a, 0xc9, 0x45, 0x84, 0xda, 0xa8, 0x53, 0x1e, 0xf8, 0x25,
	0xbc, 0xc0, 0x75, 0xf8, 0x46, 0x66, 0xb9, 0xb1, 0x72, 0x14, 0x95, 0xda, 0xa8, 0xf3, 0x7f, 0x50,
	0xcb, 0xb5, 0x67, 0x2f, 0xe5, 0x88, 0xd5, 0x96, 0x4f, 0x87, 0x72, 0xa6, 0xc5, 0x24, 0x8b, 0xfe,
	0x81, 0xbf, 0x06, 0xda, 0x03, 0x48, 0xe1, 0x15, 0x3e, 0x12, 0xce, 0x18, 0xa9, 0xac, 0x87, 0xa2,
	0x32, 0x30, 0xf5, 0x42, 0x04, 0x2a, 0x7c, 0xc2, 0x27, 0x3e, 0xe7, 0xa0, 0x77, 0x54, 0x69, 0xa3,
	0x4e, 0xed, 0xe6, 0x8c, 0xfa, 0x86, 0x34, 0x6f, 0xb8, 0xab, 0x4d, 0xef, 0x75, 0xa2, 0xfa, 0xe5,
	0xd5, 0xd7, 0x79, 0x30, 0x38, 0x06, 0xe7, 0x41, 0x9b, 0x17, 0x7c, 0x2a, 0x5c, 0xea, 0xa6, 0xdc,
	0x26, 0x73, 0x09, 0x89, 0x76, 0x17, 0x59, 0xfd, 0x5b, 0x64, 0x73, 0xef, 0xcf, 0x73, 0xad, 0x0f,
	0xee, 0x3f, 0xae, 0x36, 0x04, 0xad, 0x37, 0x04, 0x7d, 0x6f, 0x08, 0x7a, 0xdf, 0x92, 0x60, 0xbd,
	0x25, 0xc1, 0xe7, 0x96, 0x04, 0xaf, 0xdd, 0x71, 0x62, 0x27, 0x2e, 0xa6, 0x42, 0xa7, 0x0c, 0xee,
	0xd4, 0x1d, 0xe9, 0x94, 0x27, 0xca, 0x2f, 0x6c, 0xf1, 0x7b, 0x5f, 0xbb, 0x9c, 0xc9, 0x2c, 0xae,
	0xc2, 0x15, 0x6e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xdb, 0xb1, 0x3b, 0xfd, 0x01, 0x00,
	0x00,
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
