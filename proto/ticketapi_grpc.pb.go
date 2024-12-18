// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: ticketapi.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TicketAPI_BookTicket_FullMethodName   = "/proto.TicketAPI/BookTicket"
	TicketAPI_CancelTicket_FullMethodName = "/proto.TicketAPI/CancelTicket"
	TicketAPI_ModifyTicket_FullMethodName = "/proto.TicketAPI/ModifyTicket"
	TicketAPI_GetTicket_FullMethodName    = "/proto.TicketAPI/GetTicket"
	TicketAPI_ListTickets_FullMethodName  = "/proto.TicketAPI/ListTickets"
)

// TicketAPIClient is the client API for TicketAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketAPIClient interface {
	BookTicket(ctx context.Context, in *BookTicketRequest, opts ...grpc.CallOption) (*Ticket, error)
	CancelTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ModifyTicket(ctx context.Context, in *ModifyTicketRequest, opts ...grpc.CallOption) (*Ticket, error)
	GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*Ticket, error)
	ListTickets(ctx context.Context, in *ListTicketRequest, opts ...grpc.CallOption) (*Tickets, error)
}

type ticketAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketAPIClient(cc grpc.ClientConnInterface) TicketAPIClient {
	return &ticketAPIClient{cc}
}

func (c *ticketAPIClient) BookTicket(ctx context.Context, in *BookTicketRequest, opts ...grpc.CallOption) (*Ticket, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Ticket)
	err := c.cc.Invoke(ctx, TicketAPI_BookTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketAPIClient) CancelTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TicketAPI_CancelTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketAPIClient) ModifyTicket(ctx context.Context, in *ModifyTicketRequest, opts ...grpc.CallOption) (*Ticket, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Ticket)
	err := c.cc.Invoke(ctx, TicketAPI_ModifyTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketAPIClient) GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*Ticket, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Ticket)
	err := c.cc.Invoke(ctx, TicketAPI_GetTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketAPIClient) ListTickets(ctx context.Context, in *ListTicketRequest, opts ...grpc.CallOption) (*Tickets, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Tickets)
	err := c.cc.Invoke(ctx, TicketAPI_ListTickets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketAPIServer is the server API for TicketAPI service.
// All implementations must embed UnimplementedTicketAPIServer
// for forward compatibility.
type TicketAPIServer interface {
	BookTicket(context.Context, *BookTicketRequest) (*Ticket, error)
	CancelTicket(context.Context, *GetTicketRequest) (*emptypb.Empty, error)
	ModifyTicket(context.Context, *ModifyTicketRequest) (*Ticket, error)
	GetTicket(context.Context, *GetTicketRequest) (*Ticket, error)
	ListTickets(context.Context, *ListTicketRequest) (*Tickets, error)
	mustEmbedUnimplementedTicketAPIServer()
}

// UnimplementedTicketAPIServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTicketAPIServer struct{}

func (UnimplementedTicketAPIServer) BookTicket(context.Context, *BookTicketRequest) (*Ticket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookTicket not implemented")
}
func (UnimplementedTicketAPIServer) CancelTicket(context.Context, *GetTicketRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTicket not implemented")
}
func (UnimplementedTicketAPIServer) ModifyTicket(context.Context, *ModifyTicketRequest) (*Ticket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyTicket not implemented")
}
func (UnimplementedTicketAPIServer) GetTicket(context.Context, *GetTicketRequest) (*Ticket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTicket not implemented")
}
func (UnimplementedTicketAPIServer) ListTickets(context.Context, *ListTicketRequest) (*Tickets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTickets not implemented")
}
func (UnimplementedTicketAPIServer) mustEmbedUnimplementedTicketAPIServer() {}
func (UnimplementedTicketAPIServer) testEmbeddedByValue()                   {}

// UnsafeTicketAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketAPIServer will
// result in compilation errors.
type UnsafeTicketAPIServer interface {
	mustEmbedUnimplementedTicketAPIServer()
}

func RegisterTicketAPIServer(s grpc.ServiceRegistrar, srv TicketAPIServer) {
	// If the following call pancis, it indicates UnimplementedTicketAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TicketAPI_ServiceDesc, srv)
}

func _TicketAPI_BookTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketAPIServer).BookTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketAPI_BookTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketAPIServer).BookTicket(ctx, req.(*BookTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketAPI_CancelTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketAPIServer).CancelTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketAPI_CancelTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketAPIServer).CancelTicket(ctx, req.(*GetTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketAPI_ModifyTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketAPIServer).ModifyTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketAPI_ModifyTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketAPIServer).ModifyTicket(ctx, req.(*ModifyTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketAPI_GetTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketAPIServer).GetTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketAPI_GetTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketAPIServer).GetTicket(ctx, req.(*GetTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketAPI_ListTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketAPIServer).ListTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketAPI_ListTickets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketAPIServer).ListTickets(ctx, req.(*ListTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketAPI_ServiceDesc is the grpc.ServiceDesc for TicketAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TicketAPI",
	HandlerType: (*TicketAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BookTicket",
			Handler:    _TicketAPI_BookTicket_Handler,
		},
		{
			MethodName: "CancelTicket",
			Handler:    _TicketAPI_CancelTicket_Handler,
		},
		{
			MethodName: "ModifyTicket",
			Handler:    _TicketAPI_ModifyTicket_Handler,
		},
		{
			MethodName: "GetTicket",
			Handler:    _TicketAPI_GetTicket_Handler,
		},
		{
			MethodName: "ListTickets",
			Handler:    _TicketAPI_ListTickets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ticketapi.proto",
}
