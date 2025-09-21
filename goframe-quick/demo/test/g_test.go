package test

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
	"time"
)

// 核心组件-对象管理: https://goframe.org.cn/docs/core/g

// 数据类型
func Test_gMap(t *testing.T) {
	t.Log(g.Map{
		"key": "val",
	})

	t.Log(g.MapStrInt{
		"key": 99,
	})
}

// HTTP 客户端对象 - 【非单例】
func Test_gClient(t *testing.T) {
	// 1 秒超时 context! 【重要】
	// 在 Get() 的 err 拿到错误 'failed: Get "http://www.google.com": context deadline exceeded'
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// 其它 Context 则毫无用处, 一直等待请求!!

	url := "http://www.google.com"
	resp, err := g.Client().Get(ctx, url)
	if err != nil {
		t.Errorf("client get err=%v", err)
		return
	}

	t.Logf("resp StatusCode=%v", resp.StatusCode)
}

// 日志管理对象 - 【单例】 - 按 Log("name") 传入的 name 单例
func Test_gLog(t *testing.T) {
	timeFormat := "2006-01-02 15:04:05"
	g.Log().SetTimeFormat(timeFormat)
	g.Log().Info(context.Background(), "info-绿色-无调用栈")
	g.Log().Error(context.Background(), "error-红色-有调用栈-Stack!")
}

// WEB Server - 【单例】 - 按 Server("name") 传入的 name 单例
func Test_gServer(t *testing.T) {
	g.Server().Run()
}

// TCP Server - 【单例】 - 按 TCPServer("name") 传入的 name 单例
func Test_gTcpServer(t *testing.T) {
	err := g.TCPServer().Run()
	if err != nil {
		t.Errorf("run err=%v", err)
		return
	}
}
