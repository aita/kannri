// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: kannri/v1/manager.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateServiceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Command       string                 `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
	Args          []string               `protobuf:"bytes,4,rep,name=args,proto3" json:"args,omitempty"`
	Environment   []string               `protobuf:"bytes,5,rep,name=environment,proto3" json:"environment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateServiceRequest) Reset() {
	*x = CreateServiceRequest{}
	mi := &file_kannri_v1_manager_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServiceRequest) ProtoMessage() {}

func (x *CreateServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kannri_v1_manager_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServiceRequest.ProtoReflect.Descriptor instead.
func (*CreateServiceRequest) Descriptor() ([]byte, []int) {
	return file_kannri_v1_manager_proto_rawDescGZIP(), []int{0}
}

func (x *CreateServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateServiceRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateServiceRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *CreateServiceRequest) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *CreateServiceRequest) GetEnvironment() []string {
	if x != nil {
		return x.Environment
	}
	return nil
}

type CreateServiceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Service       *Service               `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateServiceResponse) Reset() {
	*x = CreateServiceResponse{}
	mi := &file_kannri_v1_manager_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServiceResponse) ProtoMessage() {}

func (x *CreateServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kannri_v1_manager_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServiceResponse.ProtoReflect.Descriptor instead.
func (*CreateServiceResponse) Descriptor() ([]byte, []int) {
	return file_kannri_v1_manager_proto_rawDescGZIP(), []int{1}
}

func (x *CreateServiceResponse) GetService() *Service {
	if x != nil {
		return x.Service
	}
	return nil
}

var File_kannri_v1_manager_proto protoreflect.FileDescriptor

var file_kannri_v1_manager_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x6b, 0x61, 0x6e, 0x6e, 0x72, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6b, 0x61, 0x6e, 0x6e, 0x72,
	0x69, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x6b, 0x61, 0x6e, 0x6e, 0x72, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x45, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x72, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x32, 0x66, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x72, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x72, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x8c, 0x01, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x72, 0x69, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x74, 0x61, 0x2f, 0x6b,
	0x61, 0x6e, 0x6e, 0x72, 0x69, 0x2f, 0x6b, 0x61, 0x6e, 0x6e, 0x72, 0x69, 0x70, 0x62, 0x2f, 0x6b,
	0x61, 0x6e, 0x72, 0x69, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4b, 0x58, 0x58, 0xaa, 0x02, 0x09,
	0x4b, 0x61, 0x6e, 0x6e, 0x72, 0x69, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x4b, 0x61, 0x6e, 0x6e,
	0x72, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x4b, 0x61, 0x6e, 0x6e, 0x72, 0x69, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a,
	0x4b, 0x61, 0x6e, 0x6e, 0x72, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_kannri_v1_manager_proto_rawDescOnce sync.Once
	file_kannri_v1_manager_proto_rawDescData []byte
)

func file_kannri_v1_manager_proto_rawDescGZIP() []byte {
	file_kannri_v1_manager_proto_rawDescOnce.Do(func() {
		file_kannri_v1_manager_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_kannri_v1_manager_proto_rawDesc), len(file_kannri_v1_manager_proto_rawDesc)))
	})
	return file_kannri_v1_manager_proto_rawDescData
}

var file_kannri_v1_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kannri_v1_manager_proto_goTypes = []any{
	(*CreateServiceRequest)(nil),  // 0: kannri.v1.CreateServiceRequest
	(*CreateServiceResponse)(nil), // 1: kannri.v1.CreateServiceResponse
	(*Service)(nil),               // 2: kannri.v1.Service
}
var file_kannri_v1_manager_proto_depIdxs = []int32{
	2, // 0: kannri.v1.CreateServiceResponse.service:type_name -> kannri.v1.Service
	0, // 1: kannri.v1.ManagerService.CreateService:input_type -> kannri.v1.CreateServiceRequest
	1, // 2: kannri.v1.ManagerService.CreateService:output_type -> kannri.v1.CreateServiceResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kannri_v1_manager_proto_init() }
func file_kannri_v1_manager_proto_init() {
	if File_kannri_v1_manager_proto != nil {
		return
	}
	file_kannri_v1_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_kannri_v1_manager_proto_rawDesc), len(file_kannri_v1_manager_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kannri_v1_manager_proto_goTypes,
		DependencyIndexes: file_kannri_v1_manager_proto_depIdxs,
		MessageInfos:      file_kannri_v1_manager_proto_msgTypes,
	}.Build()
	File_kannri_v1_manager_proto = out.File
	file_kannri_v1_manager_proto_goTypes = nil
	file_kannri_v1_manager_proto_depIdxs = nil
}
