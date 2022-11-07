// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/flow/grpc/logs.proto

package grpc

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	T   *timestamp.Timestamp `protobuf:"bytes,1,opt,name=t,proto3" json:"t,omitempty"`
	Msg string               `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_logs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_logs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_logs_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetT() *timestamp.Timestamp {
	if x != nil {
		return x.T
	}
	return nil
}

func (x *Log) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ServerLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ServerLogsRequest) Reset() {
	*x = ServerLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_logs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerLogsRequest) ProtoMessage() {}

func (x *ServerLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_logs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerLogsRequest.ProtoReflect.Descriptor instead.
func (*ServerLogsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_logs_proto_rawDescGZIP(), []int{1}
}

func (x *ServerLogsRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ServerLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageInfo *PageInfo `protobuf:"bytes,2,opt,name=pageInfo,proto3" json:"pageInfo,omitempty"`
	Results  []*Log    `protobuf:"bytes,4,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ServerLogsResponse) Reset() {
	*x = ServerLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_logs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerLogsResponse) ProtoMessage() {}

func (x *ServerLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_logs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerLogsResponse.ProtoReflect.Descriptor instead.
func (*ServerLogsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_logs_proto_rawDescGZIP(), []int{2}
}

func (x *ServerLogsResponse) GetPageInfo() *PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

func (x *ServerLogsResponse) GetResults() []*Log {
	if x != nil {
		return x.Results
	}
	return nil
}

type NamespaceLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Namespace  string      `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *NamespaceLogsRequest) Reset() {
	*x = NamespaceLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_logs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceLogsRequest) ProtoMessage() {}

func (x *NamespaceLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_logs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceLogsRequest.ProtoReflect.Descriptor instead.
func (*NamespaceLogsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_logs_proto_rawDescGZIP(), []int{3}
}

func (x *NamespaceLogsRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *NamespaceLogsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type NamespaceLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageInfo  *PageInfo `protobuf:"bytes,2,opt,name=pageInfo,proto3" json:"pageInfo,omitempty"`
	Namespace string    `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Results   []*Log    `protobuf:"bytes,5,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *NamespaceLogsResponse) Reset() {
	*x = NamespaceLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_logs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceLogsResponse) ProtoMessage() {}

func (x *NamespaceLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_logs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceLogsResponse.ProtoReflect.Descriptor instead.
func (*NamespaceLogsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_logs_proto_rawDescGZIP(), []int{4}
}

func (x *NamespaceLogsResponse) GetPageInfo() *PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

func (x *NamespaceLogsResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *NamespaceLogsResponse) GetResults() []*Log {
	if x != nil {
		return x.Results
	}
	return nil
}

type WorkflowLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Namespace  string      `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Path       string      `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *WorkflowLogsRequest) Reset() {
	*x = WorkflowLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_logs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowLogsRequest) ProtoMessage() {}

func (x *WorkflowLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_logs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowLogsRequest.ProtoReflect.Descriptor instead.
func (*WorkflowLogsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_logs_proto_rawDescGZIP(), []int{5}
}

func (x *WorkflowLogsRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *WorkflowLogsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *WorkflowLogsRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type WorkflowLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageInfo  *PageInfo `protobuf:"bytes,2,opt,name=pageInfo,proto3" json:"pageInfo,omitempty"`
	Namespace string    `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Path      string    `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	Results   []*Log    `protobuf:"bytes,6,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *WorkflowLogsResponse) Reset() {
	*x = WorkflowLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_logs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowLogsResponse) ProtoMessage() {}

func (x *WorkflowLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_logs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowLogsResponse.ProtoReflect.Descriptor instead.
func (*WorkflowLogsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_logs_proto_rawDescGZIP(), []int{6}
}

func (x *WorkflowLogsResponse) GetPageInfo() *PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

func (x *WorkflowLogsResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *WorkflowLogsResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *WorkflowLogsResponse) GetResults() []*Log {
	if x != nil {
		return x.Results
	}
	return nil
}

type InstanceLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Namespace  string      `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Instance   string      `protobuf:"bytes,3,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *InstanceLogsRequest) Reset() {
	*x = InstanceLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_logs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceLogsRequest) ProtoMessage() {}

func (x *InstanceLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_logs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceLogsRequest.ProtoReflect.Descriptor instead.
func (*InstanceLogsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_logs_proto_rawDescGZIP(), []int{7}
}

func (x *InstanceLogsRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *InstanceLogsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *InstanceLogsRequest) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

type InstanceLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageInfo  *PageInfo `protobuf:"bytes,2,opt,name=pageInfo,proto3" json:"pageInfo,omitempty"`
	Namespace string    `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Instance  string    `protobuf:"bytes,5,opt,name=instance,proto3" json:"instance,omitempty"`
	Results   []*Log    `protobuf:"bytes,6,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *InstanceLogsResponse) Reset() {
	*x = InstanceLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_logs_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceLogsResponse) ProtoMessage() {}

func (x *InstanceLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_logs_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceLogsResponse.ProtoReflect.Descriptor instead.
func (*InstanceLogsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_logs_proto_rawDescGZIP(), []int{8}
}

func (x *InstanceLogsResponse) GetPageInfo() *PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

func (x *InstanceLogsResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *InstanceLogsResponse) GetInstance() string {
	if x != nil {
		return x.Instance
	}
	return ""
}

func (x *InstanceLogsResponse) GetResults() []*Log {
	if x != nil {
		return x.Results
	}
	return nil
}

type MirrorActivityLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Namespace  string      `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Activity   string      `protobuf:"bytes,3,opt,name=activity,proto3" json:"activity,omitempty"`
}

func (x *MirrorActivityLogsRequest) Reset() {
	*x = MirrorActivityLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_logs_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MirrorActivityLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MirrorActivityLogsRequest) ProtoMessage() {}

func (x *MirrorActivityLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_logs_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MirrorActivityLogsRequest.ProtoReflect.Descriptor instead.
func (*MirrorActivityLogsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_logs_proto_rawDescGZIP(), []int{9}
}

func (x *MirrorActivityLogsRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *MirrorActivityLogsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *MirrorActivityLogsRequest) GetActivity() string {
	if x != nil {
		return x.Activity
	}
	return ""
}

type MirrorActivityLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageInfo  *PageInfo `protobuf:"bytes,2,opt,name=pageInfo,proto3" json:"pageInfo,omitempty"`
	Namespace string    `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Activity  string    `protobuf:"bytes,5,opt,name=activity,proto3" json:"activity,omitempty"`
	Results   []*Log    `protobuf:"bytes,6,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MirrorActivityLogsResponse) Reset() {
	*x = MirrorActivityLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_flow_grpc_logs_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MirrorActivityLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MirrorActivityLogsResponse) ProtoMessage() {}

func (x *MirrorActivityLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_flow_grpc_logs_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MirrorActivityLogsResponse.ProtoReflect.Descriptor instead.
func (*MirrorActivityLogsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_flow_grpc_logs_proto_rawDescGZIP(), []int{10}
}

func (x *MirrorActivityLogsResponse) GetPageInfo() *PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

func (x *MirrorActivityLogsResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *MirrorActivityLogsResponse) GetActivity() string {
	if x != nil {
		return x.Activity
	}
	return ""
}

func (x *MirrorActivityLogsResponse) GetResults() []*Log {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_pkg_flow_grpc_logs_proto protoreflect.FileDescriptor

var file_pkg_flow_grpc_logs_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x64, 0x69, 0x72, 0x65,
	0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x6b, 0x67, 0x2f,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x03, 0x4c, 0x6f,
	0x67, 0x12, 0x28, 0x0a, 0x01, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x01, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x4e, 0x0a,
	0x11, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69,
	0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x77, 0x0a,
	0x12, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76,
	0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x69, 0x72, 0x65,
	0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x6f, 0x0a, 0x14, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x98, 0x01, 0x0a, 0x15, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76,
	0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x13, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0xab, 0x01, 0x0a, 0x14, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x33, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x2c, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b,
	0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x13, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x22, 0xb3, 0x01, 0x0a, 0x14, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x69,
	0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4c, 0x6f, 0x67, 0x52,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x19, 0x4d, 0x69, 0x72,
	0x72, 0x6f, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x69, 0x72,
	0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x22, 0xb9, 0x01, 0x0a, 0x1a,
	0x4d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64,
	0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x69, 0x72,
	0x65, 0x6b, 0x74, 0x69, 0x76, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x64,
	0x69, 0x72, 0x65, 0x6b, 0x74, 0x69, 0x76, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_flow_grpc_logs_proto_rawDescOnce sync.Once
	file_pkg_flow_grpc_logs_proto_rawDescData = file_pkg_flow_grpc_logs_proto_rawDesc
)

func file_pkg_flow_grpc_logs_proto_rawDescGZIP() []byte {
	file_pkg_flow_grpc_logs_proto_rawDescOnce.Do(func() {
		file_pkg_flow_grpc_logs_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_flow_grpc_logs_proto_rawDescData)
	})
	return file_pkg_flow_grpc_logs_proto_rawDescData
}

var file_pkg_flow_grpc_logs_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pkg_flow_grpc_logs_proto_goTypes = []interface{}{
	(*Log)(nil),                        // 0: direktiv_flow.Log
	(*ServerLogsRequest)(nil),          // 1: direktiv_flow.ServerLogsRequest
	(*ServerLogsResponse)(nil),         // 2: direktiv_flow.ServerLogsResponse
	(*NamespaceLogsRequest)(nil),       // 3: direktiv_flow.NamespaceLogsRequest
	(*NamespaceLogsResponse)(nil),      // 4: direktiv_flow.NamespaceLogsResponse
	(*WorkflowLogsRequest)(nil),        // 5: direktiv_flow.WorkflowLogsRequest
	(*WorkflowLogsResponse)(nil),       // 6: direktiv_flow.WorkflowLogsResponse
	(*InstanceLogsRequest)(nil),        // 7: direktiv_flow.InstanceLogsRequest
	(*InstanceLogsResponse)(nil),       // 8: direktiv_flow.InstanceLogsResponse
	(*MirrorActivityLogsRequest)(nil),  // 9: direktiv_flow.MirrorActivityLogsRequest
	(*MirrorActivityLogsResponse)(nil), // 10: direktiv_flow.MirrorActivityLogsResponse
	(*timestamp.Timestamp)(nil),        // 11: google.protobuf.Timestamp
	(*Pagination)(nil),                 // 12: direktiv_flow.Pagination
	(*PageInfo)(nil),                   // 13: direktiv_flow.PageInfo
}
var file_pkg_flow_grpc_logs_proto_depIdxs = []int32{
	11, // 0: direktiv_flow.Log.t:type_name -> google.protobuf.Timestamp
	12, // 1: direktiv_flow.ServerLogsRequest.pagination:type_name -> direktiv_flow.Pagination
	13, // 2: direktiv_flow.ServerLogsResponse.pageInfo:type_name -> direktiv_flow.PageInfo
	0,  // 3: direktiv_flow.ServerLogsResponse.results:type_name -> direktiv_flow.Log
	12, // 4: direktiv_flow.NamespaceLogsRequest.pagination:type_name -> direktiv_flow.Pagination
	13, // 5: direktiv_flow.NamespaceLogsResponse.pageInfo:type_name -> direktiv_flow.PageInfo
	0,  // 6: direktiv_flow.NamespaceLogsResponse.results:type_name -> direktiv_flow.Log
	12, // 7: direktiv_flow.WorkflowLogsRequest.pagination:type_name -> direktiv_flow.Pagination
	13, // 8: direktiv_flow.WorkflowLogsResponse.pageInfo:type_name -> direktiv_flow.PageInfo
	0,  // 9: direktiv_flow.WorkflowLogsResponse.results:type_name -> direktiv_flow.Log
	12, // 10: direktiv_flow.InstanceLogsRequest.pagination:type_name -> direktiv_flow.Pagination
	13, // 11: direktiv_flow.InstanceLogsResponse.pageInfo:type_name -> direktiv_flow.PageInfo
	0,  // 12: direktiv_flow.InstanceLogsResponse.results:type_name -> direktiv_flow.Log
	12, // 13: direktiv_flow.MirrorActivityLogsRequest.pagination:type_name -> direktiv_flow.Pagination
	13, // 14: direktiv_flow.MirrorActivityLogsResponse.pageInfo:type_name -> direktiv_flow.PageInfo
	0,  // 15: direktiv_flow.MirrorActivityLogsResponse.results:type_name -> direktiv_flow.Log
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_pkg_flow_grpc_logs_proto_init() }
func file_pkg_flow_grpc_logs_proto_init() {
	if File_pkg_flow_grpc_logs_proto != nil {
		return
	}
	file_pkg_flow_grpc_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_flow_grpc_logs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_pkg_flow_grpc_logs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerLogsRequest); i {
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
		file_pkg_flow_grpc_logs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerLogsResponse); i {
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
		file_pkg_flow_grpc_logs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceLogsRequest); i {
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
		file_pkg_flow_grpc_logs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceLogsResponse); i {
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
		file_pkg_flow_grpc_logs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowLogsRequest); i {
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
		file_pkg_flow_grpc_logs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowLogsResponse); i {
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
		file_pkg_flow_grpc_logs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceLogsRequest); i {
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
		file_pkg_flow_grpc_logs_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceLogsResponse); i {
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
		file_pkg_flow_grpc_logs_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MirrorActivityLogsRequest); i {
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
		file_pkg_flow_grpc_logs_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MirrorActivityLogsResponse); i {
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
			RawDescriptor: file_pkg_flow_grpc_logs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_flow_grpc_logs_proto_goTypes,
		DependencyIndexes: file_pkg_flow_grpc_logs_proto_depIdxs,
		MessageInfos:      file_pkg_flow_grpc_logs_proto_msgTypes,
	}.Build()
	File_pkg_flow_grpc_logs_proto = out.File
	file_pkg_flow_grpc_logs_proto_rawDesc = nil
	file_pkg_flow_grpc_logs_proto_goTypes = nil
	file_pkg_flow_grpc_logs_proto_depIdxs = nil
}
