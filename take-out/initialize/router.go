package initialize

import (
	"github.com/Meng-Xin/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"take-out/global"
	"take-out/internal/router"
	"time"
)

func routerInit() *gin.Engine {
	r := gin.Default()
	allRouter := router.AllRouter

	// 配置 CORS - 2025年8月13日 为了 client 能访问添加
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8000"}, // 允许前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // 如果需要携带 cookie
		MaxAge:           12 * time.Hour,
	}))

	// 链路追踪日志中间件
	r.Use(logger.GinMiddleware(global.Log, "takeout"))

	// admin
	admin := r.Group("/admin")
	{
		allRouter.EmployeeRouter.InitApiRouter(admin)
		allRouter.CategoryRouter.InitApiRouter(admin)
		allRouter.DishRouter.InitApiRouter(admin)
		allRouter.CommonRouter.InitApiRouter(admin)
		allRouter.SetMealRouter.InitApiRouter(admin)
	}
	return r
}
