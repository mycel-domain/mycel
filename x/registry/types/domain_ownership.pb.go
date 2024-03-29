// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mycel/registry/v1beta1/domain_ownership.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type OwnedDomain struct {
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Parent string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (m *OwnedDomain) Reset()         { *m = OwnedDomain{} }
func (m *OwnedDomain) String() string { return proto.CompactTextString(m) }
func (*OwnedDomain) ProtoMessage()    {}
func (*OwnedDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_33a14b0e0536070b, []int{0}
}
func (m *OwnedDomain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OwnedDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OwnedDomain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OwnedDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnedDomain.Merge(m, src)
}
func (m *OwnedDomain) XXX_Size() int {
	return m.Size()
}
func (m *OwnedDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnedDomain.DiscardUnknown(m)
}

var xxx_messageInfo_OwnedDomain proto.InternalMessageInfo

func (m *OwnedDomain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OwnedDomain) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

type DomainOwnership struct {
	Owner   string         `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Domains []*OwnedDomain `protobuf:"bytes,2,rep,name=domains,proto3" json:"domains,omitempty"`
}

func (m *DomainOwnership) Reset()         { *m = DomainOwnership{} }
func (m *DomainOwnership) String() string { return proto.CompactTextString(m) }
func (*DomainOwnership) ProtoMessage()    {}
func (*DomainOwnership) Descriptor() ([]byte, []int) {
	return fileDescriptor_33a14b0e0536070b, []int{1}
}
func (m *DomainOwnership) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DomainOwnership) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DomainOwnership.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DomainOwnership) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DomainOwnership.Merge(m, src)
}
func (m *DomainOwnership) XXX_Size() int {
	return m.Size()
}
func (m *DomainOwnership) XXX_DiscardUnknown() {
	xxx_messageInfo_DomainOwnership.DiscardUnknown(m)
}

var xxx_messageInfo_DomainOwnership proto.InternalMessageInfo

func (m *DomainOwnership) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *DomainOwnership) GetDomains() []*OwnedDomain {
	if m != nil {
		return m.Domains
	}
	return nil
}

func init() {
	proto.RegisterType((*OwnedDomain)(nil), "mycel.registry.v1beta1.OwnedDomain")
	proto.RegisterType((*DomainOwnership)(nil), "mycel.registry.v1beta1.DomainOwnership")
}

func init() {
	proto.RegisterFile("mycel/registry/v1beta1/domain_ownership.proto", fileDescriptor_33a14b0e0536070b)
}

var fileDescriptor_33a14b0e0536070b = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcd, 0xad, 0x4c, 0x4e,
	0xcd, 0xd1, 0x2f, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x29, 0xaa, 0xd4, 0x2f, 0x33, 0x4c, 0x4a, 0x2d,
	0x49, 0x34, 0xd4, 0x4f, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0x8b, 0xcf, 0x2f, 0xcf, 0x4b, 0x2d, 0x2a,
	0xce, 0xc8, 0x2c, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x03, 0x2b, 0xd7, 0x83, 0x29,
	0xd7, 0x83, 0x2a, 0x97, 0x92, 0x4c, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0x8e, 0x07, 0xab, 0xd2, 0x87,
	0x70, 0x20, 0x5a, 0x94, 0x2c, 0xb9, 0xb8, 0xfd, 0xcb, 0xf3, 0x52, 0x53, 0x5c, 0xc0, 0x26, 0x0a,
	0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9,
	0x42, 0x62, 0x5c, 0x6c, 0x05, 0x89, 0x45, 0xa9, 0x79, 0x25, 0x12, 0x4c, 0x60, 0x51, 0x28, 0x4f,
	0xa9, 0x81, 0x91, 0x8b, 0x1f, 0xa2, 0xcd, 0x1f, 0xe6, 0x0e, 0x21, 0x3d, 0x2e, 0x56, 0xb0, 0xa3,
	0x20, 0x06, 0x38, 0x49, 0x5c, 0xda, 0xa2, 0x2b, 0x02, 0xb5, 0xcf, 0x31, 0x25, 0xa5, 0x28, 0xb5,
	0xb8, 0x38, 0xb8, 0xa4, 0x28, 0x33, 0x2f, 0x3d, 0x08, 0xa2, 0x4c, 0xc8, 0x96, 0x8b, 0x1d, 0xe2,
	0x97, 0x62, 0x09, 0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x65, 0x3d, 0xec, 0x7e, 0xd0, 0x43, 0x72,
	0x65, 0x10, 0x4c, 0x8f, 0x93, 0xc7, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78,
	0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44,
	0xe9, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x83, 0x4d, 0xd4, 0x85,
	0xe8, 0x81, 0x70, 0xf4, 0x2b, 0x10, 0x61, 0x5a, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x0e,
	0x0e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x96, 0x81, 0x9c, 0xe3, 0x72, 0x01, 0x00, 0x00,
}

func (m *OwnedDomain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OwnedDomain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OwnedDomain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Parent) > 0 {
		i -= len(m.Parent)
		copy(dAtA[i:], m.Parent)
		i = encodeVarintDomainOwnership(dAtA, i, uint64(len(m.Parent)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDomainOwnership(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DomainOwnership) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DomainOwnership) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DomainOwnership) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Domains) > 0 {
		for iNdEx := len(m.Domains) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Domains[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDomainOwnership(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintDomainOwnership(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDomainOwnership(dAtA []byte, offset int, v uint64) int {
	offset -= sovDomainOwnership(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OwnedDomain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDomainOwnership(uint64(l))
	}
	l = len(m.Parent)
	if l > 0 {
		n += 1 + l + sovDomainOwnership(uint64(l))
	}
	return n
}

func (m *DomainOwnership) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovDomainOwnership(uint64(l))
	}
	if len(m.Domains) > 0 {
		for _, e := range m.Domains {
			l = e.Size()
			n += 1 + l + sovDomainOwnership(uint64(l))
		}
	}
	return n
}

func sovDomainOwnership(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDomainOwnership(x uint64) (n int) {
	return sovDomainOwnership(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OwnedDomain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDomainOwnership
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
			return fmt.Errorf("proto: OwnedDomain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OwnedDomain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomainOwnership
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
				return ErrInvalidLengthDomainOwnership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomainOwnership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomainOwnership
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
				return ErrInvalidLengthDomainOwnership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomainOwnership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDomainOwnership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDomainOwnership
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
func (m *DomainOwnership) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDomainOwnership
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
			return fmt.Errorf("proto: DomainOwnership: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DomainOwnership: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomainOwnership
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
				return ErrInvalidLengthDomainOwnership
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDomainOwnership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domains", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDomainOwnership
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
				return ErrInvalidLengthDomainOwnership
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDomainOwnership
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Domains = append(m.Domains, &OwnedDomain{})
			if err := m.Domains[len(m.Domains)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDomainOwnership(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDomainOwnership
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
func skipDomainOwnership(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDomainOwnership
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
					return 0, ErrIntOverflowDomainOwnership
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
					return 0, ErrIntOverflowDomainOwnership
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
				return 0, ErrInvalidLengthDomainOwnership
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDomainOwnership
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDomainOwnership
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDomainOwnership        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDomainOwnership          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDomainOwnership = fmt.Errorf("proto: unexpected end of group")
)
