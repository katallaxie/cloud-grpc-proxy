// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: fields.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Fields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Types that are assignable to Pattern:
	//	*Fields_Resolver
	Pattern isFields_Pattern `protobuf_oneof:"pattern"`
}

func (x *Fields) Reset() {
	*x = Fields{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fields_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fields) ProtoMessage() {}

func (x *Fields) ProtoReflect() protoreflect.Message {
	mi := &file_fields_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fields.ProtoReflect.Descriptor instead.
func (*Fields) Descriptor() ([]byte, []int) {
	return file_fields_proto_rawDescGZIP(), []int{0}
}

func (x *Fields) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

func (m *Fields) GetPattern() isFields_Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (x *Fields) GetResolver() *ResolverRule {
	if x, ok := x.GetPattern().(*Fields_Resolver); ok {
		return x.Resolver
	}
	return nil
}

type isFields_Pattern interface {
	isFields_Pattern()
}

type Fields_Resolver struct {
	Resolver *ResolverRule `protobuf:"bytes,10,opt,name=resolver,proto3,oneof"`
}

func (*Fields_Resolver) isFields_Pattern() {}

var File_fields_proto protoreflect.FileDescriptor

var file_fields_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x63, 0x67, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x0e, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x06,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x38, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x67, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x52, 0x75, 0x6c, 0x65,
	0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x42, 0x09, 0x0a, 0x07,
	0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x42, 0x39, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x74, 0x61, 0x6c, 0x6c, 0x61, 0x78, 0x69, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0xf8,
	0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fields_proto_rawDescOnce sync.Once
	file_fields_proto_rawDescData = file_fields_proto_rawDesc
)

func file_fields_proto_rawDescGZIP() []byte {
	file_fields_proto_rawDescOnce.Do(func() {
		file_fields_proto_rawDescData = protoimpl.X.CompressGZIP(file_fields_proto_rawDescData)
	})
	return file_fields_proto_rawDescData
}

var file_fields_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fields_proto_goTypes = []interface{}{
	(*Fields)(nil),       // 0: cgentron.api.Fields
	(*ResolverRule)(nil), // 1: cgentron.api.ResolverRule
}
var file_fields_proto_depIdxs = []int32{
	1, // 0: cgentron.api.Fields.resolver:type_name -> cgentron.api.ResolverRule
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_fields_proto_init() }
func file_fields_proto_init() {
	if File_fields_proto != nil {
		return
	}
	file_resolver_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_fields_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fields); i {
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
	file_fields_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Fields_Resolver)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fields_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fields_proto_goTypes,
		DependencyIndexes: file_fields_proto_depIdxs,
		MessageInfos:      file_fields_proto_msgTypes,
	}.Build()
	File_fields_proto = out.File
	file_fields_proto_rawDesc = nil
	file_fields_proto_goTypes = nil
	file_fields_proto_depIdxs = nil
}
