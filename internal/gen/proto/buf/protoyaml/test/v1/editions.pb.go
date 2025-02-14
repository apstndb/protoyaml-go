// Copyright 2023-2024 Buf Technologies, Inc.
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
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: buf/protoyaml/test/v1/editions.proto

package testv1

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

type OpenEnum int32

const (
	OpenEnum_OPEN_ENUM_UNSPECIFIED OpenEnum = 0
)

// Enum value maps for OpenEnum.
var (
	OpenEnum_name = map[int32]string{
		0: "OPEN_ENUM_UNSPECIFIED",
	}
	OpenEnum_value = map[string]int32{
		"OPEN_ENUM_UNSPECIFIED": 0,
	}
)

func (x OpenEnum) Enum() *OpenEnum {
	p := new(OpenEnum)
	*p = x
	return p
}

func (x OpenEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpenEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_protoyaml_test_v1_editions_proto_enumTypes[0].Descriptor()
}

func (OpenEnum) Type() protoreflect.EnumType {
	return &file_buf_protoyaml_test_v1_editions_proto_enumTypes[0]
}

func (x OpenEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpenEnum.Descriptor instead.
func (OpenEnum) EnumDescriptor() ([]byte, []int) {
	return file_buf_protoyaml_test_v1_editions_proto_rawDescGZIP(), []int{0}
}

type ClosedEnum int32

const (
	ClosedEnum_CLOSED_ENUM_UNSPECIFIED ClosedEnum = 0
)

// Enum value maps for ClosedEnum.
var (
	ClosedEnum_name = map[int32]string{
		0: "CLOSED_ENUM_UNSPECIFIED",
	}
	ClosedEnum_value = map[string]int32{
		"CLOSED_ENUM_UNSPECIFIED": 0,
	}
)

func (x ClosedEnum) Enum() *ClosedEnum {
	p := new(ClosedEnum)
	*p = x
	return p
}

func (x ClosedEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClosedEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_protoyaml_test_v1_editions_proto_enumTypes[1].Descriptor()
}

func (ClosedEnum) Type() protoreflect.EnumType {
	return &file_buf_protoyaml_test_v1_editions_proto_enumTypes[1]
}

func (x ClosedEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClosedEnum.Descriptor instead.
func (ClosedEnum) EnumDescriptor() ([]byte, []int) {
	return file_buf_protoyaml_test_v1_editions_proto_rawDescGZIP(), []int{1}
}

type EditionsTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   *string              `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Nested *EditionsTest_Nested `protobuf:"group,2,opt,name=Nested,json=nested" json:"nested,omitempty"`
	Enum   OpenEnum             `protobuf:"varint,3,opt,name=enum,enum=buf.protoyaml.test.v1.OpenEnum" json:"enum,omitempty"`
}

func (x *EditionsTest) Reset() {
	*x = EditionsTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_protoyaml_test_v1_editions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditionsTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditionsTest) ProtoMessage() {}

func (x *EditionsTest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_protoyaml_test_v1_editions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditionsTest.ProtoReflect.Descriptor instead.
func (*EditionsTest) Descriptor() ([]byte, []int) {
	return file_buf_protoyaml_test_v1_editions_proto_rawDescGZIP(), []int{0}
}

func (x *EditionsTest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *EditionsTest) GetNested() *EditionsTest_Nested {
	if x != nil {
		return x.Nested
	}
	return nil
}

func (x *EditionsTest) GetEnum() OpenEnum {
	if x != nil {
		return x.Enum
	}
	return OpenEnum_OPEN_ENUM_UNSPECIFIED
}

type EditionsTest_Nested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,name=ids" json:"ids,omitempty"`
}

func (x *EditionsTest_Nested) Reset() {
	*x = EditionsTest_Nested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_protoyaml_test_v1_editions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditionsTest_Nested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditionsTest_Nested) ProtoMessage() {}

func (x *EditionsTest_Nested) ProtoReflect() protoreflect.Message {
	mi := &file_buf_protoyaml_test_v1_editions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditionsTest_Nested.ProtoReflect.Descriptor instead.
func (*EditionsTest_Nested) Descriptor() ([]byte, []int) {
	return file_buf_protoyaml_test_v1_editions_proto_rawDescGZIP(), []int{0, 0}
}

func (x *EditionsTest_Nested) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

var File_buf_protoyaml_test_v1_editions_proto protoreflect.FileDescriptor

var file_buf_protoyaml_test_v1_editions_proto_rawDesc = []byte{
	0x0a, 0x24, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x79, 0x61, 0x6d, 0x6c, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x79, 0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x22, 0xd3, 0x01,
	0x0a, 0x0c, 0x45, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xaa, 0x01,
	0x02, 0x08, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x6e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x79, 0x61, 0x6d, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x28, 0x02, 0x52, 0x06, 0x6e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x79, 0x61,
	0x6d, 0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x45,
	0x6e, 0x75, 0x6d, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x08, 0x02, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d,
	0x1a, 0x21, 0x0a, 0x06, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x18, 0x02, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x2a, 0x25, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x12,
	0x19, 0x0a, 0x15, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x2a, 0x2f, 0x0a, 0x0a, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4c, 0x4f, 0x53,
	0x45, 0x44, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x04, 0x3a, 0x02, 0x10, 0x02, 0x42, 0xe9, 0x01, 0x0a, 0x19,
	0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x79, 0x61, 0x6d,
	0x6c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x45, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x62, 0x75, 0x66, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x79, 0x61,
	0x6d, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x79,
	0x61, 0x6d, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x65, 0x73, 0x74,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x50, 0x54, 0xaa, 0x02, 0x15, 0x42, 0x75, 0x66, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x79, 0x61, 0x6d, 0x6c, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x15, 0x42, 0x75, 0x66, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x79, 0x61, 0x6d, 0x6c,
	0x5c, 0x54, 0x65, 0x73, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21, 0x42, 0x75, 0x66, 0x5c, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x79, 0x61, 0x6d, 0x6c, 0x5c, 0x54, 0x65, 0x73, 0x74, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x42,
	0x75, 0x66, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x3a, 0x54,
	0x65, 0x73, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x70, 0xe8, 0x07,
}

var (
	file_buf_protoyaml_test_v1_editions_proto_rawDescOnce sync.Once
	file_buf_protoyaml_test_v1_editions_proto_rawDescData = file_buf_protoyaml_test_v1_editions_proto_rawDesc
)

func file_buf_protoyaml_test_v1_editions_proto_rawDescGZIP() []byte {
	file_buf_protoyaml_test_v1_editions_proto_rawDescOnce.Do(func() {
		file_buf_protoyaml_test_v1_editions_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_protoyaml_test_v1_editions_proto_rawDescData)
	})
	return file_buf_protoyaml_test_v1_editions_proto_rawDescData
}

var file_buf_protoyaml_test_v1_editions_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_buf_protoyaml_test_v1_editions_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_buf_protoyaml_test_v1_editions_proto_goTypes = []interface{}{
	(OpenEnum)(0),               // 0: buf.protoyaml.test.v1.OpenEnum
	(ClosedEnum)(0),             // 1: buf.protoyaml.test.v1.ClosedEnum
	(*EditionsTest)(nil),        // 2: buf.protoyaml.test.v1.EditionsTest
	(*EditionsTest_Nested)(nil), // 3: buf.protoyaml.test.v1.EditionsTest.Nested
}
var file_buf_protoyaml_test_v1_editions_proto_depIdxs = []int32{
	3, // 0: buf.protoyaml.test.v1.EditionsTest.nested:type_name -> buf.protoyaml.test.v1.EditionsTest.Nested
	0, // 1: buf.protoyaml.test.v1.EditionsTest.enum:type_name -> buf.protoyaml.test.v1.OpenEnum
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_buf_protoyaml_test_v1_editions_proto_init() }
func file_buf_protoyaml_test_v1_editions_proto_init() {
	if File_buf_protoyaml_test_v1_editions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_protoyaml_test_v1_editions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditionsTest); i {
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
		file_buf_protoyaml_test_v1_editions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditionsTest_Nested); i {
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
			RawDescriptor: file_buf_protoyaml_test_v1_editions_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_protoyaml_test_v1_editions_proto_goTypes,
		DependencyIndexes: file_buf_protoyaml_test_v1_editions_proto_depIdxs,
		EnumInfos:         file_buf_protoyaml_test_v1_editions_proto_enumTypes,
		MessageInfos:      file_buf_protoyaml_test_v1_editions_proto_msgTypes,
	}.Build()
	File_buf_protoyaml_test_v1_editions_proto = out.File
	file_buf_protoyaml_test_v1_editions_proto_rawDesc = nil
	file_buf_protoyaml_test_v1_editions_proto_goTypes = nil
	file_buf_protoyaml_test_v1_editions_proto_depIdxs = nil
}
