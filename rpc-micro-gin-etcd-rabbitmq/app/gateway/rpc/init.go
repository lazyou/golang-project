package rpc

import (
	"go-micro.dev/v4"

	"github.com/CocaineCong/micro-todoList/app/gateway/wrappers"
	"github.com/CocaineCong/micro-todoList/idl/pb"
)

var (
	UserService pb.UserService
	TaskService pb.TaskService
)

func InitRPC() {
	// 用户服务
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		// wrappers 包设置了【服务熔断】的配置
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	// 用户服务调用实例
	userService := pb.NewUserService("rpcUserService", userMicroService.Client())

	// 任务服务
	taskMicroService := micro.NewService(
		micro.Name("taskService.client"),
		micro.WrapClient(wrappers.NewTaskWrapper),
	)
	// 任务服务调用实例
	taskService := pb.NewTaskService("rpcTaskService", taskMicroService.Client())

	UserService = userService
	TaskService = taskService
}
