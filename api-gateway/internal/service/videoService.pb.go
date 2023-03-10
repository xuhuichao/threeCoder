// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: videoService.proto

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

type VideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         []uint32 `protobuf:"varint,1,rep,packed,name=Id,proto3" json:"Id,omitempty"`
	AuthorId   uint32   `protobuf:"varint,2,opt,name=AuthorId,proto3" json:"AuthorId,omitempty"`
	Title      string   `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	PlayUrl    string   `protobuf:"bytes,4,opt,name=PlayUrl,proto3" json:"PlayUrl,omitempty"`
	CoverUrl   string   `protobuf:"bytes,5,opt,name=CoverUrl,proto3" json:"CoverUrl,omitempty"`
	CreateTime uint32   `protobuf:"varint,6,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
}

func (x *VideoRequest) Reset() {
	*x = VideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoRequest) ProtoMessage() {}

func (x *VideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_videoService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoRequest.ProtoReflect.Descriptor instead.
func (*VideoRequest) Descriptor() ([]byte, []int) {
	return file_videoService_proto_rawDescGZIP(), []int{0}
}

func (x *VideoRequest) GetId() []uint32 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *VideoRequest) GetAuthorId() uint32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *VideoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *VideoRequest) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *VideoRequest) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *VideoRequest) GetCreateTime() uint32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

type VideoDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoDetail []*VideoModel `protobuf:"bytes,1,rep,name=VideoDetail,proto3" json:"VideoDetail,omitempty"`
	Code        uint32        `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *VideoDetailResponse) Reset() {
	*x = VideoDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoDetailResponse) ProtoMessage() {}

func (x *VideoDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_videoService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoDetailResponse.ProtoReflect.Descriptor instead.
func (*VideoDetailResponse) Descriptor() ([]byte, []int) {
	return file_videoService_proto_rawDescGZIP(), []int{1}
}

func (x *VideoDetailResponse) GetVideoDetail() []*VideoModel {
	if x != nil {
		return x.VideoDetail
	}
	return nil
}

func (x *VideoDetailResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Data string `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_videoService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_videoService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_videoService_proto_rawDescGZIP(), []int{2}
}

func (x *CommonResponse) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CommonResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CommonResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_videoService_proto protoreflect.FileDescriptor

var file_videoService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x11, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x0c,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x76, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x76, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x5b, 0x0a, 0x13, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x52, 0x0b, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0x4a, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x32, 0xff, 0x01,
	0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d,
	0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x42, 0x79, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a,
	0x10, 0x46, 0x69, 0x6e, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0f,
	0x46, 0x69, 0x6e, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12,
	0x10, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x1b, 0x5a, 0x19, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_videoService_proto_rawDescOnce sync.Once
	file_videoService_proto_rawDescData = file_videoService_proto_rawDesc
)

func file_videoService_proto_rawDescGZIP() []byte {
	file_videoService_proto_rawDescOnce.Do(func() {
		file_videoService_proto_rawDescData = protoimpl.X.CompressGZIP(file_videoService_proto_rawDescData)
	})
	return file_videoService_proto_rawDescData
}

var file_videoService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_videoService_proto_goTypes = []interface{}{
	(*VideoRequest)(nil),        // 0: pb.VideoRequest
	(*VideoDetailResponse)(nil), // 1: pb.VideoDetailResponse
	(*CommonResponse)(nil),      // 2: pb.CommonResponse
	(*VideoModel)(nil),          // 3: pb.VideoModel
}
var file_videoService_proto_depIdxs = []int32{
	3, // 0: pb.VideoDetailResponse.VideoDetail:type_name -> pb.VideoModel
	0, // 1: pb.VideoService.FindVideosByTime:input_type -> pb.VideoRequest
	0, // 2: pb.VideoService.FindVideosByUser:input_type -> pb.VideoRequest
	0, // 3: pb.VideoService.FindVideosByIds:input_type -> pb.VideoRequest
	0, // 4: pb.VideoService.CreateVideo:input_type -> pb.VideoRequest
	1, // 5: pb.VideoService.FindVideosByTime:output_type -> pb.VideoDetailResponse
	1, // 6: pb.VideoService.FindVideosByUser:output_type -> pb.VideoDetailResponse
	1, // 7: pb.VideoService.FindVideosByIds:output_type -> pb.VideoDetailResponse
	2, // 8: pb.VideoService.CreateVideo:output_type -> pb.CommonResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_videoService_proto_init() }
func file_videoService_proto_init() {
	if File_videoService_proto != nil {
		return
	}
	file_videoModels_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_videoService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoRequest); i {
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
		file_videoService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoDetailResponse); i {
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
		file_videoService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
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
			RawDescriptor: file_videoService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_videoService_proto_goTypes,
		DependencyIndexes: file_videoService_proto_depIdxs,
		MessageInfos:      file_videoService_proto_msgTypes,
	}.Build()
	File_videoService_proto = out.File
	file_videoService_proto_rawDesc = nil
	file_videoService_proto_goTypes = nil
	file_videoService_proto_depIdxs = nil
}
