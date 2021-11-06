// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: iris-classification/v1/service.proto

package irisclassificationpb

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

// petal length, petal width, sepal length, sepal width
type PredictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PetalLength float32 `protobuf:"fixed32,1,opt,name=petal_length,json=petalLength,proto3" json:"petal_length,omitempty"`
	PetalWidth  float32 `protobuf:"fixed32,2,opt,name=petal_width,json=petalWidth,proto3" json:"petal_width,omitempty"`
	SepalLength float32 `protobuf:"fixed32,3,opt,name=sepal_length,json=sepalLength,proto3" json:"sepal_length,omitempty"`
	SepalWidth  float32 `protobuf:"fixed32,4,opt,name=sepal_width,json=sepalWidth,proto3" json:"sepal_width,omitempty"`
}

func (x *PredictRequest) Reset() {
	*x = PredictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iris_classification_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictRequest) ProtoMessage() {}

func (x *PredictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iris_classification_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictRequest.ProtoReflect.Descriptor instead.
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return file_iris_classification_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *PredictRequest) GetPetalLength() float32 {
	if x != nil {
		return x.PetalLength
	}
	return 0
}

func (x *PredictRequest) GetPetalWidth() float32 {
	if x != nil {
		return x.PetalWidth
	}
	return 0
}

func (x *PredictRequest) GetSepalLength() float32 {
	if x != nil {
		return x.SepalLength
	}
	return 0
}

func (x *PredictRequest) GetSepalWidth() float32 {
	if x != nil {
		return x.SepalWidth
	}
	return 0
}

type PredictResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Predicition string `protobuf:"bytes,1,opt,name=predicition,proto3" json:"predicition,omitempty"`
}

func (x *PredictResponse) Reset() {
	*x = PredictResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iris_classification_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictResponse) ProtoMessage() {}

func (x *PredictResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iris_classification_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictResponse.ProtoReflect.Descriptor instead.
func (*PredictResponse) Descriptor() ([]byte, []int) {
	return file_iris_classification_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *PredictResponse) GetPredicition() string {
	if x != nil {
		return x.Predicition
	}
	return ""
}

var File_iris_classification_v1_service_proto protoreflect.FileDescriptor

var file_iris_classification_v1_service_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x72, 0x69, 0x73, 0x2d, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x69, 0x72, 0x69, 0x73, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x98, 0x01, 0x0a, 0x0e, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x65, 0x74, 0x61, 0x6c, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0b, 0x70, 0x65, 0x74, 0x61, 0x6c, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x65, 0x74, 0x61, 0x6c, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x65, 0x74, 0x61, 0x6c, 0x57, 0x69, 0x64, 0x74,
	0x68, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x70, 0x61, 0x6c, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x73, 0x65, 0x70, 0x61, 0x6c, 0x4c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x70, 0x61, 0x6c, 0x5f, 0x77, 0x69,
	0x64, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x73, 0x65, 0x70, 0x61, 0x6c,
	0x57, 0x69, 0x64, 0x74, 0x68, 0x22, 0x33, 0x0a, 0x0f, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x6f, 0x0a, 0x19, 0x49, 0x72,
	0x69, 0x73, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x12, 0x22, 0x2e, 0x69, 0x72, 0x69, 0x73, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x72, 0x69, 0x73, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4a, 0x5a, 0x48, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x65, 0x6c, 0x6f, 0x72,
	0x65, 0x2f, 0x69, 0x72, 0x69, 0x73, 0x2d, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x3b, 0x69, 0x72, 0x69, 0x73, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iris_classification_v1_service_proto_rawDescOnce sync.Once
	file_iris_classification_v1_service_proto_rawDescData = file_iris_classification_v1_service_proto_rawDesc
)

func file_iris_classification_v1_service_proto_rawDescGZIP() []byte {
	file_iris_classification_v1_service_proto_rawDescOnce.Do(func() {
		file_iris_classification_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_iris_classification_v1_service_proto_rawDescData)
	})
	return file_iris_classification_v1_service_proto_rawDescData
}

var file_iris_classification_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_iris_classification_v1_service_proto_goTypes = []interface{}{
	(*PredictRequest)(nil),  // 0: irisclassification.PredictRequest
	(*PredictResponse)(nil), // 1: irisclassification.PredictResponse
}
var file_iris_classification_v1_service_proto_depIdxs = []int32{
	0, // 0: irisclassification.IrisClassificationService.Predict:input_type -> irisclassification.PredictRequest
	1, // 1: irisclassification.IrisClassificationService.Predict:output_type -> irisclassification.PredictResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_iris_classification_v1_service_proto_init() }
func file_iris_classification_v1_service_proto_init() {
	if File_iris_classification_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iris_classification_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictRequest); i {
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
		file_iris_classification_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictResponse); i {
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
			RawDescriptor: file_iris_classification_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iris_classification_v1_service_proto_goTypes,
		DependencyIndexes: file_iris_classification_v1_service_proto_depIdxs,
		MessageInfos:      file_iris_classification_v1_service_proto_msgTypes,
	}.Build()
	File_iris_classification_v1_service_proto = out.File
	file_iris_classification_v1_service_proto_rawDesc = nil
	file_iris_classification_v1_service_proto_goTypes = nil
	file_iris_classification_v1_service_proto_depIdxs = nil
}
