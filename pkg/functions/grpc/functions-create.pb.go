// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: pkg/functions/grpc/functions-create.proto

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

type FunctionsCreateFunctionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *FunctionsBaseInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof" json:"info,omitempty"`
}

func (x *FunctionsCreateFunctionRequest) Reset() {
	*x = FunctionsCreateFunctionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_functions_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionsCreateFunctionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionsCreateFunctionRequest) ProtoMessage() {}

func (x *FunctionsCreateFunctionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_functions_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionsCreateFunctionRequest.ProtoReflect.Descriptor instead.
func (*FunctionsCreateFunctionRequest) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_functions_create_proto_rawDescGZIP(), []int{0}
}

func (x *FunctionsCreateFunctionRequest) GetInfo() *FunctionsBaseInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_pkg_functions_grpc_functions_create_proto protoreflect.FileDescriptor

var file_pkg_functions_grpc_functions_create_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x69, 0x72,
	0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0x29, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x1e, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64, 0x69, 0x72,
	0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x64, 0x69, 0x72,
	0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_functions_grpc_functions_create_proto_rawDescOnce sync.Once
	file_pkg_functions_grpc_functions_create_proto_rawDescData = file_pkg_functions_grpc_functions_create_proto_rawDesc
)

func file_pkg_functions_grpc_functions_create_proto_rawDescGZIP() []byte {
	file_pkg_functions_grpc_functions_create_proto_rawDescOnce.Do(func() {
		file_pkg_functions_grpc_functions_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_functions_grpc_functions_create_proto_rawDescData)
	})
	return file_pkg_functions_grpc_functions_create_proto_rawDescData
}

var file_pkg_functions_grpc_functions_create_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_functions_grpc_functions_create_proto_goTypes = []interface{}{
	(*FunctionsCreateFunctionRequest)(nil), // 0: direktiv_functions.FunctionsCreateFunctionRequest
	(*FunctionsBaseInfo)(nil),              // 1: direktiv_functions.FunctionsBaseInfo
}
var file_pkg_functions_grpc_functions_create_proto_depIdxs = []int32{
	1, // 0: direktiv_functions.FunctionsCreateFunctionRequest.info:type_name -> direktiv_functions.FunctionsBaseInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_functions_grpc_functions_create_proto_init() }
func file_pkg_functions_grpc_functions_create_proto_init() {
	if File_pkg_functions_grpc_functions_create_proto != nil {
		return
	}
	file_pkg_functions_grpc_functions_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_functions_grpc_functions_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionsCreateFunctionRequest); i {
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
	file_pkg_functions_grpc_functions_create_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_functions_grpc_functions_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_functions_grpc_functions_create_proto_goTypes,
		DependencyIndexes: file_pkg_functions_grpc_functions_create_proto_depIdxs,
		MessageInfos:      file_pkg_functions_grpc_functions_create_proto_msgTypes,
	}.Build()
	File_pkg_functions_grpc_functions_create_proto = out.File
	file_pkg_functions_grpc_functions_create_proto_rawDesc = nil
	file_pkg_functions_grpc_functions_create_proto_goTypes = nil
	file_pkg_functions_grpc_functions_create_proto_depIdxs = nil
}
