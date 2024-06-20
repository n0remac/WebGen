// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: proto/sensor/sensor.proto

package sensor

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateSensorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sensor *Sensor `protobuf:"bytes,1,opt,name=sensor,proto3" json:"sensor,omitempty"`
}

func (x *CreateSensorRequest) Reset() {
	*x = CreateSensorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sensor_sensor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSensorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSensorRequest) ProtoMessage() {}

func (x *CreateSensorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sensor_sensor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSensorRequest.ProtoReflect.Descriptor instead.
func (*CreateSensorRequest) Descriptor() ([]byte, []int) {
	return file_proto_sensor_sensor_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSensorRequest) GetSensor() *Sensor {
	if x != nil {
		return x.Sensor
	}
	return nil
}

type CreateSensorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sensor *Sensor `protobuf:"bytes,1,opt,name=sensor,proto3" json:"sensor,omitempty"`
}

func (x *CreateSensorResponse) Reset() {
	*x = CreateSensorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sensor_sensor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSensorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSensorResponse) ProtoMessage() {}

func (x *CreateSensorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sensor_sensor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSensorResponse.ProtoReflect.Descriptor instead.
func (*CreateSensorResponse) Descriptor() ([]byte, []int) {
	return file_proto_sensor_sensor_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSensorResponse) GetSensor() *Sensor {
	if x != nil {
		return x.Sensor
	}
	return nil
}

type GetSensorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSensorRequest) Reset() {
	*x = GetSensorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sensor_sensor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSensorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSensorRequest) ProtoMessage() {}

func (x *GetSensorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sensor_sensor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSensorRequest.ProtoReflect.Descriptor instead.
func (*GetSensorRequest) Descriptor() ([]byte, []int) {
	return file_proto_sensor_sensor_proto_rawDescGZIP(), []int{2}
}

func (x *GetSensorRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetSensorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sensor *Sensor `protobuf:"bytes,1,opt,name=sensor,proto3" json:"sensor,omitempty"`
}

func (x *GetSensorResponse) Reset() {
	*x = GetSensorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sensor_sensor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSensorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSensorResponse) ProtoMessage() {}

func (x *GetSensorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sensor_sensor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSensorResponse.ProtoReflect.Descriptor instead.
func (*GetSensorResponse) Descriptor() ([]byte, []int) {
	return file_proto_sensor_sensor_proto_rawDescGZIP(), []int{3}
}

func (x *GetSensorResponse) GetSensor() *Sensor {
	if x != nil {
		return x.Sensor
	}
	return nil
}

type UpdateSensorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sensor *Sensor `protobuf:"bytes,1,opt,name=sensor,proto3" json:"sensor,omitempty"`
}

func (x *UpdateSensorRequest) Reset() {
	*x = UpdateSensorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sensor_sensor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSensorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSensorRequest) ProtoMessage() {}

func (x *UpdateSensorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sensor_sensor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSensorRequest.ProtoReflect.Descriptor instead.
func (*UpdateSensorRequest) Descriptor() ([]byte, []int) {
	return file_proto_sensor_sensor_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSensorRequest) GetSensor() *Sensor {
	if x != nil {
		return x.Sensor
	}
	return nil
}

type UpdateSensorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sensor *Sensor `protobuf:"bytes,1,opt,name=sensor,proto3" json:"sensor,omitempty"`
}

func (x *UpdateSensorResponse) Reset() {
	*x = UpdateSensorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sensor_sensor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSensorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSensorResponse) ProtoMessage() {}

func (x *UpdateSensorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sensor_sensor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSensorResponse.ProtoReflect.Descriptor instead.
func (*UpdateSensorResponse) Descriptor() ([]byte, []int) {
	return file_proto_sensor_sensor_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateSensorResponse) GetSensor() *Sensor {
	if x != nil {
		return x.Sensor
	}
	return nil
}

type DeleteSensorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSensorRequest) Reset() {
	*x = DeleteSensorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sensor_sensor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSensorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSensorRequest) ProtoMessage() {}

func (x *DeleteSensorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sensor_sensor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSensorRequest.ProtoReflect.Descriptor instead.
func (*DeleteSensorRequest) Descriptor() ([]byte, []int) {
	return file_proto_sensor_sensor_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteSensorRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteSensorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteSensorResponse) Reset() {
	*x = DeleteSensorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sensor_sensor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSensorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSensorResponse) ProtoMessage() {}

func (x *DeleteSensorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sensor_sensor_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSensorResponse.ProtoReflect.Descriptor instead.
func (*DeleteSensorResponse) Descriptor() ([]byte, []int) {
	return file_proto_sensor_sensor_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteSensorResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// The Sensor message represents a sensor in the system.
type Sensor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Temperature int32 `protobuf:"varint,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Humidity    int32 `protobuf:"varint,3,opt,name=humidity,proto3" json:"humidity,omitempty"`
	Distance    int32 `protobuf:"varint,4,opt,name=distance,proto3" json:"distance,omitempty"`
	Light       int32 `protobuf:"varint,5,opt,name=light,proto3" json:"light,omitempty"`
	Sound       int32 `protobuf:"varint,6,opt,name=sound,proto3" json:"sound,omitempty"`
}

func (x *Sensor) Reset() {
	*x = Sensor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sensor_sensor_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sensor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sensor) ProtoMessage() {}

func (x *Sensor) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sensor_sensor_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sensor.ProtoReflect.Descriptor instead.
func (*Sensor) Descriptor() ([]byte, []int) {
	return file_proto_sensor_sensor_proto_rawDescGZIP(), []int{8}
}

func (x *Sensor) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Sensor) GetTemperature() int32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *Sensor) GetHumidity() int32 {
	if x != nil {
		return x.Humidity
	}
	return 0
}

func (x *Sensor) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *Sensor) GetLight() int32 {
	if x != nil {
		return x.Light
	}
	return 0
}

func (x *Sensor) GetSound() int32 {
	if x != nil {
		return x.Sound
	}
	return 0
}

var File_proto_sensor_sensor_proto protoreflect.FileDescriptor

var file_proto_sensor_sensor_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2f, 0x73,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x22, 0x3d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x22, 0x3e, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x22, 0x3d, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x22, 0x3e, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x06,
	0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x75, 0x6d, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x68, 0x75, 0x6d, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x32, 0xb2, 0x02, 0x0a,
	0x0d, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x1b,
	0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x1b, 0x2e, 0x73, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x79, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x42,
	0x0b, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26,
	0x47, 0x6f, 0x2d, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x52, 0x65, 0x61, 0x63, 0x74, 0x2d, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x53,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0xca, 0x02, 0x06, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0xe2, 0x02,
	0x12, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x53, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_sensor_sensor_proto_rawDescOnce sync.Once
	file_proto_sensor_sensor_proto_rawDescData = file_proto_sensor_sensor_proto_rawDesc
)

func file_proto_sensor_sensor_proto_rawDescGZIP() []byte {
	file_proto_sensor_sensor_proto_rawDescOnce.Do(func() {
		file_proto_sensor_sensor_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sensor_sensor_proto_rawDescData)
	})
	return file_proto_sensor_sensor_proto_rawDescData
}

var file_proto_sensor_sensor_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_sensor_sensor_proto_goTypes = []any{
	(*CreateSensorRequest)(nil),  // 0: sensor.CreateSensorRequest
	(*CreateSensorResponse)(nil), // 1: sensor.CreateSensorResponse
	(*GetSensorRequest)(nil),     // 2: sensor.GetSensorRequest
	(*GetSensorResponse)(nil),    // 3: sensor.GetSensorResponse
	(*UpdateSensorRequest)(nil),  // 4: sensor.UpdateSensorRequest
	(*UpdateSensorResponse)(nil), // 5: sensor.UpdateSensorResponse
	(*DeleteSensorRequest)(nil),  // 6: sensor.DeleteSensorRequest
	(*DeleteSensorResponse)(nil), // 7: sensor.DeleteSensorResponse
	(*Sensor)(nil),               // 8: sensor.Sensor
}
var file_proto_sensor_sensor_proto_depIdxs = []int32{
	8, // 0: sensor.CreateSensorRequest.sensor:type_name -> sensor.Sensor
	8, // 1: sensor.CreateSensorResponse.sensor:type_name -> sensor.Sensor
	8, // 2: sensor.GetSensorResponse.sensor:type_name -> sensor.Sensor
	8, // 3: sensor.UpdateSensorRequest.sensor:type_name -> sensor.Sensor
	8, // 4: sensor.UpdateSensorResponse.sensor:type_name -> sensor.Sensor
	0, // 5: sensor.SensorService.CreateSensor:input_type -> sensor.CreateSensorRequest
	2, // 6: sensor.SensorService.GetSensor:input_type -> sensor.GetSensorRequest
	4, // 7: sensor.SensorService.UpdateSensor:input_type -> sensor.UpdateSensorRequest
	6, // 8: sensor.SensorService.DeleteSensor:input_type -> sensor.DeleteSensorRequest
	1, // 9: sensor.SensorService.CreateSensor:output_type -> sensor.CreateSensorResponse
	3, // 10: sensor.SensorService.GetSensor:output_type -> sensor.GetSensorResponse
	5, // 11: sensor.SensorService.UpdateSensor:output_type -> sensor.UpdateSensorResponse
	7, // 12: sensor.SensorService.DeleteSensor:output_type -> sensor.DeleteSensorResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_sensor_sensor_proto_init() }
func file_proto_sensor_sensor_proto_init() {
	if File_proto_sensor_sensor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_sensor_sensor_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateSensorRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_sensor_sensor_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateSensorResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_sensor_sensor_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetSensorRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_sensor_sensor_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetSensorResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_sensor_sensor_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateSensorRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_sensor_sensor_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateSensorResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_sensor_sensor_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteSensorRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_sensor_sensor_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteSensorResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_sensor_sensor_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Sensor); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_sensor_sensor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_sensor_sensor_proto_goTypes,
		DependencyIndexes: file_proto_sensor_sensor_proto_depIdxs,
		MessageInfos:      file_proto_sensor_sensor_proto_msgTypes,
	}.Build()
	File_proto_sensor_sensor_proto = out.File
	file_proto_sensor_sensor_proto_rawDesc = nil
	file_proto_sensor_sensor_proto_goTypes = nil
	file_proto_sensor_sensor_proto_depIdxs = nil
}