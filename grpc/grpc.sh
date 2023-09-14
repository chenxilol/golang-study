protoc --go_out=plugins=grpc:.  hello.proto
#protoc -I  ./many_proto --go_out=plugins=grpc:./many_proto  ./many_proto/order.proto
#protoc -I  ./many_proto --go_out=plugins=grpc:./many_proto  ./many_proto/video.proto
#protoc -I  ./many_proto --go_out=plugins=grpc:./many_proto  ./many_proto/common.proto

