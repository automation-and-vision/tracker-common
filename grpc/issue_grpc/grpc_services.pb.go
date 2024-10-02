// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: grpc_services.proto

package issue_grpc

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

type GetProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=projectId,proto3" json:"projectId,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetProjectResponse) Reset() {
	*x = GetProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProjectResponse) ProtoMessage() {}

func (x *GetProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProjectResponse.ProtoReflect.Descriptor instead.
func (*GetProjectResponse) Descriptor() ([]byte, []int) {
	return file_grpc_services_proto_rawDescGZIP(), []int{0}
}

func (x *GetProjectResponse) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *GetProjectResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserProjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowAccess bool `protobuf:"varint,1,opt,name=allowAccess,proto3" json:"allowAccess,omitempty"`
}

func (x *UserProjectResponse) Reset() {
	*x = UserProjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProjectResponse) ProtoMessage() {}

func (x *UserProjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProjectResponse.ProtoReflect.Descriptor instead.
func (*UserProjectResponse) Descriptor() ([]byte, []int) {
	return file_grpc_services_proto_rawDescGZIP(), []int{1}
}

func (x *UserProjectResponse) GetAllowAccess() bool {
	if x != nil {
		return x.AllowAccess
	}
	return false
}

var File_grpc_services_proto protoreflect.FileDescriptor

var file_grpc_services_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x22, 0x4a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x37, 0x0a,
	0x13, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x65, 0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x2e, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x1f, 0x2e, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a,
	0x0f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_services_proto_rawDescOnce sync.Once
	file_grpc_services_proto_rawDescData = file_grpc_services_proto_rawDesc
)

func file_grpc_services_proto_rawDescGZIP() []byte {
	file_grpc_services_proto_rawDescOnce.Do(func() {
		file_grpc_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_services_proto_rawDescData)
	})
	return file_grpc_services_proto_rawDescData
}

var file_grpc_services_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_services_proto_goTypes = []any{
	(*GetProjectResponse)(nil),  // 0: issue_grpc.GetProjectResponse
	(*UserProjectResponse)(nil), // 1: issue_grpc.UserProjectResponse
}
var file_grpc_services_proto_depIdxs = []int32{
	0, // 0: issue_grpc.WorkspaceService.GetUserProject:input_type -> issue_grpc.GetProjectResponse
	1, // 1: issue_grpc.WorkspaceService.GetUserProject:output_type -> issue_grpc.UserProjectResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_services_proto_init() }
func file_grpc_services_proto_init() {
	if File_grpc_services_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_services_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetProjectResponse); i {
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
		file_grpc_services_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UserProjectResponse); i {
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
			RawDescriptor: file_grpc_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_services_proto_goTypes,
		DependencyIndexes: file_grpc_services_proto_depIdxs,
		MessageInfos:      file_grpc_services_proto_msgTypes,
	}.Build()
	File_grpc_services_proto = out.File
	file_grpc_services_proto_rawDesc = nil
	file_grpc_services_proto_goTypes = nil
	file_grpc_services_proto_depIdxs = nil
}
