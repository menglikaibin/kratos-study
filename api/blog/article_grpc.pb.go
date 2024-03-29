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

// ArticleClient is the client API for Article service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleClient interface {
	CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleReply, error)
	UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*UpdateArticleReply, error)
	DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleReply, error)
	GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*GetArticleReply, error)
}

type articleClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleClient(cc grpc.ClientConnInterface) ArticleClient {
	return &articleClient{cc}
}

func (c *articleClient) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleReply, error) {
	out := new(CreateArticleReply)
	err := c.cc.Invoke(ctx, "/api.blog.Article/CreateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*UpdateArticleReply, error) {
	out := new(UpdateArticleReply)
	err := c.cc.Invoke(ctx, "/api.blog.Article/UpdateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*DeleteArticleReply, error) {
	out := new(DeleteArticleReply)
	err := c.cc.Invoke(ctx, "/api.blog.Article/DeleteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleClient) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*GetArticleReply, error) {
	out := new(GetArticleReply)
	err := c.cc.Invoke(ctx, "/api.blog.Article/GetArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServer is the server API for Article service.
// All implementations must embed UnimplementedArticleServer
// for forward compatibility
type ArticleServer interface {
	CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleReply, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*UpdateArticleReply, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleReply, error)
	GetArticle(context.Context, *GetArticleRequest) (*GetArticleReply, error)
	mustEmbedUnimplementedArticleServer()
}

// UnimplementedArticleServer must be embedded to have forward compatible implementations.
type UnimplementedArticleServer struct {
}

func (UnimplementedArticleServer) CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedArticleServer) UpdateArticle(context.Context, *UpdateArticleRequest) (*UpdateArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (UnimplementedArticleServer) DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedArticleServer) GetArticle(context.Context, *GetArticleRequest) (*GetArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedArticleServer) mustEmbedUnimplementedArticleServer() {}

// UnsafeArticleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServer will
// result in compilation errors.
type UnsafeArticleServer interface {
	mustEmbedUnimplementedArticleServer()
}

func RegisterArticleServer(s grpc.ServiceRegistrar, srv ArticleServer) {
	s.RegisterService(&Article_ServiceDesc, srv)
}

func _Article_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.blog.Article/CreateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).CreateArticle(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.blog.Article/UpdateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).UpdateArticle(ctx, req.(*UpdateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.blog.Article/DeleteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).DeleteArticle(ctx, req.(*DeleteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Article_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.blog.Article/GetArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServer).GetArticle(ctx, req.(*GetArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Article_ServiceDesc is the grpc.ServiceDesc for Article service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Article_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.blog.Article",
	HandlerType: (*ArticleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArticle",
			Handler:    _Article_CreateArticle_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _Article_UpdateArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _Article_DeleteArticle_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _Article_GetArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article.proto",
}
