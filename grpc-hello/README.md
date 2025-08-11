### 0. gRPC 和 
* gRPC: 这是一个高性能、开源的`远程过程调用` (RPC) 框架, 它支持多种编程语言, 并且默认使用 **Protocol Buffers 作为其接口定义语言 (IDL) 和消息交换格式**. 
  * https://grpc.io/docs/languages/go/quickstart/

* Protocol Buffers (protobuf): 这是一种语言中立、平台中立、可扩展的序列化数据格式, 用于`序列化结构化数据`. 开发者可以定义数据结构（称为消息类型）在一个 `.proto` 文件中, 然后**使用 `protobuf` 编译器生成对应编程语言的代码（`服务端和客户端`的代码）**, 以实现对这些数据结构的序列化和反序列化操作. 
  * https://github.com/protocolbuffers/protobuf
  * https://protobuf.dev/
  * https://protobuf.dev/installation/ 
  * 【proto 文件语法】https://protobuf.dev/programming-guides/proto3/
  * 【go 案例】https://protobuf.dev/getting-started/gotutorial/
  
* 【重要-大前提】安装 Protocol Buffers 编译器:
    ```shell
    # 下载 https://github.com/protocolbuffers/protobuf/releases/download/v31.1/protoc-31.1-win64.zip
    
    # 把解压得到的 protoc.exe 文件放到 C:\Users\lin\go\bin （不是非放这里, 能运行就行）
    ``` 

* 安装完 Protocol Buffers 编译器, 后面生成各个平台代码才有 `protoc` 程序使用


### 1. 自签证书（Self-signed Certificate）
* 生成一个`私钥（如 server.key）`, 然后使用该`私钥`创建一个`证书（如 server.pem）`

* 使用 OpenSSL 生成带有 SAN 的证书:
    * Go 1.15 版本开始, Go 的 crypto/x509 包加强了对证书的验证规则, 不再允许仅通过 Common Name 来匹配主机名, 必须使用 SAN（Subject Alternative Name）字段. 
    ```shell
    cd ./keys
    
    # server.conf
    # server.conf 内容见 ./keys/server.conf
    
    # 生成【私钥】
    openssl genrsa -out server.key 2048
    
    # 生成【证书】
    openssl req -x509 -new -nodes -key server.key -sha256 -days 3650 -out server.pem -config server.conf -extensions v3_req
    ```


### 2. Protocol Buffers 定义服务 和 消息数据(请求、响应)
* 见目录 `./protobuf` 下的文件

* 【定义消息】（数据结构）:
  * 公共响应数据结构: `common_messages.proto`
  * hello服务 的请求和响应: `hello_messages.proto`, 这里【嵌入】了 `common_messages.proto`.

* 【定义服务】hello 服务：一个 RPC 服务通过参数和返回类型来指定可以远程调用的方法;
  * hello 服务: `hello_services.proto`
  * 【一元RPC】(单个请求, 单个响应)和【服务端流RPC】(单个请求, 多个响应) 调用

* 上面 proto 文件小结:
  * `syntax` 定义 protobuf 版本
  * `package` 定义 protobuf 的包名，不是 Go 的
  * `option go_package="gen_dir;gen_package_name";` 编译后文件所在目录、以及包名
  * `import` 引用其它 protobuf 包(抽出公共部分)

* 客户端和服务端均使用服务定义生成的接口代码:
    ```shell
    # 大前提, 先安装 Protocol Buffers 编译器 -- `protoc`

    # protoc-gen-go: Google 提供, 根据 .proto 文件生成 Go 的【数据结构（结构体）和序列化/反序列化代码】
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    # protoc-gen-go-grpc: Google 提供, 根据 .proto 文件中定义的 service, 生成 【gRPC 的`客户端和服务端`的接口（interface）和 stub 代码】. 
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

    # 生成 go 的 gRPC 代码
    protoc --go_out=. .\common_messages.proto
    protoc --go_out=. .\hello_messages.proto
    # 【注意】服务的生成选项略有不同 
    protoc --go_out=. --go-grpc_out=. .\hello_services.proto
    ```


### 3. 实现服务 - 补充业务逻辑 
* 上面直到生成go的代码, 仅仅是生成 **请求、响应数据的结构体** 和 **服务的接口**

* 服务的接口实现 `services/hello_service.go`
  * 实现接口的方法，补充上具体的业务逻辑! 

  * 接收数据 --> 业务处理 --> 响应数据.


### 4. 写一个服务器(服务端)
* `main/server.go`:
  * 监听本地网络(tcp)、端口.

  * 设置 TLS 认证证书.

  * `NewServer` 创建了一个未注册服务 且 尚未开始接受请求的gRPC服务器.

  * 注册 `hello_service.go` 服务.

  * 开启gRPC服务器.


### 5. 写一个客户端
* `main/client.go`:
  * TLS证书.

  * 创建一个到远程 gRPC 服务器的客户端连接(`conn := grpc.NewClient`).

  * 创建一个"客户端存根(Client Stub)" (`helloClient := protobuf_gen.NewHelloClient(conn)`).

  * `context` 进行超时控制: 防止请求无限挂起(`context.WithTimeout`).

  * 调用rpc服务01: 客户端调用 【一元RPC 方法】

  * 调用rpc服务02: 客户端调用 【服务端流RPC 方法】


### 6. 调用
```shell
cd main
go run .\server.go
go run .\client.go
```