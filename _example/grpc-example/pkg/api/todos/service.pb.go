// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/todos/service.proto

package todos

import (
	external "github.com/0B1t322/zero-validation/grpc-example/pkg/api/external"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Enum int32

const (
	Enum_ENUM_UNKNOWN Enum = 0
	Enum_ENUM_FIRST   Enum = 1
	Enum_ENUM_SECOND  Enum = 2
)

// Enum value maps for Enum.
var (
	Enum_name = map[int32]string{
		0: "ENUM_UNKNOWN",
		1: "ENUM_FIRST",
		2: "ENUM_SECOND",
	}
	Enum_value = map[string]int32{
		"ENUM_UNKNOWN": 0,
		"ENUM_FIRST":   1,
		"ENUM_SECOND":  2,
	}
)

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

func (x Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_api_todos_service_proto_enumTypes[0].Descriptor()
}

func (Enum) Type() protoreflect.EnumType {
	return &file_api_todos_service_proto_enumTypes[0]
}

func (x Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enum.Descriptor instead.
func (Enum) EnumDescriptor() ([]byte, []int) {
	return file_api_todos_service_proto_rawDescGZIP(), []int{0}
}

type CreateSomeRequest_InnerEnum int32

const (
	CreateSomeRequest_InnerEnumValue CreateSomeRequest_InnerEnum = 0
)

// Enum value maps for CreateSomeRequest_InnerEnum.
var (
	CreateSomeRequest_InnerEnum_name = map[int32]string{
		0: "InnerEnumValue",
	}
	CreateSomeRequest_InnerEnum_value = map[string]int32{
		"InnerEnumValue": 0,
	}
)

func (x CreateSomeRequest_InnerEnum) Enum() *CreateSomeRequest_InnerEnum {
	p := new(CreateSomeRequest_InnerEnum)
	*p = x
	return p
}

func (x CreateSomeRequest_InnerEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateSomeRequest_InnerEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_api_todos_service_proto_enumTypes[1].Descriptor()
}

func (CreateSomeRequest_InnerEnum) Type() protoreflect.EnumType {
	return &file_api_todos_service_proto_enumTypes[1]
}

func (x CreateSomeRequest_InnerEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateSomeRequest_InnerEnum.Descriptor instead.
func (CreateSomeRequest_InnerEnum) EnumDescriptor() ([]byte, []int) {
	return file_api_todos_service_proto_rawDescGZIP(), []int{0, 0}
}

type CreateSomeRequest struct {
	state          protoimpl.MessageState            `protogen:"open.v1"`
	BaseType       uint64                            `protobuf:"varint,1,opt,name=base_type,json=baseType,proto3" json:"base_type,omitempty"`
	BaseTypeArray  []uint64                          `protobuf:"varint,2,rep,packed,name=base_type_array,json=baseTypeArray,proto3" json:"base_type_array,omitempty"`
	OptBaseType    *uint64                           `protobuf:"varint,3,opt,name=opt_base_type,json=optBaseType,proto3,oneof" json:"opt_base_type,omitempty"`
	RepeatedBytes  [][]byte                          `protobuf:"bytes,4,rep,name=repeated_bytes,json=repeatedBytes,proto3" json:"repeated_bytes,omitempty"`
	Enum           Enum                              `protobuf:"varint,5,opt,name=enum,proto3,enum=zero_validation.example.todos.Enum" json:"enum,omitempty"`
	ExternalEnum   external.Enum                     `protobuf:"varint,6,opt,name=external_enum,json=externalEnum,proto3,enum=zero_validation.example.external.Enum" json:"external_enum,omitempty"`
	Entity         *Entity                           `protobuf:"bytes,7,opt,name=entity,proto3" json:"entity,omitempty"`
	Entities       []*Entity                         `protobuf:"bytes,8,rep,name=entities,proto3" json:"entities,omitempty"`
	Enums          []Enum                            `protobuf:"varint,9,rep,packed,name=enums,proto3,enum=zero_validation.example.todos.Enum" json:"enums,omitempty"`
	ExternalEnums  []external.Enum                   `protobuf:"varint,10,rep,packed,name=external_enums,json=externalEnums,proto3,enum=zero_validation.example.external.Enum" json:"external_enums,omitempty"`
	InnterMessage  *CreateSomeRequest_InnerMessage   `protobuf:"bytes,11,opt,name=innter_message,json=innterMessage,proto3" json:"innter_message,omitempty"`
	InnterMessages []*CreateSomeRequest_InnerMessage `protobuf:"bytes,12,rep,name=innter_messages,json=innterMessages,proto3" json:"innter_messages,omitempty"`
	// Types that are valid to be assigned to OneofExample:
	//
	//	*CreateSomeRequest_Uint
	//	*CreateSomeRequest_EnumType
	//	*CreateSomeRequest_InnerMessage_
	OneofExample  isCreateSomeRequest_OneofExample `protobuf_oneof:"oneof_example"`
	InnerEnum     CreateSomeRequest_InnerEnum      `protobuf:"varint,16,opt,name=inner_enum,json=innerEnum,proto3,enum=zero_validation.example.todos.CreateSomeRequest_InnerEnum" json:"inner_enum,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSomeRequest) Reset() {
	*x = CreateSomeRequest{}
	mi := &file_api_todos_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSomeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSomeRequest) ProtoMessage() {}

func (x *CreateSomeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_todos_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSomeRequest.ProtoReflect.Descriptor instead.
func (*CreateSomeRequest) Descriptor() ([]byte, []int) {
	return file_api_todos_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSomeRequest) GetBaseType() uint64 {
	if x != nil {
		return x.BaseType
	}
	return 0
}

func (x *CreateSomeRequest) GetBaseTypeArray() []uint64 {
	if x != nil {
		return x.BaseTypeArray
	}
	return nil
}

func (x *CreateSomeRequest) GetOptBaseType() uint64 {
	if x != nil && x.OptBaseType != nil {
		return *x.OptBaseType
	}
	return 0
}

func (x *CreateSomeRequest) GetRepeatedBytes() [][]byte {
	if x != nil {
		return x.RepeatedBytes
	}
	return nil
}

func (x *CreateSomeRequest) GetEnum() Enum {
	if x != nil {
		return x.Enum
	}
	return Enum_ENUM_UNKNOWN
}

func (x *CreateSomeRequest) GetExternalEnum() external.Enum {
	if x != nil {
		return x.ExternalEnum
	}
	return external.Enum(0)
}

func (x *CreateSomeRequest) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CreateSomeRequest) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *CreateSomeRequest) GetEnums() []Enum {
	if x != nil {
		return x.Enums
	}
	return nil
}

func (x *CreateSomeRequest) GetExternalEnums() []external.Enum {
	if x != nil {
		return x.ExternalEnums
	}
	return nil
}

func (x *CreateSomeRequest) GetInnterMessage() *CreateSomeRequest_InnerMessage {
	if x != nil {
		return x.InnterMessage
	}
	return nil
}

func (x *CreateSomeRequest) GetInnterMessages() []*CreateSomeRequest_InnerMessage {
	if x != nil {
		return x.InnterMessages
	}
	return nil
}

func (x *CreateSomeRequest) GetOneofExample() isCreateSomeRequest_OneofExample {
	if x != nil {
		return x.OneofExample
	}
	return nil
}

func (x *CreateSomeRequest) GetUint() uint64 {
	if x != nil {
		if x, ok := x.OneofExample.(*CreateSomeRequest_Uint); ok {
			return x.Uint
		}
	}
	return 0
}

func (x *CreateSomeRequest) GetEnumType() Enum {
	if x != nil {
		if x, ok := x.OneofExample.(*CreateSomeRequest_EnumType); ok {
			return x.EnumType
		}
	}
	return Enum_ENUM_UNKNOWN
}

func (x *CreateSomeRequest) GetInnerMessage() *CreateSomeRequest_InnerMessage {
	if x != nil {
		if x, ok := x.OneofExample.(*CreateSomeRequest_InnerMessage_); ok {
			return x.InnerMessage
		}
	}
	return nil
}

func (x *CreateSomeRequest) GetInnerEnum() CreateSomeRequest_InnerEnum {
	if x != nil {
		return x.InnerEnum
	}
	return CreateSomeRequest_InnerEnumValue
}

type isCreateSomeRequest_OneofExample interface {
	isCreateSomeRequest_OneofExample()
}

type CreateSomeRequest_Uint struct {
	Uint uint64 `protobuf:"varint,13,opt,name=uint,proto3,oneof"`
}

type CreateSomeRequest_EnumType struct {
	EnumType Enum `protobuf:"varint,14,opt,name=enum_type,json=enumType,proto3,enum=zero_validation.example.todos.Enum,oneof"`
}

type CreateSomeRequest_InnerMessage_ struct {
	InnerMessage *CreateSomeRequest_InnerMessage `protobuf:"bytes,15,opt,name=inner_message,json=innerMessage,proto3,oneof"`
}

func (*CreateSomeRequest_Uint) isCreateSomeRequest_OneofExample() {}

func (*CreateSomeRequest_EnumType) isCreateSomeRequest_OneofExample() {}

func (*CreateSomeRequest_InnerMessage_) isCreateSomeRequest_OneofExample() {}

type SomeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SomeResponse) Reset() {
	*x = SomeResponse{}
	mi := &file_api_todos_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SomeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SomeResponse) ProtoMessage() {}

func (x *SomeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_todos_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SomeResponse.ProtoReflect.Descriptor instead.
func (*SomeResponse) Descriptor() ([]byte, []int) {
	return file_api_todos_service_proto_rawDescGZIP(), []int{1}
}

func (x *SomeResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Entity struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Entity) Reset() {
	*x = Entity{}
	mi := &file_api_todos_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_api_todos_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_api_todos_service_proto_rawDescGZIP(), []int{2}
}

func (x *Entity) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateSomeRequest_InnerMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Some          string                 `protobuf:"bytes,1,opt,name=some,proto3" json:"some,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateSomeRequest_InnerMessage) Reset() {
	*x = CreateSomeRequest_InnerMessage{}
	mi := &file_api_todos_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateSomeRequest_InnerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSomeRequest_InnerMessage) ProtoMessage() {}

func (x *CreateSomeRequest_InnerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_todos_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSomeRequest_InnerMessage.ProtoReflect.Descriptor instead.
func (*CreateSomeRequest_InnerMessage) Descriptor() ([]byte, []int) {
	return file_api_todos_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CreateSomeRequest_InnerMessage) GetSome() string {
	if x != nil {
		return x.Some
	}
	return ""
}

var File_api_todos_service_proto protoreflect.FileDescriptor

var file_api_todos_service_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x7a, 0x65, 0x72, 0x6f, 0x5f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x09,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x26, 0x0a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x72,
	0x72, 0x61, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x27, 0x0a, 0x0d, 0x6f, 0x70, 0x74, 0x5f,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x48,
	0x01, 0x52, 0x0b, 0x6f, 0x70, 0x74, 0x42, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x5f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x65, 0x6e, 0x75,
	0x6d, 0x12, 0x4b, 0x0a, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x6e,
	0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x5f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x52, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x3d,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x41, 0x0a,
	0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x39, 0x0a, 0x05, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x23, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x52, 0x05, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x12, 0x4d, 0x0a, 0x0e, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0d, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x12, 0x64, 0x0a, 0x0e, 0x69, 0x6e,
	0x6e, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x0d, 0x69, 0x6e, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x66, 0x0a, 0x0f, 0x69, 0x6e, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x7a, 0x65, 0x72, 0x6f,
	0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x69, 0x6e, 0x6e, 0x74, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x04, 0x75, 0x69, 0x6e, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x04, 0x75, 0x69, 0x6e, 0x74, 0x12, 0x42,
	0x0a, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x23, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x73, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x48, 0x00, 0x52, 0x08, 0x65, 0x6e, 0x75, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x64, 0x0a, 0x0d, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x7a, 0x65, 0x72, 0x6f,
	0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x6e, 0x65,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x59, 0x0a, 0x0a, 0x69, 0x6e, 0x6e, 0x65,
	0x72, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x7a,
	0x65, 0x72, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x49,
	0x6e, 0x6e, 0x65, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x09, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x45,
	0x6e, 0x75, 0x6d, 0x1a, 0x22, 0x0a, 0x0c, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x6f, 0x6d, 0x65, 0x22, 0x1f, 0x0a, 0x09, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x45, 0x6e, 0x75,
	0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x10, 0x00, 0x42, 0x0f, 0x0a, 0x0d, 0x6f, 0x6e, 0x65, 0x6f,
	0x66, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6f, 0x70,
	0x74, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x1e, 0x0a, 0x0c, 0x53,
	0x6f, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x18, 0x0a, 0x06, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x39, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x10, 0x0a,
	0x0c, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x0e, 0x0a, 0x0a, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x10, 0x02,
	0x32, 0x61, 0x0a, 0x07, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x56, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6d, 0x65, 0x12, 0x30, 0x2e, 0x7a, 0x65, 0x72, 0x6f,
	0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x6f, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x81, 0x02, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x7a, 0x65, 0x72, 0x6f,
	0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x30, 0x42, 0x31, 0x74, 0x33, 0x32, 0x32, 0x2f, 0x7a, 0x65,
	0x72, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0xa2, 0x02, 0x03, 0x5a, 0x45, 0x54, 0xaa, 0x02, 0x1c,
	0x5a, 0x65, 0x72, 0x6f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0xca, 0x02, 0x1c, 0x5a,
	0x65, 0x72, 0x6f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0xe2, 0x02, 0x28, 0x5a, 0x65,
	0x72, 0x6f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x5a, 0x65, 0x72, 0x6f, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x3a, 0x3a, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_todos_service_proto_rawDescOnce sync.Once
	file_api_todos_service_proto_rawDescData []byte
)

func file_api_todos_service_proto_rawDescGZIP() []byte {
	file_api_todos_service_proto_rawDescOnce.Do(func() {
		file_api_todos_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_todos_service_proto_rawDesc), len(file_api_todos_service_proto_rawDesc)))
	})
	return file_api_todos_service_proto_rawDescData
}

var file_api_todos_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_todos_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_todos_service_proto_goTypes = []any{
	(Enum)(0),                              // 0: zero_validation.example.todos.Enum
	(CreateSomeRequest_InnerEnum)(0),       // 1: zero_validation.example.todos.CreateSomeRequest.InnerEnum
	(*CreateSomeRequest)(nil),              // 2: zero_validation.example.todos.CreateSomeRequest
	(*SomeResponse)(nil),                   // 3: zero_validation.example.todos.SomeResponse
	(*Entity)(nil),                         // 4: zero_validation.example.todos.Entity
	(*CreateSomeRequest_InnerMessage)(nil), // 5: zero_validation.example.todos.CreateSomeRequest.InnerMessage
	(external.Enum)(0),                     // 6: zero_validation.example.external.Enum
	(*emptypb.Empty)(nil),                  // 7: google.protobuf.Empty
}
var file_api_todos_service_proto_depIdxs = []int32{
	0,  // 0: zero_validation.example.todos.CreateSomeRequest.enum:type_name -> zero_validation.example.todos.Enum
	6,  // 1: zero_validation.example.todos.CreateSomeRequest.external_enum:type_name -> zero_validation.example.external.Enum
	4,  // 2: zero_validation.example.todos.CreateSomeRequest.entity:type_name -> zero_validation.example.todos.Entity
	4,  // 3: zero_validation.example.todos.CreateSomeRequest.entities:type_name -> zero_validation.example.todos.Entity
	0,  // 4: zero_validation.example.todos.CreateSomeRequest.enums:type_name -> zero_validation.example.todos.Enum
	6,  // 5: zero_validation.example.todos.CreateSomeRequest.external_enums:type_name -> zero_validation.example.external.Enum
	5,  // 6: zero_validation.example.todos.CreateSomeRequest.innter_message:type_name -> zero_validation.example.todos.CreateSomeRequest.InnerMessage
	5,  // 7: zero_validation.example.todos.CreateSomeRequest.innter_messages:type_name -> zero_validation.example.todos.CreateSomeRequest.InnerMessage
	0,  // 8: zero_validation.example.todos.CreateSomeRequest.enum_type:type_name -> zero_validation.example.todos.Enum
	5,  // 9: zero_validation.example.todos.CreateSomeRequest.inner_message:type_name -> zero_validation.example.todos.CreateSomeRequest.InnerMessage
	1,  // 10: zero_validation.example.todos.CreateSomeRequest.inner_enum:type_name -> zero_validation.example.todos.CreateSomeRequest.InnerEnum
	2,  // 11: zero_validation.example.todos.Example.CreateSome:input_type -> zero_validation.example.todos.CreateSomeRequest
	7,  // 12: zero_validation.example.todos.Example.CreateSome:output_type -> google.protobuf.Empty
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_api_todos_service_proto_init() }
func file_api_todos_service_proto_init() {
	if File_api_todos_service_proto != nil {
		return
	}
	file_api_todos_service_proto_msgTypes[0].OneofWrappers = []any{
		(*CreateSomeRequest_Uint)(nil),
		(*CreateSomeRequest_EnumType)(nil),
		(*CreateSomeRequest_InnerMessage_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_todos_service_proto_rawDesc), len(file_api_todos_service_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_todos_service_proto_goTypes,
		DependencyIndexes: file_api_todos_service_proto_depIdxs,
		EnumInfos:         file_api_todos_service_proto_enumTypes,
		MessageInfos:      file_api_todos_service_proto_msgTypes,
	}.Build()
	File_api_todos_service_proto = out.File
	file_api_todos_service_proto_goTypes = nil
	file_api_todos_service_proto_depIdxs = nil
}
