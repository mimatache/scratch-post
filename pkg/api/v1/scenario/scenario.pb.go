// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: scenario.proto

package scenario

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

// Represents a step that has to be completed in order to complete the test
type Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Used to order step execution
	Position int32 `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
	// Name of the step
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Describe what the step intention is
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Describe what needs to be done in order to perform the step
	Action string `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	// Describe what you expect the resoult of the action to be
	ExpectedOutcome string `protobuf:"bytes,5,opt,name=expectedOutcome,proto3" json:"expectedOutcome,omitempty"`
}

func (x *Step) Reset() {
	*x = Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scenario_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Step) ProtoMessage() {}

func (x *Step) ProtoReflect() protoreflect.Message {
	mi := &file_scenario_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Step.ProtoReflect.Descriptor instead.
func (*Step) Descriptor() ([]byte, []int) {
	return file_scenario_proto_rawDescGZIP(), []int{0}
}

func (x *Step) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *Step) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Step) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Step) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Step) GetExpectedOutcome() string {
	if x != nil {
		return x.ExpectedOutcome
	}
	return ""
}

type Scenario struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity      *metadata.Identity      `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	ProjectId     string                  `protobuf:"bytes,2,opt,name=projectId,proto3" json:"projectId,omitempty"`
	Name          string                  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Prerequisites string                  `protobuf:"bytes,5,opt,name=prerequisites,proto3" json:"prerequisites,omitempty"`
	Steps         []*Step                 `protobuf:"bytes,6,rep,name=steps,proto3" json:"steps,omitempty"`
	Issues        []*metadata.LinkedIssue `protobuf:"bytes,7,rep,name=issues,proto3" json:"issues,omitempty"`
	Labels        []string                `protobuf:"bytes,8,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *Scenario) Reset() {
	*x = Scenario{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scenario_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scenario) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scenario) ProtoMessage() {}

func (x *Scenario) ProtoReflect() protoreflect.Message {
	mi := &file_scenario_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scenario.ProtoReflect.Descriptor instead.
func (*Scenario) Descriptor() ([]byte, []int) {
	return file_scenario_proto_rawDescGZIP(), []int{1}
}

func (x *Scenario) GetIdentity() *metadata.Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *Scenario) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Scenario) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Scenario) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Scenario) GetPrerequisites() string {
	if x != nil {
		return x.Prerequisites
	}
	return ""
}

func (x *Scenario) GetSteps() []*Step {
	if x != nil {
		return x.Steps
	}
	return nil
}

func (x *Scenario) GetIssues() []*metadata.LinkedIssue {
	if x != nil {
		return x.Issues
	}
	return nil
}

func (x *Scenario) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

var File_scenario_proto protoreflect.FileDescriptor

var file_scenario_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x22, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x74,
	0x63, 0x68, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x69, 0x6f, 0x75, 0x73, 0x6b, 0x69,
	0x74, 0x74, 0x65, 0x6e, 0x1a, 0x17, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01,
	0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x28, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x63,
	0x6f, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x22, 0xef, 0x02, 0x0a, 0x08, 0x53,
	0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x48, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x74, 0x63, 0x68, 0x70, 0x6f, 0x73, 0x74,
	0x2e, 0x63, 0x75, 0x72, 0x69, 0x6f, 0x75, 0x73, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x2e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x73, 0x69, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72,
	0x65, 0x72, 0x65, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x05, 0x73,
	0x74, 0x65, 0x70, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x63, 0x65,
	0x6e, 0x61, 0x72, 0x69, 0x6f, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x74, 0x63, 0x68, 0x70, 0x6f, 0x73,
	0x74, 0x2e, 0x63, 0x75, 0x72, 0x69, 0x6f, 0x75, 0x73, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x2e,
	0x53, 0x74, 0x65, 0x70, 0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x12, 0x47, 0x0a, 0x06, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x74, 0x63, 0x68, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x69, 0x6f, 0x75, 0x73, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x6e,
	0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x06, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x42, 0x3c, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x75, 0x72, 0x69, 0x6f,
	0x75, 0x73, 0x2d, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x2f, 0x73, 0x63, 0x72, 0x61, 0x74, 0x63,
	0x68, 0x2d, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x63, 0x65, 0x6e, 0x61, 0x72, 0x69, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_scenario_proto_rawDescOnce sync.Once
	file_scenario_proto_rawDescData = file_scenario_proto_rawDesc
)

func file_scenario_proto_rawDescGZIP() []byte {
	file_scenario_proto_rawDescOnce.Do(func() {
		file_scenario_proto_rawDescData = protoimpl.X.CompressGZIP(file_scenario_proto_rawDescData)
	})
	return file_scenario_proto_rawDescData
}

var file_scenario_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_scenario_proto_goTypes = []interface{}{
	(*Step)(nil),                 // 0: scenario.scratchpost.curiouskitten.Step
	(*Scenario)(nil),             // 1: scenario.scratchpost.curiouskitten.Scenario
	(*metadata.Identity)(nil),    // 2: metadata.scratchpost.curiouskitten.Identity
	(*metadata.LinkedIssue)(nil), // 3: metadata.scratchpost.curiouskitten.LinkedIssue
}
var file_scenario_proto_depIdxs = []int32{
	2, // 0: scenario.scratchpost.curiouskitten.Scenario.identity:type_name -> metadata.scratchpost.curiouskitten.Identity
	0, // 1: scenario.scratchpost.curiouskitten.Scenario.steps:type_name -> scenario.scratchpost.curiouskitten.Step
	3, // 2: scenario.scratchpost.curiouskitten.Scenario.issues:type_name -> metadata.scratchpost.curiouskitten.LinkedIssue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_scenario_proto_init() }
func file_scenario_proto_init() {
	if File_scenario_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scenario_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Step); i {
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
		file_scenario_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scenario); i {
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
			RawDescriptor: file_scenario_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_scenario_proto_goTypes,
		DependencyIndexes: file_scenario_proto_depIdxs,
		MessageInfos:      file_scenario_proto_msgTypes,
	}.Build()
	File_scenario_proto = out.File
	file_scenario_proto_rawDesc = nil
	file_scenario_proto_goTypes = nil
	file_scenario_proto_depIdxs = nil
}
