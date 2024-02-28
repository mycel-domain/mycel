// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: mycel/registry/v1beta1/tx.proto

package registryv1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_UpdateWalletRecord_FullMethodName                     = "/mycel.registry.v1beta1.Msg/UpdateWalletRecord"
	Msg_UpdateDnsRecord_FullMethodName                        = "/mycel.registry.v1beta1.Msg/UpdateDnsRecord"
	Msg_RegisterSecondLevelDomain_FullMethodName              = "/mycel.registry.v1beta1.Msg/RegisterSecondLevelDomain"
	Msg_RegisterTopLevelDomain_FullMethodName                 = "/mycel.registry.v1beta1.Msg/RegisterTopLevelDomain"
	Msg_WithdrawRegistrationFee_FullMethodName                = "/mycel.registry.v1beta1.Msg/WithdrawRegistrationFee"
	Msg_ExtendTopLevelDomainExpirationDate_FullMethodName     = "/mycel.registry.v1beta1.Msg/ExtendTopLevelDomainExpirationDate"
	Msg_UpdateTextRecord_FullMethodName                       = "/mycel.registry.v1beta1.Msg/UpdateTextRecord"
	Msg_UpdateTopLevelDomainRegistrationPolicy_FullMethodName = "/mycel.registry.v1beta1.Msg/UpdateTopLevelDomainRegistrationPolicy"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	UpdateWalletRecord(ctx context.Context, in *MsgUpdateWalletRecord, opts ...grpc.CallOption) (*MsgUpdateWalletRecordResponse, error)
	UpdateDnsRecord(ctx context.Context, in *MsgUpdateDnsRecord, opts ...grpc.CallOption) (*MsgUpdateDnsRecordResponse, error)
	RegisterSecondLevelDomain(ctx context.Context, in *MsgRegisterSecondLevelDomain, opts ...grpc.CallOption) (*MsgRegisterSecondLevelDomainResponse, error)
	RegisterTopLevelDomain(ctx context.Context, in *MsgRegisterTopLevelDomain, opts ...grpc.CallOption) (*MsgRegisterTopLevelDomainResponse, error)
	WithdrawRegistrationFee(ctx context.Context, in *MsgWithdrawRegistrationFee, opts ...grpc.CallOption) (*MsgWithdrawRegistrationFeeResponse, error)
	ExtendTopLevelDomainExpirationDate(ctx context.Context, in *MsgExtendTopLevelDomainExpirationDate, opts ...grpc.CallOption) (*MsgExtendTopLevelDomainExpirationDateResponse, error)
	UpdateTextRecord(ctx context.Context, in *MsgUpdateTextRecord, opts ...grpc.CallOption) (*MsgUpdateTextRecordResponse, error)
	UpdateTopLevelDomainRegistrationPolicy(ctx context.Context, in *MsgUpdateTopLevelDomainRegistrationPolicy, opts ...grpc.CallOption) (*MsgUpdateTopLevelDomainRegistrationPolicyResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateWalletRecord(ctx context.Context, in *MsgUpdateWalletRecord, opts ...grpc.CallOption) (*MsgUpdateWalletRecordResponse, error) {
	out := new(MsgUpdateWalletRecordResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateWalletRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateDnsRecord(ctx context.Context, in *MsgUpdateDnsRecord, opts ...grpc.CallOption) (*MsgUpdateDnsRecordResponse, error) {
	out := new(MsgUpdateDnsRecordResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateDnsRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RegisterSecondLevelDomain(ctx context.Context, in *MsgRegisterSecondLevelDomain, opts ...grpc.CallOption) (*MsgRegisterSecondLevelDomainResponse, error) {
	out := new(MsgRegisterSecondLevelDomainResponse)
	err := c.cc.Invoke(ctx, Msg_RegisterSecondLevelDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RegisterTopLevelDomain(ctx context.Context, in *MsgRegisterTopLevelDomain, opts ...grpc.CallOption) (*MsgRegisterTopLevelDomainResponse, error) {
	out := new(MsgRegisterTopLevelDomainResponse)
	err := c.cc.Invoke(ctx, Msg_RegisterTopLevelDomain_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawRegistrationFee(ctx context.Context, in *MsgWithdrawRegistrationFee, opts ...grpc.CallOption) (*MsgWithdrawRegistrationFeeResponse, error) {
	out := new(MsgWithdrawRegistrationFeeResponse)
	err := c.cc.Invoke(ctx, Msg_WithdrawRegistrationFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ExtendTopLevelDomainExpirationDate(ctx context.Context, in *MsgExtendTopLevelDomainExpirationDate, opts ...grpc.CallOption) (*MsgExtendTopLevelDomainExpirationDateResponse, error) {
	out := new(MsgExtendTopLevelDomainExpirationDateResponse)
	err := c.cc.Invoke(ctx, Msg_ExtendTopLevelDomainExpirationDate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateTextRecord(ctx context.Context, in *MsgUpdateTextRecord, opts ...grpc.CallOption) (*MsgUpdateTextRecordResponse, error) {
	out := new(MsgUpdateTextRecordResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateTextRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateTopLevelDomainRegistrationPolicy(ctx context.Context, in *MsgUpdateTopLevelDomainRegistrationPolicy, opts ...grpc.CallOption) (*MsgUpdateTopLevelDomainRegistrationPolicyResponse, error) {
	out := new(MsgUpdateTopLevelDomainRegistrationPolicyResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateTopLevelDomainRegistrationPolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	UpdateWalletRecord(context.Context, *MsgUpdateWalletRecord) (*MsgUpdateWalletRecordResponse, error)
	UpdateDnsRecord(context.Context, *MsgUpdateDnsRecord) (*MsgUpdateDnsRecordResponse, error)
	RegisterSecondLevelDomain(context.Context, *MsgRegisterSecondLevelDomain) (*MsgRegisterSecondLevelDomainResponse, error)
	RegisterTopLevelDomain(context.Context, *MsgRegisterTopLevelDomain) (*MsgRegisterTopLevelDomainResponse, error)
	WithdrawRegistrationFee(context.Context, *MsgWithdrawRegistrationFee) (*MsgWithdrawRegistrationFeeResponse, error)
	ExtendTopLevelDomainExpirationDate(context.Context, *MsgExtendTopLevelDomainExpirationDate) (*MsgExtendTopLevelDomainExpirationDateResponse, error)
	UpdateTextRecord(context.Context, *MsgUpdateTextRecord) (*MsgUpdateTextRecordResponse, error)
	UpdateTopLevelDomainRegistrationPolicy(context.Context, *MsgUpdateTopLevelDomainRegistrationPolicy) (*MsgUpdateTopLevelDomainRegistrationPolicyResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateWalletRecord(context.Context, *MsgUpdateWalletRecord) (*MsgUpdateWalletRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWalletRecord not implemented")
}
func (UnimplementedMsgServer) UpdateDnsRecord(context.Context, *MsgUpdateDnsRecord) (*MsgUpdateDnsRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDnsRecord not implemented")
}
func (UnimplementedMsgServer) RegisterSecondLevelDomain(context.Context, *MsgRegisterSecondLevelDomain) (*MsgRegisterSecondLevelDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterSecondLevelDomain not implemented")
}
func (UnimplementedMsgServer) RegisterTopLevelDomain(context.Context, *MsgRegisterTopLevelDomain) (*MsgRegisterTopLevelDomainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterTopLevelDomain not implemented")
}
func (UnimplementedMsgServer) WithdrawRegistrationFee(context.Context, *MsgWithdrawRegistrationFee) (*MsgWithdrawRegistrationFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawRegistrationFee not implemented")
}
func (UnimplementedMsgServer) ExtendTopLevelDomainExpirationDate(context.Context, *MsgExtendTopLevelDomainExpirationDate) (*MsgExtendTopLevelDomainExpirationDateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtendTopLevelDomainExpirationDate not implemented")
}
func (UnimplementedMsgServer) UpdateTextRecord(context.Context, *MsgUpdateTextRecord) (*MsgUpdateTextRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTextRecord not implemented")
}
func (UnimplementedMsgServer) UpdateTopLevelDomainRegistrationPolicy(context.Context, *MsgUpdateTopLevelDomainRegistrationPolicy) (*MsgUpdateTopLevelDomainRegistrationPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTopLevelDomainRegistrationPolicy not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateWalletRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateWalletRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateWalletRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateWalletRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateWalletRecord(ctx, req.(*MsgUpdateWalletRecord))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateDnsRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateDnsRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateDnsRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateDnsRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateDnsRecord(ctx, req.(*MsgUpdateDnsRecord))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RegisterSecondLevelDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterSecondLevelDomain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterSecondLevelDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RegisterSecondLevelDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterSecondLevelDomain(ctx, req.(*MsgRegisterSecondLevelDomain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RegisterTopLevelDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterTopLevelDomain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterTopLevelDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RegisterTopLevelDomain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterTopLevelDomain(ctx, req.(*MsgRegisterTopLevelDomain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawRegistrationFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawRegistrationFee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawRegistrationFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_WithdrawRegistrationFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawRegistrationFee(ctx, req.(*MsgWithdrawRegistrationFee))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ExtendTopLevelDomainExpirationDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExtendTopLevelDomainExpirationDate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ExtendTopLevelDomainExpirationDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ExtendTopLevelDomainExpirationDate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ExtendTopLevelDomainExpirationDate(ctx, req.(*MsgExtendTopLevelDomainExpirationDate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateTextRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateTextRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateTextRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateTextRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateTextRecord(ctx, req.(*MsgUpdateTextRecord))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateTopLevelDomainRegistrationPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateTopLevelDomainRegistrationPolicy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateTopLevelDomainRegistrationPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateTopLevelDomainRegistrationPolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateTopLevelDomainRegistrationPolicy(ctx, req.(*MsgUpdateTopLevelDomainRegistrationPolicy))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mycel.registry.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateWalletRecord",
			Handler:    _Msg_UpdateWalletRecord_Handler,
		},
		{
			MethodName: "UpdateDnsRecord",
			Handler:    _Msg_UpdateDnsRecord_Handler,
		},
		{
			MethodName: "RegisterSecondLevelDomain",
			Handler:    _Msg_RegisterSecondLevelDomain_Handler,
		},
		{
			MethodName: "RegisterTopLevelDomain",
			Handler:    _Msg_RegisterTopLevelDomain_Handler,
		},
		{
			MethodName: "WithdrawRegistrationFee",
			Handler:    _Msg_WithdrawRegistrationFee_Handler,
		},
		{
			MethodName: "ExtendTopLevelDomainExpirationDate",
			Handler:    _Msg_ExtendTopLevelDomainExpirationDate_Handler,
		},
		{
			MethodName: "UpdateTextRecord",
			Handler:    _Msg_UpdateTextRecord_Handler,
		},
		{
			MethodName: "UpdateTopLevelDomainRegistrationPolicy",
			Handler:    _Msg_UpdateTopLevelDomainRegistrationPolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mycel/registry/v1beta1/tx.proto",
}
