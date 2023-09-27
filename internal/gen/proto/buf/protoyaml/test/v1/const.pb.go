// Copyright 2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: buf/protoyaml/test/v1/const.proto

package testv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConstValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`       // file name, relative to root of source tree
	Package    string                     `protobuf:"bytes,2,opt,name=package,proto3" json:"package,omitempty"` // e.g. "foo", "foo.bar", etc.
	Dependency []string                   `protobuf:"bytes,3,rep,name=dependency,proto3" json:"dependency,omitempty"`
	Options    *descriptorpb.FileOptions  `protobuf:"bytes,8,opt,name=options,proto3" json:"options,omitempty"`
	Values     map[string]*structpb.Value `protobuf:"bytes,4,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConstValues) Reset() {
	*x = ConstValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_protoyaml_test_v1_const_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConstValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConstValues) ProtoMessage() {}

func (x *ConstValues) ProtoReflect() protoreflect.Message {
	mi := &file_buf_protoyaml_test_v1_const_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConstValues.ProtoReflect.Descriptor instead.
func (*ConstValues) Descriptor() ([]byte, []int) {
	return file_buf_protoyaml_test_v1_const_proto_rawDescGZIP(), []int{0}
}

func (x *ConstValues) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConstValues) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

func (x *ConstValues) GetDependency() []string {
	if x != nil {
		return x.Dependency
	}
	return nil
}

func (x *ConstValues) GetOptions() *descriptorpb.FileOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *ConstValues) GetValues() map[string]*structpb.Value {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_buf_protoyaml_test_v1_const_proto protoreflect.FileDescriptor

var file_buf_protoyaml_test_v1_const_proto_rawDesc = []byte{
	0x0a, 0x21, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x79, 0x61, 0x6d, 0x6c, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x79, 0x61,
	0x6d, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x02, 0x0a, 0x0b, 0x43,
	0x6f, 0x6e, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65,
	0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x36, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x46, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x79, 0x61, 0x6d, 0x6c,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x51, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xf0, 0x01, 0x0a, 0x19,
	0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x79, 0x61, 0x6d,
	0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x79, 0x61, 0x6d, 0x6c, 0x2d, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x66,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x79, 0x61, 0x6d, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f,
	0x76, 0x31, 0x3b, 0x74, 0x65, 0x73, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x50, 0x54, 0xaa,
	0x02, 0x15, 0x42, 0x75, 0x66, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x79, 0x61, 0x6d, 0x6c, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x42, 0x75, 0x66, 0x5c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x79, 0x61, 0x6d, 0x6c, 0x5c, 0x54, 0x65, 0x73, 0x74, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x21, 0x42, 0x75, 0x66, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x79, 0x61, 0x6d, 0x6c, 0x5c,
	0x54, 0x65, 0x73, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x3a, 0x54, 0x65, 0x73, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_protoyaml_test_v1_const_proto_rawDescOnce sync.Once
	file_buf_protoyaml_test_v1_const_proto_rawDescData = file_buf_protoyaml_test_v1_const_proto_rawDesc
)

func file_buf_protoyaml_test_v1_const_proto_rawDescGZIP() []byte {
	file_buf_protoyaml_test_v1_const_proto_rawDescOnce.Do(func() {
		file_buf_protoyaml_test_v1_const_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_protoyaml_test_v1_const_proto_rawDescData)
	})
	return file_buf_protoyaml_test_v1_const_proto_rawDescData
}

var file_buf_protoyaml_test_v1_const_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_buf_protoyaml_test_v1_const_proto_goTypes = []interface{}{
	(*ConstValues)(nil),              // 0: buf.protoyaml.test.v1.ConstValues
	nil,                              // 1: buf.protoyaml.test.v1.ConstValues.ValuesEntry
	(*descriptorpb.FileOptions)(nil), // 2: google.protobuf.FileOptions
	(*structpb.Value)(nil),           // 3: google.protobuf.Value
}
var file_buf_protoyaml_test_v1_const_proto_depIdxs = []int32{
	2, // 0: buf.protoyaml.test.v1.ConstValues.options:type_name -> google.protobuf.FileOptions
	1, // 1: buf.protoyaml.test.v1.ConstValues.values:type_name -> buf.protoyaml.test.v1.ConstValues.ValuesEntry
	3, // 2: buf.protoyaml.test.v1.ConstValues.ValuesEntry.value:type_name -> google.protobuf.Value
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_buf_protoyaml_test_v1_const_proto_init() }
func file_buf_protoyaml_test_v1_const_proto_init() {
	if File_buf_protoyaml_test_v1_const_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_protoyaml_test_v1_const_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConstValues); i {
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
			RawDescriptor: file_buf_protoyaml_test_v1_const_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_protoyaml_test_v1_const_proto_goTypes,
		DependencyIndexes: file_buf_protoyaml_test_v1_const_proto_depIdxs,
		MessageInfos:      file_buf_protoyaml_test_v1_const_proto_msgTypes,
	}.Build()
	File_buf_protoyaml_test_v1_const_proto = out.File
	file_buf_protoyaml_test_v1_const_proto_rawDesc = nil
	file_buf_protoyaml_test_v1_const_proto_goTypes = nil
	file_buf_protoyaml_test_v1_const_proto_depIdxs = nil
}
