// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: pkg/functions/grpc/protocol.proto

package grpc

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_pkg_functions_grpc_protocol_proto protoreflect.FileDescriptor

var file_pkg_functions_grpc_protocol_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x2d, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70,
	0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2d, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2a, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2d,
	0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x6b,
	0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x67, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x72, 0x65,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x27, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x28, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6f,
	0x64, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x6b,
	0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2d, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x2d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2d, 0x70, 0x6f, 0x64, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x2d,
	0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x6b, 0x67, 0x2f,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe3, 0x11, 0x0a, 0x10, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x53, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x12, 0x28, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x28, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76,
	0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x13,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x12, 0x2e, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x47, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5a, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x29, 0x2e, 0x64,
	0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x12, 0x2f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x1a,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x2e, 0x2e, 0x64, 0x69, 0x72,
	0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x47, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x29, 0x2e, 0x64, 0x69, 0x72,
	0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x2f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69,
	0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x55, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x12, 0x29, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2e, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x64, 0x69,
	0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x55, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x2e, 0x64, 0x69, 0x72,
	0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x66,
	0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x28, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x64, 0x69, 0x72, 0x65,
	0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76,
	0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x64, 0x69, 0x72,
	0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x13,
	0x53, 0x65, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x12, 0x25, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69,
	0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x0e, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x2e,
	0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b,
	0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x5c, 0x0a, 0x09, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x50, 0x6f, 0x64, 0x73, 0x12, 0x24, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76,
	0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x50, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x64, 0x69,
	0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x6b, 0x0a, 0x0e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b,
	0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x5c, 0x0a, 0x09, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x67, 0x73,
	0x12, 0x24, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69,
	0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x57, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x64, 0x73, 0x12, 0x23, 0x2e,
	0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x05, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x64, 0x69,
	0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_pkg_functions_grpc_protocol_proto_goTypes = []interface{}{
	(*StoreRegistryRequest)(nil),        // 0: direktiv_functions.StoreRegistryRequest
	(*GetRegistriesRequest)(nil),        // 1: direktiv_functions.GetRegistriesRequest
	(*StoreGlobalRegistryRequest)(nil),  // 2: direktiv_functions.StoreGlobalRegistryRequest
	(*empty.Empty)(nil),                 // 3: google.protobuf.Empty
	(*DeleteGlobalRegistryRequest)(nil), // 4: direktiv_functions.DeleteGlobalRegistryRequest
	(*DeleteRegistryRequest)(nil),       // 5: direktiv_functions.DeleteRegistryRequest
	(*ReconstructFunctionRequest)(nil),  // 6: direktiv_functions.ReconstructFunctionRequest
	(*UpdateFunctionRequest)(nil),       // 7: direktiv_functions.UpdateFunctionRequest
	(*CreateFunctionRequest)(nil),       // 8: direktiv_functions.CreateFunctionRequest
	(*ListFunctionsRequest)(nil),        // 9: direktiv_functions.ListFunctionsRequest
	(*GetFunctionRequest)(nil),          // 10: direktiv_functions.GetFunctionRequest
	(*SetTrafficRequest)(nil),           // 11: direktiv_functions.SetTrafficRequest
	(*DeleteRevisionRequest)(nil),       // 12: direktiv_functions.DeleteRevisionRequest
	(*WatchFunctionsRequest)(nil),       // 13: direktiv_functions.WatchFunctionsRequest
	(*WatchPodsRequest)(nil),            // 14: direktiv_functions.WatchPodsRequest
	(*WatchRevisionsRequest)(nil),       // 15: direktiv_functions.WatchRevisionsRequest
	(*WatchLogsRequest)(nil),            // 16: direktiv_functions.WatchLogsRequest
	(*ListPodsRequest)(nil),             // 17: direktiv_functions.ListPodsRequest
	(*GetRegistriesResponse)(nil),       // 18: direktiv_functions.GetRegistriesResponse
	(*ListFunctionsResponse)(nil),       // 19: direktiv_functions.ListFunctionsResponse
	(*GetFunctionResponse)(nil),         // 20: direktiv_functions.GetFunctionResponse
	(*WatchFunctionsResponse)(nil),      // 21: direktiv_functions.WatchFunctionsResponse
	(*WatchPodsResponse)(nil),           // 22: direktiv_functions.WatchPodsResponse
	(*WatchRevisionsResponse)(nil),      // 23: direktiv_functions.WatchRevisionsResponse
	(*WatchLogsResponse)(nil),           // 24: direktiv_functions.WatchLogsResponse
	(*ListPodsResponse)(nil),            // 25: direktiv_functions.ListPodsResponse
	(*BuildResponse)(nil),               // 26: direktiv_functions.BuildResponse
}
var file_pkg_functions_grpc_protocol_proto_depIdxs = []int32{
	0,  // 0: direktiv_functions.FunctionsService.StoreRegistry:input_type -> direktiv_functions.StoreRegistryRequest
	1,  // 1: direktiv_functions.FunctionsService.GetRegistries:input_type -> direktiv_functions.GetRegistriesRequest
	2,  // 2: direktiv_functions.FunctionsService.StoreGlobalRegistry:input_type -> direktiv_functions.StoreGlobalRegistryRequest
	3,  // 3: direktiv_functions.FunctionsService.GetGlobalRegistries:input_type -> google.protobuf.Empty
	4,  // 4: direktiv_functions.FunctionsService.DeleteGlobalRegistry:input_type -> direktiv_functions.DeleteGlobalRegistryRequest
	2,  // 5: direktiv_functions.FunctionsService.StoreGlobalPrivateRegistry:input_type -> direktiv_functions.StoreGlobalRegistryRequest
	3,  // 6: direktiv_functions.FunctionsService.GetGlobalPrivateRegistries:input_type -> google.protobuf.Empty
	4,  // 7: direktiv_functions.FunctionsService.DeleteGlobalPrivateRegistry:input_type -> direktiv_functions.DeleteGlobalRegistryRequest
	5,  // 8: direktiv_functions.FunctionsService.DeleteRegistry:input_type -> direktiv_functions.DeleteRegistryRequest
	6,  // 9: direktiv_functions.FunctionsService.ReconstructFunction:input_type -> direktiv_functions.ReconstructFunctionRequest
	7,  // 10: direktiv_functions.FunctionsService.UpdateFunction:input_type -> direktiv_functions.UpdateFunctionRequest
	8,  // 11: direktiv_functions.FunctionsService.CreateFunction:input_type -> direktiv_functions.CreateFunctionRequest
	9,  // 12: direktiv_functions.FunctionsService.DeleteFunctions:input_type -> direktiv_functions.ListFunctionsRequest
	9,  // 13: direktiv_functions.FunctionsService.ListFunctions:input_type -> direktiv_functions.ListFunctionsRequest
	10, // 14: direktiv_functions.FunctionsService.GetFunction:input_type -> direktiv_functions.GetFunctionRequest
	10, // 15: direktiv_functions.FunctionsService.DeleteFunction:input_type -> direktiv_functions.GetFunctionRequest
	11, // 16: direktiv_functions.FunctionsService.SetFunctionsTraffic:input_type -> direktiv_functions.SetTrafficRequest
	12, // 17: direktiv_functions.FunctionsService.DeleteRevision:input_type -> direktiv_functions.DeleteRevisionRequest
	13, // 18: direktiv_functions.FunctionsService.WatchFunctions:input_type -> direktiv_functions.WatchFunctionsRequest
	14, // 19: direktiv_functions.FunctionsService.WatchPods:input_type -> direktiv_functions.WatchPodsRequest
	15, // 20: direktiv_functions.FunctionsService.WatchRevisions:input_type -> direktiv_functions.WatchRevisionsRequest
	16, // 21: direktiv_functions.FunctionsService.WatchLogs:input_type -> direktiv_functions.WatchLogsRequest
	17, // 22: direktiv_functions.FunctionsService.ListPods:input_type -> direktiv_functions.ListPodsRequest
	3,  // 23: direktiv_functions.FunctionsService.Build:input_type -> google.protobuf.Empty
	3,  // 24: direktiv_functions.FunctionsService.StoreRegistry:output_type -> google.protobuf.Empty
	18, // 25: direktiv_functions.FunctionsService.GetRegistries:output_type -> direktiv_functions.GetRegistriesResponse
	3,  // 26: direktiv_functions.FunctionsService.StoreGlobalRegistry:output_type -> google.protobuf.Empty
	18, // 27: direktiv_functions.FunctionsService.GetGlobalRegistries:output_type -> direktiv_functions.GetRegistriesResponse
	3,  // 28: direktiv_functions.FunctionsService.DeleteGlobalRegistry:output_type -> google.protobuf.Empty
	3,  // 29: direktiv_functions.FunctionsService.StoreGlobalPrivateRegistry:output_type -> google.protobuf.Empty
	18, // 30: direktiv_functions.FunctionsService.GetGlobalPrivateRegistries:output_type -> direktiv_functions.GetRegistriesResponse
	3,  // 31: direktiv_functions.FunctionsService.DeleteGlobalPrivateRegistry:output_type -> google.protobuf.Empty
	3,  // 32: direktiv_functions.FunctionsService.DeleteRegistry:output_type -> google.protobuf.Empty
	3,  // 33: direktiv_functions.FunctionsService.ReconstructFunction:output_type -> google.protobuf.Empty
	3,  // 34: direktiv_functions.FunctionsService.UpdateFunction:output_type -> google.protobuf.Empty
	3,  // 35: direktiv_functions.FunctionsService.CreateFunction:output_type -> google.protobuf.Empty
	3,  // 36: direktiv_functions.FunctionsService.DeleteFunctions:output_type -> google.protobuf.Empty
	19, // 37: direktiv_functions.FunctionsService.ListFunctions:output_type -> direktiv_functions.ListFunctionsResponse
	20, // 38: direktiv_functions.FunctionsService.GetFunction:output_type -> direktiv_functions.GetFunctionResponse
	3,  // 39: direktiv_functions.FunctionsService.DeleteFunction:output_type -> google.protobuf.Empty
	3,  // 40: direktiv_functions.FunctionsService.SetFunctionsTraffic:output_type -> google.protobuf.Empty
	3,  // 41: direktiv_functions.FunctionsService.DeleteRevision:output_type -> google.protobuf.Empty
	21, // 42: direktiv_functions.FunctionsService.WatchFunctions:output_type -> direktiv_functions.WatchFunctionsResponse
	22, // 43: direktiv_functions.FunctionsService.WatchPods:output_type -> direktiv_functions.WatchPodsResponse
	23, // 44: direktiv_functions.FunctionsService.WatchRevisions:output_type -> direktiv_functions.WatchRevisionsResponse
	24, // 45: direktiv_functions.FunctionsService.WatchLogs:output_type -> direktiv_functions.WatchLogsResponse
	25, // 46: direktiv_functions.FunctionsService.ListPods:output_type -> direktiv_functions.ListPodsResponse
	26, // 47: direktiv_functions.FunctionsService.Build:output_type -> direktiv_functions.BuildResponse
	24, // [24:48] is the sub-list for method output_type
	0,  // [0:24] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_functions_grpc_protocol_proto_init() }
func file_pkg_functions_grpc_protocol_proto_init() {
	if File_pkg_functions_grpc_protocol_proto != nil {
		return
	}
	file_pkg_functions_grpc_registries_get_proto_init()
	file_pkg_functions_grpc_registry_delete_proto_init()
	file_pkg_functions_grpc_registry_store_proto_init()
	file_pkg_functions_grpc_registries_global_proto_init()
	file_pkg_functions_grpc_functions_get_proto_init()
	file_pkg_functions_grpc_functions_create_proto_init()
	file_pkg_functions_grpc_functions_reconstruct_proto_init()
	file_pkg_functions_grpc_functions_list_proto_init()
	file_pkg_functions_grpc_functions_traffic_proto_init()
	file_pkg_functions_grpc_functions_update_proto_init()
	file_pkg_functions_grpc_revision_delete_proto_init()
	file_pkg_functions_grpc_pod_list_proto_init()
	file_pkg_functions_grpc_watch_revisions_proto_init()
	file_pkg_functions_grpc_watch_functions_proto_init()
	file_pkg_functions_grpc_watch_pods_proto_init()
	file_pkg_functions_grpc_watch_logs_proto_init()
	file_pkg_functions_grpc_build_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_functions_grpc_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_functions_grpc_protocol_proto_goTypes,
		DependencyIndexes: file_pkg_functions_grpc_protocol_proto_depIdxs,
	}.Build()
	File_pkg_functions_grpc_protocol_proto = out.File
	file_pkg_functions_grpc_protocol_proto_rawDesc = nil
	file_pkg_functions_grpc_protocol_proto_goTypes = nil
	file_pkg_functions_grpc_protocol_proto_depIdxs = nil
}
