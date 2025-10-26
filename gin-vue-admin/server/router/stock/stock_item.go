package stock

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StockItemRouter struct {}

// InitStockItemRouter 初始化 股票 路由信息
func (s *StockItemRouter) InitStockItemRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	stockItemRouter := Router.Group("stockItem").Use(middleware.OperationRecord())
	stockItemRouterWithoutRecord := Router.Group("stockItem")
	stockItemRouterWithoutAuth := PublicRouter.Group("stockItem")
	{
		stockItemRouter.POST("createStockItem", stockItemApi.CreateStockItem)   // 新建股票
		stockItemRouter.DELETE("deleteStockItem", stockItemApi.DeleteStockItem) // 删除股票
		stockItemRouter.DELETE("deleteStockItemByIds", stockItemApi.DeleteStockItemByIds) // 批量删除股票
		stockItemRouter.PUT("updateStockItem", stockItemApi.UpdateStockItem)    // 更新股票
	}
	{
		stockItemRouterWithoutRecord.GET("findStockItem", stockItemApi.FindStockItem)        // 根据ID获取股票
		stockItemRouterWithoutRecord.GET("getStockItemList", stockItemApi.GetStockItemList)  // 获取股票列表
	}
	{
	    stockItemRouterWithoutAuth.GET("getStockItemDataSource", stockItemApi.GetStockItemDataSource)  // 获取股票数据源
	    stockItemRouterWithoutAuth.GET("getStockItemPublic", stockItemApi.GetStockItemPublic)  // 股票开放接口
	}
}
