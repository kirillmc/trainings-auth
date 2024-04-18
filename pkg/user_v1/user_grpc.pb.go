// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: user.proto

package user_v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserV1Client is the client API for UserV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserV1Client interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	LockUser(ctx context.Context, in *LockUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UnlockUser(ctx context.Context, in *UnlockUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type userV1Client struct {
	cc grpc.ClientConnInterface
}

func NewUserV1Client(cc grpc.ClientConnInterface) UserV1Client {
	return &userV1Client{cc}
}

func (c *userV1Client) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/user_v1.UserV1/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userV1Client) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/user_v1.UserV1/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userV1Client) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/user_v1.UserV1/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userV1Client) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/user_v1.UserV1/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userV1Client) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/user_v1.UserV1/UpdateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userV1Client) LockUser(ctx context.Context, in *LockUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/user_v1.UserV1/LockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userV1Client) UnlockUser(ctx context.Context, in *UnlockUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/user_v1.UserV1/UnlockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userV1Client) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/user_v1.UserV1/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserV1Server is the server API for UserV1 service.
// All implementations must embed UnimplementedUserV1Server
// for forward compatibility
type UserV1Server interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Update(context.Context, *UpdateRequest) (*empty.Empty, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*empty.Empty, error)
	UpdateRole(context.Context, *UpdateRoleRequest) (*empty.Empty, error)
	LockUser(context.Context, *LockUserRequest) (*empty.Empty, error)
	UnlockUser(context.Context, *UnlockUserRequest) (*empty.Empty, error)
	Delete(context.Context, *DeleteRequest) (*empty.Empty, error)
	mustEmbedUnimplementedUserV1Server()
}

// UnimplementedUserV1Server must be embedded to have forward compatible implementations.
type UnimplementedUserV1Server struct {
}

func (UnimplementedUserV1Server) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserV1Server) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserV1Server) Update(context.Context, *UpdateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserV1Server) UpdatePassword(context.Context, *UpdatePasswordRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedUserV1Server) UpdateRole(context.Context, *UpdateRoleRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedUserV1Server) LockUser(context.Context, *LockUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockUser not implemented")
}
func (UnimplementedUserV1Server) UnlockUser(context.Context, *UnlockUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlockUser not implemented")
}
func (UnimplementedUserV1Server) Delete(context.Context, *DeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserV1Server) mustEmbedUnimplementedUserV1Server() {}

// UnsafeUserV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserV1Server will
// result in compilation errors.
type UnsafeUserV1Server interface {
	mustEmbedUnimplementedUserV1Server()
}

func RegisterUserV1Server(s grpc.ServiceRegistrar, srv UserV1Server) {
	s.RegisterService(&UserV1_ServiceDesc, srv)
}

func _UserV1_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserV1Server).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_v1.UserV1/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserV1Server).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserV1_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserV1Server).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_v1.UserV1/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserV1Server).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserV1_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserV1Server).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_v1.UserV1/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserV1Server).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserV1_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserV1Server).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_v1.UserV1/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserV1Server).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserV1_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserV1Server).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_v1.UserV1/UpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserV1Server).UpdateRole(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserV1_LockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserV1Server).LockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_v1.UserV1/LockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserV1Server).LockUser(ctx, req.(*LockUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserV1_UnlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserV1Server).UnlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_v1.UserV1/UnlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserV1Server).UnlockUser(ctx, req.(*UnlockUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserV1_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserV1Server).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_v1.UserV1/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserV1Server).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserV1_ServiceDesc is the grpc.ServiceDesc for UserV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_v1.UserV1",
	HandlerType: (*UserV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserV1_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserV1_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserV1_Update_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _UserV1_UpdatePassword_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _UserV1_UpdateRole_Handler,
		},
		{
			MethodName: "LockUser",
			Handler:    _UserV1_LockUser_Handler,
		},
		{
			MethodName: "UnlockUser",
			Handler:    _UserV1_UnlockUser_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserV1_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}