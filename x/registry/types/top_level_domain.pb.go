// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mycel/registry/v1beta1/top_level_domain.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
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

type TopLevelDomain struct {
	Name                  string                                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ExpirationDate        time.Time                                `protobuf:"bytes,2,opt,name=expirationDate,proto3,stdtime" json:"expirationDate"`
	SubdomainConfig       *SubdomainConfig                         `protobuf:"bytes,3,opt,name=subdomainConfig,proto3" json:"subdomainConfig,omitempty"`
	SubdomainCount        uint64                                   `protobuf:"varint,4,opt,name=subdomainCount,proto3" json:"subdomainCount,omitempty"`
	AccessControl         []*AccessControl                         `protobuf:"bytes,5,rep,name=accessControl,proto3" json:"accessControl,omitempty"`
	TotalWithdrawalAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=totalWithdrawalAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"totalWithdrawalAmount"`
}

func (m *TopLevelDomain) Reset()         { *m = TopLevelDomain{} }
func (m *TopLevelDomain) String() string { return proto.CompactTextString(m) }
func (*TopLevelDomain) ProtoMessage()    {}
func (*TopLevelDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_82805ec1d6b2bfd3, []int{0}
}
func (m *TopLevelDomain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TopLevelDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TopLevelDomain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TopLevelDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopLevelDomain.Merge(m, src)
}
func (m *TopLevelDomain) XXX_Size() int {
	return m.Size()
}
func (m *TopLevelDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_TopLevelDomain.DiscardUnknown(m)
}

var xxx_messageInfo_TopLevelDomain proto.InternalMessageInfo

func (m *TopLevelDomain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TopLevelDomain) GetExpirationDate() time.Time {
	if m != nil {
		return m.ExpirationDate
	}
	return time.Time{}
}

func (m *TopLevelDomain) GetSubdomainConfig() *SubdomainConfig {
	if m != nil {
		return m.SubdomainConfig
	}
	return nil
}

func (m *TopLevelDomain) GetSubdomainCount() uint64 {
	if m != nil {
		return m.SubdomainCount
	}
	return 0
}

func (m *TopLevelDomain) GetAccessControl() []*AccessControl {
	if m != nil {
		return m.AccessControl
	}
	return nil
}

func (m *TopLevelDomain) GetTotalWithdrawalAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TotalWithdrawalAmount
	}
	return nil
}

type TopLevelDomainFee struct {
	TotalFee      github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=totalFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"totalFee"`
	BurnWeight    string                                   `protobuf:"bytes,2,opt,name=burnWeight,proto3" json:"burnWeight,omitempty"`
	FeeToBurn     types.Coin                               `protobuf:"bytes,3,opt,name=feeToBurn,proto3,castvalue=github.com/cosmos/cosmos-sdk/types.Coin" json:"feeToBurn"`
	FeeToTreasury types.Coin                               `protobuf:"bytes,4,opt,name=feeToTreasury,proto3,castvalue=github.com/cosmos/cosmos-sdk/types.Coin" json:"feeToTreasury"`
}

func (m *TopLevelDomainFee) Reset()         { *m = TopLevelDomainFee{} }
func (m *TopLevelDomainFee) String() string { return proto.CompactTextString(m) }
func (*TopLevelDomainFee) ProtoMessage()    {}
func (*TopLevelDomainFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_82805ec1d6b2bfd3, []int{1}
}
func (m *TopLevelDomainFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TopLevelDomainFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TopLevelDomainFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TopLevelDomainFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopLevelDomainFee.Merge(m, src)
}
func (m *TopLevelDomainFee) XXX_Size() int {
	return m.Size()
}
func (m *TopLevelDomainFee) XXX_DiscardUnknown() {
	xxx_messageInfo_TopLevelDomainFee.DiscardUnknown(m)
}

var xxx_messageInfo_TopLevelDomainFee proto.InternalMessageInfo

func (m *TopLevelDomainFee) GetTotalFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TotalFee
	}
	return nil
}

func (m *TopLevelDomainFee) GetBurnWeight() string {
	if m != nil {
		return m.BurnWeight
	}
	return ""
}

func (m *TopLevelDomainFee) GetFeeToBurn() types.Coin {
	if m != nil {
		return m.FeeToBurn
	}
	return types.Coin{}
}

func (m *TopLevelDomainFee) GetFeeToTreasury() types.Coin {
	if m != nil {
		return m.FeeToTreasury
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*TopLevelDomain)(nil), "mycel.registry.v1beta1.TopLevelDomain")
	proto.RegisterType((*TopLevelDomainFee)(nil), "mycel.registry.v1beta1.TopLevelDomainFee")
}

func init() {
	proto.RegisterFile("mycel/registry/v1beta1/top_level_domain.proto", fileDescriptor_82805ec1d6b2bfd3)
}

var fileDescriptor_82805ec1d6b2bfd3 = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0xdb, 0x50, 0x35, 0x5b, 0x35, 0x88, 0x15, 0xa0, 0x90, 0x83, 0x1d, 0x55, 0x82, 0x5a,
	0x42, 0xd9, 0xa5, 0xe5, 0x0b, 0x9a, 0x54, 0x15, 0x12, 0xbd, 0x60, 0x22, 0x55, 0xe2, 0x12, 0xad,
	0x9d, 0x89, 0xb3, 0xc2, 0xf6, 0x5a, 0xbb, 0xeb, 0xd2, 0x1c, 0xb9, 0x72, 0xea, 0x77, 0xf0, 0x11,
	0x9c, 0x7b, 0xec, 0x91, 0x13, 0x41, 0xc9, 0x4f, 0x70, 0x44, 0x5e, 0x3b, 0x69, 0x12, 0x35, 0x88,
	0x43, 0x4f, 0xde, 0x1d, 0xbf, 0x79, 0xf3, 0x66, 0xe6, 0x2d, 0x6a, 0xc7, 0xe3, 0x00, 0x22, 0x2a,
	0x21, 0xe4, 0x4a, 0xcb, 0x31, 0xbd, 0x3c, 0xf2, 0x41, 0xb3, 0x23, 0xaa, 0x45, 0xda, 0x8f, 0xe0,
	0x12, 0xa2, 0xfe, 0x40, 0xc4, 0x8c, 0x27, 0x24, 0x95, 0x42, 0x0b, 0xfc, 0xdc, 0xc0, 0xc9, 0x1c,
	0x4e, 0x4a, 0x78, 0xd3, 0x0e, 0x84, 0x8a, 0x85, 0xa2, 0x3e, 0x53, 0xb0, 0xe0, 0x08, 0xc4, 0x3c,
	0xaf, 0xf9, 0x34, 0x14, 0xa1, 0x30, 0x47, 0x9a, 0x9f, 0xca, 0xa8, 0x13, 0x0a, 0x11, 0x46, 0x40,
	0xcd, 0xcd, 0xcf, 0x86, 0x54, 0xf3, 0x18, 0x94, 0x66, 0x71, 0x5a, 0x02, 0x5e, 0x6f, 0x50, 0xc7,
	0x82, 0x00, 0x94, 0xea, 0x07, 0x22, 0xd1, 0x52, 0x44, 0x25, 0x78, 0x53, 0x2b, 0x2a, 0xf3, 0x8b,
	0x1e, 0x72, 0xfc, 0x90, 0x87, 0x05, 0xfc, 0xe0, 0xc7, 0x36, 0xaa, 0xf7, 0x44, 0x7a, 0x9e, 0x37,
	0x79, 0x6a, 0xfe, 0x63, 0x8c, 0xaa, 0x09, 0x8b, 0xa1, 0x61, 0xb5, 0x2c, 0xb7, 0xe6, 0x99, 0x33,
	0x3e, 0x47, 0x75, 0xb8, 0x4a, 0xb9, 0x64, 0x9a, 0x8b, 0xe4, 0x94, 0x69, 0x68, 0x6c, 0xb5, 0x2c,
	0x77, 0xef, 0xb8, 0x49, 0x0a, 0xf1, 0x64, 0x2e, 0x9e, 0xf4, 0xe6, 0xe2, 0x3b, 0xbb, 0x37, 0xbf,
	0x9c, 0xca, 0xf5, 0xc4, 0xb1, 0xbc, 0xb5, 0x5c, 0xfc, 0x01, 0x3d, 0x5e, 0xc8, 0xe9, 0x1a, 0x35,
	0x8d, 0x6d, 0x43, 0x77, 0x48, 0xee, 0x9f, 0x2c, 0xf9, 0xb8, 0x0a, 0xf7, 0xd6, 0xf3, 0xf1, 0x2b,
	0x54, 0x5f, 0x0a, 0x65, 0x89, 0x6e, 0x54, 0x5b, 0x96, 0x5b, 0xf5, 0xd6, 0xa2, 0xf8, 0x3d, 0xda,
	0x2f, 0xc6, 0xd6, 0x2d, 0xa6, 0xd6, 0x78, 0xd4, 0xda, 0x76, 0xf7, 0x8e, 0x5f, 0x6e, 0x2a, 0x7c,
	0xb2, 0x0c, 0xf6, 0x56, 0x73, 0xf1, 0x57, 0x0b, 0x3d, 0xd3, 0x42, 0xb3, 0xe8, 0x82, 0xeb, 0xd1,
	0x40, 0xb2, 0x2f, 0x2c, 0x3a, 0x89, 0x4d, 0xf1, 0x1d, 0xc3, 0xfa, 0x82, 0x14, 0x86, 0x20, 0xb9,
	0x21, 0x16, 0x94, 0x5d, 0xc1, 0x93, 0xce, 0x9b, 0x7c, 0x38, 0xdf, 0x27, 0x8e, 0x1b, 0x72, 0x3d,
	0xca, 0x7c, 0x12, 0x88, 0x98, 0x96, 0xee, 0x29, 0x3e, 0x6d, 0x35, 0xf8, 0x4c, 0xf5, 0x38, 0x05,
	0x65, 0x12, 0x94, 0x77, 0x7f, 0xa5, 0x83, 0x3f, 0x5b, 0xe8, 0xc9, 0xea, 0x02, 0xcf, 0x00, 0x70,
	0x88, 0x76, 0x0d, 0xfc, 0x0c, 0xf2, 0x3d, 0x3e, 0xb8, 0x96, 0x05, 0x39, 0xb6, 0x11, 0xf2, 0x33,
	0x99, 0x5c, 0x00, 0x0f, 0x47, 0xda, 0x98, 0xa2, 0xe6, 0x2d, 0x45, 0xf0, 0x08, 0xd5, 0x86, 0x00,
	0x3d, 0xd1, 0xc9, 0x64, 0x52, 0x2e, 0xf9, 0x1f, 0x4a, 0x68, 0xae, 0xe4, 0xdb, 0xc4, 0x39, 0xfc,
	0x4f, 0x25, 0xde, 0x1d, 0x39, 0x4e, 0xd1, 0xbe, 0xb9, 0xf4, 0x24, 0x30, 0x95, 0xc9, 0xb1, 0x31,
	0xc0, 0xc3, 0x56, 0x5b, 0x2d, 0xd0, 0x79, 0x77, 0x33, 0xb5, 0xad, 0xdb, 0xa9, 0x6d, 0xfd, 0x9e,
	0xda, 0xd6, 0xf5, 0xcc, 0xae, 0xdc, 0xce, 0xec, 0xca, 0xcf, 0x99, 0x5d, 0xf9, 0x44, 0x96, 0x18,
	0x8d, 0xb1, 0xda, 0x85, 0x07, 0x8b, 0x0b, 0xbd, 0xba, 0x7b, 0x9e, 0x86, 0xdd, 0xdf, 0x31, 0xcf,
	0xe7, 0xed, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0x13, 0x23, 0xb6, 0x88, 0x04, 0x00, 0x00,
}

func (m *TopLevelDomain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TopLevelDomain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TopLevelDomain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TotalWithdrawalAmount) > 0 {
		for iNdEx := len(m.TotalWithdrawalAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TotalWithdrawalAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTopLevelDomain(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.AccessControl) > 0 {
		for iNdEx := len(m.AccessControl) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccessControl[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTopLevelDomain(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.SubdomainCount != 0 {
		i = encodeVarintTopLevelDomain(dAtA, i, uint64(m.SubdomainCount))
		i--
		dAtA[i] = 0x20
	}
	if m.SubdomainConfig != nil {
		{
			size, err := m.SubdomainConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTopLevelDomain(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.ExpirationDate, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.ExpirationDate):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTopLevelDomain(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTopLevelDomain(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TopLevelDomainFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TopLevelDomainFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TopLevelDomainFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.FeeToTreasury.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTopLevelDomain(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.FeeToBurn.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTopLevelDomain(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.BurnWeight) > 0 {
		i -= len(m.BurnWeight)
		copy(dAtA[i:], m.BurnWeight)
		i = encodeVarintTopLevelDomain(dAtA, i, uint64(len(m.BurnWeight)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TotalFee) > 0 {
		for iNdEx := len(m.TotalFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TotalFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTopLevelDomain(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTopLevelDomain(dAtA []byte, offset int, v uint64) int {
	offset -= sovTopLevelDomain(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TopLevelDomain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTopLevelDomain(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.ExpirationDate)
	n += 1 + l + sovTopLevelDomain(uint64(l))
	if m.SubdomainConfig != nil {
		l = m.SubdomainConfig.Size()
		n += 1 + l + sovTopLevelDomain(uint64(l))
	}
	if m.SubdomainCount != 0 {
		n += 1 + sovTopLevelDomain(uint64(m.SubdomainCount))
	}
	if len(m.AccessControl) > 0 {
		for _, e := range m.AccessControl {
			l = e.Size()
			n += 1 + l + sovTopLevelDomain(uint64(l))
		}
	}
	if len(m.TotalWithdrawalAmount) > 0 {
		for _, e := range m.TotalWithdrawalAmount {
			l = e.Size()
			n += 1 + l + sovTopLevelDomain(uint64(l))
		}
	}
	return n
}

func (m *TopLevelDomainFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TotalFee) > 0 {
		for _, e := range m.TotalFee {
			l = e.Size()
			n += 1 + l + sovTopLevelDomain(uint64(l))
		}
	}
	l = len(m.BurnWeight)
	if l > 0 {
		n += 1 + l + sovTopLevelDomain(uint64(l))
	}
	l = m.FeeToBurn.Size()
	n += 1 + l + sovTopLevelDomain(uint64(l))
	l = m.FeeToTreasury.Size()
	n += 1 + l + sovTopLevelDomain(uint64(l))
	return n
}

func sovTopLevelDomain(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTopLevelDomain(x uint64) (n int) {
	return sovTopLevelDomain(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TopLevelDomain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTopLevelDomain
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
			return fmt.Errorf("proto: TopLevelDomain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TopLevelDomain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopLevelDomain
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
				return ErrInvalidLengthTopLevelDomain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopLevelDomain
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
				return ErrInvalidLengthTopLevelDomain
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.ExpirationDate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubdomainConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopLevelDomain
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
				return ErrInvalidLengthTopLevelDomain
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SubdomainConfig == nil {
				m.SubdomainConfig = &SubdomainConfig{}
			}
			if err := m.SubdomainConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubdomainCount", wireType)
			}
			m.SubdomainCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopLevelDomain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubdomainCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessControl", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopLevelDomain
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
				return ErrInvalidLengthTopLevelDomain
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessControl = append(m.AccessControl, &AccessControl{})
			if err := m.AccessControl[len(m.AccessControl)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalWithdrawalAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopLevelDomain
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
				return ErrInvalidLengthTopLevelDomain
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalWithdrawalAmount = append(m.TotalWithdrawalAmount, types.Coin{})
			if err := m.TotalWithdrawalAmount[len(m.TotalWithdrawalAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTopLevelDomain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTopLevelDomain
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
func (m *TopLevelDomainFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTopLevelDomain
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
			return fmt.Errorf("proto: TopLevelDomainFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TopLevelDomainFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopLevelDomain
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
				return ErrInvalidLengthTopLevelDomain
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalFee = append(m.TotalFee, types.Coin{})
			if err := m.TotalFee[len(m.TotalFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopLevelDomain
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
				return ErrInvalidLengthTopLevelDomain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BurnWeight = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeToBurn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopLevelDomain
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
				return ErrInvalidLengthTopLevelDomain
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeeToBurn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeToTreasury", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopLevelDomain
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
				return ErrInvalidLengthTopLevelDomain
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeeToTreasury.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTopLevelDomain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTopLevelDomain
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
func skipTopLevelDomain(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTopLevelDomain
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
					return 0, ErrIntOverflowTopLevelDomain
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
					return 0, ErrIntOverflowTopLevelDomain
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
				return 0, ErrInvalidLengthTopLevelDomain
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTopLevelDomain
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTopLevelDomain
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTopLevelDomain        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTopLevelDomain          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTopLevelDomain = fmt.Errorf("proto: unexpected end of group")
)
