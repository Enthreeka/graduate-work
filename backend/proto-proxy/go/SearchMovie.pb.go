// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: SearchMovie.proto

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

type SearchMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query    string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Cache    bool   `protobuf:"varint,2,opt,name=cache,proto3" json:"cache,omitempty"`
	RedisKey string `protobuf:"bytes,3,opt,name=redis_key,proto3" json:"redis_key,omitempty"`
}

func (x *SearchMovieRequest) Reset() {
	*x = SearchMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SearchMovie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMovieRequest) ProtoMessage() {}

func (x *SearchMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_SearchMovie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMovieRequest.ProtoReflect.Descriptor instead.
func (*SearchMovieRequest) Descriptor() ([]byte, []int) {
	return file_SearchMovie_proto_rawDescGZIP(), []int{0}
}

func (x *SearchMovieRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchMovieRequest) GetCache() bool {
	if x != nil {
		return x.Cache
	}
	return false
}

func (x *SearchMovieRequest) GetRedisKey() string {
	if x != nil {
		return x.RedisKey
	}
	return ""
}

type SearchMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hits    *TotalHits `protobuf:"bytes,1,opt,name=hits,proto3" json:"hits,omitempty"`
	Suggest *Suggest   `protobuf:"bytes,2,opt,name=suggest,proto3" json:"suggest,omitempty"`
	Status  string     `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SearchMovieResponse) Reset() {
	*x = SearchMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SearchMovie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMovieResponse) ProtoMessage() {}

func (x *SearchMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_SearchMovie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMovieResponse.ProtoReflect.Descriptor instead.
func (*SearchMovieResponse) Descriptor() ([]byte, []int) {
	return file_SearchMovie_proto_rawDescGZIP(), []int{1}
}

func (x *SearchMovieResponse) GetHits() *TotalHits {
	if x != nil {
		return x.Hits
	}
	return nil
}

func (x *SearchMovieResponse) GetSuggest() *Suggest {
	if x != nil {
		return x.Suggest
	}
	return nil
}

func (x *SearchMovieResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_SearchMovie_proto protoreflect.FileDescriptor

var file_SearchMovie_proto_rawDesc = []byte{
	0x0a, 0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x64, 0x69, 0x73,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x64, 0x69,
	0x73, 0x5f, 0x6b, 0x65, 0x79, 0x22, 0x89, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a,
	0x04, 0x68, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x48,
	0x69, 0x74, 0x73, 0x52, 0x04, 0x68, 0x69, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x75, 0x67,
	0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x52, 0x07, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SearchMovie_proto_rawDescOnce sync.Once
	file_SearchMovie_proto_rawDescData = file_SearchMovie_proto_rawDesc
)

func file_SearchMovie_proto_rawDescGZIP() []byte {
	file_SearchMovie_proto_rawDescOnce.Do(func() {
		file_SearchMovie_proto_rawDescData = protoimpl.X.CompressGZIP(file_SearchMovie_proto_rawDescData)
	})
	return file_SearchMovie_proto_rawDescData
}

var file_SearchMovie_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_SearchMovie_proto_goTypes = []interface{}{
	(*SearchMovieRequest)(nil),  // 0: proxy_proto.SearchMovieRequest
	(*SearchMovieResponse)(nil), // 1: proxy_proto.SearchMovieResponse
	(*TotalHits)(nil),           // 2: proxy_proto.TotalHits
	(*Suggest)(nil),             // 3: proxy_proto.Suggest
}
var file_SearchMovie_proto_depIdxs = []int32{
	2, // 0: proxy_proto.SearchMovieResponse.hits:type_name -> proxy_proto.TotalHits
	3, // 1: proxy_proto.SearchMovieResponse.suggest:type_name -> proxy_proto.Suggest
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_SearchMovie_proto_init() }
func file_SearchMovie_proto_init() {
	if File_SearchMovie_proto != nil {
		return
	}
	file_elasticsearch_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SearchMovie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMovieRequest); i {
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
		file_SearchMovie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMovieResponse); i {
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
			RawDescriptor: file_SearchMovie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SearchMovie_proto_goTypes,
		DependencyIndexes: file_SearchMovie_proto_depIdxs,
		MessageInfos:      file_SearchMovie_proto_msgTypes,
	}.Build()
	File_SearchMovie_proto = out.File
	file_SearchMovie_proto_rawDesc = nil
	file_SearchMovie_proto_goTypes = nil
	file_SearchMovie_proto_depIdxs = nil
}
