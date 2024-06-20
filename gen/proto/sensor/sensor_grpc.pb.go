// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/sensor/sensor.proto

package sensor

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

// SensorServiceClient is the client API for SensorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SensorServiceClient interface {
	CreateSensor(ctx context.Context, in *CreateSensorRequest, opts ...grpc.CallOption) (*CreateSensorResponse, error)
	GetSensor(ctx context.Context, in *GetSensorRequest, opts ...grpc.CallOption) (*GetSensorResponse, error)
	UpdateSensor(ctx context.Context, in *UpdateSensorRequest, opts ...grpc.CallOption) (*UpdateSensorResponse, error)
	DeleteSensor(ctx context.Context, in *DeleteSensorRequest, opts ...grpc.CallOption) (*DeleteSensorResponse, error)
}

type sensorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSensorServiceClient(cc grpc.ClientConnInterface) SensorServiceClient {
	return &sensorServiceClient{cc}
}

func (c *sensorServiceClient) CreateSensor(ctx context.Context, in *CreateSensorRequest, opts ...grpc.CallOption) (*CreateSensorResponse, error) {
	out := new(CreateSensorResponse)
	err := c.cc.Invoke(ctx, "/sensor.SensorService/CreateSensor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensorServiceClient) GetSensor(ctx context.Context, in *GetSensorRequest, opts ...grpc.CallOption) (*GetSensorResponse, error) {
	out := new(GetSensorResponse)
	err := c.cc.Invoke(ctx, "/sensor.SensorService/GetSensor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensorServiceClient) UpdateSensor(ctx context.Context, in *UpdateSensorRequest, opts ...grpc.CallOption) (*UpdateSensorResponse, error) {
	out := new(UpdateSensorResponse)
	err := c.cc.Invoke(ctx, "/sensor.SensorService/UpdateSensor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sensorServiceClient) DeleteSensor(ctx context.Context, in *DeleteSensorRequest, opts ...grpc.CallOption) (*DeleteSensorResponse, error) {
	out := new(DeleteSensorResponse)
	err := c.cc.Invoke(ctx, "/sensor.SensorService/DeleteSensor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SensorServiceServer is the server API for SensorService service.
// All implementations should embed UnimplementedSensorServiceServer
// for forward compatibility
type SensorServiceServer interface {
	CreateSensor(context.Context, *CreateSensorRequest) (*CreateSensorResponse, error)
	GetSensor(context.Context, *GetSensorRequest) (*GetSensorResponse, error)
	UpdateSensor(context.Context, *UpdateSensorRequest) (*UpdateSensorResponse, error)
	DeleteSensor(context.Context, *DeleteSensorRequest) (*DeleteSensorResponse, error)
}

// UnimplementedSensorServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSensorServiceServer struct {
}

func (UnimplementedSensorServiceServer) CreateSensor(context.Context, *CreateSensorRequest) (*CreateSensorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSensor not implemented")
}
func (UnimplementedSensorServiceServer) GetSensor(context.Context, *GetSensorRequest) (*GetSensorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSensor not implemented")
}
func (UnimplementedSensorServiceServer) UpdateSensor(context.Context, *UpdateSensorRequest) (*UpdateSensorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSensor not implemented")
}
func (UnimplementedSensorServiceServer) DeleteSensor(context.Context, *DeleteSensorRequest) (*DeleteSensorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSensor not implemented")
}

// UnsafeSensorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SensorServiceServer will
// result in compilation errors.
type UnsafeSensorServiceServer interface {
	mustEmbedUnimplementedSensorServiceServer()
}

func RegisterSensorServiceServer(s grpc.ServiceRegistrar, srv SensorServiceServer) {
	s.RegisterService(&SensorService_ServiceDesc, srv)
}

func _SensorService_CreateSensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSensorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorServiceServer).CreateSensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sensor.SensorService/CreateSensor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorServiceServer).CreateSensor(ctx, req.(*CreateSensorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SensorService_GetSensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSensorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorServiceServer).GetSensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sensor.SensorService/GetSensor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorServiceServer).GetSensor(ctx, req.(*GetSensorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SensorService_UpdateSensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSensorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorServiceServer).UpdateSensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sensor.SensorService/UpdateSensor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorServiceServer).UpdateSensor(ctx, req.(*UpdateSensorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SensorService_DeleteSensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSensorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SensorServiceServer).DeleteSensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sensor.SensorService/DeleteSensor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SensorServiceServer).DeleteSensor(ctx, req.(*DeleteSensorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SensorService_ServiceDesc is the grpc.ServiceDesc for SensorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SensorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sensor.SensorService",
	HandlerType: (*SensorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSensor",
			Handler:    _SensorService_CreateSensor_Handler,
		},
		{
			MethodName: "GetSensor",
			Handler:    _SensorService_GetSensor_Handler,
		},
		{
			MethodName: "UpdateSensor",
			Handler:    _SensorService_UpdateSensor_Handler,
		},
		{
			MethodName: "DeleteSensor",
			Handler:    _SensorService_DeleteSensor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/sensor/sensor.proto",
}
