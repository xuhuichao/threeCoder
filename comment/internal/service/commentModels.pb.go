// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: commentModels.proto

package service

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

type CommentModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	Id uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// @inject_tag: json:"user_id"
	UserId uint32 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	// @inject_tag: json:"video_id"
	VideoId uint32 `protobuf:"varint,3,opt,name=VideoId,proto3" json:"VideoId,omitempty"`
	// @inject_tag: json:"content"
	Content string `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	// @inject_tag: json:"create_time"
	CreateTime uint32 `protobuf:"varint,5,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
}

func (x *CommentModel) Reset() {
	*x = CommentModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commentModels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentModel) ProtoMessage() {}

func (x *CommentModel) ProtoReflect() protoreflect.Message {
	mi := &file_commentModels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentModel.ProtoReflect.Descriptor instead.
func (*CommentModel) Descriptor() ([]byte, []int) {
	return file_commentModels_proto_rawDescGZIP(), []int{0}
}

func (x *CommentModel) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommentModel) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CommentModel) GetVideoId() uint32 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *CommentModel) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CommentModel) GetCreateTime() uint32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

var File_commentModels_proto protoreflect.FileDescriptor

var file_commentModels_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x8a, 0x01, 0x0a, 0x0c, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commentModels_proto_rawDescOnce sync.Once
	file_commentModels_proto_rawDescData = file_commentModels_proto_rawDesc
)

func file_commentModels_proto_rawDescGZIP() []byte {
	file_commentModels_proto_rawDescOnce.Do(func() {
		file_commentModels_proto_rawDescData = protoimpl.X.CompressGZIP(file_commentModels_proto_rawDescData)
	})
	return file_commentModels_proto_rawDescData
}

var file_commentModels_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_commentModels_proto_goTypes = []interface{}{
	(*CommentModel)(nil), // 0: pb.CommentModel
}
var file_commentModels_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_commentModels_proto_init() }
func file_commentModels_proto_init() {
	if File_commentModels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commentModels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentModel); i {
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
			RawDescriptor: file_commentModels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commentModels_proto_goTypes,
		DependencyIndexes: file_commentModels_proto_depIdxs,
		MessageInfos:      file_commentModels_proto_msgTypes,
	}.Build()
	File_commentModels_proto = out.File
	file_commentModels_proto_rawDesc = nil
	file_commentModels_proto_goTypes = nil
	file_commentModels_proto_depIdxs = nil
}
