// Copyright 2016 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.20.1
// source: storage.proto

package storagepb

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

// SubtreeProto contains nodes of a subtree.
type SubtreeProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// subtree's prefix (must be a multiple of 8 bits)
	Prefix []byte `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// subtree's depth
	Depth int32 `protobuf:"varint,2,opt,name=depth,proto3" json:"depth,omitempty"`
	// map of suffix (within subtree) to subtree-leaf node hash
	Leaves map[string][]byte `protobuf:"bytes,4,rep,name=leaves,proto3" json:"leaves,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Map of suffix (within subtree) to subtree-internal node hash.
	// This structure is usually used in RAM as a cache, the internal nodes of
	// the subtree are not generally stored. However internal nodes are stored for
	// partially filled log subtrees.
	InternalNodes map[string][]byte `protobuf:"bytes,5,rep,name=internal_nodes,json=internalNodes,proto3" json:"internal_nodes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Used as a crosscheck on the internal node map by recording its expected
	// size after loading and repopulation.
	InternalNodeCount uint32 `protobuf:"varint,6,opt,name=internal_node_count,json=internalNodeCount,proto3" json:"internal_node_count,omitempty"`
}

func (x *SubtreeProto) Reset() {
	*x = SubtreeProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubtreeProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubtreeProto) ProtoMessage() {}

func (x *SubtreeProto) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubtreeProto.ProtoReflect.Descriptor instead.
func (*SubtreeProto) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{0}
}

func (x *SubtreeProto) GetPrefix() []byte {
	if x != nil {
		return x.Prefix
	}
	return nil
}

func (x *SubtreeProto) GetDepth() int32 {
	if x != nil {
		return x.Depth
	}
	return 0
}

func (x *SubtreeProto) GetLeaves() map[string][]byte {
	if x != nil {
		return x.Leaves
	}
	return nil
}

func (x *SubtreeProto) GetInternalNodes() map[string][]byte {
	if x != nil {
		return x.InternalNodes
	}
	return nil
}

func (x *SubtreeProto) GetInternalNodeCount() uint32 {
	if x != nil {
		return x.InternalNodeCount
	}
	return 0
}

var File_storage_proto protoreflect.FileDescriptor

var file_storage_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x70, 0x62, 0x22, 0xff, 0x02, 0x0a, 0x0c, 0x53,
	0x75, 0x62, 0x74, 0x72, 0x65, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68, 0x12, 0x3b, 0x0a, 0x06, 0x6c, 0x65, 0x61,
	0x76, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x74, 0x72, 0x65, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x6c, 0x65, 0x61, 0x76, 0x65, 0x73, 0x12, 0x51, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x74, 0x72,
	0x65, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x65, 0x61,
	0x76, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x40, 0x0a, 0x12, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x4e, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x42, 0x2e, 0x5a, 0x2c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x74, 0x72, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_proto_rawDescOnce sync.Once
	file_storage_proto_rawDescData = file_storage_proto_rawDesc
)

func file_storage_proto_rawDescGZIP() []byte {
	file_storage_proto_rawDescOnce.Do(func() {
		file_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_proto_rawDescData)
	})
	return file_storage_proto_rawDescData
}

var file_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_storage_proto_goTypes = []interface{}{
	(*SubtreeProto)(nil), // 0: storagepb.SubtreeProto
	nil,                  // 1: storagepb.SubtreeProto.LeavesEntry
	nil,                  // 2: storagepb.SubtreeProto.InternalNodesEntry
}
var file_storage_proto_depIdxs = []int32{
	1, // 0: storagepb.SubtreeProto.leaves:type_name -> storagepb.SubtreeProto.LeavesEntry
	2, // 1: storagepb.SubtreeProto.internal_nodes:type_name -> storagepb.SubtreeProto.InternalNodesEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_storage_proto_init() }
func file_storage_proto_init() {
	if File_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubtreeProto); i {
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
			RawDescriptor: file_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_proto_goTypes,
		DependencyIndexes: file_storage_proto_depIdxs,
		MessageInfos:      file_storage_proto_msgTypes,
	}.Build()
	File_storage_proto = out.File
	file_storage_proto_rawDesc = nil
	file_storage_proto_goTypes = nil
	file_storage_proto_depIdxs = nil
}
