package test

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func Test_gcfg(t *testing.T) {
	ctx := gctx.New()
	// 单例对象, 按名字单例
	t.Log(gcfg.Instance().Get(ctx, "database"))
	t.Log(g.Cfg().Get(ctx, "database.default.link"))
	t.Log(g.Cfg().MustGet(ctx, "viewpath"))

	gcfg.NewAdapterFile("file01", "file02")
}
