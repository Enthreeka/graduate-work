// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: elasticsearch.proto

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

type Shards struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total      int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Successful int32 `protobuf:"varint,2,opt,name=successful,proto3" json:"successful,omitempty"`
	Skipped    int32 `protobuf:"varint,3,opt,name=skipped,proto3" json:"skipped,omitempty"`
	Failed     int32 `protobuf:"varint,4,opt,name=failed,proto3" json:"failed,omitempty"`
}

func (x *Shards) Reset() {
	*x = Shards{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elasticsearch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shards) ProtoMessage() {}

func (x *Shards) ProtoReflect() protoreflect.Message {
	mi := &file_elasticsearch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shards.ProtoReflect.Descriptor instead.
func (*Shards) Descriptor() ([]byte, []int) {
	return file_elasticsearch_proto_rawDescGZIP(), []int{0}
}

func (x *Shards) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Shards) GetSuccessful() int32 {
	if x != nil {
		return x.Successful
	}
	return 0
}

func (x *Shards) GetSkipped() int32 {
	if x != nil {
		return x.Skipped
	}
	return 0
}

func (x *Shards) GetFailed() int32 {
	if x != nil {
		return x.Failed
	}
	return 0
}

type Total struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value    int64  `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Relation string `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *Total) Reset() {
	*x = Total{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elasticsearch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Total) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Total) ProtoMessage() {}

func (x *Total) ProtoReflect() protoreflect.Message {
	mi := &file_elasticsearch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Total.ProtoReflect.Descriptor instead.
func (*Total) Descriptor() ([]byte, []int) {
	return file_elasticsearch_proto_rawDescGZIP(), []int{1}
}

func (x *Total) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Total) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

type Hits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XIndex  string  `protobuf:"bytes,1,opt,name=_index,proto3" json:"_index,omitempty"`
	XId     string  `protobuf:"bytes,2,opt,name=_id,proto3" json:"_id,omitempty"`
	XScore  float32 `protobuf:"fixed32,3,opt,name=_score,proto3" json:"_score,omitempty"`
	XSource *Movie  `protobuf:"bytes,4,opt,name=_source,proto3" json:"_source,omitempty"`
}

func (x *Hits) Reset() {
	*x = Hits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elasticsearch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hits) ProtoMessage() {}

func (x *Hits) ProtoReflect() protoreflect.Message {
	mi := &file_elasticsearch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hits.ProtoReflect.Descriptor instead.
func (*Hits) Descriptor() ([]byte, []int) {
	return file_elasticsearch_proto_rawDescGZIP(), []int{2}
}

func (x *Hits) GetXIndex() string {
	if x != nil {
		return x.XIndex
	}
	return ""
}

func (x *Hits) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *Hits) GetXScore() float32 {
	if x != nil {
		return x.XScore
	}
	return 0
}

func (x *Hits) GetXSource() *Movie {
	if x != nil {
		return x.XSource
	}
	return nil
}

type TotalHits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total    *Total  `protobuf:"bytes,1,opt,name=total,proto3" json:"total,omitempty"`
	MaxScore float32 `protobuf:"fixed32,2,opt,name=max_score,proto3" json:"max_score,omitempty"`
	Hits     []*Hits `protobuf:"bytes,3,rep,name=hits,proto3" json:"hits,omitempty"`
}

func (x *TotalHits) Reset() {
	*x = TotalHits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elasticsearch_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TotalHits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TotalHits) ProtoMessage() {}

func (x *TotalHits) ProtoReflect() protoreflect.Message {
	mi := &file_elasticsearch_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TotalHits.ProtoReflect.Descriptor instead.
func (*TotalHits) Descriptor() ([]byte, []int) {
	return file_elasticsearch_proto_rawDescGZIP(), []int{3}
}

func (x *TotalHits) GetTotal() *Total {
	if x != nil {
		return x.Total
	}
	return nil
}

func (x *TotalHits) GetMaxScore() float32 {
	if x != nil {
		return x.MaxScore
	}
	return 0
}

func (x *TotalHits) GetHits() []*Hits {
	if x != nil {
		return x.Hits
	}
	return nil
}

type Suggest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SimplePhrase []*SimplePhrase `protobuf:"bytes,1,rep,name=simple_phrase,proto3" json:"simple_phrase,omitempty"`
}

func (x *Suggest) Reset() {
	*x = Suggest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elasticsearch_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Suggest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Suggest) ProtoMessage() {}

func (x *Suggest) ProtoReflect() protoreflect.Message {
	mi := &file_elasticsearch_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Suggest.ProtoReflect.Descriptor instead.
func (*Suggest) Descriptor() ([]byte, []int) {
	return file_elasticsearch_proto_rawDescGZIP(), []int{4}
}

func (x *Suggest) GetSimplePhrase() []*SimplePhrase {
	if x != nil {
		return x.SimplePhrase
	}
	return nil
}

type SimplePhrase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Length  int64      `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Options []*Options `protobuf:"bytes,2,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *SimplePhrase) Reset() {
	*x = SimplePhrase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elasticsearch_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimplePhrase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimplePhrase) ProtoMessage() {}

func (x *SimplePhrase) ProtoReflect() protoreflect.Message {
	mi := &file_elasticsearch_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimplePhrase.ProtoReflect.Descriptor instead.
func (*SimplePhrase) Descriptor() ([]byte, []int) {
	return file_elasticsearch_proto_rawDescGZIP(), []int{5}
}

func (x *SimplePhrase) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *SimplePhrase) GetOptions() []*Options {
	if x != nil {
		return x.Options
	}
	return nil
}

type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  string  `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_elasticsearch_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_elasticsearch_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_elasticsearch_proto_rawDescGZIP(), []int{6}
}

func (x *Options) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Options) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_elasticsearch_proto protoreflect.FileDescriptor

var file_elasticsearch_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x70, 0x0a, 0x06, 0x53, 0x68, 0x61, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x22, 0x39, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x76, 0x0a, 0x04,
	0x48, 0x69, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x5f, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06,
	0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x07, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x22, 0x7a, 0x0a, 0x09, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x48, 0x69, 0x74,
	0x73, 0x12, 0x28, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x6d,
	0x61, 0x78, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09,
	0x6d, 0x61, 0x78, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x68, 0x69, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x69, 0x74, 0x73, 0x52, 0x04, 0x68, 0x69, 0x74, 0x73,
	0x22, 0x4a, 0x0a, 0x07, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0d, 0x73,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x52, 0x0d, 0x73,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x22, 0x56, 0x0a, 0x0c,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x12, 0x2e, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x33, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_elasticsearch_proto_rawDescOnce sync.Once
	file_elasticsearch_proto_rawDescData = file_elasticsearch_proto_rawDesc
)

func file_elasticsearch_proto_rawDescGZIP() []byte {
	file_elasticsearch_proto_rawDescOnce.Do(func() {
		file_elasticsearch_proto_rawDescData = protoimpl.X.CompressGZIP(file_elasticsearch_proto_rawDescData)
	})
	return file_elasticsearch_proto_rawDescData
}

var file_elasticsearch_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_elasticsearch_proto_goTypes = []interface{}{
	(*Shards)(nil),       // 0: proxy_proto.Shards
	(*Total)(nil),        // 1: proxy_proto.Total
	(*Hits)(nil),         // 2: proxy_proto.Hits
	(*TotalHits)(nil),    // 3: proxy_proto.TotalHits
	(*Suggest)(nil),      // 4: proxy_proto.Suggest
	(*SimplePhrase)(nil), // 5: proxy_proto.SimplePhrase
	(*Options)(nil),      // 6: proxy_proto.Options
	(*Movie)(nil),        // 7: proxy_proto.Movie
}
var file_elasticsearch_proto_depIdxs = []int32{
	7, // 0: proxy_proto.Hits._source:type_name -> proxy_proto.Movie
	1, // 1: proxy_proto.TotalHits.total:type_name -> proxy_proto.Total
	2, // 2: proxy_proto.TotalHits.hits:type_name -> proxy_proto.Hits
	5, // 3: proxy_proto.Suggest.simple_phrase:type_name -> proxy_proto.SimplePhrase
	6, // 4: proxy_proto.SimplePhrase.options:type_name -> proxy_proto.Options
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_elasticsearch_proto_init() }
func file_elasticsearch_proto_init() {
	if File_elasticsearch_proto != nil {
		return
	}
	file_movie_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_elasticsearch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shards); i {
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
		file_elasticsearch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Total); i {
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
		file_elasticsearch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hits); i {
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
		file_elasticsearch_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TotalHits); i {
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
		file_elasticsearch_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Suggest); i {
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
		file_elasticsearch_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimplePhrase); i {
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
		file_elasticsearch_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
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
			RawDescriptor: file_elasticsearch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_elasticsearch_proto_goTypes,
		DependencyIndexes: file_elasticsearch_proto_depIdxs,
		MessageInfos:      file_elasticsearch_proto_msgTypes,
	}.Build()
	File_elasticsearch_proto = out.File
	file_elasticsearch_proto_rawDesc = nil
	file_elasticsearch_proto_goTypes = nil
	file_elasticsearch_proto_depIdxs = nil
}
