// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: space.proto

package __

import (
	context "context"
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Owner    string `protobuf:"bytes,2,opt,name=Owner,proto3" json:"Owner,omitempty"`
	Manager  bool   `protobuf:"varint,3,opt,name=Manager,proto3" json:"Manager,omitempty"`
	Seat     int32  `protobuf:"varint,4,opt,name=Seat,proto3" json:"Seat,omitempty"`
	Capacity int32  `protobuf:"varint,5,opt,name=Capacity,proto3" json:"Capacity,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_space_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_space_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_space_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Message) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Message) GetManager() bool {
	if x != nil {
		return x.Manager
	}
	return false
}

func (x *Message) GetSeat() int32 {
	if x != nil {
		return x.Seat
	}
	return 0
}

func (x *Message) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

var File_space_proto protoreflect.FileDescriptor

var file_space_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x22, 0x7d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x65, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x53, 0x65, 0x61, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x43, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x32, 0x3b, 0x0a, 0x0c, 0x53, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0e,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_space_proto_rawDescOnce sync.Once
	file_space_proto_rawDescData = file_space_proto_rawDesc
)

func file_space_proto_rawDescGZIP() []byte {
	file_space_proto_rawDescOnce.Do(func() {
		file_space_proto_rawDescData = protoimpl.X.CompressGZIP(file_space_proto_rawDescData)
	})
	return file_space_proto_rawDescData
}

var file_space_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_space_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: space.Message
}
var file_space_proto_depIdxs = []int32{
	0, // 0: space.SpaceService.GetList:input_type -> space.Message
	0, // 1: space.SpaceService.GetList:output_type -> space.Message
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_space_proto_init() }
func file_space_proto_init() {
	if File_space_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_space_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_space_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_space_proto_goTypes,
		DependencyIndexes: file_space_proto_depIdxs,
		MessageInfos:      file_space_proto_msgTypes,
	}.Build()
	File_space_proto = out.File
	file_space_proto_rawDesc = nil
	file_space_proto_goTypes = nil
	file_space_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SpaceServiceClient is the client API for SpaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpaceServiceClient interface {
	GetList(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type spaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpaceServiceClient(cc grpc.ClientConnInterface) SpaceServiceClient {
	return &spaceServiceClient{cc}
}

func (c *spaceServiceClient) GetList(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/space.SpaceService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpaceServiceServer is the server API for SpaceService service.
type SpaceServiceServer interface {
	GetList(context.Context, *Message) (*Message, error)
}

// UnimplementedSpaceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSpaceServiceServer struct {
}

func (*UnimplementedSpaceServiceServer) GetList(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}

func RegisterSpaceServiceServer(s *grpc.Server, srv SpaceServiceServer) {
	s.RegisterService(&_SpaceService_serviceDesc, srv)
}

func _SpaceService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpaceServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/space.SpaceService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpaceServiceServer).GetList(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _SpaceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "space.SpaceService",
	HandlerType: (*SpaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _SpaceService_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "space.proto",
}
