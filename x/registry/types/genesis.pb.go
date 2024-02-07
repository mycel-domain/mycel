// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mycel/registry/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the registry module's genesis state.
type GenesisState struct {
	Params             Params              `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	TopLevelDomains    []TopLevelDomain    `protobuf:"bytes,2,rep,name=top_level_domains,json=topLevelDomains,proto3" json:"top_level_domains"`
	SecondLevelDomains []SecondLevelDomain `protobuf:"bytes,3,rep,name=second_level_domains,json=secondLevelDomains,proto3" json:"second_level_domains"`
	DomainOwnerships   []DomainOwnership   `protobuf:"bytes,4,rep,name=domain_ownerships,json=domainOwnerships,proto3" json:"domain_ownerships"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae1403d0e8816949, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetTopLevelDomains() []TopLevelDomain {
	if m != nil {
		return m.TopLevelDomains
	}
	return nil
}

func (m *GenesisState) GetSecondLevelDomains() []SecondLevelDomain {
	if m != nil {
		return m.SecondLevelDomains
	}
	return nil
}

func (m *GenesisState) GetDomainOwnerships() []DomainOwnership {
	if m != nil {
		return m.DomainOwnerships
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "mycel.registry.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("mycel/registry/v1beta1/genesis.proto", fileDescriptor_ae1403d0e8816949)
}

var fileDescriptor_ae1403d0e8816949 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4b, 0xf3, 0x30,
	0x1c, 0xc7, 0xdb, 0x6d, 0xec, 0x90, 0x3d, 0xf0, 0x48, 0x10, 0x19, 0x3b, 0xc4, 0xa1, 0xa2, 0xf3,
	0xb0, 0xc4, 0xcd, 0xab, 0xa7, 0x31, 0xd0, 0x83, 0xa0, 0x38, 0x11, 0xf4, 0x52, 0xba, 0x2d, 0x74,
	0x85, 0xb5, 0x09, 0x4d, 0x9c, 0xf6, 0x5d, 0xf8, 0xb2, 0x76, 0xdc, 0xd1, 0x93, 0x8c, 0xf6, 0x8d,
	0xc8, 0x92, 0x0c, 0xe9, 0x6c, 0xf0, 0x96, 0xc0, 0xe7, 0xf7, 0xf9, 0xfe, 0xfe, 0x80, 0x93, 0x28,
	0x9d, 0xd0, 0x39, 0x49, 0x68, 0x10, 0x0a, 0x99, 0xa4, 0x64, 0xd1, 0x1b, 0x53, 0xe9, 0xf7, 0x48,
	0x40, 0x63, 0x2a, 0x42, 0x81, 0x79, 0xc2, 0x24, 0x83, 0x07, 0x8a, 0xc2, 0x5b, 0x0a, 0x1b, 0xaa,
	0xb5, 0x1f, 0xb0, 0x80, 0x29, 0x84, 0x6c, 0x5e, 0x9a, 0x6e, 0x75, 0x2d, 0xce, 0x29, 0x8b, 0xfc,
	0x30, 0xf6, 0xd8, 0x5b, 0x4c, 0x13, 0x31, 0x0b, 0xb9, 0xc1, 0x8f, 0x2d, 0x38, 0xf7, 0x13, 0x3f,
	0x32, 0x1d, 0xb4, 0x2e, 0x2c, 0x90, 0xa0, 0x13, 0x16, 0x4f, 0xbd, 0x39, 0x5d, 0xd0, 0xb9, 0xa7,
	0x03, 0xfe, 0xe8, 0x42, 0x32, 0x5e, 0x82, 0x1f, 0xad, 0x2b, 0xe0, 0xdf, 0xb5, 0x1e, 0x7a, 0x24,
	0x7d, 0x49, 0xe1, 0x15, 0xa8, 0xeb, 0x0e, 0x9a, 0x6e, 0xdb, 0xed, 0x34, 0xfa, 0x08, 0x97, 0x2f,
	0x01, 0xdf, 0x2b, 0x6a, 0x50, 0x5b, 0x7e, 0x1d, 0x3a, 0x0f, 0xa6, 0x06, 0x3e, 0x81, 0xff, 0x92,
	0xf1, 0xdb, 0x4d, 0xce, 0x50, 0xc5, 0x88, 0x66, 0xa5, 0x5d, 0xed, 0x34, 0xfa, 0xa7, 0x36, 0xcd,
	0x63, 0x01, 0x37, 0xba, 0x5d, 0x09, 0xf4, 0x00, 0xd4, 0x23, 0x17, 0xd4, 0x55, 0xa5, 0x3e, 0xb7,
	0xa9, 0x47, 0xbb, 0x15, 0xc6, 0x5e, 0xa2, 0x82, 0xcf, 0x60, 0x4f, 0xef, 0xe5, 0x6e, 0x7b, 0x26,
	0xd1, 0xac, 0x29, 0xfd, 0x99, 0x4d, 0x3f, 0x2c, 0xf2, 0x46, 0xfe, 0x4b, 0x33, 0xb8, 0x59, 0x66,
	0xc8, 0x5d, 0x65, 0xc8, 0x5d, 0x67, 0xc8, 0xfd, 0xc8, 0x91, 0xb3, 0xca, 0x91, 0xf3, 0x99, 0x23,
	0xe7, 0x05, 0x07, 0xa1, 0x9c, 0xbd, 0x8e, 0xf1, 0x84, 0x45, 0x44, 0x85, 0x74, 0x75, 0xb1, 0xfe,
	0x90, 0xf7, 0x9f, 0x2b, 0xca, 0x94, 0x53, 0x31, 0xae, 0xab, 0x9b, 0x5d, 0x7e, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x34, 0xdf, 0xf4, 0x1c, 0xbe, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DomainOwnerships) > 0 {
		for iNdEx := len(m.DomainOwnerships) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DomainOwnerships[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.SecondLevelDomains) > 0 {
		for iNdEx := len(m.SecondLevelDomains) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SecondLevelDomains[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.TopLevelDomains) > 0 {
		for iNdEx := len(m.TopLevelDomains) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TopLevelDomains[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.TopLevelDomains) > 0 {
		for _, e := range m.TopLevelDomains {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SecondLevelDomains) > 0 {
		for _, e := range m.SecondLevelDomains {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DomainOwnerships) > 0 {
		for _, e := range m.DomainOwnerships {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TopLevelDomains", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TopLevelDomains = append(m.TopLevelDomains, TopLevelDomain{})
			if err := m.TopLevelDomains[len(m.TopLevelDomains)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecondLevelDomains", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecondLevelDomains = append(m.SecondLevelDomains, SecondLevelDomain{})
			if err := m.SecondLevelDomains[len(m.SecondLevelDomains)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DomainOwnerships", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DomainOwnerships = append(m.DomainOwnerships, DomainOwnership{})
			if err := m.DomainOwnerships[len(m.DomainOwnerships)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
