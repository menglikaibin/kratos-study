// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package blog

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

// BlogClient is the client API for Blog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogClient interface {
	CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogReply, error)
	UpdateBlog(ctx context.Context, in *UpdateBlogRequest, opts ...grpc.CallOption) (*UpdateBlogReply, error)
	DeleteBlog(ctx context.Context, in *DeleteBlogRequest, opts ...grpc.CallOption) (*DeleteBlogReply, error)
	GetBlog(ctx context.Context, in *GetBlogRequest, opts ...grpc.CallOption) (*GetBlogReply, error)
	ListBlog(ctx context.Context, in *ListBlogRequest, opts ...grpc.CallOption) (*ListBlogReply, error)
}

type blogClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogClient(cc grpc.ClientConnInterface) BlogClient {
	return &blogClient{cc}
}

func (c *blogClient) CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogReply, error) {
	out := new(CreateBlogReply)
	err := c.cc.Invoke(ctx, "/api.blog.Blog/CreateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) UpdateBlog(ctx context.Context, in *UpdateBlogRequest, opts ...grpc.CallOption) (*UpdateBlogReply, error) {
	out := new(UpdateBlogReply)
	err := c.cc.Invoke(ctx, "/api.blog.Blog/UpdateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) DeleteBlog(ctx context.Context, in *DeleteBlogRequest, opts ...grpc.CallOption) (*DeleteBlogReply, error) {
	out := new(DeleteBlogReply)
	err := c.cc.Invoke(ctx, "/api.blog.Blog/DeleteBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) GetBlog(ctx context.Context, in *GetBlogRequest, opts ...grpc.CallOption) (*GetBlogReply, error) {
	out := new(GetBlogReply)
	err := c.cc.Invoke(ctx, "/api.blog.Blog/GetBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) ListBlog(ctx context.Context, in *ListBlogRequest, opts ...grpc.CallOption) (*ListBlogReply, error) {
	out := new(ListBlogReply)
	err := c.cc.Invoke(ctx, "/api.blog.Blog/ListBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServer is the server API for Blog service.
// All implementations must embed UnimplementedBlogServer
// for forward compatibility
type BlogServer interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*CreateBlogReply, error)
	UpdateBlog(context.Context, *UpdateBlogRequest) (*UpdateBlogReply, error)
	DeleteBlog(context.Context, *DeleteBlogRequest) (*DeleteBlogReply, error)
	GetBlog(context.Context, *GetBlogRequest) (*GetBlogReply, error)
	ListBlog(context.Context, *ListBlogRequest) (*ListBlogReply, error)
	mustEmbedUnimplementedBlogServer()
}

// UnimplementedBlogServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServer struct {
}

func (UnimplementedBlogServer) CreateBlog(context.Context, *CreateBlogRequest) (*CreateBlogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}
func (UnimplementedBlogServer) UpdateBlog(context.Context, *UpdateBlogRequest) (*UpdateBlogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlog not implemented")
}
func (UnimplementedBlogServer) DeleteBlog(context.Context, *DeleteBlogRequest) (*DeleteBlogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlog not implemented")
}
func (UnimplementedBlogServer) GetBlog(context.Context, *GetBlogRequest) (*GetBlogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlog not implemented")
}
func (UnimplementedBlogServer) ListBlog(context.Context, *ListBlogRequest) (*ListBlogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBlog not implemented")
}
func (UnimplementedBlogServer) mustEmbedUnimplementedBlogServer() {}

// UnsafeBlogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogServer will
// result in compilation errors.
type UnsafeBlogServer interface {
	mustEmbedUnimplementedBlogServer()
}

func RegisterBlogServer(s grpc.ServiceRegistrar, srv BlogServer) {
	s.RegisterService(&Blog_ServiceDesc, srv)
}

func _Blog_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.blog.Blog/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).CreateBlog(ctx, req.(*CreateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_UpdateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).UpdateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.blog.Blog/UpdateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).UpdateBlog(ctx, req.(*UpdateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_DeleteBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).DeleteBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.blog.Blog/DeleteBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).DeleteBlog(ctx, req.(*DeleteBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_GetBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).GetBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.blog.Blog/GetBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).GetBlog(ctx, req.(*GetBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_ListBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).ListBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.blog.Blog/ListBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).ListBlog(ctx, req.(*ListBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Blog_ServiceDesc is the grpc.ServiceDesc for Blog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Blog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.blog.Blog",
	HandlerType: (*BlogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _Blog_CreateBlog_Handler,
		},
		{
			MethodName: "UpdateBlog",
			Handler:    _Blog_UpdateBlog_Handler,
		},
		{
			MethodName: "DeleteBlog",
			Handler:    _Blog_DeleteBlog_Handler,
		},
		{
			MethodName: "GetBlog",
			Handler:    _Blog_GetBlog_Handler,
		},
		{
			MethodName: "ListBlog",
			Handler:    _Blog_ListBlog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blog.proto",
}
