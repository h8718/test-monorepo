// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: pkg/functions/grpc/registries-get.proto

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

type FunctionsRegistry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Id   *string `protobuf:"bytes,2,opt,name=id,proto3,oneof" json:"id,omitempty"`
	User *string `protobuf:"bytes,3,opt,name=user,proto3,oneof" json:"user,omitempty"`
}

func (x *FunctionsRegistry) Reset() {
	*x = FunctionsRegistry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_registries_get_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionsRegistry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionsRegistry) ProtoMessage() {}

func (x *FunctionsRegistry) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_registries_get_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionsRegistry.ProtoReflect.Descriptor instead.
func (*FunctionsRegistry) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_registries_get_proto_rawDescGZIP(), []int{0}
}

func (x *FunctionsRegistry) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *FunctionsRegistry) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *FunctionsRegistry) GetUser() string {
	if x != nil && x.User != nil {
		return *x.User
	}
	return ""
}

type FunctionsGetRegistriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace *string `protobuf:"bytes,1,opt,name=namespace,proto3,oneof" json:"namespace,omitempty"`
}

func (x *FunctionsGetRegistriesRequest) Reset() {
	*x = FunctionsGetRegistriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_registries_get_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionsGetRegistriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionsGetRegistriesRequest) ProtoMessage() {}

func (x *FunctionsGetRegistriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_registries_get_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionsGetRegistriesRequest.ProtoReflect.Descriptor instead.
func (*FunctionsGetRegistriesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_registries_get_proto_rawDescGZIP(), []int{1}
}

func (x *FunctionsGetRegistriesRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

type FunctionsGetRegistriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Registries []*FunctionsRegistry `protobuf:"bytes,1,rep,name=registries,proto3" json:"registries,omitempty"`
}

func (x *FunctionsGetRegistriesResponse) Reset() {
	*x = FunctionsGetRegistriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_functions_grpc_registries_get_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionsGetRegistriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionsGetRegistriesResponse) ProtoMessage() {}

func (x *FunctionsGetRegistriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_functions_grpc_registries_get_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionsGetRegistriesResponse.ProtoReflect.Descriptor instead.
func (*FunctionsGetRegistriesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_functions_grpc_registries_get_proto_rawDescGZIP(), []int{2}
}

func (x *FunctionsGetRegistriesResponse) GetRegistries() []*FunctionsRegistry {
	if x != nil {
		return x.Registries
	}
	return nil
}

var File_pkg_functions_grpc_registries_get_proto protoreflect.FileDescriptor

var file_pkg_functions_grpc_registries_get_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2d,
	0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x69, 0x72, 0x65, 0x6b,
	0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x73, 0x0a,
	0x11, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x13, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x50, 0x0a, 0x1d, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x22, 0x67, 0x0a, 0x1e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64, 0x69, 0x72,
	0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x42, 0x31, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x72, 0x65,
	0x6b, 0x74, 0x69, 0x76, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_functions_grpc_registries_get_proto_rawDescOnce sync.Once
	file_pkg_functions_grpc_registries_get_proto_rawDescData = file_pkg_functions_grpc_registries_get_proto_rawDesc
)

func file_pkg_functions_grpc_registries_get_proto_rawDescGZIP() []byte {
	file_pkg_functions_grpc_registries_get_proto_rawDescOnce.Do(func() {
		file_pkg_functions_grpc_registries_get_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_functions_grpc_registries_get_proto_rawDescData)
	})
	return file_pkg_functions_grpc_registries_get_proto_rawDescData
}

var file_pkg_functions_grpc_registries_get_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_functions_grpc_registries_get_proto_goTypes = []interface{}{
	(*FunctionsRegistry)(nil),              // 0: direktiv_functions.FunctionsRegistry
	(*FunctionsGetRegistriesRequest)(nil),  // 1: direktiv_functions.FunctionsGetRegistriesRequest
	(*FunctionsGetRegistriesResponse)(nil), // 2: direktiv_functions.FunctionsGetRegistriesResponse
}
var file_pkg_functions_grpc_registries_get_proto_depIdxs = []int32{
	0, // 0: direktiv_functions.FunctionsGetRegistriesResponse.registries:type_name -> direktiv_functions.FunctionsRegistry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_functions_grpc_registries_get_proto_init() }
func file_pkg_functions_grpc_registries_get_proto_init() {
	if File_pkg_functions_grpc_registries_get_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_functions_grpc_registries_get_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionsRegistry); i {
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
		file_pkg_functions_grpc_registries_get_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionsGetRegistriesRequest); i {
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
		file_pkg_functions_grpc_registries_get_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionsGetRegistriesResponse); i {
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
	file_pkg_functions_grpc_registries_get_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pkg_functions_grpc_registries_get_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_functions_grpc_registries_get_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_functions_grpc_registries_get_proto_goTypes,
		DependencyIndexes: file_pkg_functions_grpc_registries_get_proto_depIdxs,
		MessageInfos:      file_pkg_functions_grpc_registries_get_proto_msgTypes,
	}.Build()
	File_pkg_functions_grpc_registries_get_proto = out.File
	file_pkg_functions_grpc_registries_get_proto_rawDesc = nil
	file_pkg_functions_grpc_registries_get_proto_goTypes = nil
	file_pkg_functions_grpc_registries_get_proto_depIdxs = nil
}
