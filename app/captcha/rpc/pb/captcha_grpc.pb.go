// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: captcha.proto

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
	CaptchaService_GetCaptchaByEmail_FullMethodName = "/pb.CaptchaService/GetCaptchaByEmail"
	CaptchaService_VerifyCaptcha_FullMethodName     = "/pb.CaptchaService/VerifyCaptcha"
)

// CaptchaServiceClient is the client API for CaptchaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaptchaServiceClient interface {
	GetCaptchaByEmail(ctx context.Context, in *GetCaptchaByEmailReq, opts ...grpc.CallOption) (*GetCaptchaByEmailResp, error)
	VerifyCaptcha(ctx context.Context, in *VerifyCaptchaReq, opts ...grpc.CallOption) (*VerifyCaptchaResp, error)
}

type captchaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCaptchaServiceClient(cc grpc.ClientConnInterface) CaptchaServiceClient {
	return &captchaServiceClient{cc}
}

func (c *captchaServiceClient) GetCaptchaByEmail(ctx context.Context, in *GetCaptchaByEmailReq, opts ...grpc.CallOption) (*GetCaptchaByEmailResp, error) {
	out := new(GetCaptchaByEmailResp)
	err := c.cc.Invoke(ctx, CaptchaService_GetCaptchaByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captchaServiceClient) VerifyCaptcha(ctx context.Context, in *VerifyCaptchaReq, opts ...grpc.CallOption) (*VerifyCaptchaResp, error) {
	out := new(VerifyCaptchaResp)
	err := c.cc.Invoke(ctx, CaptchaService_VerifyCaptcha_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaptchaServiceServer is the server API for CaptchaService service.
// All implementations must embed UnimplementedCaptchaServiceServer
// for forward compatibility
type CaptchaServiceServer interface {
	GetCaptchaByEmail(context.Context, *GetCaptchaByEmailReq) (*GetCaptchaByEmailResp, error)
	VerifyCaptcha(context.Context, *VerifyCaptchaReq) (*VerifyCaptchaResp, error)
	mustEmbedUnimplementedCaptchaServiceServer()
}

// UnimplementedCaptchaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCaptchaServiceServer struct {
}

func (UnimplementedCaptchaServiceServer) GetCaptchaByEmail(context.Context, *GetCaptchaByEmailReq) (*GetCaptchaByEmailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCaptchaByEmail not implemented")
}
func (UnimplementedCaptchaServiceServer) VerifyCaptcha(context.Context, *VerifyCaptchaReq) (*VerifyCaptchaResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCaptcha not implemented")
}
func (UnimplementedCaptchaServiceServer) mustEmbedUnimplementedCaptchaServiceServer() {}

// UnsafeCaptchaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaptchaServiceServer will
// result in compilation errors.
type UnsafeCaptchaServiceServer interface {
	mustEmbedUnimplementedCaptchaServiceServer()
}

func RegisterCaptchaServiceServer(s grpc.ServiceRegistrar, srv CaptchaServiceServer) {
	s.RegisterService(&CaptchaService_ServiceDesc, srv)
}

func _CaptchaService_GetCaptchaByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaptchaByEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptchaServiceServer).GetCaptchaByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptchaService_GetCaptchaByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptchaServiceServer).GetCaptchaByEmail(ctx, req.(*GetCaptchaByEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptchaService_VerifyCaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCaptchaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptchaServiceServer).VerifyCaptcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptchaService_VerifyCaptcha_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptchaServiceServer).VerifyCaptcha(ctx, req.(*VerifyCaptchaReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CaptchaService_ServiceDesc is the grpc.ServiceDesc for CaptchaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CaptchaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CaptchaService",
	HandlerType: (*CaptchaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCaptchaByEmail",
			Handler:    _CaptchaService_GetCaptchaByEmail_Handler,
		},
		{
			MethodName: "VerifyCaptcha",
			Handler:    _CaptchaService_VerifyCaptcha_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "captcha.proto",
}
