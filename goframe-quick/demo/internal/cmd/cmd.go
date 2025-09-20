package cmd

import (
	"context"
	"demo/internal/controller/user"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"demo/internal/controller/hello"
)

var (
	// 命令管理: https://goframe.org.cn/docs/core/gcmd

	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 路由管理、注册 和 中间件: https://goframe.org.cn/docs/web/router
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
					user.NewV1(),
				)
			})
			// 阻塞运行接收客户端请求, 并监听进程信号(
			// 优雅关闭: 停止接收新请求、处理完旧请求再关闭, 为防止处理旧请求太久加上超时上下文 Context
			s.Run()
			return nil
		},
	}
)
