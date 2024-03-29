// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: proto/sso/sso.proto

package sso

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Gender int32

const (
	Gender_GENDER_UNSPECIFIED Gender = 0
	Gender_GENDER_MALE        Gender = 1
	Gender_GENDER_FEMALE      Gender = 2
	Gender_GENDER_Others      Gender = 3
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "GENDER_UNSPECIFIED",
		1: "GENDER_MALE",
		2: "GENDER_FEMALE",
		3: "GENDER_Others",
	}
	Gender_value = map[string]int32{
		"GENDER_UNSPECIFIED": 0,
		"GENDER_MALE":        1,
		"GENDER_FEMALE":      2,
		"GENDER_Others":      3,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_sso_sso_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_proto_sso_sso_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_proto_sso_sso_proto_rawDescGZIP(), []int{0}
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action   string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Resource string `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sso_sso_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sso_sso_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_proto_sso_sso_proto_rawDescGZIP(), []int{0}
}

func (x *Object) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Object) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

type CheckPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    string  `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Object *Object `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
}

func (x *CheckPermissionRequest) Reset() {
	*x = CheckPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sso_sso_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPermissionRequest) ProtoMessage() {}

func (x *CheckPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sso_sso_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPermissionRequest.ProtoReflect.Descriptor instead.
func (*CheckPermissionRequest) Descriptor() ([]byte, []int) {
	return file_proto_sso_sso_proto_rawDescGZIP(), []int{1}
}

func (x *CheckPermissionRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *CheckPermissionRequest) GetObject() *Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type CheckPermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CheckPermissionResponse) Reset() {
	*x = CheckPermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sso_sso_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPermissionResponse) ProtoMessage() {}

func (x *CheckPermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sso_sso_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPermissionResponse.ProtoReflect.Descriptor instead.
func (*CheckPermissionResponse) Descriptor() ([]byte, []int) {
	return file_proto_sso_sso_proto_rawDescGZIP(), []int{2}
}

type GetUserByUIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetUserByUIDRequest) Reset() {
	*x = GetUserByUIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sso_sso_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByUIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByUIDRequest) ProtoMessage() {}

func (x *GetUserByUIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sso_sso_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByUIDRequest.ProtoReflect.Descriptor instead.
func (*GetUserByUIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_sso_sso_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserByUIDRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type GetUserByUIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Name  string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// google.protobuf.Timestamp join_time = 5;
	JoinTime    string   `protobuf:"bytes,5,opt,name=join_time,json=joinTime,proto3" json:"join_time,omitempty"`
	AvatarUrl   string   `protobuf:"bytes,6,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Gender      Gender   `protobuf:"varint,7,opt,name=gender,proto3,enum=sso.v1.Gender" json:"gender,omitempty"`
	Groups      []string `protobuf:"bytes,8,rep,name=groups,proto3" json:"groups,omitempty"`
	LarkUnionId string   `protobuf:"bytes,9,opt,name=lark_union_id,json=larkUnionId,proto3" json:"lark_union_id,omitempty"`
}

func (x *GetUserByUIDResponse) Reset() {
	*x = GetUserByUIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sso_sso_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByUIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByUIDResponse) ProtoMessage() {}

func (x *GetUserByUIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sso_sso_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByUIDResponse.ProtoReflect.Descriptor instead.
func (*GetUserByUIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_sso_sso_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserByUIDResponse) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *GetUserByUIDResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetUserByUIDResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetUserByUIDResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUserByUIDResponse) GetJoinTime() string {
	if x != nil {
		return x.JoinTime
	}
	return ""
}

func (x *GetUserByUIDResponse) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *GetUserByUIDResponse) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_GENDER_UNSPECIFIED
}

func (x *GetUserByUIDResponse) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *GetUserByUIDResponse) GetLarkUnionId() string {
	if x != nil {
		return x.LarkUnionId
	}
	return ""
}

type GetRolesByUIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetRolesByUIDRequest) Reset() {
	*x = GetRolesByUIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sso_sso_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRolesByUIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRolesByUIDRequest) ProtoMessage() {}

func (x *GetRolesByUIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sso_sso_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRolesByUIDRequest.ProtoReflect.Descriptor instead.
func (*GetRolesByUIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_sso_sso_proto_rawDescGZIP(), []int{5}
}

func (x *GetRolesByUIDRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type GetRolesByUIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roles []string `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *GetRolesByUIDResponse) Reset() {
	*x = GetRolesByUIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sso_sso_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRolesByUIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRolesByUIDResponse) ProtoMessage() {}

func (x *GetRolesByUIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sso_sso_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRolesByUIDResponse.ProtoReflect.Descriptor instead.
func (*GetRolesByUIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_sso_sso_proto_rawDescGZIP(), []int{6}
}

func (x *GetRolesByUIDResponse) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type GetUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid []string `protobuf:"bytes,1,rep,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetUsersRequest) Reset() {
	*x = GetUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sso_sso_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersRequest) ProtoMessage() {}

func (x *GetUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sso_sso_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersRequest.ProtoReflect.Descriptor instead.
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return file_proto_sso_sso_proto_rawDescGZIP(), []int{7}
}

func (x *GetUsersRequest) GetUid() []string {
	if x != nil {
		return x.Uid
	}
	return nil
}

type GetUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*GetUserByUIDResponse `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetUsersResponse) Reset() {
	*x = GetUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sso_sso_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersResponse) ProtoMessage() {}

func (x *GetUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sso_sso_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersResponse.ProtoReflect.Descriptor instead.
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return file_proto_sso_sso_proto_rawDescGZIP(), []int{8}
}

func (x *GetUsersResponse) GetUsers() []*GetUserByUIDResponse {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetGroupsDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups *structpb.Struct `protobuf:"bytes,1,opt,name=groups,proto3" json:"groups,omitempty"`
}

func (x *GetGroupsDetailResponse) Reset() {
	*x = GetGroupsDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sso_sso_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupsDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupsDetailResponse) ProtoMessage() {}

func (x *GetGroupsDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sso_sso_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupsDetailResponse.ProtoReflect.Descriptor instead.
func (*GetGroupsDetailResponse) Descriptor() ([]byte, []int) {
	return file_proto_sso_sso_proto_rawDescGZIP(), []int{9}
}

func (x *GetGroupsDetailResponse) GetGroups() *structpb.Struct {
	if x != nil {
		return x.Groups
	}
	return nil
}

var File_proto_sso_sso_proto protoreflect.FileDescriptor

var file_proto_sso_sso_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x73, 0x6f, 0x2f, 0x73, 0x73, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x52, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x79, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x88,
	0x02, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x6f, 0x69,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x6f,
	0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x26, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x61, 0x72, 0x6b, 0x5f, 0x75, 0x6e,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61,
	0x72, 0x6b, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x42,
	0x79, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x73, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22,
	0x4a, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2a, 0x57, 0x0a, 0x06, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x11,
	0x0a, 0x0d, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10,
	0x02, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4f, 0x74, 0x68, 0x65,
	0x72, 0x73, 0x10, 0x03, 0x32, 0x84, 0x03, 0x0a, 0x0a, 0x53, 0x53, 0x4f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x79, 0x55, 0x49, 0x44, 0x12, 0x1b, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x42, 0x79,
	0x55, 0x49, 0x44, 0x12, 0x1c, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x42, 0x79, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x42, 0x79, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x17, 0x2e, 0x73,
	0x73, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x73, 0x73, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x53, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x44, 0x4c,
	0x2f, 0x73, 0x73, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_sso_sso_proto_rawDescOnce sync.Once
	file_proto_sso_sso_proto_rawDescData = file_proto_sso_sso_proto_rawDesc
)

func file_proto_sso_sso_proto_rawDescGZIP() []byte {
	file_proto_sso_sso_proto_rawDescOnce.Do(func() {
		file_proto_sso_sso_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sso_sso_proto_rawDescData)
	})
	return file_proto_sso_sso_proto_rawDescData
}

var file_proto_sso_sso_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_sso_sso_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_sso_sso_proto_goTypes = []interface{}{
	(Gender)(0),                     // 0: sso.v1.Gender
	(*Object)(nil),                  // 1: sso.v1.Object
	(*CheckPermissionRequest)(nil),  // 2: sso.v1.CheckPermissionRequest
	(*CheckPermissionResponse)(nil), // 3: sso.v1.CheckPermissionResponse
	(*GetUserByUIDRequest)(nil),     // 4: sso.v1.GetUserByUIDRequest
	(*GetUserByUIDResponse)(nil),    // 5: sso.v1.GetUserByUIDResponse
	(*GetRolesByUIDRequest)(nil),    // 6: sso.v1.GetRolesByUIDRequest
	(*GetRolesByUIDResponse)(nil),   // 7: sso.v1.GetRolesByUIDResponse
	(*GetUsersRequest)(nil),         // 8: sso.v1.GetUsersRequest
	(*GetUsersResponse)(nil),        // 9: sso.v1.GetUsersResponse
	(*GetGroupsDetailResponse)(nil), // 10: sso.v1.GetGroupsDetailResponse
	(*structpb.Struct)(nil),         // 11: google.protobuf.Struct
	(*emptypb.Empty)(nil),           // 12: google.protobuf.Empty
}
var file_proto_sso_sso_proto_depIdxs = []int32{
	1,  // 0: sso.v1.CheckPermissionRequest.object:type_name -> sso.v1.Object
	0,  // 1: sso.v1.GetUserByUIDResponse.gender:type_name -> sso.v1.Gender
	5,  // 2: sso.v1.GetUsersResponse.users:type_name -> sso.v1.GetUserByUIDResponse
	11, // 3: sso.v1.GetGroupsDetailResponse.groups:type_name -> google.protobuf.Struct
	2,  // 4: sso.v1.SSOService.CheckPermission:input_type -> sso.v1.CheckPermissionRequest
	4,  // 5: sso.v1.SSOService.GetUserByUID:input_type -> sso.v1.GetUserByUIDRequest
	6,  // 6: sso.v1.SSOService.GetRolesByUID:input_type -> sso.v1.GetRolesByUIDRequest
	8,  // 7: sso.v1.SSOService.GetUsers:input_type -> sso.v1.GetUsersRequest
	12, // 8: sso.v1.SSOService.GetGroupsDetail:input_type -> google.protobuf.Empty
	3,  // 9: sso.v1.SSOService.CheckPermission:output_type -> sso.v1.CheckPermissionResponse
	5,  // 10: sso.v1.SSOService.GetUserByUID:output_type -> sso.v1.GetUserByUIDResponse
	7,  // 11: sso.v1.SSOService.GetRolesByUID:output_type -> sso.v1.GetRolesByUIDResponse
	9,  // 12: sso.v1.SSOService.GetUsers:output_type -> sso.v1.GetUsersResponse
	10, // 13: sso.v1.SSOService.GetGroupsDetail:output_type -> sso.v1.GetGroupsDetailResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_sso_sso_proto_init() }
func file_proto_sso_sso_proto_init() {
	if File_proto_sso_sso_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_sso_sso_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_proto_sso_sso_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPermissionRequest); i {
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
		file_proto_sso_sso_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPermissionResponse); i {
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
		file_proto_sso_sso_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByUIDRequest); i {
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
		file_proto_sso_sso_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByUIDResponse); i {
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
		file_proto_sso_sso_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRolesByUIDRequest); i {
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
		file_proto_sso_sso_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRolesByUIDResponse); i {
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
		file_proto_sso_sso_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersRequest); i {
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
		file_proto_sso_sso_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersResponse); i {
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
		file_proto_sso_sso_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupsDetailResponse); i {
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
			RawDescriptor: file_proto_sso_sso_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_sso_sso_proto_goTypes,
		DependencyIndexes: file_proto_sso_sso_proto_depIdxs,
		EnumInfos:         file_proto_sso_sso_proto_enumTypes,
		MessageInfos:      file_proto_sso_sso_proto_msgTypes,
	}.Build()
	File_proto_sso_sso_proto = out.File
	file_proto_sso_sso_proto_rawDesc = nil
	file_proto_sso_sso_proto_goTypes = nil
	file_proto_sso_sso_proto_depIdxs = nil
}
