// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: user.proto

package user

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Gender int32

const (
	Gender_UNKNOWN Gender = 0
	Gender_MALE    Gender = 1
	Gender_FEMALE  Gender = 2
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "UNKNOWN",
		1: "MALE",
		2: "FEMALE",
	}
	Gender_value = map[string]int32{
		"UNKNOWN": 0,
		"MALE":    1,
		"FEMALE":  2,
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
	return file_user_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

type UserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserReq) Reset() {
	*x = UserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReq) ProtoMessage() {}

func (x *UserReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserReq.ProtoReflect.Descriptor instead.
func (*UserReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 性别，0-未知，1-男，2-女
	Gender Gender `protobuf:"varint,3,opt,name=gender,proto3,enum=user.Gender" json:"gender,omitempty"`
	// 角色，teacher-教师，student-学生
	Role       string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	CreateTime int64  `protobuf:"varint,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime int64  `protobuf:"varint,6,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *UserReply) Reset() {
	*x = UserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReply) ProtoMessage() {}

func (x *UserReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserReply.ProtoReflect.Descriptor instead.
func (*UserReply) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserReply) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_UNKNOWN
}

func (x *UserReply) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UserReply) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *UserReply) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type IdsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *IdsReq) Reset() {
	*x = IdsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdsReq) ProtoMessage() {}

func (x *IdsReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use IdsReq.ProtoReflect.Descriptor instead.
func (*IdsReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *IdsReq) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type UserListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*UserReply `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *UserListReply) Reset() {
	*x = UserListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListReply) ProtoMessage() {}

func (x *UserListReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UserListReply.ProtoReflect.Descriptor instead.
func (*UserListReply) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserListReply) GetList() []*UserReply {
	if x != nil {
		return x.List
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x19, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa9, 0x01,
	0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x1a, 0x0a, 0x06, 0x49, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x34, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x23, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x2a, 0x2b, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x32, 0x68, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x66, 0x69, 0x6e, 0x64, 0x4f,
	0x6e, 0x65, 0x12, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x09, 0x66, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12,
	0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_proto_goTypes = []interface{}{
	(Gender)(0),           // 0: user.gender
	(*UserReq)(nil),       // 1: user.UserReq
	(*UserReply)(nil),     // 2: user.UserReply
	(*IdsReq)(nil),        // 3: user.IdsReq
	(*UserListReply)(nil), // 4: user.UserListReply
}
var file_user_proto_depIdxs = []int32{
	0, // 0: user.UserReply.gender:type_name -> user.gender
	2, // 1: user.UserListReply.list:type_name -> user.UserReply
	1, // 2: user.UserService.findOne:input_type -> user.UserReq
	3, // 3: user.UserService.findByIds:input_type -> user.IdsReq
	2, // 4: user.UserService.findOne:output_type -> user.UserReply
	4, // 5: user.UserService.findByIds:output_type -> user.UserListReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReq); i {
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
			switch v := v.(*UserReply); i {
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
			switch v := v.(*IdsReq); i {
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
			switch v := v.(*UserListReply); i {
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
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		EnumInfos:         file_user_proto_enumTypes,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	// findone
	FindOne(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserReply, error)
	// findByIds
	FindByIds(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*UserListReply, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) FindOne(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/user.UserService/findOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FindByIds(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*UserListReply, error) {
	out := new(UserListReply)
	err := c.cc.Invoke(ctx, "/user.UserService/findByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	// findone
	FindOne(context.Context, *UserReq) (*UserReply, error)
	// findByIds
	FindByIds(context.Context, *IdsReq) (*UserListReply, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) FindOne(context.Context, *UserReq) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOne not implemented")
}
func (*UnimplementedUserServiceServer) FindByIds(context.Context, *IdsReq) (*UserListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByIds not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/FindOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindOne(ctx, req.(*UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FindByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/FindByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindByIds(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "findOne",
			Handler:    _UserService_FindOne_Handler,
		},
		{
			MethodName: "findByIds",
			Handler:    _UserService_FindByIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
