# proto 语法
1. service对应的就是go里面的接口，可以作为服务器，客户端
2. rpc 对应的就是结构体中的方法
3. message 对应的也是结构体
4. reserved 标识号的保留,标识号[1 - 15]占一个字节，[16 - 2047] 占两个字节，所以尽量控制在15以内

# grpc 应用流程
1. 先编写proto文件，然后在protoc --go_out=plugins=grpc:.  manyService.proto 转化为go文件
2. protoc --go_out=. hello.proto
   protoc --go-grpc_out=. hello.proto 可能还有一部分的人看网上的操作是这样的 protoc --go_out=plugins=grpc:. hello.proto
2. sever:
   先监听端口，然后实例化grpc.NewServer()，在实例化结构体，结构体需要被方法继承，然后在进行注册服务，传入实例化的grpc.server,和结构体
   然后在调用服务的server，把listen传进去
3. client:
   先建立连接通过grpc.dial,传入地址和grpc.WithTransportCredentials(insecure.NewCredentials())，这代表的是不安全的连接
   然后实例化对应的服务，
# GRPC 流式传输
1. UnaryAPI ：普通rpc ： 一问一答模式，前面做的所有操作都是一问一答模式
2. ServerStreaming : 服务端推送流
3. ClientStreaming : 客户端推送流
4. BidirectionalStreaming : 双向推送流

# 思考：
1. 因为如果在传输大文件的场景下，需要长时间的传输，如果使用一问一答模式，需要多次建立连接，以保证文件传入成功，这也导致了大量的时间浪费在了建立连接上
所以我们用Stream流进行传输，他可以源源不断的推送数据，很适合大数据传输，或者服务端和客户端长时间数据交互的场景。Stream API 和 Unary API 相比，因为省掉了中间每次建立连接的花费，所以效率上会提升一些
2. 鉴权问题
3. grpc数据传递，类似http header
4. 拦截器
5. 客户端负载均衡(如果服务端已经部署为负载均衡，那么无需客户端负债均衡)
6. 服务的健康检查
7. 数据传输的方式，(一元请求或者流失请求)
8. 服务之间的认证问题
9. 服务限流的问题(这个需要自己进行操作)，服务接口限流
10. 服务的熔断，通过判断发生错误的次数，对服务做降级
11. 日志追踪 

# 流失传输的 proto 语法

```proto
service Echo {
// UnaryAPI
rpc UnaryEcho(EchoRequest) returns (EchoResponse) {}
// SServerStreaming
rpc ServerStreamingEcho(EchoRequest) returns (stream EchoResponse) {}
// ClientStreamingE
rpc ClientStreamingEcho(stream EchoRequest) returns (EchoResponse) {}
// BidirectionalStreaming
rpc BidirectionalStreamingEcho(stream EchoRequest) returns (stream EchoResponse) {}
}
```


