// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: video.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_video_proto protoreflect.FileDescriptor

var file_video_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x3a, 0x0a, 0x0a, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x2c, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08,
	0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_video_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: proto.Request
	(*Response)(nil), // 1: proto.Response
}
var file_video_proto_depIdxs = []int32{
	0, // 0: proto.VideoProto.Message:input_type -> proto.Request
	1, // 1: proto.VideoProto.Message:output_type -> proto.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_video_proto_init() }
func file_video_proto_init() {
	if File_video_proto != nil {
		return
	}
	file_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_proto_goTypes,
		DependencyIndexes: file_video_proto_depIdxs,
	}.Build()
	File_video_proto = out.File
	file_video_proto_rawDesc = nil
	file_video_proto_goTypes = nil
	file_video_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VideoProtoClient is the client API for VideoProto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VideoProtoClient interface {
	Message(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type videoProtoClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoProtoClient(cc grpc.ClientConnInterface) VideoProtoClient {
	return &videoProtoClient{cc}
}

func (c *videoProtoClient) Message(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.VideoProto/Message", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoProtoServer is the server API for VideoProto service.
type VideoProtoServer interface {
	Message(context.Context, *Request) (*Response, error)
}

// UnimplementedVideoProtoServer can be embedded to have forward compatible implementations.
type UnimplementedVideoProtoServer struct {
}

func (*UnimplementedVideoProtoServer) Message(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Message not implemented")
}

func RegisterVideoProtoServer(s *grpc.Server, srv VideoProtoServer) {
	s.RegisterService(&_VideoProto_serviceDesc, srv)
}

func _VideoProto_Message_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoProtoServer).Message(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VideoProto/Message",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoProtoServer).Message(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _VideoProto_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.VideoProto",
	HandlerType: (*VideoProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Message",
			Handler:    _VideoProto_Message_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}
