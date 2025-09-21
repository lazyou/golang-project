package test

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"os"
	"testing"
)

func Test_gcmdInit(t *testing.T) {
	gcmd.Init("gf", "run", "../main.go", "arg", "--gf.debug=true", "-opt")
	t.Logf("GetArgAll=%v", gcmd.GetArgAll())
	t.Logf("GetOptAll=%v", gcmd.GetOptAll())
}

func Test_gcmdParse(t *testing.T) {
	os.Args = []string{"gf", "build", "main.go", "-o=gf.exe", "-y"}
	p, err := gcmd.Parse(g.MapStrBool{
		"o,output": true,
		"y,yes":    false,
	})
	if err != nil {
		panic(err)
	}

	t.Logf("%v", p.GetOpt("o"))
	t.Logf("%v", p.GetOpt("output"))
	t.Logf("%v", p.GetOpt("y") != nil)
	t.Logf("%v", p.GetOpt("yes") != nil)
	t.Logf("%v", p.GetOpt("none") != nil)
}

func Test_gcmdCommand(t *testing.T) {
	// 见文档 https://goframe.org.cn/docs/core/gcmd-command
}

func Test_gcmdScan(t *testing.T) {
	// 【重要】单元测试 里无法触发终端输入数据的交互！
	name := gcmd.Scan("input your name?\n")
	t.Logf("name = %v", name)
}
