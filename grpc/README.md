# proto 语法
1. service对应的就是go里面的接口，可以作为服务器，客户端
2. rpc 对应的就是结构体中的方法
3. message 对应的也是结构体

# grpc 应用流程
1. 先编写proto文件，然后在protoc --go_out=plugins=grpc:.  manyService.proto 转化为go文件
2. sever:
   先监听端口，然后实例化grpc.NewServer()，在实例化结构体，结构体需要被方法继承，然后在进行注册服务，传入实例化的grpc.server,和结构体
   然后在调用服务的server，把listen传进去
3. client:
   先建立连接通过grpc.dial,传入地址和grpc.WithTransportCredentials(insecure.NewCredentials())，这代表的是不安全的连接
   然后实例化对应的服务，