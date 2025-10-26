package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/stock"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(stock.StockCategory{}, stock.StockItem{})
	if err != nil {
		return err
	}
	return nil
}
