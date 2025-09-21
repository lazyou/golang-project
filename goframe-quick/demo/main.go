package main

import (
	_ "demo/internal/packed"
	// 引入数据库驱动
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"demo/internal/cmd"
)

func main() {
	// g.SetDebug(true) // 调试模式运行, 日志会相当的多!
	cmd.Main.Run(gctx.GetInitCtx())
}
