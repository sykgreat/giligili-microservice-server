// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: resource.proto

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

const (
	UploadService_UploadVideo_FullMethodName         = "/pb.UploadService/UploadVideo"
	UploadService_UploadImg_FullMethodName           = "/pb.UploadService/UploadImg"
	UploadService_ChangeResourceTitle_FullMethodName = "/pb.UploadService/ChangeResourceTitle"
	UploadService_DeleteResource_FullMethodName      = "/pb.UploadService/DeleteResource"
)

// UploadServiceClient is the client API for UploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadServiceClient interface {
	UploadVideo(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadVideoClient, error)
	UploadImg(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadImgClient, error)
	ChangeResourceTitle(ctx context.Context, in *ChangeResourceTitleRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...grpc.CallOption) (*BaseResponse, error)
}

type uploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadServiceClient(cc grpc.ClientConnInterface) UploadServiceClient {
	return &uploadServiceClient{cc}
}

func (c *uploadServiceClient) UploadVideo(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadVideoClient, error) {
	stream, err := c.cc.NewStream(ctx, &UploadService_ServiceDesc.Streams[0], UploadService_UploadVideo_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadServiceUploadVideoClient{stream}
	return x, nil
}

type UploadService_UploadVideoClient interface {
	Send(*UploadVideoRequest) error
	CloseAndRecv() (*BaseResponse, error)
	grpc.ClientStream
}

type uploadServiceUploadVideoClient struct {
	grpc.ClientStream
}

func (x *uploadServiceUploadVideoClient) Send(m *UploadVideoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploadServiceUploadVideoClient) CloseAndRecv() (*BaseResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BaseResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uploadServiceClient) UploadImg(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadImgClient, error) {
	stream, err := c.cc.NewStream(ctx, &UploadService_ServiceDesc.Streams[1], UploadService_UploadImg_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadServiceUploadImgClient{stream}
	return x, nil
}

type UploadService_UploadImgClient interface {
	Send(*UploadImgRequest) error
	CloseAndRecv() (*UploadImgResponse, error)
	grpc.ClientStream
}

type uploadServiceUploadImgClient struct {
	grpc.ClientStream
}

func (x *uploadServiceUploadImgClient) Send(m *UploadImgRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploadServiceUploadImgClient) CloseAndRecv() (*UploadImgResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadImgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uploadServiceClient) ChangeResourceTitle(ctx context.Context, in *ChangeResourceTitleRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, UploadService_ChangeResourceTitle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadServiceClient) DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, UploadService_DeleteResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UploadServiceServer is the server API for UploadService service.
// All implementations must embed UnimplementedUploadServiceServer
// for forward compatibility
type UploadServiceServer interface {
	UploadVideo(UploadService_UploadVideoServer) error
	UploadImg(UploadService_UploadImgServer) error
	ChangeResourceTitle(context.Context, *ChangeResourceTitleRequest) (*BaseResponse, error)
	DeleteResource(context.Context, *DeleteResourceRequest) (*BaseResponse, error)
	mustEmbedUnimplementedUploadServiceServer()
}

// UnimplementedUploadServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUploadServiceServer struct {
}

func (UnimplementedUploadServiceServer) UploadVideo(UploadService_UploadVideoServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadVideo not implemented")
}
func (UnimplementedUploadServiceServer) UploadImg(UploadService_UploadImgServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadImg not implemented")
}
func (UnimplementedUploadServiceServer) ChangeResourceTitle(context.Context, *ChangeResourceTitleRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeResourceTitle not implemented")
}
func (UnimplementedUploadServiceServer) DeleteResource(context.Context, *DeleteResourceRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResource not implemented")
}
func (UnimplementedUploadServiceServer) mustEmbedUnimplementedUploadServiceServer() {}

// UnsafeUploadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadServiceServer will
// result in compilation errors.
type UnsafeUploadServiceServer interface {
	mustEmbedUnimplementedUploadServiceServer()
}

func RegisterUploadServiceServer(s grpc.ServiceRegistrar, srv UploadServiceServer) {
	s.RegisterService(&UploadService_ServiceDesc, srv)
}

func _UploadService_UploadVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadServiceServer).UploadVideo(&uploadServiceUploadVideoServer{stream})
}

type UploadService_UploadVideoServer interface {
	SendAndClose(*BaseResponse) error
	Recv() (*UploadVideoRequest, error)
	grpc.ServerStream
}

type uploadServiceUploadVideoServer struct {
	grpc.ServerStream
}

func (x *uploadServiceUploadVideoServer) SendAndClose(m *BaseResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploadServiceUploadVideoServer) Recv() (*UploadVideoRequest, error) {
	m := new(UploadVideoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UploadService_UploadImg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadServiceServer).UploadImg(&uploadServiceUploadImgServer{stream})
}

type UploadService_UploadImgServer interface {
	SendAndClose(*UploadImgResponse) error
	Recv() (*UploadImgRequest, error)
	grpc.ServerStream
}

type uploadServiceUploadImgServer struct {
	grpc.ServerStream
}

func (x *uploadServiceUploadImgServer) SendAndClose(m *UploadImgResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploadServiceUploadImgServer) Recv() (*UploadImgRequest, error) {
	m := new(UploadImgRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UploadService_ChangeResourceTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeResourceTitleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadServiceServer).ChangeResourceTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadService_ChangeResourceTitle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadServiceServer).ChangeResourceTitle(ctx, req.(*ChangeResourceTitleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UploadService_DeleteResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadServiceServer).DeleteResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UploadService_DeleteResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadServiceServer).DeleteResource(ctx, req.(*DeleteResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UploadService_ServiceDesc is the grpc.ServiceDesc for UploadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UploadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UploadService",
	HandlerType: (*UploadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChangeResourceTitle",
			Handler:    _UploadService_ChangeResourceTitle_Handler,
		},
		{
			MethodName: "DeleteResource",
			Handler:    _UploadService_DeleteResource_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadVideo",
			Handler:       _UploadService_UploadVideo_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UploadImg",
			Handler:       _UploadService_UploadImg_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "resource.proto",
}
