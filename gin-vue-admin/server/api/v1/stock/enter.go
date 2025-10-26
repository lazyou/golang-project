package stock

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ StockCategoryApi }

var stockCategoryService = service.ServiceGroupApp.StockServiceGroup.StockCategoryService
