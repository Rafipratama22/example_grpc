// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: post_update.proto

package pb

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

type UpdatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string  `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title   *string `protobuf:"bytes,2,opt,name=Title,proto3,oneof" json:"Title,omitempty"`
	Content *string `protobuf:"bytes,3,opt,name=Content,proto3,oneof" json:"Content,omitempty"`
	Image   *string `protobuf:"bytes,4,opt,name=Image,proto3,oneof" json:"Image,omitempty"`
	User    *string `protobuf:"bytes,5,opt,name=User,proto3,oneof" json:"User,omitempty"`
}

func (x *UpdatePostRequest) Reset() {
	*x = UpdatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostRequest) ProtoMessage() {}

func (x *UpdatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostRequest.ProtoReflect.Descriptor instead.
func (*UpdatePostRequest) Descriptor() ([]byte, []int) {
	return file_post_update_proto_rawDescGZIP(), []int{0}
}

func (x *UpdatePostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePostRequest) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *UpdatePostRequest) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

func (x *UpdatePostRequest) GetImage() string {
	if x != nil && x.Image != nil {
		return *x.Image
	}
	return ""
}

func (x *UpdatePostRequest) GetUser() string {
	if x != nil && x.User != nil {
		return *x.User
	}
	return ""
}

var File_post_update_proto protoreflect.FileDescriptor

var file_post_update_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xba, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x03, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_post_update_proto_rawDescOnce sync.Once
	file_post_update_proto_rawDescData = file_post_update_proto_rawDesc
)

func file_post_update_proto_rawDescGZIP() []byte {
	file_post_update_proto_rawDescOnce.Do(func() {
		file_post_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_update_proto_rawDescData)
	})
	return file_post_update_proto_rawDescData
}

var file_post_update_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_post_update_proto_goTypes = []interface{}{
	(*UpdatePostRequest)(nil), // 0: pb.UpdatePostRequest
}
var file_post_update_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_post_update_proto_init() }
func file_post_update_proto_init() {
	if File_post_update_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_post_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostRequest); i {
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
	file_post_update_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_post_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_post_update_proto_goTypes,
		DependencyIndexes: file_post_update_proto_depIdxs,
		MessageInfos:      file_post_update_proto_msgTypes,
	}.Build()
	File_post_update_proto = out.File
	file_post_update_proto_rawDesc = nil
	file_post_update_proto_goTypes = nil
	file_post_update_proto_depIdxs = nil
}
