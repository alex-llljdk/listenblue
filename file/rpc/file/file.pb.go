// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.19.4
// source: file.proto

package file

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

// 注册
type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg      string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	RecordId int32  `protobuf:"varint,3,opt,name=RecordId,proto3" json:"RecordId,omitempty"`
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{0}
}

func (x *UploadResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UploadResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *UploadResponse) GetRecordId() int32 {
	if x != nil {
		return x.RecordId
	}
	return 0
}

type FileRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                      // 记录
	Path          string `protobuf:"bytes,2,opt,name=Path,proto3" json:"Path,omitempty"`                   //文件路径
	Title         string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`                 //文件名称和标题
	UserId        uint64 `protobuf:"varint,4,opt,name=UserId,proto3" json:"UserId,omitempty"`              //用户ID
	DataType      uint32 `protobuf:"varint,5,opt,name=DataType,proto3" json:"DataType,omitempty"`          //文件类型 0为mp3,1为mp4  后续添加文档等文件
	Keyword       string `protobuf:"bytes,6,opt,name=Keyword,proto3" json:"Keyword,omitempty"`             //关键要点
	Summary       string `protobuf:"bytes,7,opt,name=Summary,proto3" json:"Summary,omitempty"`             //全文概述
	Scanning      string `protobuf:"bytes,8,opt,name=Scanning,proto3" json:"Scanning,omitempty"`           //全文速读
	SpeechSummary string `protobuf:"bytes,9,opt,name=SpeechSummary,proto3" json:"SpeechSummary,omitempty"` //发言总结
	Review        string `protobuf:"bytes,10,opt,name=Review,proto3" json:"Review,omitempty"`              //回顾
	TransText     string `protobuf:"bytes,11,opt,name=TransText,proto3" json:"TransText,omitempty"`        //原文
}

func (x *FileRecord) Reset() {
	*x = FileRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRecord) ProtoMessage() {}

func (x *FileRecord) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRecord.ProtoReflect.Descriptor instead.
func (*FileRecord) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{1}
}

func (x *FileRecord) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FileRecord) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileRecord) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FileRecord) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FileRecord) GetDataType() uint32 {
	if x != nil {
		return x.DataType
	}
	return 0
}

func (x *FileRecord) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *FileRecord) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *FileRecord) GetScanning() string {
	if x != nil {
		return x.Scanning
	}
	return ""
}

func (x *FileRecord) GetSpeechSummary() string {
	if x != nil {
		return x.SpeechSummary
	}
	return ""
}

func (x *FileRecord) GetReview() string {
	if x != nil {
		return x.Review
	}
	return ""
}

func (x *FileRecord) GetTransText() string {
	if x != nil {
		return x.TransText
	}
	return ""
}

var File_file_proto protoreflect.FileDescriptor

var file_file_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x22, 0x52, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x22, 0xa6, 0x02, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x63, 0x61, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x63, 0x61, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x53, 0x70, 0x65, 0x65,
	0x63, 0x68, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x65, 0x78, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x65, 0x78, 0x74, 0x32,
	0x3c, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x1a, 0x14, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_proto_rawDescOnce sync.Once
	file_file_proto_rawDescData = file_file_proto_rawDesc
)

func file_file_proto_rawDescGZIP() []byte {
	file_file_proto_rawDescOnce.Do(func() {
		file_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_proto_rawDescData)
	})
	return file_file_proto_rawDescData
}

var file_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_file_proto_goTypes = []interface{}{
	(*UploadResponse)(nil), // 0: file.UploadResponse
	(*FileRecord)(nil),     // 1: file.FileRecord
}
var file_file_proto_depIdxs = []int32{
	1, // 0: file.File.UploadFile:input_type -> file.FileRecord
	0, // 1: file.File.UploadFile:output_type -> file.UploadResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file_proto_init() }
func file_file_proto_init() {
	if File_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResponse); i {
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
		file_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRecord); i {
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
			RawDescriptor: file_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_proto_goTypes,
		DependencyIndexes: file_file_proto_depIdxs,
		MessageInfos:      file_file_proto_msgTypes,
	}.Build()
	File_file_proto = out.File
	file_file_proto_rawDesc = nil
	file_file_proto_goTypes = nil
	file_file_proto_depIdxs = nil
}