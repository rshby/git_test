// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: proto/student.proto

package pb

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
	StudentService_RegisterStudent_FullMethodName      = "/git_test.proto.StudentService/RegisterStudent"
	StudentService_GetStudentByID_FullMethodName       = "/git_test.proto.StudentService/GetStudentByID"
	StudentService_GetStudentBySchoolID_FullMethodName = "/git_test.proto.StudentService/GetStudentBySchoolID"
)

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentServiceClient interface {
	RegisterStudent(ctx context.Context, in *CreateStudentRequestDTO, opts ...grpc.CallOption) (*CreateStudentResponseDTO, error)
	GetStudentByID(ctx context.Context, in *GetStudentByIDRequestDTO, opts ...grpc.CallOption) (*Student, error)
	GetStudentBySchoolID(ctx context.Context, in *GetStudentBySchoolIDRequestDTO, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Student], error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) RegisterStudent(ctx context.Context, in *CreateStudentRequestDTO, opts ...grpc.CallOption) (*CreateStudentResponseDTO, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateStudentResponseDTO)
	err := c.cc.Invoke(ctx, StudentService_RegisterStudent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetStudentByID(ctx context.Context, in *GetStudentByIDRequestDTO, opts ...grpc.CallOption) (*Student, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Student)
	err := c.cc.Invoke(ctx, StudentService_GetStudentByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetStudentBySchoolID(ctx context.Context, in *GetStudentBySchoolIDRequestDTO, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Student], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StudentService_ServiceDesc.Streams[0], StudentService_GetStudentBySchoolID_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetStudentBySchoolIDRequestDTO, Student]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StudentService_GetStudentBySchoolIDClient = grpc.ServerStreamingClient[Student]

// StudentServiceServer is the server API for StudentService service.
// All implementations must embed UnimplementedStudentServiceServer
// for forward compatibility.
type StudentServiceServer interface {
	RegisterStudent(context.Context, *CreateStudentRequestDTO) (*CreateStudentResponseDTO, error)
	GetStudentByID(context.Context, *GetStudentByIDRequestDTO) (*Student, error)
	GetStudentBySchoolID(*GetStudentBySchoolIDRequestDTO, grpc.ServerStreamingServer[Student]) error
	mustEmbedUnimplementedStudentServiceServer()
}

// UnimplementedStudentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStudentServiceServer struct{}

func (UnimplementedStudentServiceServer) RegisterStudent(context.Context, *CreateStudentRequestDTO) (*CreateStudentResponseDTO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterStudent not implemented")
}
func (UnimplementedStudentServiceServer) GetStudentByID(context.Context, *GetStudentByIDRequestDTO) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentByID not implemented")
}
func (UnimplementedStudentServiceServer) GetStudentBySchoolID(*GetStudentBySchoolIDRequestDTO, grpc.ServerStreamingServer[Student]) error {
	return status.Errorf(codes.Unimplemented, "method GetStudentBySchoolID not implemented")
}
func (UnimplementedStudentServiceServer) mustEmbedUnimplementedStudentServiceServer() {}
func (UnimplementedStudentServiceServer) testEmbeddedByValue()                        {}

// UnsafeStudentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServiceServer will
// result in compilation errors.
type UnsafeStudentServiceServer interface {
	mustEmbedUnimplementedStudentServiceServer()
}

func RegisterStudentServiceServer(s grpc.ServiceRegistrar, srv StudentServiceServer) {
	// If the following call pancis, it indicates UnimplementedStudentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StudentService_ServiceDesc, srv)
}

func _StudentService_RegisterStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStudentRequestDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).RegisterStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_RegisterStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).RegisterStudent(ctx, req.(*CreateStudentRequestDTO))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetStudentByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentByIDRequestDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetStudentByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_GetStudentByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetStudentByID(ctx, req.(*GetStudentByIDRequestDTO))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetStudentBySchoolID_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetStudentBySchoolIDRequestDTO)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StudentServiceServer).GetStudentBySchoolID(m, &grpc.GenericServerStream[GetStudentBySchoolIDRequestDTO, Student]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StudentService_GetStudentBySchoolIDServer = grpc.ServerStreamingServer[Student]

// StudentService_ServiceDesc is the grpc.ServiceDesc for StudentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "git_test.proto.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterStudent",
			Handler:    _StudentService_RegisterStudent_Handler,
		},
		{
			MethodName: "GetStudentByID",
			Handler:    _StudentService_GetStudentByID_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStudentBySchoolID",
			Handler:       _StudentService_GetStudentBySchoolID_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/student.proto",
}
