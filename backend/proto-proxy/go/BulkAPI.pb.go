// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: BulkAPI.proto

package proto

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

type BulkAPIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie []*Movie `protobuf:"bytes,1,rep,name=movie,proto3" json:"movie,omitempty"`
}

func (x *BulkAPIRequest) Reset() {
	*x = BulkAPIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BulkAPI_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkAPIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkAPIRequest) ProtoMessage() {}

func (x *BulkAPIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_BulkAPI_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkAPIRequest.ProtoReflect.Descriptor instead.
func (*BulkAPIRequest) Descriptor() ([]byte, []int) {
	return file_BulkAPI_proto_rawDescGZIP(), []int{0}
}

func (x *BulkAPIRequest) GetMovie() []*Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

type BulkAPIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *BulkAPIResponse) Reset() {
	*x = BulkAPIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BulkAPI_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkAPIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkAPIResponse) ProtoMessage() {}

func (x *BulkAPIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_BulkAPI_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkAPIResponse.ProtoReflect.Descriptor instead.
func (*BulkAPIResponse) Descriptor() ([]byte, []int) {
	return file_BulkAPI_proto_rawDescGZIP(), []int{1}
}

func (x *BulkAPIResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_BulkAPI_proto protoreflect.FileDescriptor

var file_BulkAPI_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x42, 0x75, 0x6c, 0x6b, 0x41, 0x50, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x0e, 0x42, 0x75, 0x6c,
	0x6b, 0x41, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x29, 0x0a, 0x0f, 0x42, 0x75, 0x6c, 0x6b, 0x41, 0x50, 0x49,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BulkAPI_proto_rawDescOnce sync.Once
	file_BulkAPI_proto_rawDescData = file_BulkAPI_proto_rawDesc
)

func file_BulkAPI_proto_rawDescGZIP() []byte {
	file_BulkAPI_proto_rawDescOnce.Do(func() {
		file_BulkAPI_proto_rawDescData = protoimpl.X.CompressGZIP(file_BulkAPI_proto_rawDescData)
	})
	return file_BulkAPI_proto_rawDescData
}

var file_BulkAPI_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_BulkAPI_proto_goTypes = []interface{}{
	(*BulkAPIRequest)(nil),  // 0: proxy_proto.BulkAPIRequest
	(*BulkAPIResponse)(nil), // 1: proxy_proto.BulkAPIResponse
	(*Movie)(nil),           // 2: proxy_proto.Movie
}
var file_BulkAPI_proto_depIdxs = []int32{
	2, // 0: proxy_proto.BulkAPIRequest.movie:type_name -> proxy_proto.Movie
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_BulkAPI_proto_init() }
func file_BulkAPI_proto_init() {
	if File_BulkAPI_proto != nil {
		return
	}
	file_movie_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BulkAPI_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkAPIRequest); i {
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
		file_BulkAPI_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkAPIResponse); i {
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
			RawDescriptor: file_BulkAPI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BulkAPI_proto_goTypes,
		DependencyIndexes: file_BulkAPI_proto_depIdxs,
		MessageInfos:      file_BulkAPI_proto_msgTypes,
	}.Build()
	File_BulkAPI_proto = out.File
	file_BulkAPI_proto_rawDesc = nil
	file_BulkAPI_proto_goTypes = nil
	file_BulkAPI_proto_depIdxs = nil
}
