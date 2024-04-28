// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: GetMovieByID.proto

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

type GetMovieByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId int64 `protobuf:"varint,1,opt,name=movie_id,json=movieID,proto3" json:"movie_id,omitempty"`
}

func (x *GetMovieByIDRequest) Reset() {
	*x = GetMovieByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetMovieByID_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieByIDRequest) ProtoMessage() {}

func (x *GetMovieByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_GetMovieByID_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieByIDRequest.ProtoReflect.Descriptor instead.
func (*GetMovieByIDRequest) Descriptor() ([]byte, []int) {
	return file_GetMovieByID_proto_rawDescGZIP(), []int{0}
}

func (x *GetMovieByIDRequest) GetMovieId() int64 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

type GetMovieByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie *Movie `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
}

func (x *GetMovieByIDResponse) Reset() {
	*x = GetMovieByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetMovieByID_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieByIDResponse) ProtoMessage() {}

func (x *GetMovieByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_GetMovieByID_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieByIDResponse.ProtoReflect.Descriptor instead.
func (*GetMovieByIDResponse) Descriptor() ([]byte, []int) {
	return file_GetMovieByID_proto_rawDescGZIP(), []int{1}
}

func (x *GetMovieByIDResponse) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

var File_GetMovieByID_proto protoreflect.FileDescriptor

var file_GetMovieByID_proto_rawDesc = []byte{
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x79, 0x49, 0x44, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x44,
	0x22, 0x40, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetMovieByID_proto_rawDescOnce sync.Once
	file_GetMovieByID_proto_rawDescData = file_GetMovieByID_proto_rawDesc
)

func file_GetMovieByID_proto_rawDescGZIP() []byte {
	file_GetMovieByID_proto_rawDescOnce.Do(func() {
		file_GetMovieByID_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetMovieByID_proto_rawDescData)
	})
	return file_GetMovieByID_proto_rawDescData
}

var file_GetMovieByID_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_GetMovieByID_proto_goTypes = []interface{}{
	(*GetMovieByIDRequest)(nil),  // 0: proxy_proto.GetMovieByIDRequest
	(*GetMovieByIDResponse)(nil), // 1: proxy_proto.GetMovieByIDResponse
	(*Movie)(nil),                // 2: proxy_proto.Movie
}
var file_GetMovieByID_proto_depIdxs = []int32{
	2, // 0: proxy_proto.GetMovieByIDResponse.movie:type_name -> proxy_proto.Movie
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetMovieByID_proto_init() }
func file_GetMovieByID_proto_init() {
	if File_GetMovieByID_proto != nil {
		return
	}
	file_movie_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetMovieByID_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieByIDRequest); i {
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
		file_GetMovieByID_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieByIDResponse); i {
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
			RawDescriptor: file_GetMovieByID_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetMovieByID_proto_goTypes,
		DependencyIndexes: file_GetMovieByID_proto_depIdxs,
		MessageInfos:      file_GetMovieByID_proto_msgTypes,
	}.Build()
	File_GetMovieByID_proto = out.File
	file_GetMovieByID_proto_rawDesc = nil
	file_GetMovieByID_proto_goTypes = nil
	file_GetMovieByID_proto_depIdxs = nil
}
