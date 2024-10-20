// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: communication/fuse.proto

package communication

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FileSystem_AddDirectory_FullMethodName    = "/fuse.FileSystem/AddDirectory"
	FileSystem_RenameDirectory_FullMethodName = "/fuse.FileSystem/RenameDirectory"
	FileSystem_RemoveDirectory_FullMethodName = "/fuse.FileSystem/RemoveDirectory"
	FileSystem_AddFile_FullMethodName         = "/fuse.FileSystem/AddFile"
	FileSystem_RenameFile_FullMethodName      = "/fuse.FileSystem/RenameFile"
	FileSystem_RemoveFile_FullMethodName      = "/fuse.FileSystem/RemoveFile"
)

// FileSystemClient is the client API for FileSystem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileSystemClient interface {
	AddDirectory(ctx context.Context, in *AddDirectoryRequest, opts ...grpc.CallOption) (*DirectoryResponse, error)
	RenameDirectory(ctx context.Context, in *RenameDirectoryRequest, opts ...grpc.CallOption) (*DirectoryResponse, error)
	RemoveDirectory(ctx context.Context, in *RemoveDirectoryRequest, opts ...grpc.CallOption) (*DirectoryResponse, error)
	AddFile(ctx context.Context, in *AddFileRequest, opts ...grpc.CallOption) (*FileResponse, error)
	RenameFile(ctx context.Context, in *RenameFileRequest, opts ...grpc.CallOption) (*FileResponse, error)
	RemoveFile(ctx context.Context, in *RemoveFileRequest, opts ...grpc.CallOption) (*FileResponse, error)
}

type fileSystemClient struct {
	cc grpc.ClientConnInterface
}

func NewFileSystemClient(cc grpc.ClientConnInterface) FileSystemClient {
	return &fileSystemClient{cc}
}

func (c *fileSystemClient) AddDirectory(ctx context.Context, in *AddDirectoryRequest, opts ...grpc.CallOption) (*DirectoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DirectoryResponse)
	err := c.cc.Invoke(ctx, FileSystem_AddDirectory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemClient) RenameDirectory(ctx context.Context, in *RenameDirectoryRequest, opts ...grpc.CallOption) (*DirectoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DirectoryResponse)
	err := c.cc.Invoke(ctx, FileSystem_RenameDirectory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemClient) RemoveDirectory(ctx context.Context, in *RemoveDirectoryRequest, opts ...grpc.CallOption) (*DirectoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DirectoryResponse)
	err := c.cc.Invoke(ctx, FileSystem_RemoveDirectory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemClient) AddFile(ctx context.Context, in *AddFileRequest, opts ...grpc.CallOption) (*FileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileResponse)
	err := c.cc.Invoke(ctx, FileSystem_AddFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemClient) RenameFile(ctx context.Context, in *RenameFileRequest, opts ...grpc.CallOption) (*FileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileResponse)
	err := c.cc.Invoke(ctx, FileSystem_RenameFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileSystemClient) RemoveFile(ctx context.Context, in *RemoveFileRequest, opts ...grpc.CallOption) (*FileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FileResponse)
	err := c.cc.Invoke(ctx, FileSystem_RemoveFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileSystemServer is the server API for FileSystem service.
// All implementations must embed UnimplementedFileSystemServer
// for forward compatibility.
type FileSystemServer interface {
	AddDirectory(context.Context, *AddDirectoryRequest) (*DirectoryResponse, error)
	RenameDirectory(context.Context, *RenameDirectoryRequest) (*DirectoryResponse, error)
	RemoveDirectory(context.Context, *RemoveDirectoryRequest) (*DirectoryResponse, error)
	AddFile(context.Context, *AddFileRequest) (*FileResponse, error)
	RenameFile(context.Context, *RenameFileRequest) (*FileResponse, error)
	RemoveFile(context.Context, *RemoveFileRequest) (*FileResponse, error)
	mustEmbedUnimplementedFileSystemServer()
}

// UnimplementedFileSystemServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFileSystemServer struct{}

func (UnimplementedFileSystemServer) AddDirectory(context.Context, *AddDirectoryRequest) (*DirectoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDirectory not implemented")
}
func (UnimplementedFileSystemServer) RenameDirectory(context.Context, *RenameDirectoryRequest) (*DirectoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameDirectory not implemented")
}
func (UnimplementedFileSystemServer) RemoveDirectory(context.Context, *RemoveDirectoryRequest) (*DirectoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDirectory not implemented")
}
func (UnimplementedFileSystemServer) AddFile(context.Context, *AddFileRequest) (*FileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFile not implemented")
}
func (UnimplementedFileSystemServer) RenameFile(context.Context, *RenameFileRequest) (*FileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameFile not implemented")
}
func (UnimplementedFileSystemServer) RemoveFile(context.Context, *RemoveFileRequest) (*FileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFile not implemented")
}
func (UnimplementedFileSystemServer) mustEmbedUnimplementedFileSystemServer() {}
func (UnimplementedFileSystemServer) testEmbeddedByValue()                    {}

// UnsafeFileSystemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileSystemServer will
// result in compilation errors.
type UnsafeFileSystemServer interface {
	mustEmbedUnimplementedFileSystemServer()
}

func RegisterFileSystemServer(s grpc.ServiceRegistrar, srv FileSystemServer) {
	// If the following call panics, it indicates UnimplementedFileSystemServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FileSystem_ServiceDesc, srv)
}

func _FileSystem_AddDirectory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSystemServer).AddDirectory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileSystem_AddDirectory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSystemServer).AddDirectory(ctx, req.(*AddDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSystem_RenameDirectory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSystemServer).RenameDirectory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileSystem_RenameDirectory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSystemServer).RenameDirectory(ctx, req.(*RenameDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSystem_RemoveDirectory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSystemServer).RemoveDirectory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileSystem_RemoveDirectory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSystemServer).RemoveDirectory(ctx, req.(*RemoveDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSystem_AddFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSystemServer).AddFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileSystem_AddFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSystemServer).AddFile(ctx, req.(*AddFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSystem_RenameFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSystemServer).RenameFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileSystem_RenameFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSystemServer).RenameFile(ctx, req.(*RenameFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileSystem_RemoveFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSystemServer).RemoveFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileSystem_RemoveFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSystemServer).RemoveFile(ctx, req.(*RemoveFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileSystem_ServiceDesc is the grpc.ServiceDesc for FileSystem service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileSystem_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fuse.FileSystem",
	HandlerType: (*FileSystemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDirectory",
			Handler:    _FileSystem_AddDirectory_Handler,
		},
		{
			MethodName: "RenameDirectory",
			Handler:    _FileSystem_RenameDirectory_Handler,
		},
		{
			MethodName: "RemoveDirectory",
			Handler:    _FileSystem_RemoveDirectory_Handler,
		},
		{
			MethodName: "AddFile",
			Handler:    _FileSystem_AddFile_Handler,
		},
		{
			MethodName: "RenameFile",
			Handler:    _FileSystem_RenameFile_Handler,
		},
		{
			MethodName: "RemoveFile",
			Handler:    _FileSystem_RemoveFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "communication/fuse.proto",
}
