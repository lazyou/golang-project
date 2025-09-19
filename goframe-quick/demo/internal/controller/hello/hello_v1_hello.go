package hello

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"demo/api/hello/v1"
)

// 路由对象的具体实现: TODO: 为什么独立文件, 而不放 hello_new.go.

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
