// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: internal/testprotos/test/test_import.proto

package test

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

type ImportEnum int32

const (
	ImportEnum_IMPORT_ZERO ImportEnum = 0
)

// Enum value maps for ImportEnum.
var (
	ImportEnum_name = map[int32]string{
		0: "IMPORT_ZERO",
	}
	ImportEnum_value = map[string]int32{
		"IMPORT_ZERO": 0,
	}
)

func (x ImportEnum) Enum() *ImportEnum {
	p := new(ImportEnum)
	*p = x
	return p
}

func (x ImportEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImportEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_testprotos_test_test_import_proto_enumTypes[0].Descriptor()
}

func (ImportEnum) Type() protoreflect.EnumType {
	return &file_internal_testprotos_test_test_import_proto_enumTypes[0]
}

func (x ImportEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ImportEnum) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ImportEnum(num)
	return nil
}

// Deprecated: Use ImportEnum.Descriptor instead.
func (ImportEnum) EnumDescriptor() ([]byte, []int) {
	return file_internal_testprotos_test_test_import_proto_rawDescGZIP(), []int{0}
}

type ImportMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ImportMessage) Reset() {
	*x = ImportMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_test_test_import_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportMessage) ProtoMessage() {}

func (x *ImportMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_test_test_import_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportMessage.ProtoReflect.Descriptor instead.
func (*ImportMessage) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_test_test_import_proto_rawDescGZIP(), []int{0}
}

var File_internal_testprotos_test_test_import_proto protoreflect.FileDescriptor

var file_internal_testprotos_test_test_import_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x22, 0x0f, 0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2a, 0x1d, 0x0a, 0x0a, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x12,
	0x0f, 0x0a, 0x0b, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00,
	0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x65, 0x6c, 0x69, 0x61, 0x73, 0x31, 0x32, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x74, 0x65, 0x73, 0x74,
}

var (
	file_internal_testprotos_test_test_import_proto_rawDescOnce sync.Once
	file_internal_testprotos_test_test_import_proto_rawDescData = file_internal_testprotos_test_test_import_proto_rawDesc
)

func file_internal_testprotos_test_test_import_proto_rawDescGZIP() []byte {
	file_internal_testprotos_test_test_import_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_test_test_import_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_testprotos_test_test_import_proto_rawDescData)
	})
	return file_internal_testprotos_test_test_import_proto_rawDescData
}

var file_internal_testprotos_test_test_import_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_testprotos_test_test_import_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_testprotos_test_test_import_proto_goTypes = []interface{}{
	(ImportEnum)(0),       // 0: goproto.proto.test.ImportEnum
	(*ImportMessage)(nil), // 1: goproto.proto.test.ImportMessage
}
var file_internal_testprotos_test_test_import_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_testprotos_test_test_import_proto_init() }
func file_internal_testprotos_test_test_import_proto_init() {
	if File_internal_testprotos_test_test_import_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_testprotos_test_test_import_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportMessage); i {
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
			RawDescriptor: file_internal_testprotos_test_test_import_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_test_test_import_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_test_test_import_proto_depIdxs,
		EnumInfos:         file_internal_testprotos_test_test_import_proto_enumTypes,
		MessageInfos:      file_internal_testprotos_test_test_import_proto_msgTypes,
	}.Build()
	File_internal_testprotos_test_test_import_proto = out.File
	file_internal_testprotos_test_test_import_proto_rawDesc = nil
	file_internal_testprotos_test_test_import_proto_goTypes = nil
	file_internal_testprotos_test_test_import_proto_depIdxs = nil
}
