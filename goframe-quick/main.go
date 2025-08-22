package main

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// HelloReq 【请求】的数据结构来接收客户端提交的参数信息
type HelloReq struct {
	// 【规范路由】注册路由
	// 接口文档生成: 标签都是【OpenAPIv3】标准接口协议的规范字段。tags: 该接口属于哪个分类，或者接口模块。 summary: 接口描述。
	g.Meta `path:"/" method:"get" tags:"Test" summary:"接口描述-Hello world test case"`
	// v：表示校验规则，为valid的缩写，用于自动校验该参数。
	// dc：表示参数描述信息，为description的缩写，用于描述该参数的含义。
	// TODO: 这个案例里并没有把验证失败响应出去
	Name string `v:"required" dc:"姓名"`
	Age  int    `v:"required" dc:"年龄"`
}

// HelloRes 【响应】
type HelloRes struct {
	Content string `json:"content" dc:"返回结果"`
}

// Response【统一响应数据结构体】
type Response struct {
	Message string      `json:"message" dc:"消息提示"`
	Data    interface{} `json:"data"    dc:"执行结果"`
}

// Hello 【路由对象管理】
type Hello struct{}

func (Hello) Say(ctx context.Context, req *HelloReq) (res *HelloRes, err error) {
	// g.RequestFromCtx 方法从 ctx 获取原始的 *ghttp.Request 请求对象
	//r := g.RequestFromCtx(ctx)
	//r.Response.Writef(
	//	"【路由对象管理】Hello %s! Your Age is %d",
	//	req.Name,
	//	req.Age,
	//)

	// 通过 HelloRes 返回数据结构返回执行结果，而不是 r.Response.Write 来响应
	res = &HelloRes{
		Content: fmt.Sprintf(
			"Hello %s! Your Age is %d",
			req.Name,
			req.Age,
		),
	}

	return
}

func ResponseMiddleware(r *ghttp.Request) {
	r.Middleware.Next()

	var (
		msg = "OK"
		err = r.GetError()
	)

	if err != nil {
		// r.GetError()获取路由函数的执行状态，即路由函数返回的【第二个结果参数】error
		msg = err.Error()
	}

	// 【统一处理响应】
	// r.Response.WriteJson 将结果整合到统一的返回数据结构 Response，并编码为json格式返回给调用端
	r.Response.WriteJson(Response{
		Message: msg,
		// r.GetHandlerResponse() 方法获取路由函数的执行结果，即路由函数返回的【第一个结果参数】 *HelloRes
		Data: r.GetHandlerResponse(),
	})
}

// ErrorHandler 用来作为中间件的处理函数
func ErrorHandler(r *ghttp.Request) {
	// 【前置中间件】处理逻辑
	r.Middleware.Next()
	// 【后置中间件】处理逻辑
	// 在该中间件中我们先通过 r.Middleware.Next() 执行路由函数流程， 随后通过 r.GetError() 获取路由函数是否有错误产生
	if err := r.GetError(); err != nil {
		r.Response.Write("error occurs: ", err.Error())
		return
	}
}

func main() {
	fmt.Println("Hello GoFrame:", gf.VERSION)

	// g.Server() 方法获得一个默认的 Server 对象，该方法采用【单例模式】设计
	s := g.Server()

	s.BindHandler("/", func(r *ghttp.Request) {
		var req HelloReq
		// r.Parse 方法将请求参数映射到请求对象
		if err := r.Parse(&req); err != nil {
			r.Response.Writefln(err.Error())
			return
		}

		r.Response.Writef(
			"Hello %s! Your Age is %d",
			req.Name,
			req.Age,
		)
	})

	// s.Group 的分组路由方式定义一组路由注册
	s.Group("/hello", func(group *ghttp.RouterGroup) {
		// 错误处理的【后置中间件】ErrorHandler -- 【请求结构体表单验证失败会在这里被拦截!】
		// group.Middleware(ErrorHandler)
		group.Middleware(ResponseMiddleware)
		// group.Bind 方法注册路由对象，该方法将会【遍历路由对象的所有公开方法】，读取方法的输入输出结构体定义，并对其执行【路由注册】
		group.Bind(
			new(Hello),
		)
	})

	// 启用OpenAPIv3的接口文档生成，并指定生成的文件路径 /api.json
	// OpenAPIv3是目前业内的接口文档标准协议，用于接口文档的定义，通常使用json格式生成。
	// 该接口文档 json 文件可以用许多接口管理工具打开，例如【Swagger UI/PostMan/APIFox】等等。
	s.SetOpenApiPath("/api.json")
	// 启用内置的Swagger接口文档UI，并指定客访问的UI地址为 /swagger。内置的Swagger UI可自定义修改，具体可参考开发手册相应章节。
	s.SetSwaggerPath("/swagger")

	s.SetPort(8000)
	s.Run()
}
