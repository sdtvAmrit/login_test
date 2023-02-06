// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: service_1/proto/add_msg.proto

package pb

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

// AdderClient is the client API for Adder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdderClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
}

type adderClient struct {
	cc grpc.ClientConnInterface
}

func NewAdderClient(cc grpc.ClientConnInterface) AdderClient {
	return &adderClient{cc}
}

func (c *adderClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/pb.Adder/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdderServer is the server API for Adder service.
// All implementations should embed UnimplementedAdderServer
// for forward compatibility
type AdderServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
}

// UnimplementedAdderServer should be embedded to have forward compatible implementations.
type UnimplementedAdderServer struct {
}

func (UnimplementedAdderServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

// UnsafeAdderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdderServer will
// result in compilation errors.
type UnsafeAdderServer interface {
	mustEmbedUnimplementedAdderServer()
}

func RegisterAdderServer(s grpc.ServiceRegistrar, srv AdderServer) {
	s.RegisterService(&Adder_ServiceDesc, srv)
}

func _Adder_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdderServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Adder/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdderServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Adder_ServiceDesc is the grpc.ServiceDesc for Adder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Adder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Adder",
	HandlerType: (*AdderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Adder_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_1/proto/add_msg.proto",
}

// PrinterClient is the client API for Printer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrinterClient interface {
	Print(ctx context.Context, in *PrintRequest, opts ...grpc.CallOption) (*PrintResponse, error)
}

type printerClient struct {
	cc grpc.ClientConnInterface
}

func NewPrinterClient(cc grpc.ClientConnInterface) PrinterClient {
	return &printerClient{cc}
}

func (c *printerClient) Print(ctx context.Context, in *PrintRequest, opts ...grpc.CallOption) (*PrintResponse, error) {
	out := new(PrintResponse)
	err := c.cc.Invoke(ctx, "/pb.Printer/Print", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrinterServer is the server API for Printer service.
// All implementations should embed UnimplementedPrinterServer
// for forward compatibility
type PrinterServer interface {
	Print(context.Context, *PrintRequest) (*PrintResponse, error)
}

// UnimplementedPrinterServer should be embedded to have forward compatible implementations.
type UnimplementedPrinterServer struct {
}

func (UnimplementedPrinterServer) Print(context.Context, *PrintRequest) (*PrintResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Print not implemented")
}

// UnsafePrinterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrinterServer will
// result in compilation errors.
type UnsafePrinterServer interface {
	mustEmbedUnimplementedPrinterServer()
}

func RegisterPrinterServer(s grpc.ServiceRegistrar, srv PrinterServer) {
	s.RegisterService(&Printer_ServiceDesc, srv)
}

func _Printer_Print_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrintRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrinterServer).Print(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Printer/Print",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrinterServer).Print(ctx, req.(*PrintRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Printer_ServiceDesc is the grpc.ServiceDesc for Printer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Printer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Printer",
	HandlerType: (*PrinterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Print",
			Handler:    _Printer_Print_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_1/proto/add_msg.proto",
}