// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mycel/epochs/epoch_info.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type EpochInfo struct {
	Identifier              string        `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	StartTime               time.Time     `protobuf:"bytes,2,opt,name=startTime,proto3,stdtime" json:"startTime" yaml:"start_time"`
	Duration                time.Duration `protobuf:"bytes,3,opt,name=duration,proto3,stdduration" json:"duration,omitempty" yaml:"duration"`
	CurrentEpoch            uint64        `protobuf:"varint,4,opt,name=currentEpoch,proto3" json:"currentEpoch,omitempty"`
	CurrentEpochStartTime   time.Time     `protobuf:"bytes,5,opt,name=currentEpochStartTime,proto3,stdtime" json:"currentEpochStartTime" yaml:"current_epoch_start_time"`
	EpochCountingStarted    bool          `protobuf:"varint,6,opt,name=epochCountingStarted,proto3" json:"epochCountingStarted,omitempty"`
	CurrentEpochStartHeight uint64        `protobuf:"varint,7,opt,name=currentEpochStartHeight,proto3" json:"currentEpochStartHeight,omitempty"`
}

func (m *EpochInfo) Reset()         { *m = EpochInfo{} }
func (m *EpochInfo) String() string { return proto.CompactTextString(m) }
func (*EpochInfo) ProtoMessage()    {}
func (*EpochInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9f0c1175734b392, []int{0}
}
func (m *EpochInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochInfo.Merge(m, src)
}
func (m *EpochInfo) XXX_Size() int {
	return m.Size()
}
func (m *EpochInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EpochInfo proto.InternalMessageInfo

func (m *EpochInfo) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *EpochInfo) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *EpochInfo) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *EpochInfo) GetCurrentEpoch() uint64 {
	if m != nil {
		return m.CurrentEpoch
	}
	return 0
}

func (m *EpochInfo) GetCurrentEpochStartTime() time.Time {
	if m != nil {
		return m.CurrentEpochStartTime
	}
	return time.Time{}
}

func (m *EpochInfo) GetEpochCountingStarted() bool {
	if m != nil {
		return m.EpochCountingStarted
	}
	return false
}

func (m *EpochInfo) GetCurrentEpochStartHeight() uint64 {
	if m != nil {
		return m.CurrentEpochStartHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*EpochInfo)(nil), "mycel.epochs.EpochInfo")
}

func init() { proto.RegisterFile("mycel/epochs/epoch_info.proto", fileDescriptor_e9f0c1175734b392) }

var fileDescriptor_e9f0c1175734b392 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6f, 0xda, 0x30,
	0x1c, 0xc5, 0xe3, 0xc1, 0x18, 0x78, 0x48, 0xd3, 0x2c, 0xa6, 0x65, 0x48, 0x38, 0x51, 0x4e, 0x91,
	0x36, 0x25, 0x1a, 0xbb, 0x4c, 0x3b, 0xb2, 0x4d, 0xda, 0xae, 0x29, 0x52, 0xa5, 0x5e, 0x50, 0x00,
	0x27, 0x58, 0x22, 0x71, 0x94, 0x38, 0x52, 0x73, 0xe8, 0x77, 0xe0, 0xd8, 0x8f, 0xc4, 0x91, 0x63,
	0x4f, 0x69, 0x05, 0xb7, 0x1e, 0xf9, 0x02, 0xad, 0x62, 0x27, 0x94, 0x02, 0x55, 0x4f, 0x89, 0xfd,
	0x7b, 0x7e, 0xef, 0xff, 0x12, 0xc3, 0x5e, 0x90, 0x4d, 0xc8, 0xdc, 0x26, 0x11, 0x9b, 0xcc, 0x12,
	0xf9, 0x18, 0xd1, 0xd0, 0x63, 0x56, 0x14, 0x33, 0xce, 0x50, 0x5b, 0x60, 0x4b, 0xe2, 0x6e, 0xc7,
	0x67, 0x3e, 0x13, 0xc0, 0x2e, 0xde, 0xa4, 0xa6, 0x8b, 0x7d, 0xc6, 0xfc, 0x39, 0xb1, 0xc5, 0x6a,
	0x9c, 0x7a, 0xf6, 0x34, 0x8d, 0x5d, 0x4e, 0x59, 0x58, 0x72, 0xed, 0x90, 0x73, 0x1a, 0x90, 0x84,
	0xbb, 0x41, 0x24, 0x05, 0xc6, 0x43, 0x0d, 0xb6, 0xfe, 0x16, 0x09, 0xff, 0x43, 0x8f, 0x21, 0x0c,
	0x21, 0x9d, 0x92, 0x90, 0x53, 0x8f, 0x92, 0x58, 0x05, 0x3a, 0x30, 0x5b, 0xce, 0xde, 0x0e, 0x3a,
	0x87, 0xad, 0x84, 0xbb, 0x31, 0x1f, 0xd2, 0x80, 0xa8, 0x6f, 0x74, 0x60, 0xbe, 0xef, 0x77, 0x2d,
	0x19, 0x61, 0x55, 0x11, 0xd6, 0xb0, 0x8a, 0x18, 0xf4, 0x96, 0xb9, 0xa6, 0x6c, 0x73, 0xed, 0x63,
	0xe6, 0x06, 0xf3, 0x5f, 0x86, 0x38, 0x3a, 0x2a, 0x26, 0x30, 0x16, 0xb7, 0x1a, 0x70, 0x9e, 0xbc,
	0xd0, 0x0c, 0x36, 0xab, 0xc9, 0xd5, 0x9a, 0xf0, 0xfd, 0x72, 0xe4, 0xfb, 0xa7, 0x14, 0x0c, 0xbe,
	0x17, 0xb6, 0xf7, 0xb9, 0x86, 0xaa, 0x23, 0xdf, 0x58, 0x40, 0x39, 0x09, 0x22, 0x9e, 0x6d, 0x73,
	0xed, 0x83, 0x0c, 0xab, 0x98, 0x71, 0x5d, 0x44, 0xed, 0xdc, 0x91, 0x01, 0xdb, 0x93, 0x34, 0x8e,
	0x49, 0xc8, 0x45, 0x6d, 0xb5, 0xae, 0x03, 0xb3, 0xee, 0x3c, 0xdb, 0x43, 0x57, 0xf0, 0xd3, 0xfe,
	0xfa, 0x6c, 0x57, 0xf9, 0xed, 0xab, 0x95, 0xbf, 0x96, 0x95, 0x35, 0x39, 0x45, 0x69, 0x33, 0x92,
	0x3f, 0xf7, 0xf0, 0x03, 0x9c, 0x4e, 0x41, 0x7d, 0xd8, 0x11, 0xfa, 0xdf, 0x2c, 0x0d, 0x39, 0x0d,
	0x7d, 0x41, 0xc8, 0x54, 0x6d, 0xe8, 0xc0, 0x6c, 0x3a, 0x27, 0x19, 0xfa, 0x09, 0x3f, 0x1f, 0x99,
	0xfd, 0x23, 0xd4, 0x9f, 0x71, 0xf5, 0x9d, 0x68, 0xf8, 0x12, 0x1e, 0x58, 0xcb, 0x35, 0x06, 0xab,
	0x35, 0x06, 0x77, 0x6b, 0x0c, 0x16, 0x1b, 0xac, 0xac, 0x36, 0x58, 0xb9, 0xd9, 0x60, 0xe5, 0xa2,
	0x23, 0xef, 0xe7, 0x65, 0x75, 0x43, 0x79, 0x16, 0x91, 0x64, 0xdc, 0x10, 0xad, 0x7f, 0x3c, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xe5, 0xa2, 0x9d, 0xd7, 0xbe, 0x02, 0x00, 0x00,
}

func (m *EpochInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CurrentEpochStartHeight != 0 {
		i = encodeVarintEpochInfo(dAtA, i, uint64(m.CurrentEpochStartHeight))
		i--
		dAtA[i] = 0x38
	}
	if m.EpochCountingStarted {
		i--
		if m.EpochCountingStarted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CurrentEpochStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CurrentEpochStartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintEpochInfo(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if m.CurrentEpoch != 0 {
		i = encodeVarintEpochInfo(dAtA, i, uint64(m.CurrentEpoch))
		i--
		dAtA[i] = 0x20
	}
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintEpochInfo(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintEpochInfo(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	if len(m.Identifier) > 0 {
		i -= len(m.Identifier)
		copy(dAtA[i:], m.Identifier)
		i = encodeVarintEpochInfo(dAtA, i, uint64(len(m.Identifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEpochInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovEpochInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EpochInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identifier)
	if l > 0 {
		n += 1 + l + sovEpochInfo(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovEpochInfo(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovEpochInfo(uint64(l))
	if m.CurrentEpoch != 0 {
		n += 1 + sovEpochInfo(uint64(m.CurrentEpoch))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CurrentEpochStartTime)
	n += 1 + l + sovEpochInfo(uint64(l))
	if m.EpochCountingStarted {
		n += 2
	}
	if m.CurrentEpochStartHeight != 0 {
		n += 1 + sovEpochInfo(uint64(m.CurrentEpochStartHeight))
	}
	return n
}

func sovEpochInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEpochInfo(x uint64) (n int) {
	return sovEpochInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EpochInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEpochInfo
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
			return fmt.Errorf("proto: EpochInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochInfo
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
				return ErrInvalidLengthEpochInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEpochInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochInfo
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
				return ErrInvalidLengthEpochInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpochInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochInfo
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
				return ErrInvalidLengthEpochInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpochInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpoch", wireType)
			}
			m.CurrentEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochInfo
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
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochInfo
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
				return ErrInvalidLengthEpochInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpochInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CurrentEpochStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochCountingStarted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochInfo
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
			m.EpochCountingStarted = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochStartHeight", wireType)
			}
			m.CurrentEpochStartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpochInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpochStartHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEpochInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEpochInfo
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
func skipEpochInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEpochInfo
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
					return 0, ErrIntOverflowEpochInfo
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
					return 0, ErrIntOverflowEpochInfo
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
				return 0, ErrInvalidLengthEpochInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEpochInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEpochInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEpochInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEpochInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEpochInfo = fmt.Errorf("proto: unexpected end of group")
)