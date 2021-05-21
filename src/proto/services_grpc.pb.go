// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package _

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

// EmailTokenClient is the client API for EmailToken service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailTokenClient interface {
	GetEmailAndToken(ctx context.Context, in *EmailTokenRequest, opts ...grpc.CallOption) (*EmailTokenResponse, error)
}

type emailTokenClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailTokenClient(cc grpc.ClientConnInterface) EmailTokenClient {
	return &emailTokenClient{cc}
}

func (c *emailTokenClient) GetEmailAndToken(ctx context.Context, in *EmailTokenRequest, opts ...grpc.CallOption) (*EmailTokenResponse, error) {
	out := new(EmailTokenResponse)
	err := c.cc.Invoke(ctx, "/EmailToken.EmailToken/getEmailAndToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailTokenServer is the server API for EmailToken service.
// All implementations must embed UnimplementedEmailTokenServer
// for forward compatibility
type EmailTokenServer interface {
	GetEmailAndToken(context.Context, *EmailTokenRequest) (*EmailTokenResponse, error)
	mustEmbedUnimplementedEmailTokenServer()
}

// UnimplementedEmailTokenServer must be embedded to have forward compatible implementations.
type UnimplementedEmailTokenServer struct {
}

func (UnimplementedEmailTokenServer) GetEmailAndToken(context.Context, *EmailTokenRequest) (*EmailTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmailAndToken not implemented")
}
func (UnimplementedEmailTokenServer) mustEmbedUnimplementedEmailTokenServer() {}

// UnsafeEmailTokenServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailTokenServer will
// result in compilation errors.
type UnsafeEmailTokenServer interface {
	mustEmbedUnimplementedEmailTokenServer()
}

func RegisterEmailTokenServer(s grpc.ServiceRegistrar, srv EmailTokenServer) {
	s.RegisterService(&EmailToken_ServiceDesc, srv)
}

func _EmailToken_GetEmailAndToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailTokenServer).GetEmailAndToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmailToken.EmailToken/getEmailAndToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailTokenServer).GetEmailAndToken(ctx, req.(*EmailTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmailToken_ServiceDesc is the grpc.ServiceDesc for EmailToken service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailToken_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EmailToken.EmailToken",
	HandlerType: (*EmailTokenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getEmailAndToken",
			Handler:    _EmailToken_GetEmailAndToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}