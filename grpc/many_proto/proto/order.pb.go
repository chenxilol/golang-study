// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: order.proto

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

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x3a, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08,
	0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_order_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: proto.Request
	(*Response)(nil), // 1: proto.Response
}
var file_order_proto_depIdxs = []int32{
	0, // 0: proto.OrderMessage.Print:input_type -> proto.Request
	1, // 1: proto.OrderMessage.Print:output_type -> proto.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	file_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OrderMessageClient is the client API for OrderMessage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderMessageClient interface {
	Print(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type orderMessageClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderMessageClient(cc grpc.ClientConnInterface) OrderMessageClient {
	return &orderMessageClient{cc}
}

func (c *orderMessageClient) Print(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.OrderMessage/Print", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderMessageServer is the server API for OrderMessage service.
type OrderMessageServer interface {
	Print(context.Context, *Request) (*Response, error)
}

// UnimplementedOrderMessageServer can be embedded to have forward compatible implementations.
type UnimplementedOrderMessageServer struct {
}

func (*UnimplementedOrderMessageServer) Print(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Print not implemented")
}

func RegisterOrderMessageServer(s *grpc.Server, srv OrderMessageServer) {
	s.RegisterService(&_OrderMessage_serviceDesc, srv)
}

func _OrderMessage_Print_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderMessageServer).Print(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OrderMessage/Print",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderMessageServer).Print(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderMessage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.OrderMessage",
	HandlerType: (*OrderMessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Print",
			Handler:    _OrderMessage_Print_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
