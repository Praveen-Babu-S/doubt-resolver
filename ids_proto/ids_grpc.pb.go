// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: ids_proto/ids.proto

package backend_ids

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

// StudentCRUDClient is the client API for StudentCRUD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentCRUDClient interface {
	CreateQuestion(ctx context.Context, in *Question, opts ...grpc.CallOption) (*Status, error)
	CreateComment(ctx context.Context, in *SolId, opts ...grpc.CallOption) (*Comment, error)
}

type studentCRUDClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentCRUDClient(cc grpc.ClientConnInterface) StudentCRUDClient {
	return &studentCRUDClient{cc}
}

func (c *studentCRUDClient) CreateQuestion(ctx context.Context, in *Question, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ids_proto.StudentCRUD/CreateQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentCRUDClient) CreateComment(ctx context.Context, in *SolId, opts ...grpc.CallOption) (*Comment, error) {
	out := new(Comment)
	err := c.cc.Invoke(ctx, "/ids_proto.StudentCRUD/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentCRUDServer is the server API for StudentCRUD service.
// All implementations must embed UnimplementedStudentCRUDServer
// for forward compatibility
type StudentCRUDServer interface {
	CreateQuestion(context.Context, *Question) (*Status, error)
	CreateComment(context.Context, *SolId) (*Comment, error)
	mustEmbedUnimplementedStudentCRUDServer()
}

// UnimplementedStudentCRUDServer must be embedded to have forward compatible implementations.
type UnimplementedStudentCRUDServer struct {
}

func (UnimplementedStudentCRUDServer) CreateQuestion(context.Context, *Question) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuestion not implemented")
}
func (UnimplementedStudentCRUDServer) CreateComment(context.Context, *SolId) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedStudentCRUDServer) mustEmbedUnimplementedStudentCRUDServer() {}

// UnsafeStudentCRUDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentCRUDServer will
// result in compilation errors.
type UnsafeStudentCRUDServer interface {
	mustEmbedUnimplementedStudentCRUDServer()
}

func RegisterStudentCRUDServer(s grpc.ServiceRegistrar, srv StudentCRUDServer) {
	s.RegisterService(&StudentCRUD_ServiceDesc, srv)
}

func _StudentCRUD_CreateQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Question)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentCRUDServer).CreateQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ids_proto.StudentCRUD/CreateQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentCRUDServer).CreateQuestion(ctx, req.(*Question))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentCRUD_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentCRUDServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ids_proto.StudentCRUD/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentCRUDServer).CreateComment(ctx, req.(*SolId))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentCRUD_ServiceDesc is the grpc.ServiceDesc for StudentCRUD service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentCRUD_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ids_proto.StudentCRUD",
	HandlerType: (*StudentCRUDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQuestion",
			Handler:    _StudentCRUD_CreateQuestion_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _StudentCRUD_CreateComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ids_proto/ids.proto",
}

// MentorCRUDClient is the client API for MentorCRUD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MentorCRUDClient interface {
	CreateSolution(ctx context.Context, in *Solution, opts ...grpc.CallOption) (*Status, error)
	CreateComment(ctx context.Context, in *SolId, opts ...grpc.CallOption) (*Comment, error)
}

type mentorCRUDClient struct {
	cc grpc.ClientConnInterface
}

func NewMentorCRUDClient(cc grpc.ClientConnInterface) MentorCRUDClient {
	return &mentorCRUDClient{cc}
}

func (c *mentorCRUDClient) CreateSolution(ctx context.Context, in *Solution, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ids_proto.MentorCRUD/CreateSolution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mentorCRUDClient) CreateComment(ctx context.Context, in *SolId, opts ...grpc.CallOption) (*Comment, error) {
	out := new(Comment)
	err := c.cc.Invoke(ctx, "/ids_proto.MentorCRUD/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MentorCRUDServer is the server API for MentorCRUD service.
// All implementations must embed UnimplementedMentorCRUDServer
// for forward compatibility
type MentorCRUDServer interface {
	CreateSolution(context.Context, *Solution) (*Status, error)
	CreateComment(context.Context, *SolId) (*Comment, error)
	mustEmbedUnimplementedMentorCRUDServer()
}

// UnimplementedMentorCRUDServer must be embedded to have forward compatible implementations.
type UnimplementedMentorCRUDServer struct {
}

func (UnimplementedMentorCRUDServer) CreateSolution(context.Context, *Solution) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSolution not implemented")
}
func (UnimplementedMentorCRUDServer) CreateComment(context.Context, *SolId) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedMentorCRUDServer) mustEmbedUnimplementedMentorCRUDServer() {}

// UnsafeMentorCRUDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MentorCRUDServer will
// result in compilation errors.
type UnsafeMentorCRUDServer interface {
	mustEmbedUnimplementedMentorCRUDServer()
}

func RegisterMentorCRUDServer(s grpc.ServiceRegistrar, srv MentorCRUDServer) {
	s.RegisterService(&MentorCRUD_ServiceDesc, srv)
}

func _MentorCRUD_CreateSolution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Solution)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MentorCRUDServer).CreateSolution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ids_proto.MentorCRUD/CreateSolution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MentorCRUDServer).CreateSolution(ctx, req.(*Solution))
	}
	return interceptor(ctx, in, info, handler)
}

func _MentorCRUD_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MentorCRUDServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ids_proto.MentorCRUD/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MentorCRUDServer).CreateComment(ctx, req.(*SolId))
	}
	return interceptor(ctx, in, info, handler)
}

// MentorCRUD_ServiceDesc is the grpc.ServiceDesc for MentorCRUD service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MentorCRUD_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ids_proto.MentorCRUD",
	HandlerType: (*MentorCRUDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSolution",
			Handler:    _MentorCRUD_CreateSolution_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _MentorCRUD_CreateComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ids_proto/ids.proto",
}