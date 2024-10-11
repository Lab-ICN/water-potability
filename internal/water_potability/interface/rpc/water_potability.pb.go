// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: water_potability.proto

package rpc

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

type PredictWaterPotabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ph                  float64 `protobuf:"fixed64,1,opt,name=ph,proto3" json:"ph,omitempty"`
	TotalDissolveSolids float64 `protobuf:"fixed64,2,opt,name=totalDissolveSolids,proto3" json:"totalDissolveSolids,omitempty"`
	Turbidity           float64 `protobuf:"fixed64,3,opt,name=turbidity,proto3" json:"turbidity,omitempty"`
}

func (x *PredictWaterPotabilityRequest) Reset() {
	*x = PredictWaterPotabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_water_potability_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictWaterPotabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictWaterPotabilityRequest) ProtoMessage() {}

func (x *PredictWaterPotabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_water_potability_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictWaterPotabilityRequest.ProtoReflect.Descriptor instead.
func (*PredictWaterPotabilityRequest) Descriptor() ([]byte, []int) {
	return file_water_potability_proto_rawDescGZIP(), []int{0}
}

func (x *PredictWaterPotabilityRequest) GetPh() float64 {
	if x != nil {
		return x.Ph
	}
	return 0
}

func (x *PredictWaterPotabilityRequest) GetTotalDissolveSolids() float64 {
	if x != nil {
		return x.TotalDissolveSolids
	}
	return 0
}

func (x *PredictWaterPotabilityRequest) GetTurbidity() float64 {
	if x != nil {
		return x.Turbidity
	}
	return 0
}

type PredictWaterPotabilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prediction float64 `protobuf:"fixed64,1,opt,name=prediction,proto3" json:"prediction,omitempty"`
	Level      string  `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	Error      *Error  `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *PredictWaterPotabilityResponse) Reset() {
	*x = PredictWaterPotabilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_water_potability_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictWaterPotabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictWaterPotabilityResponse) ProtoMessage() {}

func (x *PredictWaterPotabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_water_potability_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictWaterPotabilityResponse.ProtoReflect.Descriptor instead.
func (*PredictWaterPotabilityResponse) Descriptor() ([]byte, []int) {
	return file_water_potability_proto_rawDescGZIP(), []int{1}
}

func (x *PredictWaterPotabilityResponse) GetPrediction() float64 {
	if x != nil {
		return x.Prediction
	}
	return 0
}

func (x *PredictWaterPotabilityResponse) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *PredictWaterPotabilityResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_water_potability_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_water_potability_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_water_potability_proto_rawDescGZIP(), []int{2}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_water_potability_proto protoreflect.FileDescriptor

var file_water_potability_proto_rawDesc = []byte{
	0x0a, 0x16, 0x77, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x77, 0x61, 0x74, 0x65, 0x72, 0x5f,
	0x70, 0x6f, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x7f, 0x0a, 0x1d, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x57, 0x61, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x74, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x70,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x70, 0x68, 0x12, 0x30, 0x0a, 0x13, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x53, 0x6f, 0x6c, 0x69,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44,
	0x69, 0x73, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x53, 0x6f, 0x6c, 0x69, 0x64, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x75, 0x72, 0x62, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x74, 0x75, 0x72, 0x62, 0x69, 0x64, 0x69, 0x74, 0x79, 0x22, 0x85, 0x01, 0x0a, 0x1e,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x57, 0x61, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x74, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x74, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x95, 0x01, 0x0a, 0x16, 0x57,
	0x61, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7b, 0x0a, 0x16, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x57, 0x61, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x2f, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x57, 0x61, 0x74, 0x65, 0x72, 0x50,
	0x6f, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x30, 0x2e, 0x77, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x74, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x57, 0x61, 0x74, 0x65, 0x72,
	0x50, 0x6f, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x77, 0x61, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_water_potability_proto_rawDescOnce sync.Once
	file_water_potability_proto_rawDescData = file_water_potability_proto_rawDesc
)

func file_water_potability_proto_rawDescGZIP() []byte {
	file_water_potability_proto_rawDescOnce.Do(func() {
		file_water_potability_proto_rawDescData = protoimpl.X.CompressGZIP(file_water_potability_proto_rawDescData)
	})
	return file_water_potability_proto_rawDescData
}

var file_water_potability_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_water_potability_proto_goTypes = []any{
	(*PredictWaterPotabilityRequest)(nil),  // 0: water_potability.PredictWaterPotabilityRequest
	(*PredictWaterPotabilityResponse)(nil), // 1: water_potability.PredictWaterPotabilityResponse
	(*Error)(nil),                          // 2: water_potability.Error
}
var file_water_potability_proto_depIdxs = []int32{
	2, // 0: water_potability.PredictWaterPotabilityResponse.error:type_name -> water_potability.Error
	0, // 1: water_potability.WaterPotabilityService.PredictWaterPotability:input_type -> water_potability.PredictWaterPotabilityRequest
	1, // 2: water_potability.WaterPotabilityService.PredictWaterPotability:output_type -> water_potability.PredictWaterPotabilityResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_water_potability_proto_init() }
func file_water_potability_proto_init() {
	if File_water_potability_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_water_potability_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PredictWaterPotabilityRequest); i {
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
		file_water_potability_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PredictWaterPotabilityResponse); i {
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
		file_water_potability_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_water_potability_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_water_potability_proto_goTypes,
		DependencyIndexes: file_water_potability_proto_depIdxs,
		MessageInfos:      file_water_potability_proto_msgTypes,
	}.Build()
	File_water_potability_proto = out.File
	file_water_potability_proto_rawDesc = nil
	file_water_potability_proto_goTypes = nil
	file_water_potability_proto_depIdxs = nil
}
