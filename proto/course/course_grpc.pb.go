// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package course

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

// CourseServiceClient is the client API for CourseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourseServiceClient interface {
	GetCourses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CoursesResponse, error)
	AddCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*CourseId, error)
	DeleteCourse(ctx context.Context, in *CourseId, opts ...grpc.CallOption) (*MessageResponse, error)
	AddStudentsToCourse(ctx context.Context, in *StudentCourseIdRequest, opts ...grpc.CallOption) (*MessageResponse, error)
	RemoveStudentsFromCourse(ctx context.Context, in *StudentCourseIdRequest, opts ...grpc.CallOption) (*MessageResponse, error)
}

type courseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseServiceClient(cc grpc.ClientConnInterface) CourseServiceClient {
	return &courseServiceClient{cc}
}

func (c *courseServiceClient) GetCourses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CoursesResponse, error) {
	out := new(CoursesResponse)
	err := c.cc.Invoke(ctx, "/CourseService/GetCourses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) AddCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*CourseId, error) {
	out := new(CourseId)
	err := c.cc.Invoke(ctx, "/CourseService/AddCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) DeleteCourse(ctx context.Context, in *CourseId, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/CourseService/DeleteCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) AddStudentsToCourse(ctx context.Context, in *StudentCourseIdRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/CourseService/AddStudentsToCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) RemoveStudentsFromCourse(ctx context.Context, in *StudentCourseIdRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/CourseService/RemoveStudentsFromCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseServiceServer is the server API for CourseService service.
// All implementations must embed UnimplementedCourseServiceServer
// for forward compatibility
type CourseServiceServer interface {
	GetCourses(context.Context, *Empty) (*CoursesResponse, error)
	AddCourse(context.Context, *Course) (*CourseId, error)
	DeleteCourse(context.Context, *CourseId) (*MessageResponse, error)
	AddStudentsToCourse(context.Context, *StudentCourseIdRequest) (*MessageResponse, error)
	RemoveStudentsFromCourse(context.Context, *StudentCourseIdRequest) (*MessageResponse, error)
	mustEmbedUnimplementedCourseServiceServer()
}

// UnimplementedCourseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCourseServiceServer struct {
}

func (UnimplementedCourseServiceServer) GetCourses(context.Context, *Empty) (*CoursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourses not implemented")
}
func (UnimplementedCourseServiceServer) AddCourse(context.Context, *Course) (*CourseId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCourse not implemented")
}
func (UnimplementedCourseServiceServer) DeleteCourse(context.Context, *CourseId) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourse not implemented")
}
func (UnimplementedCourseServiceServer) AddStudentsToCourse(context.Context, *StudentCourseIdRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStudentsToCourse not implemented")
}
func (UnimplementedCourseServiceServer) RemoveStudentsFromCourse(context.Context, *StudentCourseIdRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveStudentsFromCourse not implemented")
}
func (UnimplementedCourseServiceServer) mustEmbedUnimplementedCourseServiceServer() {}

// UnsafeCourseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourseServiceServer will
// result in compilation errors.
type UnsafeCourseServiceServer interface {
	mustEmbedUnimplementedCourseServiceServer()
}

func RegisterCourseServiceServer(s grpc.ServiceRegistrar, srv CourseServiceServer) {
	s.RegisterService(&CourseService_ServiceDesc, srv)
}

func _CourseService_GetCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).GetCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CourseService/GetCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).GetCourses(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_AddCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Course)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).AddCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CourseService/AddCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).AddCourse(ctx, req.(*Course))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_DeleteCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).DeleteCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CourseService/DeleteCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).DeleteCourse(ctx, req.(*CourseId))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_AddStudentsToCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentCourseIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).AddStudentsToCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CourseService/AddStudentsToCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).AddStudentsToCourse(ctx, req.(*StudentCourseIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_RemoveStudentsFromCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentCourseIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).RemoveStudentsFromCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CourseService/RemoveStudentsFromCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).RemoveStudentsFromCourse(ctx, req.(*StudentCourseIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CourseService_ServiceDesc is the grpc.ServiceDesc for CourseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CourseService",
	HandlerType: (*CourseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCourses",
			Handler:    _CourseService_GetCourses_Handler,
		},
		{
			MethodName: "AddCourse",
			Handler:    _CourseService_AddCourse_Handler,
		},
		{
			MethodName: "DeleteCourse",
			Handler:    _CourseService_DeleteCourse_Handler,
		},
		{
			MethodName: "AddStudentsToCourse",
			Handler:    _CourseService_AddStudentsToCourse_Handler,
		},
		{
			MethodName: "RemoveStudentsFromCourse",
			Handler:    _CourseService_RemoveStudentsFromCourse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "course/course.proto",
}
