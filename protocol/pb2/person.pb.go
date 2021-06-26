// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: person.proto

package hello_bigbomb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PhoneType int32

const (
	PhoneType_MOBILE PhoneType = 0
	PhoneType_HOME   PhoneType = 1
	PhoneType_WORK   PhoneType = 2
)

// Enum value maps for PhoneType.
var (
	PhoneType_name = map[int32]string{
		0: "MOBILE",
		1: "HOME",
		2: "WORK",
	}
	PhoneType_value = map[string]int32{
		"MOBILE": 0,
		"HOME":   1,
		"WORK":   2,
	}
)

func (x PhoneType) Enum() *PhoneType {
	p := new(PhoneType)
	*p = x
	return p
}

func (x PhoneType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PhoneType) Descriptor() protoreflect.EnumDescriptor {
	return file_person_proto_enumTypes[0].Descriptor()
}

func (PhoneType) Type() protoreflect.EnumType {
	return &file_person_proto_enumTypes[0]
}

func (x PhoneType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *PhoneType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = PhoneType(num)
	return nil
}

// Deprecated: Use PhoneType.Descriptor instead.
func (PhoneType) EnumDescriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{0}
}

type PhoneNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number *string    `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   *PhoneType `protobuf:"varint,2,opt,name=type,enum=com.tencent.ketang.open.bigbomb.PhoneType" json:"type,omitempty"`
}

func (x *PhoneNumber) Reset() {
	*x = PhoneNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhoneNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhoneNumber) ProtoMessage() {}

func (x *PhoneNumber) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhoneNumber.ProtoReflect.Descriptor instead.
func (*PhoneNumber) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{0}
}

func (x *PhoneNumber) GetNumber() string {
	if x != nil && x.Number != nil {
		return *x.Number
	}
	return ""
}

func (x *PhoneNumber) GetType() PhoneType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return PhoneType_MOBILE
}

// Person用户
type PersonRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  *string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id    *int32         `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email *string        `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phone []*PhoneNumber `protobuf:"bytes,4,rep,name=phone" json:"phone,omitempty"`
}

func (x *PersonRsp) Reset() {
	*x = PersonRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonRsp) ProtoMessage() {}

func (x *PersonRsp) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonRsp.ProtoReflect.Descriptor instead.
func (*PersonRsp) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{1}
}

func (x *PersonRsp) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *PersonRsp) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *PersonRsp) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *PersonRsp) GetPhone() []*PhoneNumber {
	if x != nil {
		return x.Phone
	}
	return nil
}

// 请求内容
type PersonReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *int32 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
}

func (x *PersonReq) Reset() {
	*x = PersonReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonReq) ProtoMessage() {}

func (x *PersonReq) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonReq.ProtoReflect.Descriptor instead.
func (*PersonReq) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{2}
}

func (x *PersonReq) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

var File_person_proto protoreflect.FileDescriptor

var file_person_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x2e, 0x6b, 0x65, 0x74, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6f, 0x6d, 0x62, 0x22,
	0x65, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x63, 0x65,
	0x6e, 0x74, 0x2e, 0x6b, 0x65, 0x74, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x2e, 0x62,
	0x69, 0x67, 0x62, 0x6f, 0x6d, 0x62, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x09, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x42,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x2e, 0x6b, 0x65, 0x74, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6f, 0x6d, 0x62, 0x2e,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x22, 0x1b, 0x0a, 0x09, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x2a,
	0x2b, 0x0a, 0x09, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06,
	0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x4d, 0x45,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x02, 0x32, 0x7d, 0x0a, 0x12,
	0x4b, 0x65, 0x74, 0x61, 0x6b, 0x6e, 0x67, 0x4f, 0x70, 0x65, 0x6e, 0x42, 0x69, 0x67, 0x42, 0x6f,
	0x6d, 0x62, 0x12, 0x67, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e,
	0x74, 0x2e, 0x6b, 0x65, 0x74, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x2e, 0x62, 0x69,
	0x67, 0x62, 0x6f, 0x6d, 0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a,
	0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x2e, 0x6b, 0x65,
	0x74, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6f, 0x6d,
	0x62, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x42, 0x0f, 0x5a, 0x0d, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x62, 0x69, 0x67, 0x62, 0x6f, 0x6d, 0x62,
}

var (
	file_person_proto_rawDescOnce sync.Once
	file_person_proto_rawDescData = file_person_proto_rawDesc
)

func file_person_proto_rawDescGZIP() []byte {
	file_person_proto_rawDescOnce.Do(func() {
		file_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_person_proto_rawDescData)
	})
	return file_person_proto_rawDescData
}

var file_person_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_person_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_person_proto_goTypes = []interface{}{
	(PhoneType)(0),      // 0: com.tencent.ketang.open.bigbomb.PhoneType
	(*PhoneNumber)(nil), // 1: com.tencent.ketang.open.bigbomb.PhoneNumber
	(*PersonRsp)(nil),   // 2: com.tencent.ketang.open.bigbomb.PersonRsp
	(*PersonReq)(nil),   // 3: com.tencent.ketang.open.bigbomb.PersonReq
}
var file_person_proto_depIdxs = []int32{
	0, // 0: com.tencent.ketang.open.bigbomb.PhoneNumber.type:type_name -> com.tencent.ketang.open.bigbomb.PhoneType
	1, // 1: com.tencent.ketang.open.bigbomb.PersonRsp.phone:type_name -> com.tencent.ketang.open.bigbomb.PhoneNumber
	3, // 2: com.tencent.ketang.open.bigbomb.KetakngOpenBigBomb.GetPersonById:input_type -> com.tencent.ketang.open.bigbomb.PersonReq
	2, // 3: com.tencent.ketang.open.bigbomb.KetakngOpenBigBomb.GetPersonById:output_type -> com.tencent.ketang.open.bigbomb.PersonRsp
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_person_proto_init() }
func file_person_proto_init() {
	if File_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhoneNumber); i {
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
		file_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonRsp); i {
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
		file_person_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonReq); i {
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
			RawDescriptor: file_person_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_person_proto_goTypes,
		DependencyIndexes: file_person_proto_depIdxs,
		EnumInfos:         file_person_proto_enumTypes,
		MessageInfos:      file_person_proto_msgTypes,
	}.Build()
	File_person_proto = out.File
	file_person_proto_rawDesc = nil
	file_person_proto_goTypes = nil
	file_person_proto_depIdxs = nil
}
