// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mycel/epochs/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	math "math"
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

func init() { proto.RegisterFile("mycel/epochs/v1/tx.proto", fileDescriptor_56b8aa43db691e30) }

var fileDescriptor_56b8aa43db691e30 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0xad, 0x4c, 0x4e,
	0xcd, 0xd1, 0x4f, 0x2d, 0xc8, 0x4f, 0xce, 0x28, 0xd6, 0x2f, 0x33, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x07, 0xcb, 0xe8, 0x41, 0x64, 0xf4, 0xca, 0x0c, 0xa5, 0x04,
	0x13, 0x73, 0x33, 0xf3, 0xf2, 0xf5, 0xc1, 0x24, 0x44, 0x8d, 0x94, 0x78, 0x72, 0x7e, 0x71, 0x6e,
	0x7e, 0xb1, 0x7e, 0x6e, 0x71, 0x3a, 0x48, 0x6f, 0x6e, 0x71, 0x3a, 0x54, 0x42, 0x12, 0x22, 0x11,
	0x0f, 0xe6, 0xe9, 0x43, 0x38, 0x50, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0x88, 0x38, 0x88, 0x05,
	0x11, 0x35, 0xe2, 0xe1, 0x62, 0xf6, 0x2d, 0x4e, 0x97, 0x62, 0x6d, 0x78, 0xbe, 0x41, 0x8b, 0xd1,
	0xc9, 0xed, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0,
	0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x74, 0xd2, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xc1, 0x0e, 0xd4, 0x4d, 0xc9, 0xcf, 0x4d, 0xcc,
	0xcc, 0x83, 0x70, 0xf4, 0x2b, 0x60, 0x3e, 0x29, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x1b,
	0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x6f, 0x28, 0x40, 0xe6, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mycel.epochs.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "mycel/epochs/v1/tx.proto",
}
