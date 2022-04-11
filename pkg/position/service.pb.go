// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: pkg/position/service.proto

package position

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Latitude   float64 `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude  float64 `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Altitude   float64 `protobuf:"fixed64,4,opt,name=altitude,proto3" json:"altitude,omitempty"`
	SpeedKnots float64 `protobuf:"fixed64,5,opt,name=speed_knots,json=speedKnots,proto3" json:"speed_knots,omitempty"`
	Accuracy   float64 `protobuf:"fixed64,6,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	DeviceId   string  `protobuf:"bytes,7,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_position_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_position_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_pkg_position_service_proto_rawDescGZIP(), []int{0}
}

func (x *Position) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Position) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Position) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Position) GetAltitude() float64 {
	if x != nil {
		return x.Altitude
	}
	return 0
}

func (x *Position) GetSpeedKnots() float64 {
	if x != nil {
		return x.SpeedKnots
	}
	return 0
}

func (x *Position) GetAccuracy() float64 {
	if x != nil {
		return x.Accuracy
	}
	return 0
}

func (x *Position) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type GetPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *GetPositionRequest) Reset() {
	*x = GetPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_position_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPositionRequest) ProtoMessage() {}

func (x *GetPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_position_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPositionRequest.ProtoReflect.Descriptor instead.
func (*GetPositionRequest) Descriptor() ([]byte, []int) {
	return file_pkg_position_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetPositionRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type GetPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position *Position `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *GetPositionResponse) Reset() {
	*x = GetPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_position_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPositionResponse) ProtoMessage() {}

func (x *GetPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_position_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPositionResponse.ProtoReflect.Descriptor instead.
func (*GetPositionResponse) Descriptor() ([]byte, []int) {
	return file_pkg_position_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetPositionResponse) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

type GetRecentPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *GetRecentPositionRequest) Reset() {
	*x = GetRecentPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_position_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecentPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecentPositionRequest) ProtoMessage() {}

func (x *GetRecentPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_position_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecentPositionRequest.ProtoReflect.Descriptor instead.
func (*GetRecentPositionRequest) Descriptor() ([]byte, []int) {
	return file_pkg_position_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetRecentPositionRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type GetRecentPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Positions []*Position `protobuf:"bytes,1,rep,name=positions,proto3" json:"positions,omitempty"`
}

func (x *GetRecentPositionResponse) Reset() {
	*x = GetRecentPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_position_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecentPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecentPositionResponse) ProtoMessage() {}

func (x *GetRecentPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_position_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecentPositionResponse.ProtoReflect.Descriptor instead.
func (*GetRecentPositionResponse) Descriptor() ([]byte, []int) {
	return file_pkg_position_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetRecentPositionResponse) GetPositions() []*Position {
	if x != nil {
		return x.Positions
	}
	return nil
}

type PostPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceTime string  `protobuf:"bytes,1,opt,name=device_time,json=deviceTime,proto3" json:"device_time,omitempty"`
	Latitude   float64 `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude  float64 `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Altitude   float64 `protobuf:"fixed64,4,opt,name=altitude,proto3" json:"altitude,omitempty"`
	SpeedKnts  float64 `protobuf:"fixed64,5,opt,name=speed_knts,json=speedKnts,proto3" json:"speed_knts,omitempty"`
	Accuracy   float64 `protobuf:"fixed64,6,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
}

func (x *PostPositionRequest) Reset() {
	*x = PostPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_position_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostPositionRequest) ProtoMessage() {}

func (x *PostPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_position_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostPositionRequest.ProtoReflect.Descriptor instead.
func (*PostPositionRequest) Descriptor() ([]byte, []int) {
	return file_pkg_position_service_proto_rawDescGZIP(), []int{5}
}

func (x *PostPositionRequest) GetDeviceTime() string {
	if x != nil {
		return x.DeviceTime
	}
	return ""
}

func (x *PostPositionRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *PostPositionRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *PostPositionRequest) GetAltitude() float64 {
	if x != nil {
		return x.Altitude
	}
	return 0
}

func (x *PostPositionRequest) GetSpeedKnts() float64 {
	if x != nil {
		return x.SpeedKnts
	}
	return 0
}

func (x *PostPositionRequest) GetAccuracy() float64 {
	if x != nil {
		return x.Accuracy
	}
	return 0
}

type PostPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PostPositionResponse) Reset() {
	*x = PostPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_position_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostPositionResponse) ProtoMessage() {}

func (x *PostPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_position_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostPositionResponse.ProtoReflect.Descriptor instead.
func (*PostPositionResponse) Descriptor() ([]byte, []int) {
	return file_pkg_position_service_proto_rawDescGZIP(), []int{6}
}

type GetOwnedDevicesPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOwnedDevicesPositionRequest) Reset() {
	*x = GetOwnedDevicesPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_position_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOwnedDevicesPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOwnedDevicesPositionRequest) ProtoMessage() {}

func (x *GetOwnedDevicesPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_position_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOwnedDevicesPositionRequest.ProtoReflect.Descriptor instead.
func (*GetOwnedDevicesPositionRequest) Descriptor() ([]byte, []int) {
	return file_pkg_position_service_proto_rawDescGZIP(), []int{7}
}

type GetOwnedDevicesPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Positions []*Position `protobuf:"bytes,1,rep,name=positions,proto3" json:"positions,omitempty"`
}

func (x *GetOwnedDevicesPositionResponse) Reset() {
	*x = GetOwnedDevicesPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_position_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOwnedDevicesPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOwnedDevicesPositionResponse) ProtoMessage() {}

func (x *GetOwnedDevicesPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_position_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOwnedDevicesPositionResponse.ProtoReflect.Descriptor instead.
func (*GetOwnedDevicesPositionResponse) Descriptor() ([]byte, []int) {
	return file_pkg_position_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetOwnedDevicesPositionResponse) GetPositions() []*Position {
	if x != nil {
		return x.Positions
	}
	return nil
}

var File_pkg_position_service_proto protoreflect.FileDescriptor

var file_pkg_position_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x6b,
	0x67, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x08, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x70, 0x65, 0x65, 0x64, 0x5f, 0x6b, 0x6e, 0x6f, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0a, 0x73, 0x70, 0x65, 0x65, 0x64, 0x4b, 0x6e, 0x6f, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74,
	0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0xc7, 0x01, 0x0a, 0x13, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x6b, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x09, 0x73, 0x70, 0x65, 0x65, 0x64, 0x4b, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x50, 0x6f, 0x73,
	0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x20, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x57, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x90, 0x05, 0x0a,
	0x0f, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x86, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x92, 0x41, 0x12, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a,
	0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x17, 0x12, 0x15, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x80, 0x01, 0x0a, 0x0c, 0x50, 0x6f,
	0x73, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x70, 0x6b, 0x67,
	0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x29, 0x92, 0x41, 0x12, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b,
	0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09,
	0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x9f, 0x01, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x6b, 0x67,
	0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63,
	0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x39, 0x92, 0x41, 0x12, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x41, 0x70,
	0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e,
	0x12, 0x1c, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x12, 0xce,
	0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x77,
	0x6e, 0x65, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65,
	0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x55, 0x92, 0x41, 0x3e, 0x12, 0x2a, 0x47,
	0x65, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x20, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x6f, 0x77, 0x6e, 0x65,
	0x64, 0x20, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x41,
	0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0e, 0x12, 0x0c, 0x2f, 0x6d, 0x79, 0x2d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x42,
	0x9e, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x79, 0x2d, 0x65, 0x76, 0x61, 0x6e, 0x73, 0x2f, 0x67, 0x70, 0x73, 0x2d, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x2d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x92, 0x41, 0x62, 0x12, 0x17,
	0x0a, 0x10, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x23, 0x0a, 0x21, 0x0a,
	0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x13, 0x08, 0x02, 0x1a,
	0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_position_service_proto_rawDescOnce sync.Once
	file_pkg_position_service_proto_rawDescData = file_pkg_position_service_proto_rawDesc
)

func file_pkg_position_service_proto_rawDescGZIP() []byte {
	file_pkg_position_service_proto_rawDescOnce.Do(func() {
		file_pkg_position_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_position_service_proto_rawDescData)
	})
	return file_pkg_position_service_proto_rawDescData
}

var file_pkg_position_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_position_service_proto_goTypes = []interface{}{
	(*Position)(nil),                        // 0: pkg.position.Position
	(*GetPositionRequest)(nil),              // 1: pkg.position.GetPositionRequest
	(*GetPositionResponse)(nil),             // 2: pkg.position.GetPositionResponse
	(*GetRecentPositionRequest)(nil),        // 3: pkg.position.GetRecentPositionRequest
	(*GetRecentPositionResponse)(nil),       // 4: pkg.position.GetRecentPositionResponse
	(*PostPositionRequest)(nil),             // 5: pkg.position.PostPositionRequest
	(*PostPositionResponse)(nil),            // 6: pkg.position.PostPositionResponse
	(*GetOwnedDevicesPositionRequest)(nil),  // 7: pkg.position.GetOwnedDevicesPositionRequest
	(*GetOwnedDevicesPositionResponse)(nil), // 8: pkg.position.GetOwnedDevicesPositionResponse
}
var file_pkg_position_service_proto_depIdxs = []int32{
	0, // 0: pkg.position.GetPositionResponse.position:type_name -> pkg.position.Position
	0, // 1: pkg.position.GetRecentPositionResponse.positions:type_name -> pkg.position.Position
	0, // 2: pkg.position.GetOwnedDevicesPositionResponse.positions:type_name -> pkg.position.Position
	1, // 3: pkg.position.PositionService.GetPosition:input_type -> pkg.position.GetPositionRequest
	5, // 4: pkg.position.PositionService.PostPosition:input_type -> pkg.position.PostPositionRequest
	3, // 5: pkg.position.PositionService.GetRecentPosition:input_type -> pkg.position.GetRecentPositionRequest
	7, // 6: pkg.position.PositionService.GetOwnedDevicesPositions:input_type -> pkg.position.GetOwnedDevicesPositionRequest
	2, // 7: pkg.position.PositionService.GetPosition:output_type -> pkg.position.GetPositionResponse
	6, // 8: pkg.position.PositionService.PostPosition:output_type -> pkg.position.PostPositionResponse
	4, // 9: pkg.position.PositionService.GetRecentPosition:output_type -> pkg.position.GetRecentPositionResponse
	8, // 10: pkg.position.PositionService.GetOwnedDevicesPositions:output_type -> pkg.position.GetOwnedDevicesPositionResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_position_service_proto_init() }
func file_pkg_position_service_proto_init() {
	if File_pkg_position_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_position_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
		file_pkg_position_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPositionRequest); i {
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
		file_pkg_position_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPositionResponse); i {
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
		file_pkg_position_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecentPositionRequest); i {
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
		file_pkg_position_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecentPositionResponse); i {
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
		file_pkg_position_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostPositionRequest); i {
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
		file_pkg_position_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostPositionResponse); i {
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
		file_pkg_position_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOwnedDevicesPositionRequest); i {
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
		file_pkg_position_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOwnedDevicesPositionResponse); i {
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
			RawDescriptor: file_pkg_position_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_position_service_proto_goTypes,
		DependencyIndexes: file_pkg_position_service_proto_depIdxs,
		MessageInfos:      file_pkg_position_service_proto_msgTypes,
	}.Build()
	File_pkg_position_service_proto = out.File
	file_pkg_position_service_proto_rawDesc = nil
	file_pkg_position_service_proto_goTypes = nil
	file_pkg_position_service_proto_depIdxs = nil
}
