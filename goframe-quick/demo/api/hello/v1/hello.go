package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 控制器的【请求与响应】
// 路由规范: https://goframe.org.cn/docs/web/router-registering-strict-router

type HelloReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
}
type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
