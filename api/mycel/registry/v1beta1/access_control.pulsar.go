// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package registryv1beta1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_AccessControl         protoreflect.MessageDescriptor
	fd_AccessControl_address protoreflect.FieldDescriptor
	fd_AccessControl_role    protoreflect.FieldDescriptor
)

func init() {
	file_mycel_registry_v1beta1_access_control_proto_init()
	md_AccessControl = File_mycel_registry_v1beta1_access_control_proto.Messages().ByName("AccessControl")
	fd_AccessControl_address = md_AccessControl.Fields().ByName("address")
	fd_AccessControl_role = md_AccessControl.Fields().ByName("role")
}

var _ protoreflect.Message = (*fastReflection_AccessControl)(nil)

type fastReflection_AccessControl AccessControl

func (x *AccessControl) ProtoReflect() protoreflect.Message {
	return (*fastReflection_AccessControl)(x)
}

func (x *AccessControl) slowProtoReflect() protoreflect.Message {
	mi := &file_mycel_registry_v1beta1_access_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_AccessControl_messageType fastReflection_AccessControl_messageType
var _ protoreflect.MessageType = fastReflection_AccessControl_messageType{}

type fastReflection_AccessControl_messageType struct{}

func (x fastReflection_AccessControl_messageType) Zero() protoreflect.Message {
	return (*fastReflection_AccessControl)(nil)
}
func (x fastReflection_AccessControl_messageType) New() protoreflect.Message {
	return new(fastReflection_AccessControl)
}
func (x fastReflection_AccessControl_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_AccessControl
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_AccessControl) Descriptor() protoreflect.MessageDescriptor {
	return md_AccessControl
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_AccessControl) Type() protoreflect.MessageType {
	return _fastReflection_AccessControl_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_AccessControl) New() protoreflect.Message {
	return new(fastReflection_AccessControl)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_AccessControl) Interface() protoreflect.ProtoMessage {
	return (*AccessControl)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_AccessControl) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_AccessControl_address, value) {
			return
		}
	}
	if x.Role != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Role))
		if !f(fd_AccessControl_role, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_AccessControl) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "mycel.registry.v1beta1.AccessControl.address":
		return x.Address != ""
	case "mycel.registry.v1beta1.AccessControl.role":
		return x.Role != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mycel.registry.v1beta1.AccessControl"))
		}
		panic(fmt.Errorf("message mycel.registry.v1beta1.AccessControl does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AccessControl) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "mycel.registry.v1beta1.AccessControl.address":
		x.Address = ""
	case "mycel.registry.v1beta1.AccessControl.role":
		x.Role = 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mycel.registry.v1beta1.AccessControl"))
		}
		panic(fmt.Errorf("message mycel.registry.v1beta1.AccessControl does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_AccessControl) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "mycel.registry.v1beta1.AccessControl.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "mycel.registry.v1beta1.AccessControl.role":
		value := x.Role
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mycel.registry.v1beta1.AccessControl"))
		}
		panic(fmt.Errorf("message mycel.registry.v1beta1.AccessControl does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AccessControl) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "mycel.registry.v1beta1.AccessControl.address":
		x.Address = value.Interface().(string)
	case "mycel.registry.v1beta1.AccessControl.role":
		x.Role = (DomainRole)(value.Enum())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mycel.registry.v1beta1.AccessControl"))
		}
		panic(fmt.Errorf("message mycel.registry.v1beta1.AccessControl does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AccessControl) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "mycel.registry.v1beta1.AccessControl.address":
		panic(fmt.Errorf("field address of message mycel.registry.v1beta1.AccessControl is not mutable"))
	case "mycel.registry.v1beta1.AccessControl.role":
		panic(fmt.Errorf("field role of message mycel.registry.v1beta1.AccessControl is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mycel.registry.v1beta1.AccessControl"))
		}
		panic(fmt.Errorf("message mycel.registry.v1beta1.AccessControl does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_AccessControl) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "mycel.registry.v1beta1.AccessControl.address":
		return protoreflect.ValueOfString("")
	case "mycel.registry.v1beta1.AccessControl.role":
		return protoreflect.ValueOfEnum(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: mycel.registry.v1beta1.AccessControl"))
		}
		panic(fmt.Errorf("message mycel.registry.v1beta1.AccessControl does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_AccessControl) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in mycel.registry.v1beta1.AccessControl", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_AccessControl) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AccessControl) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_AccessControl) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_AccessControl) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*AccessControl)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Role != 0 {
			n += 1 + runtime.Sov(uint64(x.Role))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*AccessControl)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Role != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Role))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*AccessControl)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AccessControl: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AccessControl: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
				}
				x.Role = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Role |= DomainRole(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: mycel/registry/v1beta1/access_control.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DomainRole int32

const (
	DomainRole_NO_ROLE DomainRole = 0
	DomainRole_OWNER   DomainRole = 1
	DomainRole_EDITOR  DomainRole = 2
)

// Enum value maps for DomainRole.
var (
	DomainRole_name = map[int32]string{
		0: "NO_ROLE",
		1: "OWNER",
		2: "EDITOR",
	}
	DomainRole_value = map[string]int32{
		"NO_ROLE": 0,
		"OWNER":   1,
		"EDITOR":  2,
	}
)

func (x DomainRole) Enum() *DomainRole {
	p := new(DomainRole)
	*p = x
	return p
}

func (x DomainRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DomainRole) Descriptor() protoreflect.EnumDescriptor {
	return file_mycel_registry_v1beta1_access_control_proto_enumTypes[0].Descriptor()
}

func (DomainRole) Type() protoreflect.EnumType {
	return &file_mycel_registry_v1beta1_access_control_proto_enumTypes[0]
}

func (x DomainRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DomainRole.Descriptor instead.
func (DomainRole) EnumDescriptor() ([]byte, []int) {
	return file_mycel_registry_v1beta1_access_control_proto_rawDescGZIP(), []int{0}
}

type AccessControl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Role    DomainRole `protobuf:"varint,2,opt,name=role,proto3,enum=mycel.registry.v1beta1.DomainRole" json:"role,omitempty"`
}

func (x *AccessControl) Reset() {
	*x = AccessControl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mycel_registry_v1beta1_access_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessControl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessControl) ProtoMessage() {}

// Deprecated: Use AccessControl.ProtoReflect.Descriptor instead.
func (*AccessControl) Descriptor() ([]byte, []int) {
	return file_mycel_registry_v1beta1_access_control_proto_rawDescGZIP(), []int{0}
}

func (x *AccessControl) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AccessControl) GetRole() DomainRole {
	if x != nil {
		return x.Role
	}
	return DomainRole_NO_ROLE
}

var File_mycel_registry_v1beta1_access_control_proto protoreflect.FileDescriptor

var file_mycel_registry_v1beta1_access_control_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6d, 0x79, 0x63, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6d,
	0x79, 0x63, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x22, 0x61, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x36, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22,
	0x2e, 0x6d, 0x79, 0x63, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x2a, 0x30, 0x0a, 0x0a, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x4f, 0x5f, 0x52, 0x4f, 0x4c,
	0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x45, 0x44, 0x49, 0x54, 0x4f, 0x52, 0x10, 0x02, 0x42, 0xe3, 0x01, 0x0a, 0x1a, 0x63,
	0x6f, 0x6d, 0x2e, 0x6d, 0x79, 0x63, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x12, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x37, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6d, 0x79, 0x63, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x52, 0x58, 0xaa, 0x02,
	0x16, 0x4d, 0x79, 0x63, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x16, 0x4d, 0x79, 0x63, 0x65, 0x6c, 0x5c,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xe2, 0x02, 0x22, 0x4d, 0x79, 0x63, 0x65, 0x6c, 0x5c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x4d, 0x79, 0x63, 0x65, 0x6c, 0x3a, 0x3a, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mycel_registry_v1beta1_access_control_proto_rawDescOnce sync.Once
	file_mycel_registry_v1beta1_access_control_proto_rawDescData = file_mycel_registry_v1beta1_access_control_proto_rawDesc
)

func file_mycel_registry_v1beta1_access_control_proto_rawDescGZIP() []byte {
	file_mycel_registry_v1beta1_access_control_proto_rawDescOnce.Do(func() {
		file_mycel_registry_v1beta1_access_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_mycel_registry_v1beta1_access_control_proto_rawDescData)
	})
	return file_mycel_registry_v1beta1_access_control_proto_rawDescData
}

var file_mycel_registry_v1beta1_access_control_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mycel_registry_v1beta1_access_control_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mycel_registry_v1beta1_access_control_proto_goTypes = []interface{}{
	(DomainRole)(0),       // 0: mycel.registry.v1beta1.DomainRole
	(*AccessControl)(nil), // 1: mycel.registry.v1beta1.AccessControl
}
var file_mycel_registry_v1beta1_access_control_proto_depIdxs = []int32{
	0, // 0: mycel.registry.v1beta1.AccessControl.role:type_name -> mycel.registry.v1beta1.DomainRole
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mycel_registry_v1beta1_access_control_proto_init() }
func file_mycel_registry_v1beta1_access_control_proto_init() {
	if File_mycel_registry_v1beta1_access_control_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mycel_registry_v1beta1_access_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessControl); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mycel_registry_v1beta1_access_control_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mycel_registry_v1beta1_access_control_proto_goTypes,
		DependencyIndexes: file_mycel_registry_v1beta1_access_control_proto_depIdxs,
		EnumInfos:         file_mycel_registry_v1beta1_access_control_proto_enumTypes,
		MessageInfos:      file_mycel_registry_v1beta1_access_control_proto_msgTypes,
	}.Build()
	File_mycel_registry_v1beta1_access_control_proto = out.File
	file_mycel_registry_v1beta1_access_control_proto_rawDesc = nil
	file_mycel_registry_v1beta1_access_control_proto_goTypes = nil
	file_mycel_registry_v1beta1_access_control_proto_depIdxs = nil
}
