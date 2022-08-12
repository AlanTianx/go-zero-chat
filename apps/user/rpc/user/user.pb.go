// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: user.proto

package user

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

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone   string `protobuf:"bytes,1,opt,name=Phone,proto3" json:"Phone,omitempty"`
	CodeKey string `protobuf:"bytes,2,opt,name=CodeKey,proto3" json:"CodeKey,omitempty"`
	Code    string `protobuf:"bytes,3,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *LoginReq) GetCodeKey() string {
	if x != nil {
		return x.CodeKey
	}
	return ""
}

func (x *LoginReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type LoginRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Uuid     string `protobuf:"bytes,2,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=Phone,proto3" json:"Phone,omitempty"`
}

func (x *LoginRes) Reset() {
	*x = LoginRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRes) ProtoMessage() {}

func (x *LoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRes.ProtoReflect.Descriptor instead.
func (*LoginRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *LoginRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LoginRes) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *LoginRes) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRes) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type LoginCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=Phone,proto3" json:"Phone,omitempty"`
}

func (x *LoginCodeReq) Reset() {
	*x = LoginCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginCodeReq) ProtoMessage() {}

func (x *LoginCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginCodeReq.ProtoReflect.Descriptor instead.
func (*LoginCodeReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *LoginCodeReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type LoginCodeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	CodeKey string `protobuf:"bytes,2,opt,name=CodeKey,proto3" json:"CodeKey,omitempty"`
}

func (x *LoginCodeRes) Reset() {
	*x = LoginCodeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginCodeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginCodeRes) ProtoMessage() {}

func (x *LoginCodeRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginCodeRes.ProtoReflect.Descriptor instead.
func (*LoginCodeRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *LoginCodeRes) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *LoginCodeRes) GetCodeKey() string {
	if x != nil {
		return x.CodeKey
	}
	return ""
}

type DetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *DetailReq) Reset() {
	*x = DetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailReq) ProtoMessage() {}

func (x *DetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailReq.ProtoReflect.Descriptor instead.
func (*DetailReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *DetailReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SaveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *SaveReq) Reset() {
	*x = SaveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveReq) ProtoMessage() {}

func (x *SaveReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveReq.ProtoReflect.Descriptor instead.
func (*SaveReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *SaveReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SaveReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type SaveRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SaveRes) Reset() {
	*x = SaveRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveRes) ProtoMessage() {}

func (x *SaveRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveRes.ProtoReflect.Descriptor instead.
func (*SaveRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

type GetUserTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetUserTokenReq) Reset() {
	*x = GetUserTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokenReq) ProtoMessage() {}

func (x *GetUserTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokenReq.ProtoReflect.Descriptor instead.
func (*GetUserTokenReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserTokenReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserTokenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	ExpireTime int64  `protobuf:"varint,2,opt,name=ExpireTime,proto3" json:"ExpireTime,omitempty"`
}

func (x *GetUserTokenRes) Reset() {
	*x = GetUserTokenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokenRes) ProtoMessage() {}

func (x *GetUserTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokenRes.ProtoReflect.Descriptor instead.
func (*GetUserTokenRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserTokenRes) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetUserTokenRes) GetExpireTime() int64 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

type CheckTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenMd5 string `protobuf:"bytes,1,opt,name=TokenMd5,proto3" json:"TokenMd5,omitempty"`
}

func (x *CheckTokenReq) Reset() {
	*x = CheckTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTokenReq) ProtoMessage() {}

func (x *CheckTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTokenReq.ProtoReflect.Descriptor instead.
func (*CheckTokenReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *CheckTokenReq) GetTokenMd5() string {
	if x != nil {
		return x.TokenMd5
	}
	return ""
}

type CheckTokenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Uuid     string `protobuf:"bytes,2,opt,name=Uuid,proto3" json:"Uuid,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *CheckTokenRes) Reset() {
	*x = CheckTokenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTokenRes) ProtoMessage() {}

func (x *CheckTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTokenRes.ProtoReflect.Descriptor instead.
func (*CheckTokenRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{10}
}

func (x *CheckTokenRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CheckTokenRes) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *CheckTokenRes) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x4e, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x64, 0x65, 0x4b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x64, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x60, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x75,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x22, 0x24, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x3c, 0x0a, 0x0c, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x64, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x43, 0x6f, 0x64, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x1b, 0x0a, 0x09, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x07, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x09, 0x0a, 0x07,
	0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x64, 0x35,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x64, 0x35,
	0x22, 0x4f, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x32, 0xaf, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x05, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x0d, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_user_proto_goTypes = []interface{}{
	(*LoginReq)(nil),        // 0: user.LoginReq
	(*LoginRes)(nil),        // 1: user.LoginRes
	(*LoginCodeReq)(nil),    // 2: user.LoginCodeReq
	(*LoginCodeRes)(nil),    // 3: user.LoginCodeRes
	(*DetailReq)(nil),       // 4: user.DetailReq
	(*SaveReq)(nil),         // 5: user.SaveReq
	(*SaveRes)(nil),         // 6: user.SaveRes
	(*GetUserTokenReq)(nil), // 7: user.GetUserTokenReq
	(*GetUserTokenRes)(nil), // 8: user.GetUserTokenRes
	(*CheckTokenReq)(nil),   // 9: user.CheckTokenReq
	(*CheckTokenRes)(nil),   // 10: user.CheckTokenRes
}
var file_user_proto_depIdxs = []int32{
	0,  // 0: user.User.Login:input_type -> user.LoginReq
	2,  // 1: user.User.LoginCode:input_type -> user.LoginCodeReq
	4,  // 2: user.User.Detail:input_type -> user.DetailReq
	5,  // 3: user.User.Save:input_type -> user.SaveReq
	7,  // 4: user.User.GetUserToken:input_type -> user.GetUserTokenReq
	9,  // 5: user.User.CheckUserToken:input_type -> user.CheckTokenReq
	1,  // 6: user.User.Login:output_type -> user.LoginRes
	3,  // 7: user.User.LoginCode:output_type -> user.LoginCodeRes
	1,  // 8: user.User.Detail:output_type -> user.LoginRes
	6,  // 9: user.User.Save:output_type -> user.SaveRes
	8,  // 10: user.User.GetUserToken:output_type -> user.GetUserTokenRes
	10, // 11: user.User.CheckUserToken:output_type -> user.CheckTokenRes
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRes); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginCodeReq); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginCodeRes); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailReq); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveReq); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveRes); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTokenReq); i {
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
		file_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTokenRes); i {
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
		file_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTokenReq); i {
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
		file_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTokenRes); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
