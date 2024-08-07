// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: pkg/api/api/api.proto

package api

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_api_proto_rawDescGZIP(), []int{0}
}

type NodeGPUResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalGpuCount int32            `protobuf:"varint,1,opt,name=total_gpu_count,json=totalGpuCount,proto3" json:"total_gpu_count,omitempty"`
	IndexUuidMap  map[string]int32 `protobuf:"bytes,2,rep,name=index_uuid_map,json=indexUuidMap,proto3" json:"index_uuid_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	NvlinkInfo    []*NVLink        `protobuf:"bytes,3,rep,name=nvlink_info,json=nvlinkInfo,proto3" json:"nvlink_info,omitempty"`
}

func (x *NodeGPUResponse) Reset() {
	*x = NodeGPUResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeGPUResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeGPUResponse) ProtoMessage() {}

func (x *NodeGPUResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeGPUResponse.ProtoReflect.Descriptor instead.
func (*NodeGPUResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *NodeGPUResponse) GetTotalGpuCount() int32 {
	if x != nil {
		return x.TotalGpuCount
	}
	return 0
}

func (x *NodeGPUResponse) GetIndexUuidMap() map[string]int32 {
	if x != nil {
		return x.IndexUuidMap
	}
	return nil
}

func (x *NodeGPUResponse) GetNvlinkInfo() []*NVLink {
	if x != nil {
		return x.NvlinkInfo
	}
	return nil
}

type NVLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gpu1Uuid  string `protobuf:"bytes,1,opt,name=gpu1uuid,proto3" json:"gpu1uuid,omitempty"`
	Gpu2Uuid  string `protobuf:"bytes,2,opt,name=gpu2uuid,proto3" json:"gpu2uuid,omitempty"`
	Lanecount int32  `protobuf:"varint,3,opt,name=lanecount,proto3" json:"lanecount,omitempty"`
}

func (x *NVLink) Reset() {
	*x = NVLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_api_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NVLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NVLink) ProtoMessage() {}

func (x *NVLink) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NVLink.ProtoReflect.Descriptor instead.
func (*NVLink) Descriptor() ([]byte, []int) {
	return file_pkg_api_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *NVLink) GetGpu1Uuid() string {
	if x != nil {
		return x.Gpu1Uuid
	}
	return ""
}

func (x *NVLink) GetGpu2Uuid() string {
	if x != nil {
		return x.Gpu2Uuid
	}
	return ""
}

func (x *NVLink) GetLanecount() int32 {
	if x != nil {
		return x.Lanecount
	}
	return 0
}

var File_pkg_api_api_api_proto protoreflect.FileDescriptor

var file_pkg_api_api_api_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x09, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xf6, 0x01, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65,
	0x47, 0x50, 0x55, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x67, 0x70, 0x75, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x47, 0x70, 0x75, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x4c, 0x0a, 0x0e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x75, 0x75, 0x69,
	0x64, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x50, 0x55, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x55, 0x75, 0x69, 0x64, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0c, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x55, 0x75, 0x69, 0x64, 0x4d, 0x61,
	0x70, 0x12, 0x2c, 0x0a, 0x0b, 0x6e, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x56, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x0a, 0x6e, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x3f, 0x0a, 0x11, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x55, 0x75, 0x69, 0x64, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x5e, 0x0a, 0x06, 0x4e, 0x56, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x70,
	0x75, 0x31, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x70,
	0x75, 0x31, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x70, 0x75, 0x32, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x70, 0x75, 0x32, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x65, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x61, 0x6e, 0x65, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x32, 0x3d, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0b,
	0x4e, 0x6f, 0x64, 0x65, 0x47, 0x50, 0x55, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x47, 0x50, 0x55, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x22, 0x5a, 0x20, 0x67, 0x70, 0x75, 0x2d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2d, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_api_api_api_proto_rawDescOnce sync.Once
	file_pkg_api_api_api_proto_rawDescData = file_pkg_api_api_api_proto_rawDesc
)

func file_pkg_api_api_api_proto_rawDescGZIP() []byte {
	file_pkg_api_api_api_proto_rawDescOnce.Do(func() {
		file_pkg_api_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_api_api_api_proto_rawDescData)
	})
	return file_pkg_api_api_api_proto_rawDescData
}

var file_pkg_api_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_api_api_api_proto_goTypes = []interface{}{
	(*Request)(nil),         // 0: api.Request
	(*NodeGPUResponse)(nil), // 1: api.NodeGPUResponse
	(*NVLink)(nil),          // 2: api.NVLink
	nil,                     // 3: api.NodeGPUResponse.IndexUuidMapEntry
}
var file_pkg_api_api_api_proto_depIdxs = []int32{
	3, // 0: api.NodeGPUResponse.index_uuid_map:type_name -> api.NodeGPUResponse.IndexUuidMapEntry
	2, // 1: api.NodeGPUResponse.nvlink_info:type_name -> api.NVLink
	0, // 2: api.Traveler.NodeGPUInfo:input_type -> api.Request
	1, // 3: api.Traveler.NodeGPUInfo:output_type -> api.NodeGPUResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_api_api_api_proto_init() }
func file_pkg_api_api_api_proto_init() {
	if File_pkg_api_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_api_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_pkg_api_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeGPUResponse); i {
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
		file_pkg_api_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NVLink); i {
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
			RawDescriptor: file_pkg_api_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_api_api_proto_goTypes,
		DependencyIndexes: file_pkg_api_api_api_proto_depIdxs,
		MessageInfos:      file_pkg_api_api_api_proto_msgTypes,
	}.Build()
	File_pkg_api_api_api_proto = out.File
	file_pkg_api_api_api_proto_rawDesc = nil
	file_pkg_api_api_api_proto_goTypes = nil
	file_pkg_api_api_api_proto_depIdxs = nil
}
