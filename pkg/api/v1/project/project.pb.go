// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: project.proto

package project

import (
	reflect "reflect"
	sync "sync"

	metadata "github.com/curious-kitten/scratch-post/pkg/api/v1/metadata"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity *metadata.Identity `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	// Name of the project
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// A description of the project
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetIdentity() *metadata.Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_project_proto protoreflect.FileDescriptor

var file_project_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x74, 0x63, 0x68,
	0x70, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x69, 0x6f, 0x75, 0x73, 0x6b, 0x69, 0x74, 0x74,
	0x65, 0x6e, 0x1a, 0x17, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x07,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x48, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x74, 0x63, 0x68, 0x70, 0x6f, 0x73, 0x74,
	0x2e, 0x63, 0x75, 0x72, 0x69, 0x6f, 0x75, 0x73, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x2e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x75, 0x72, 0x69, 0x6f, 0x75, 0x73, 0x2d, 0x6b, 0x69,
	0x74, 0x74, 0x65, 0x6e, 0x2f, 0x73, 0x63, 0x72, 0x61, 0x74, 0x63, 0x68, 0x2d, 0x70, 0x6f, 0x73,
	0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_proto_rawDescOnce sync.Once
	file_project_proto_rawDescData = file_project_proto_rawDesc
)

func file_project_proto_rawDescGZIP() []byte {
	file_project_proto_rawDescOnce.Do(func() {
		file_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_proto_rawDescData)
	})
	return file_project_proto_rawDescData
}

var file_project_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_proto_goTypes = []interface{}{
	(*Project)(nil),           // 0: project.scratchpost.curiouskitten.Project
	(*metadata.Identity)(nil), // 1: metadata.scratchpost.curiouskitten.Identity
}
var file_project_proto_depIdxs = []int32{
	1, // 0: project.scratchpost.curiouskitten.Project.identity:type_name -> metadata.scratchpost.curiouskitten.Identity
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_project_proto_init() }
func file_project_proto_init() {
	if File_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
			RawDescriptor: file_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_proto_goTypes,
		DependencyIndexes: file_project_proto_depIdxs,
		MessageInfos:      file_project_proto_msgTypes,
	}.Build()
	File_project_proto = out.File
	file_project_proto_rawDesc = nil
	file_project_proto_goTypes = nil
	file_project_proto_depIdxs = nil
}
