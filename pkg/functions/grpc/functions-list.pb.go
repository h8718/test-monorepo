// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/functions/grpc/functions-list.proto

package grpc

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

type ListFunctionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config    *FunctionsConfig `protobuf:"bytes,1,opt,name=config,proto3,oneof" json:"config,omitempty"`
	Functions []*FunctionsInfo `protobuf:"bytes,2,rep,name=functions,proto3" json:"functions,omitempty"`
}

func (x *ListFunctionsResponse) Reset() {
	*x = ListFunctionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_functions_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFunctionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFunctionsResponse) ProtoMessage() {}

func (x *ListFunctionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_functions_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFunctionsResponse.ProtoReflect.Descriptor instead.
func (*ListFunctionsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_functions_list_proto_rawDescGZIP(), []int{0}
}

func (x *ListFunctionsResponse) GetConfig() *FunctionsConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *ListFunctionsResponse) GetFunctions() []*FunctionsInfo {
	if x != nil {
		return x.Functions
	}
	return nil
}

type ListFunctionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Annotations map[string]string `protobuf:"bytes,1,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListFunctionsRequest) Reset() {
	*x = ListFunctionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_functions_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFunctionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFunctionsRequest) ProtoMessage() {}

func (x *ListFunctionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_functions_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFunctionsRequest.ProtoReflect.Descriptor instead.
func (*ListFunctionsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_functions_list_proto_rawDescGZIP(), []int{1}
}

func (x *ListFunctionsRequest) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

var File_pkg_functions_grpc_functions_list_proto protoreflect.FileDescriptor

var file_pkg_functions_grpc_functions_list_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x69, 0x72, 0x65, 0x6b,
	0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x29, 0x70,
	0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74,
	0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0xb3, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5b, 0x0a, 0x0b, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39,
	0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x64, 0x69,
	0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_functions_grpc_functions_list_proto_rawDescOnce sync.Once
	file_pkg_functions_grpc_functions_list_proto_rawDescData = file_pkg_functions_grpc_functions_list_proto_rawDesc
)

func file_pkg_functions_grpc_functions_list_proto_rawDescGZIP() []byte {
	file_pkg_functions_grpc_functions_list_proto_rawDescOnce.Do(func() {
		file_pkg_functions_grpc_functions_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_functions_grpc_functions_list_proto_rawDescData)
	})
	return file_pkg_functions_grpc_functions_list_proto_rawDescData
}

var file_pkg_functions_grpc_functions_list_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_functions_grpc_functions_list_proto_goTypes = []interface{}{
	(*ListFunctionsResponse)(nil), // 0: direktiv_functions.ListFunctionsResponse
	(*ListFunctionsRequest)(nil),  // 1: direktiv_functions.ListFunctionsRequest
	nil,                           // 2: direktiv_functions.ListFunctionsRequest.AnnotationsEntry
	(*FunctionsConfig)(nil),       // 3: direktiv_functions.FunctionsConfig
	(*FunctionsInfo)(nil),         // 4: direktiv_functions.FunctionsInfo
}
var file_pkg_functions_grpc_functions_list_proto_depIdxs = []int32{
	3, // 0: direktiv_functions.ListFunctionsResponse.config:type_name -> direktiv_functions.FunctionsConfig
	4, // 1: direktiv_functions.ListFunctionsResponse.functions:type_name -> direktiv_functions.FunctionsInfo
	2, // 2: direktiv_functions.ListFunctionsRequest.annotations:type_name -> direktiv_functions.ListFunctionsRequest.AnnotationsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_functions_grpc_functions_list_proto_init() }
func file_pkg_functions_grpc_functions_list_proto_init() {
	if File_pkg_functions_grpc_functions_list_proto != nil {
		return
	}
	file_pkg_functions_grpc_functions_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_functions_grpc_functions_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFunctionsResponse); i {
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
		file_pkg_functions_grpc_functions_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFunctionsRequest); i {
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
	file_pkg_functions_grpc_functions_list_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_functions_grpc_functions_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_functions_grpc_functions_list_proto_goTypes,
		DependencyIndexes: file_pkg_functions_grpc_functions_list_proto_depIdxs,
		MessageInfos:      file_pkg_functions_grpc_functions_list_proto_msgTypes,
	}.Build()
	File_pkg_functions_grpc_functions_list_proto = out.File
	file_pkg_functions_grpc_functions_list_proto_rawDesc = nil
	file_pkg_functions_grpc_functions_list_proto_goTypes = nil
	file_pkg_functions_grpc_functions_list_proto_depIdxs = nil
}
