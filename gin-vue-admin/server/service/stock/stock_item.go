
package stock

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/stock"
    stockReq "github.com/flipped-aurora/gin-vue-admin/server/model/stock/request"
)

type StockItemService struct {}
// CreateStockItem 创建股票记录
// Author [yourname](https://github.com/yourname)
func (stockItemService *StockItemService) CreateStockItem(ctx context.Context, stockItem *stock.StockItem) (err error) {
	err = global.GVA_DB.Create(stockItem).Error
	return err
}

// DeleteStockItem 删除股票记录
// Author [yourname](https://github.com/yourname)
func (stockItemService *StockItemService)DeleteStockItem(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&stock.StockItem{},"id = ?",ID).Error
	return err
}

// DeleteStockItemByIds 批量删除股票记录
// Author [yourname](https://github.com/yourname)
func (stockItemService *StockItemService)DeleteStockItemByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]stock.StockItem{},"id in ?",IDs).Error
	return err
}

// UpdateStockItem 更新股票记录
// Author [yourname](https://github.com/yourname)
func (stockItemService *StockItemService)UpdateStockItem(ctx context.Context, stockItem stock.StockItem) (err error) {
	err = global.GVA_DB.Model(&stock.StockItem{}).Where("id = ?",stockItem.ID).Updates(&stockItem).Error
	return err
}

// GetStockItem 根据ID获取股票记录
// Author [yourname](https://github.com/yourname)
func (stockItemService *StockItemService)GetStockItem(ctx context.Context, ID string) (stockItem stock.StockItem, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&stockItem).Error
	return
}
// GetStockItemInfoList 分页获取股票记录
// Author [yourname](https://github.com/yourname)
func (stockItemService *StockItemService)GetStockItemInfoList(ctx context.Context, info stockReq.StockItemSearch) (list []stock.StockItem, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&stock.StockItem{})
    var stockItems []stock.StockItem
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.StockCategoryId != nil {
        db = db.Where("stock_category_id = ?", *info.StockCategoryId)
    }
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
    if info.Code != nil && *info.Code != "" {
        db = db.Where("code = ?", *info.Code)
    }
    if info.CompanyName != nil && *info.CompanyName != "" {
        db = db.Where("company_name LIKE ?", "%"+ *info.CompanyName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&stockItems).Error
	return  stockItems, total, err
}
func (stockItemService *StockItemService)GetStockItemDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	
	   stockCategoryId := make([]map[string]any, 0)
	   
       
       global.GVA_DB.Table("stock_category").Where("deleted_at IS NULL").Select("name as label,id as value").Scan(&stockCategoryId)
	   res["stockCategoryId"] = stockCategoryId
	return
}
func (stockItemService *StockItemService)GetStockItemPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
