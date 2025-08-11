// hello 服务实现
package services

import (
	"context"
	"log"
	"test/protobuf_gen"
	"test/util"
	"time"
)

// HelloServer 服务的实现
type HelloServer struct {
	protobuf_gen.HelloServer // PS: struct 包含(内嵌) interface 之后，并不需要实现 interface 的接口，也能成为 interface 接口类
}

func (s *HelloServer) SayHello(ctx context.Context, request *protobuf_gen.HelloRequest) (*protobuf_gen.HelloResponse, error) {
	log.Printf("Received-SayHello %s: %s \n", util.ClientIP(ctx), request.Name)

	return &protobuf_gen.HelloResponse{
		Result: &protobuf_gen.BaseResponse{
			IsOk:    true,
			Code:    200,
			Message: "",
		},
		Message: "Hello-SayHello! " + request.Name,
	}, nil
}

func (s *HelloServer) SayHelloServerStream(request *protobuf_gen.HelloRequest, stream protobuf_gen.Hello_SayHelloServerStreamServer) error {
	log.Printf("Received-SayHelloServerStream %s: %s \n", util.ClientIP(stream.Context()), request.Name)

	// 每秒
	for {
		time.Sleep(time.Second)

		err := stream.Send(&protobuf_gen.HelloResponse{
			Result: &protobuf_gen.BaseResponse{
				IsOk:    true,
				Code:    200,
				Message: "",
			},
			Message: "Hello-SayHelloServerStream! " + util.NowTime(),
		})

		if err != nil {
			return err
		}
	}

	return nil
}
