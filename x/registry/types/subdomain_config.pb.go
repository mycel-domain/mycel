// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mycel/registry/subdomain_config.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

type RegistrationPolicyType int32

const (
	RegistrationPolicyType_PRIVATE RegistrationPolicyType = 0
	RegistrationPolicyType_PUBLIC  RegistrationPolicyType = 1
)

var RegistrationPolicyType_name = map[int32]string{
	0: "PRIVATE",
	1: "PUBLIC",
}

var RegistrationPolicyType_value = map[string]int32{
	"PRIVATE": 0,
	"PUBLIC":  1,
}

func (x RegistrationPolicyType) String() string {
	return proto.EnumName(RegistrationPolicyType_name, int32(x))
}

func (RegistrationPolicyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a1e8e4d1516ffeeb, []int{0}
}

type SubdomainRegistrationFees struct {
	FeeByLength []*FeeByLength `protobuf:"bytes,1,rep,name=fee_by_length,json=feeByLength,proto3" json:"fee_by_length,omitempty"`
	FeeByName   []*FeeByName   `protobuf:"bytes,2,rep,name=fee_by_name,json=feeByName,proto3" json:"fee_by_name,omitempty"`
	DefaultFee  *types.Coin    `protobuf:"bytes,3,opt,name=default_fee,json=defaultFee,proto3" json:"default_fee,omitempty"`
}

func (m *SubdomainRegistrationFees) Reset()         { *m = SubdomainRegistrationFees{} }
func (m *SubdomainRegistrationFees) String() string { return proto.CompactTextString(m) }
func (*SubdomainRegistrationFees) ProtoMessage()    {}
func (*SubdomainRegistrationFees) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1e8e4d1516ffeeb, []int{0}
}
func (m *SubdomainRegistrationFees) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubdomainRegistrationFees) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubdomainRegistrationFees.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubdomainRegistrationFees) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubdomainRegistrationFees.Merge(m, src)
}
func (m *SubdomainRegistrationFees) XXX_Size() int {
	return m.Size()
}
func (m *SubdomainRegistrationFees) XXX_DiscardUnknown() {
	xxx_messageInfo_SubdomainRegistrationFees.DiscardUnknown(m)
}

var xxx_messageInfo_SubdomainRegistrationFees proto.InternalMessageInfo

func (m *SubdomainRegistrationFees) GetFeeByLength() []*FeeByLength {
	if m != nil {
		return m.FeeByLength
	}
	return nil
}

func (m *SubdomainRegistrationFees) GetFeeByName() []*FeeByName {
	if m != nil {
		return m.FeeByName
	}
	return nil
}

func (m *SubdomainRegistrationFees) GetDefaultFee() *types.Coin {
	if m != nil {
		return m.DefaultFee
	}
	return nil
}

type FeeByLength struct {
	Length        uint32      `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	IsRegistrable bool        `protobuf:"varint,2,opt,name=is_registrable,json=isRegistrable,proto3" json:"is_registrable,omitempty"`
	Fee           *types.Coin `protobuf:"bytes,3,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *FeeByLength) Reset()         { *m = FeeByLength{} }
func (m *FeeByLength) String() string { return proto.CompactTextString(m) }
func (*FeeByLength) ProtoMessage()    {}
func (*FeeByLength) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1e8e4d1516ffeeb, []int{1}
}
func (m *FeeByLength) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeByLength) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeByLength.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeByLength) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeByLength.Merge(m, src)
}
func (m *FeeByLength) XXX_Size() int {
	return m.Size()
}
func (m *FeeByLength) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeByLength.DiscardUnknown(m)
}

var xxx_messageInfo_FeeByLength proto.InternalMessageInfo

func (m *FeeByLength) GetLength() uint32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *FeeByLength) GetIsRegistrable() bool {
	if m != nil {
		return m.IsRegistrable
	}
	return false
}

func (m *FeeByLength) GetFee() *types.Coin {
	if m != nil {
		return m.Fee
	}
	return nil
}

type FeeByName struct {
	Name          string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IsRegistrable bool        `protobuf:"varint,2,opt,name=is_registrable,json=isRegistrable,proto3" json:"is_registrable,omitempty"`
	Fee           *types.Coin `protobuf:"bytes,3,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *FeeByName) Reset()         { *m = FeeByName{} }
func (m *FeeByName) String() string { return proto.CompactTextString(m) }
func (*FeeByName) ProtoMessage()    {}
func (*FeeByName) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1e8e4d1516ffeeb, []int{2}
}
func (m *FeeByName) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeByName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeByName.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeByName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeByName.Merge(m, src)
}
func (m *FeeByName) XXX_Size() int {
	return m.Size()
}
func (m *FeeByName) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeByName.DiscardUnknown(m)
}

var xxx_messageInfo_FeeByName proto.InternalMessageInfo

func (m *FeeByName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FeeByName) GetIsRegistrable() bool {
	if m != nil {
		return m.IsRegistrable
	}
	return false
}

func (m *FeeByName) GetFee() *types.Coin {
	if m != nil {
		return m.Fee
	}
	return nil
}

type SubdomainConfig struct {
	MaxSubdomainRegistrations uint64                     `protobuf:"varint,1,opt,name=max_subdomain_registrations,json=maxSubdomainRegistrations,proto3" json:"max_subdomain_registrations,omitempty"`
	SubdomainRegistrationFees *SubdomainRegistrationFees `protobuf:"bytes,2,opt,name=subdomain_registration_fees,json=subdomainRegistrationFees,proto3" json:"subdomain_registration_fees,omitempty"`
	IsRegistrable             bool                       `protobuf:"varint,3,opt,name=is_registrable,json=isRegistrable,proto3" json:"is_registrable,omitempty"`
	RegistrableRole           DomainRole                 `protobuf:"varint,4,opt,name=registrable_role,json=registrableRole,proto3,enum=mycel.registry.DomainRole" json:"registrable_role,omitempty"`
	CustomExpirationDate      bool                       `protobuf:"varint,5,opt,name=custom_expiration_date,json=customExpirationDate,proto3" json:"custom_expiration_date,omitempty"`
	RegistrationPolicy        RegistrationPolicyType     `protobuf:"varint,6,opt,name=registration_policy,json=registrationPolicy,proto3,enum=mycel.registry.RegistrationPolicyType" json:"registration_policy,omitempty"`
}

func (m *SubdomainConfig) Reset()         { *m = SubdomainConfig{} }
func (m *SubdomainConfig) String() string { return proto.CompactTextString(m) }
func (*SubdomainConfig) ProtoMessage()    {}
func (*SubdomainConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1e8e4d1516ffeeb, []int{3}
}
func (m *SubdomainConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubdomainConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubdomainConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubdomainConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubdomainConfig.Merge(m, src)
}
func (m *SubdomainConfig) XXX_Size() int {
	return m.Size()
}
func (m *SubdomainConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SubdomainConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SubdomainConfig proto.InternalMessageInfo

func (m *SubdomainConfig) GetMaxSubdomainRegistrations() uint64 {
	if m != nil {
		return m.MaxSubdomainRegistrations
	}
	return 0
}

func (m *SubdomainConfig) GetSubdomainRegistrationFees() *SubdomainRegistrationFees {
	if m != nil {
		return m.SubdomainRegistrationFees
	}
	return nil
}

func (m *SubdomainConfig) GetIsRegistrable() bool {
	if m != nil {
		return m.IsRegistrable
	}
	return false
}

func (m *SubdomainConfig) GetRegistrableRole() DomainRole {
	if m != nil {
		return m.RegistrableRole
	}
	return DomainRole_NO_ROLE
}

func (m *SubdomainConfig) GetCustomExpirationDate() bool {
	if m != nil {
		return m.CustomExpirationDate
	}
	return false
}

func (m *SubdomainConfig) GetRegistrationPolicy() RegistrationPolicyType {
	if m != nil {
		return m.RegistrationPolicy
	}
	return RegistrationPolicyType_PRIVATE
}

func init() {
	proto.RegisterEnum("mycel.registry.RegistrationPolicyType", RegistrationPolicyType_name, RegistrationPolicyType_value)
	proto.RegisterType((*SubdomainRegistrationFees)(nil), "mycel.registry.SubdomainRegistrationFees")
	proto.RegisterType((*FeeByLength)(nil), "mycel.registry.FeeByLength")
	proto.RegisterType((*FeeByName)(nil), "mycel.registry.FeeByName")
	proto.RegisterType((*SubdomainConfig)(nil), "mycel.registry.SubdomainConfig")
}

func init() {
	proto.RegisterFile("mycel/registry/subdomain_config.proto", fileDescriptor_a1e8e4d1516ffeeb)
}

var fileDescriptor_a1e8e4d1516ffeeb = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x6e, 0xec, 0x5a, 0xdd, 0x09, 0xdb, 0x2d, 0xa3, 0x94, 0xb4, 0x85, 0x50, 0x2a, 0x2b, 0x55,
	0x31, 0xa1, 0xd5, 0x1b, 0xbd, 0x50, 0x6c, 0xb7, 0xc5, 0x85, 0x45, 0xca, 0xb8, 0x2a, 0x78, 0x13,
	0x26, 0xe9, 0x49, 0x77, 0x20, 0xc9, 0x94, 0xcc, 0x54, 0x9a, 0xb7, 0xf0, 0x75, 0x7c, 0x03, 0x2f,
	0x17, 0xbc, 0xf1, 0x52, 0xda, 0x17, 0x91, 0x4c, 0xd2, 0x9f, 0xed, 0x76, 0xc1, 0x1b, 0xef, 0x66,
	0xce, 0x7c, 0xe7, 0x7c, 0xdf, 0x77, 0xce, 0x61, 0xd0, 0x49, 0x98, 0x78, 0x10, 0xd8, 0x31, 0x4c,
	0x98, 0x90, 0x71, 0x62, 0x8b, 0x99, 0x3b, 0xe6, 0x21, 0x65, 0x91, 0xe3, 0xf1, 0xc8, 0x67, 0x13,
	0x6b, 0x1a, 0x73, 0xc9, 0x71, 0x59, 0xc1, 0xac, 0x15, 0xac, 0x6e, 0x7a, 0x5c, 0x84, 0x5c, 0xd8,
	0x2e, 0x15, 0x60, 0x7f, 0xeb, 0xb8, 0x20, 0x69, 0xc7, 0xf6, 0x38, 0x8b, 0x32, 0x7c, 0xfd, 0xd1,
	0x4e, 0x59, 0xea, 0x79, 0x20, 0x44, 0x5a, 0x53, 0xc6, 0x3c, 0xc8, 0x40, 0xad, 0x5f, 0x1a, 0xaa,
	0x7d, 0x5c, 0xf1, 0x91, 0x0c, 0x4a, 0x25, 0xe3, 0xd1, 0x10, 0x40, 0xe0, 0xb7, 0xe8, 0xc8, 0x07,
	0x70, 0xdc, 0xc4, 0x09, 0x20, 0x9a, 0xc8, 0x4b, 0x43, 0x6b, 0x16, 0xdb, 0x7a, 0xb7, 0x61, 0x5d,
	0x97, 0x62, 0x0d, 0x01, 0x7a, 0xc9, 0xb9, 0x82, 0x10, 0xdd, 0xdf, 0x5c, 0xf0, 0x2b, 0xa4, 0xe7,
	0x05, 0x22, 0x1a, 0x82, 0x71, 0x47, 0xa5, 0xd7, 0xf6, 0xa6, 0x7f, 0xa0, 0x21, 0x90, 0x43, 0x7f,
	0x75, 0xc4, 0xaf, 0x91, 0x3e, 0x06, 0x9f, 0xce, 0x02, 0xe9, 0xf8, 0x00, 0x46, 0xb1, 0xa9, 0xa9,
	0xd4, 0xcc, 0xb4, 0x95, 0x9a, 0xb6, 0x72, 0xd3, 0x56, 0x9f, 0xb3, 0x88, 0xa0, 0x1c, 0x3d, 0x04,
	0x68, 0x25, 0x48, 0xdf, 0x92, 0x84, 0xab, 0xa8, 0xb4, 0xd6, 0xaf, 0xb5, 0x8f, 0x48, 0x7e, 0xc3,
	0x27, 0xa8, 0xcc, 0x84, 0x93, 0xcb, 0xa0, 0x6e, 0x90, 0x0a, 0xd4, 0xda, 0xf7, 0xc9, 0x11, 0x13,
	0x64, 0x13, 0xc4, 0xcf, 0x50, 0xf1, 0x9f, 0x14, 0xa4, 0xa8, 0x96, 0x40, 0x87, 0x6b, 0x3b, 0x18,
	0xa3, 0x03, 0xe5, 0x3b, 0xa5, 0x3d, 0x24, 0xea, 0xfc, 0x5f, 0x48, 0x7f, 0x14, 0xd1, 0xf1, 0x7a,
	0x8a, 0x7d, 0xb5, 0x34, 0xf8, 0x0d, 0x6a, 0x84, 0x74, 0xee, 0x6c, 0x96, 0x29, 0xde, 0x9a, 0xae,
	0x50, 0x92, 0x0e, 0x48, 0x2d, 0xa4, 0xf3, 0xbd, 0xe3, 0x17, 0x98, 0xa1, 0xc6, 0xfe, 0xdc, 0x74,
	0x1c, 0x42, 0x89, 0xd6, 0xbb, 0x4f, 0x76, 0x47, 0x79, 0xeb, 0x2e, 0x91, 0x9a, 0xb8, 0x75, 0xcd,
	0x6e, 0xb6, 0xa4, 0xb8, 0xaf, 0x25, 0x03, 0x54, 0xd9, 0xc2, 0x38, 0x31, 0x0f, 0xc0, 0x38, 0x68,
	0x6a, 0xed, 0x72, 0xb7, 0xbe, 0x2b, 0xe3, 0x34, 0x23, 0xe2, 0x01, 0x90, 0xe3, 0xad, 0x9c, 0x34,
	0x80, 0x5f, 0xa2, 0xaa, 0x37, 0x13, 0x92, 0x87, 0x0e, 0xcc, 0xa7, 0x2c, 0xf7, 0x34, 0xa6, 0x12,
	0x8c, 0xbb, 0x8a, 0xf5, 0x61, 0xf6, 0x3a, 0x58, 0x3f, 0x9e, 0x52, 0x09, 0xf8, 0x0b, 0x7a, 0x70,
	0xad, 0x09, 0x53, 0x1e, 0x30, 0x2f, 0x31, 0x4a, 0x8a, 0xff, 0xf1, 0x2e, 0xff, 0xb6, 0xc5, 0x91,
	0x42, 0x5e, 0x24, 0x53, 0x20, 0x38, 0xbe, 0x11, 0x7f, 0xda, 0x41, 0xd5, 0xfd, 0x68, 0xac, 0xa3,
	0x7b, 0x23, 0x72, 0xf6, 0xf9, 0xdd, 0xc5, 0xa0, 0x52, 0xc0, 0x08, 0x95, 0x46, 0x9f, 0x7a, 0xe7,
	0x67, 0xfd, 0x8a, 0xd6, 0x7b, 0xff, 0x73, 0x61, 0x6a, 0x57, 0x0b, 0x53, 0xfb, 0xb3, 0x30, 0xb5,
	0xef, 0x4b, 0xb3, 0x70, 0xb5, 0x34, 0x0b, 0xbf, 0x97, 0x66, 0xe1, 0xab, 0x35, 0x61, 0xf2, 0x72,
	0xe6, 0x5a, 0x1e, 0x0f, 0x6d, 0x25, 0xe9, 0x79, 0xd6, 0xf2, 0xec, 0x62, 0xcf, 0x37, 0xbf, 0x81,
	0x4c, 0xa6, 0x20, 0xdc, 0x92, 0xfa, 0x05, 0x5e, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x60, 0x2c,
	0x0c, 0x6f, 0x83, 0x04, 0x00, 0x00,
}

func (m *SubdomainRegistrationFees) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubdomainRegistrationFees) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubdomainRegistrationFees) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DefaultFee != nil {
		{
			size, err := m.DefaultFee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSubdomainConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FeeByName) > 0 {
		for iNdEx := len(m.FeeByName) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeByName[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSubdomainConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.FeeByLength) > 0 {
		for iNdEx := len(m.FeeByLength) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeByLength[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSubdomainConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *FeeByLength) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeByLength) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeByLength) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fee != nil {
		{
			size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSubdomainConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.IsRegistrable {
		i--
		if m.IsRegistrable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Length != 0 {
		i = encodeVarintSubdomainConfig(dAtA, i, uint64(m.Length))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *FeeByName) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeByName) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeByName) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fee != nil {
		{
			size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSubdomainConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.IsRegistrable {
		i--
		if m.IsRegistrable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSubdomainConfig(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SubdomainConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubdomainConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubdomainConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RegistrationPolicy != 0 {
		i = encodeVarintSubdomainConfig(dAtA, i, uint64(m.RegistrationPolicy))
		i--
		dAtA[i] = 0x30
	}
	if m.CustomExpirationDate {
		i--
		if m.CustomExpirationDate {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.RegistrableRole != 0 {
		i = encodeVarintSubdomainConfig(dAtA, i, uint64(m.RegistrableRole))
		i--
		dAtA[i] = 0x20
	}
	if m.IsRegistrable {
		i--
		if m.IsRegistrable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.SubdomainRegistrationFees != nil {
		{
			size, err := m.SubdomainRegistrationFees.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSubdomainConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.MaxSubdomainRegistrations != 0 {
		i = encodeVarintSubdomainConfig(dAtA, i, uint64(m.MaxSubdomainRegistrations))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSubdomainConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovSubdomainConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SubdomainRegistrationFees) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FeeByLength) > 0 {
		for _, e := range m.FeeByLength {
			l = e.Size()
			n += 1 + l + sovSubdomainConfig(uint64(l))
		}
	}
	if len(m.FeeByName) > 0 {
		for _, e := range m.FeeByName {
			l = e.Size()
			n += 1 + l + sovSubdomainConfig(uint64(l))
		}
	}
	if m.DefaultFee != nil {
		l = m.DefaultFee.Size()
		n += 1 + l + sovSubdomainConfig(uint64(l))
	}
	return n
}

func (m *FeeByLength) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Length != 0 {
		n += 1 + sovSubdomainConfig(uint64(m.Length))
	}
	if m.IsRegistrable {
		n += 2
	}
	if m.Fee != nil {
		l = m.Fee.Size()
		n += 1 + l + sovSubdomainConfig(uint64(l))
	}
	return n
}

func (m *FeeByName) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSubdomainConfig(uint64(l))
	}
	if m.IsRegistrable {
		n += 2
	}
	if m.Fee != nil {
		l = m.Fee.Size()
		n += 1 + l + sovSubdomainConfig(uint64(l))
	}
	return n
}

func (m *SubdomainConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxSubdomainRegistrations != 0 {
		n += 1 + sovSubdomainConfig(uint64(m.MaxSubdomainRegistrations))
	}
	if m.SubdomainRegistrationFees != nil {
		l = m.SubdomainRegistrationFees.Size()
		n += 1 + l + sovSubdomainConfig(uint64(l))
	}
	if m.IsRegistrable {
		n += 2
	}
	if m.RegistrableRole != 0 {
		n += 1 + sovSubdomainConfig(uint64(m.RegistrableRole))
	}
	if m.CustomExpirationDate {
		n += 2
	}
	if m.RegistrationPolicy != 0 {
		n += 1 + sovSubdomainConfig(uint64(m.RegistrationPolicy))
	}
	return n
}

func sovSubdomainConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSubdomainConfig(x uint64) (n int) {
	return sovSubdomainConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SubdomainRegistrationFees) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubdomainConfig
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
			return fmt.Errorf("proto: SubdomainRegistrationFees: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubdomainRegistrationFees: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeByLength", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainConfig
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
				return ErrInvalidLengthSubdomainConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubdomainConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeByLength = append(m.FeeByLength, &FeeByLength{})
			if err := m.FeeByLength[len(m.FeeByLength)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeByName", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainConfig
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
				return ErrInvalidLengthSubdomainConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubdomainConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeByName = append(m.FeeByName, &FeeByName{})
			if err := m.FeeByName[len(m.FeeByName)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainConfig
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
				return ErrInvalidLengthSubdomainConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubdomainConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DefaultFee == nil {
				m.DefaultFee = &types.Coin{}
			}
			if err := m.DefaultFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubdomainConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubdomainConfig
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
func (m *FeeByLength) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubdomainConfig
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
			return fmt.Errorf("proto: FeeByLength: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeByLength: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length", wireType)
			}
			m.Length = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Length |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsRegistrable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainConfig
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
			m.IsRegistrable = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainConfig
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
				return ErrInvalidLengthSubdomainConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubdomainConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fee == nil {
				m.Fee = &types.Coin{}
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubdomainConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubdomainConfig
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
func (m *FeeByName) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubdomainConfig
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
			return fmt.Errorf("proto: FeeByName: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeByName: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainConfig
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
				return ErrInvalidLengthSubdomainConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubdomainConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsRegistrable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainConfig
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
			m.IsRegistrable = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainConfig
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
				return ErrInvalidLengthSubdomainConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubdomainConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fee == nil {
				m.Fee = &types.Coin{}
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubdomainConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubdomainConfig
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
func (m *SubdomainConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubdomainConfig
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
			return fmt.Errorf("proto: SubdomainConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubdomainConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSubdomainRegistrations", wireType)
			}
			m.MaxSubdomainRegistrations = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSubdomainRegistrations |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubdomainRegistrationFees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainConfig
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
				return ErrInvalidLengthSubdomainConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubdomainConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SubdomainRegistrationFees == nil {
				m.SubdomainRegistrationFees = &SubdomainRegistrationFees{}
			}
			if err := m.SubdomainRegistrationFees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsRegistrable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainConfig
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
			m.IsRegistrable = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistrableRole", wireType)
			}
			m.RegistrableRole = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegistrableRole |= DomainRole(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomExpirationDate", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainConfig
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
			m.CustomExpirationDate = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegistrationPolicy", wireType)
			}
			m.RegistrationPolicy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubdomainConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegistrationPolicy |= RegistrationPolicyType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSubdomainConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubdomainConfig
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
func skipSubdomainConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSubdomainConfig
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
					return 0, ErrIntOverflowSubdomainConfig
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
					return 0, ErrIntOverflowSubdomainConfig
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
				return 0, ErrInvalidLengthSubdomainConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSubdomainConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSubdomainConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSubdomainConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSubdomainConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSubdomainConfig = fmt.Errorf("proto: unexpected end of group")
)
