// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: sms.proto

package sms

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

type SendSmsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action string `protobuf:"bytes,1,opt,name=Action,proto3" json:"Action,omitempty"` // 发送消息类型 login -
	Phone  string `protobuf:"bytes,2,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Msg    string `protobuf:"bytes,3,opt,name=Msg,proto3" json:"Msg,omitempty"` // 消息内容
}

func (x *SendSmsReq) Reset() {
	*x = SendSmsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsReq) ProtoMessage() {}

func (x *SendSmsReq) ProtoReflect() protoreflect.Message {
	mi := &file_sms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsReq.ProtoReflect.Descriptor instead.
func (*SendSmsReq) Descriptor() ([]byte, []int) {
	return file_sms_proto_rawDescGZIP(), []int{0}
}

func (x *SendSmsReq) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *SendSmsReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SendSmsReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SendSmsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *SendSmsResp) Reset() {
	*x = SendSmsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsResp) ProtoMessage() {}

func (x *SendSmsResp) ProtoReflect() protoreflect.Message {
	mi := &file_sms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsResp.ProtoReflect.Descriptor instead.
func (*SendSmsResp) Descriptor() ([]byte, []int) {
	return file_sms_proto_rawDescGZIP(), []int{1}
}

func (x *SendSmsResp) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SendSmsResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_sms_proto protoreflect.FileDescriptor

var file_sms_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x6d, 0x73,
	0x22, 0x4c, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x4d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x33,
	0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x4d, 0x73, 0x67, 0x32, 0x33, 0x0a, 0x03, 0x53, 0x6d, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x73, 0x65,
	0x6e, 0x64, 0x53, 0x6d, 0x73, 0x12, 0x0f, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x73, 0x6d,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sms_proto_rawDescOnce sync.Once
	file_sms_proto_rawDescData = file_sms_proto_rawDesc
)

func file_sms_proto_rawDescGZIP() []byte {
	file_sms_proto_rawDescOnce.Do(func() {
		file_sms_proto_rawDescData = protoimpl.X.CompressGZIP(file_sms_proto_rawDescData)
	})
	return file_sms_proto_rawDescData
}

var file_sms_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sms_proto_goTypes = []interface{}{
	(*SendSmsReq)(nil),  // 0: sms.SendSmsReq
	(*SendSmsResp)(nil), // 1: sms.SendSmsResp
}
var file_sms_proto_depIdxs = []int32{
	0, // 0: sms.Sms.sendSms:input_type -> sms.SendSmsReq
	1, // 1: sms.Sms.sendSms:output_type -> sms.SendSmsResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sms_proto_init() }
func file_sms_proto_init() {
	if File_sms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsReq); i {
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
		file_sms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsResp); i {
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
			RawDescriptor: file_sms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sms_proto_goTypes,
		DependencyIndexes: file_sms_proto_depIdxs,
		MessageInfos:      file_sms_proto_msgTypes,
	}.Build()
	File_sms_proto = out.File
	file_sms_proto_rawDesc = nil
	file_sms_proto_goTypes = nil
	file_sms_proto_depIdxs = nil
}
