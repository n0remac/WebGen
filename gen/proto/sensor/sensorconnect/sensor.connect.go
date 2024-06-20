// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/sensor/sensor.proto

package sensorconnect

import (
	sensor "Go-gRPC-React-starter/gen/proto/sensor"
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// SensorServiceName is the fully-qualified name of the SensorService service.
	SensorServiceName = "sensor.SensorService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SensorServiceCreateSensorProcedure is the fully-qualified name of the SensorService's
	// CreateSensor RPC.
	SensorServiceCreateSensorProcedure = "/sensor.SensorService/CreateSensor"
	// SensorServiceGetSensorProcedure is the fully-qualified name of the SensorService's GetSensor RPC.
	SensorServiceGetSensorProcedure = "/sensor.SensorService/GetSensor"
	// SensorServiceUpdateSensorProcedure is the fully-qualified name of the SensorService's
	// UpdateSensor RPC.
	SensorServiceUpdateSensorProcedure = "/sensor.SensorService/UpdateSensor"
	// SensorServiceDeleteSensorProcedure is the fully-qualified name of the SensorService's
	// DeleteSensor RPC.
	SensorServiceDeleteSensorProcedure = "/sensor.SensorService/DeleteSensor"
)

// SensorServiceClient is a client for the sensor.SensorService service.
type SensorServiceClient interface {
	CreateSensor(context.Context, *connect_go.Request[sensor.CreateSensorRequest]) (*connect_go.Response[sensor.CreateSensorResponse], error)
	GetSensor(context.Context, *connect_go.Request[sensor.GetSensorRequest]) (*connect_go.Response[sensor.GetSensorResponse], error)
	UpdateSensor(context.Context, *connect_go.Request[sensor.UpdateSensorRequest]) (*connect_go.Response[sensor.UpdateSensorResponse], error)
	DeleteSensor(context.Context, *connect_go.Request[sensor.DeleteSensorRequest]) (*connect_go.Response[sensor.DeleteSensorResponse], error)
}

// NewSensorServiceClient constructs a client for the sensor.SensorService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSensorServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) SensorServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &sensorServiceClient{
		createSensor: connect_go.NewClient[sensor.CreateSensorRequest, sensor.CreateSensorResponse](
			httpClient,
			baseURL+SensorServiceCreateSensorProcedure,
			opts...,
		),
		getSensor: connect_go.NewClient[sensor.GetSensorRequest, sensor.GetSensorResponse](
			httpClient,
			baseURL+SensorServiceGetSensorProcedure,
			opts...,
		),
		updateSensor: connect_go.NewClient[sensor.UpdateSensorRequest, sensor.UpdateSensorResponse](
			httpClient,
			baseURL+SensorServiceUpdateSensorProcedure,
			opts...,
		),
		deleteSensor: connect_go.NewClient[sensor.DeleteSensorRequest, sensor.DeleteSensorResponse](
			httpClient,
			baseURL+SensorServiceDeleteSensorProcedure,
			opts...,
		),
	}
}

// sensorServiceClient implements SensorServiceClient.
type sensorServiceClient struct {
	createSensor *connect_go.Client[sensor.CreateSensorRequest, sensor.CreateSensorResponse]
	getSensor    *connect_go.Client[sensor.GetSensorRequest, sensor.GetSensorResponse]
	updateSensor *connect_go.Client[sensor.UpdateSensorRequest, sensor.UpdateSensorResponse]
	deleteSensor *connect_go.Client[sensor.DeleteSensorRequest, sensor.DeleteSensorResponse]
}

// CreateSensor calls sensor.SensorService.CreateSensor.
func (c *sensorServiceClient) CreateSensor(ctx context.Context, req *connect_go.Request[sensor.CreateSensorRequest]) (*connect_go.Response[sensor.CreateSensorResponse], error) {
	return c.createSensor.CallUnary(ctx, req)
}

// GetSensor calls sensor.SensorService.GetSensor.
func (c *sensorServiceClient) GetSensor(ctx context.Context, req *connect_go.Request[sensor.GetSensorRequest]) (*connect_go.Response[sensor.GetSensorResponse], error) {
	return c.getSensor.CallUnary(ctx, req)
}

// UpdateSensor calls sensor.SensorService.UpdateSensor.
func (c *sensorServiceClient) UpdateSensor(ctx context.Context, req *connect_go.Request[sensor.UpdateSensorRequest]) (*connect_go.Response[sensor.UpdateSensorResponse], error) {
	return c.updateSensor.CallUnary(ctx, req)
}

// DeleteSensor calls sensor.SensorService.DeleteSensor.
func (c *sensorServiceClient) DeleteSensor(ctx context.Context, req *connect_go.Request[sensor.DeleteSensorRequest]) (*connect_go.Response[sensor.DeleteSensorResponse], error) {
	return c.deleteSensor.CallUnary(ctx, req)
}

// SensorServiceHandler is an implementation of the sensor.SensorService service.
type SensorServiceHandler interface {
	CreateSensor(context.Context, *connect_go.Request[sensor.CreateSensorRequest]) (*connect_go.Response[sensor.CreateSensorResponse], error)
	GetSensor(context.Context, *connect_go.Request[sensor.GetSensorRequest]) (*connect_go.Response[sensor.GetSensorResponse], error)
	UpdateSensor(context.Context, *connect_go.Request[sensor.UpdateSensorRequest]) (*connect_go.Response[sensor.UpdateSensorResponse], error)
	DeleteSensor(context.Context, *connect_go.Request[sensor.DeleteSensorRequest]) (*connect_go.Response[sensor.DeleteSensorResponse], error)
}

// NewSensorServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSensorServiceHandler(svc SensorServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	sensorServiceCreateSensorHandler := connect_go.NewUnaryHandler(
		SensorServiceCreateSensorProcedure,
		svc.CreateSensor,
		opts...,
	)
	sensorServiceGetSensorHandler := connect_go.NewUnaryHandler(
		SensorServiceGetSensorProcedure,
		svc.GetSensor,
		opts...,
	)
	sensorServiceUpdateSensorHandler := connect_go.NewUnaryHandler(
		SensorServiceUpdateSensorProcedure,
		svc.UpdateSensor,
		opts...,
	)
	sensorServiceDeleteSensorHandler := connect_go.NewUnaryHandler(
		SensorServiceDeleteSensorProcedure,
		svc.DeleteSensor,
		opts...,
	)
	return "/sensor.SensorService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SensorServiceCreateSensorProcedure:
			sensorServiceCreateSensorHandler.ServeHTTP(w, r)
		case SensorServiceGetSensorProcedure:
			sensorServiceGetSensorHandler.ServeHTTP(w, r)
		case SensorServiceUpdateSensorProcedure:
			sensorServiceUpdateSensorHandler.ServeHTTP(w, r)
		case SensorServiceDeleteSensorProcedure:
			sensorServiceDeleteSensorHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSensorServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSensorServiceHandler struct{}

func (UnimplementedSensorServiceHandler) CreateSensor(context.Context, *connect_go.Request[sensor.CreateSensorRequest]) (*connect_go.Response[sensor.CreateSensorResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sensor.SensorService.CreateSensor is not implemented"))
}

func (UnimplementedSensorServiceHandler) GetSensor(context.Context, *connect_go.Request[sensor.GetSensorRequest]) (*connect_go.Response[sensor.GetSensorResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sensor.SensorService.GetSensor is not implemented"))
}

func (UnimplementedSensorServiceHandler) UpdateSensor(context.Context, *connect_go.Request[sensor.UpdateSensorRequest]) (*connect_go.Response[sensor.UpdateSensorResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sensor.SensorService.UpdateSensor is not implemented"))
}

func (UnimplementedSensorServiceHandler) DeleteSensor(context.Context, *connect_go.Request[sensor.DeleteSensorRequest]) (*connect_go.Response[sensor.DeleteSensorResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sensor.SensorService.DeleteSensor is not implemented"))
}
