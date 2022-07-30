// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: pkg/filebox/proto/filebox.proto

package proto

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

// FileboxServiceClient is the client API for FileboxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileboxServiceClient interface {
	SendFile(ctx context.Context, in *SendFileRequest, opts ...grpc.CallOption) (*SendFileResponse, error)
	UpdateFile(ctx context.Context, in *UpdateFileRequest, opts ...grpc.CallOption) (*UpdateFileResponse, error)
	GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*GetFileResponse, error)
	GetListOfAllFiles(ctx context.Context, in *GetListOfAllFilesRequest, opts ...grpc.CallOption) (*GetListOfAllFilesResponse, error)
	DeleteFile(ctx context.Context, in *DeleteFileRequest, opts ...grpc.CallOption) (*DeleteFileResponse, error)
	SendFileToPerson(ctx context.Context, in *SendFileToPersonRequest, opts ...grpc.CallOption) (*SendFileToPersonResponse, error)
}

type fileboxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileboxServiceClient(cc grpc.ClientConnInterface) FileboxServiceClient {
	return &fileboxServiceClient{cc}
}

func (c *fileboxServiceClient) SendFile(ctx context.Context, in *SendFileRequest, opts ...grpc.CallOption) (*SendFileResponse, error) {
	out := new(SendFileResponse)
	err := c.cc.Invoke(ctx, "/filebox.FileboxService/SendFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileboxServiceClient) UpdateFile(ctx context.Context, in *UpdateFileRequest, opts ...grpc.CallOption) (*UpdateFileResponse, error) {
	out := new(UpdateFileResponse)
	err := c.cc.Invoke(ctx, "/filebox.FileboxService/UpdateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileboxServiceClient) GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*GetFileResponse, error) {
	out := new(GetFileResponse)
	err := c.cc.Invoke(ctx, "/filebox.FileboxService/GetFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileboxServiceClient) GetListOfAllFiles(ctx context.Context, in *GetListOfAllFilesRequest, opts ...grpc.CallOption) (*GetListOfAllFilesResponse, error) {
	out := new(GetListOfAllFilesResponse)
	err := c.cc.Invoke(ctx, "/filebox.FileboxService/GetListOfAllFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileboxServiceClient) DeleteFile(ctx context.Context, in *DeleteFileRequest, opts ...grpc.CallOption) (*DeleteFileResponse, error) {
	out := new(DeleteFileResponse)
	err := c.cc.Invoke(ctx, "/filebox.FileboxService/DeleteFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileboxServiceClient) SendFileToPerson(ctx context.Context, in *SendFileToPersonRequest, opts ...grpc.CallOption) (*SendFileToPersonResponse, error) {
	out := new(SendFileToPersonResponse)
	err := c.cc.Invoke(ctx, "/filebox.FileboxService/SendFileToPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileboxServiceServer is the server API for FileboxService service.
// All implementations must embed UnimplementedFileboxServiceServer
// for forward compatibility
type FileboxServiceServer interface {
	SendFile(context.Context, *SendFileRequest) (*SendFileResponse, error)
	UpdateFile(context.Context, *UpdateFileRequest) (*UpdateFileResponse, error)
	GetFile(context.Context, *GetFileRequest) (*GetFileResponse, error)
	GetListOfAllFiles(context.Context, *GetListOfAllFilesRequest) (*GetListOfAllFilesResponse, error)
	DeleteFile(context.Context, *DeleteFileRequest) (*DeleteFileResponse, error)
	SendFileToPerson(context.Context, *SendFileToPersonRequest) (*SendFileToPersonResponse, error)
	mustEmbedUnimplementedFileboxServiceServer()
}

// UnimplementedFileboxServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileboxServiceServer struct {
}

func (UnimplementedFileboxServiceServer) SendFile(context.Context, *SendFileRequest) (*SendFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendFile not implemented")
}
func (UnimplementedFileboxServiceServer) UpdateFile(context.Context, *UpdateFileRequest) (*UpdateFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFile not implemented")
}
func (UnimplementedFileboxServiceServer) GetFile(context.Context, *GetFileRequest) (*GetFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
func (UnimplementedFileboxServiceServer) GetListOfAllFiles(context.Context, *GetListOfAllFilesRequest) (*GetListOfAllFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListOfAllFiles not implemented")
}
func (UnimplementedFileboxServiceServer) DeleteFile(context.Context, *DeleteFileRequest) (*DeleteFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedFileboxServiceServer) SendFileToPerson(context.Context, *SendFileToPersonRequest) (*SendFileToPersonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendFileToPerson not implemented")
}
func (UnimplementedFileboxServiceServer) mustEmbedUnimplementedFileboxServiceServer() {}

// UnsafeFileboxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileboxServiceServer will
// result in compilation errors.
type UnsafeFileboxServiceServer interface {
	mustEmbedUnimplementedFileboxServiceServer()
}

func RegisterFileboxServiceServer(s grpc.ServiceRegistrar, srv FileboxServiceServer) {
	s.RegisterService(&FileboxService_ServiceDesc, srv)
}

func _FileboxService_SendFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileboxServiceServer).SendFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filebox.FileboxService/SendFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileboxServiceServer).SendFile(ctx, req.(*SendFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileboxService_UpdateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileboxServiceServer).UpdateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filebox.FileboxService/UpdateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileboxServiceServer).UpdateFile(ctx, req.(*UpdateFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileboxService_GetFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileboxServiceServer).GetFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filebox.FileboxService/GetFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileboxServiceServer).GetFile(ctx, req.(*GetFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileboxService_GetListOfAllFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListOfAllFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileboxServiceServer).GetListOfAllFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filebox.FileboxService/GetListOfAllFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileboxServiceServer).GetListOfAllFiles(ctx, req.(*GetListOfAllFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileboxService_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileboxServiceServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filebox.FileboxService/DeleteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileboxServiceServer).DeleteFile(ctx, req.(*DeleteFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileboxService_SendFileToPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendFileToPersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileboxServiceServer).SendFileToPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filebox.FileboxService/SendFileToPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileboxServiceServer).SendFileToPerson(ctx, req.(*SendFileToPersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileboxService_ServiceDesc is the grpc.ServiceDesc for FileboxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileboxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "filebox.FileboxService",
	HandlerType: (*FileboxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendFile",
			Handler:    _FileboxService_SendFile_Handler,
		},
		{
			MethodName: "UpdateFile",
			Handler:    _FileboxService_UpdateFile_Handler,
		},
		{
			MethodName: "GetFile",
			Handler:    _FileboxService_GetFile_Handler,
		},
		{
			MethodName: "GetListOfAllFiles",
			Handler:    _FileboxService_GetListOfAllFiles_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _FileboxService_DeleteFile_Handler,
		},
		{
			MethodName: "SendFileToPerson",
			Handler:    _FileboxService_SendFileToPerson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/filebox/proto/filebox.proto",
}
