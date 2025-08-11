// Package main implements a client for Greeter service.
package main

import (
	"context"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"test/protobuf_gen"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// TLS连接
	credential, err := credentials.NewClientTLSFromFile("../keys/server.pem", "")
	if err != nil {
		grpclog.Fatalf("Failed to create TLS credentials %v", err)
	}

	// Set up a connection to the server.
	// 创建一个到远程 gRPC 服务器的客户端连接（Connection
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(credential))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建一个“客户端存根（Client Stub）”
	helloClient := protobuf_gen.NewHelloClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// context 进行超时控制: 防止请求无限挂起
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 客户端调用 【一元RPC 方法】
	result, err := helloClient.SayHello(ctx, &protobuf_gen.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("SayHello err: %v", err)
	}

	log.Printf("SayHello: %s", result.Message)

	// 客户端调用 【服务端流RPC 方法】
	stream, err := helloClient.SayHelloServerStream(context.Background(), &protobuf_gen.HelloRequest{Name: name})
	if err != nil {
		log.Printf("ServerStream: %s", result.Message)
	}

	for {
		streamResult, err := stream.Recv()
		if err == io.EOF {
			log.Printf("ServerStream EOF \n")
			break
		}

		if err != nil { // 服务端 ctrl+ c 中断会触发
			log.Fatalf("ServerStream err: %v\n", err)
		}

		log.Printf("ServerStream: %s", streamResult)
	}
}
