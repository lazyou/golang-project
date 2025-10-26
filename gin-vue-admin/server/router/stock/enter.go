package stock

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	StockCategoryRouter
	StockItemRouter
}

var (
	stockCategoryApi = api.ApiGroupApp.StockApiGroup.StockCategoryApi
	stockItemApi     = api.ApiGroupApp.StockApiGroup.StockItemApi
)
