package stock

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	StockCategoryApi
	StockItemApi
}

var (
	stockCategoryService = service.ServiceGroupApp.StockServiceGroup.StockCategoryService
	stockItemService     = service.ServiceGroupApp.StockServiceGroup.StockItemService
)
