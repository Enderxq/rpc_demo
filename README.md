#gRpc是Google公司开发的一个高性能、开源和通用的Rpc框架，面向移动和HTTP/2设计。

##gRpc的使用流程

    1.定义标准的proto文件
    2.生成标准代码
    3.服务端使用生成的代码提供服务
    4.客户端使用生成的代码调用服务

##安装:

###1.安装gRpc

使用go命令下载：
```
go get -u google.golang.org/grpc

```
使用git下载：
```
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
cd $GOPATH/src/
go install google.golang.org/grpc
```
###2.安装protoc-gen-go
protoc依赖该工具生成代码
``` 
go get -u github.com/golang/protobuf/protoc-gen-go
```
    

##gRPC案例

1.先使用protobuf定义服务:创建message.proto文件.

2.在执行命令生成message.pb.go文件
```  
    protoc -I ./myProto --go_out=plugins=grpc:./myProto ./myProto/message.proto
命令解释：
    protoDir="./myProto/message.proto"
    outDir="./myProto"
    protoc -I ${protoDir}/ --go_out=plugins=grpc:${outDir} ${protoDir}/*proto
```
   
protoc工具参数解释：
    -I： 指定import路径，可以指定多个-I参数，编译时按顺序查找，不指定默认当前目录
    -go_out：指定go语言的访问类
    plugins：指定依赖的插件

3.编写服务端srv/main.go,导入"rpc_demo/myProto"包，提供服务

4.编写客户端cli/main.go,导入"rpc_demo/myProto"包，调用服务


