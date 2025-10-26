package stock

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StockCategoryRouter struct{}

// InitStockCategoryRouter 初始化 股票分类 路由信息
func (s *StockCategoryRouter) InitStockCategoryRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	stockCategoryRouter := Router.Group("stockCategory").Use(middleware.OperationRecord())
	stockCategoryRouterWithoutRecord := Router.Group("stockCategory")
	stockCategoryRouterWithoutAuth := PublicRouter.Group("stockCategory")
	{
		stockCategoryRouter.POST("createStockCategory", stockCategoryApi.CreateStockCategory)             // 新建股票分类
		stockCategoryRouter.DELETE("deleteStockCategory", stockCategoryApi.DeleteStockCategory)           // 删除股票分类
		stockCategoryRouter.DELETE("deleteStockCategoryByIds", stockCategoryApi.DeleteStockCategoryByIds) // 批量删除股票分类
		stockCategoryRouter.PUT("updateStockCategory", stockCategoryApi.UpdateStockCategory)              // 更新股票分类
	}
	{
		stockCategoryRouterWithoutRecord.GET("findStockCategory", stockCategoryApi.FindStockCategory)       // 根据ID获取股票分类
		stockCategoryRouterWithoutRecord.GET("getStockCategoryList", stockCategoryApi.GetStockCategoryList) // 获取股票分类列表
	}
	{
		stockCategoryRouterWithoutAuth.GET("getStockCategoryPublic", stockCategoryApi.GetStockCategoryPublic) // 股票分类开放接口
	}
}
