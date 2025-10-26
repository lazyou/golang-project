
package stock

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/stock"
    stockReq "github.com/flipped-aurora/gin-vue-admin/server/model/stock/request"
)

type StockCategoryService struct {}
// CreateStockCategory 创建股票分类记录
// Author [yourname](https://github.com/yourname)
func (stockCategoryService *StockCategoryService) CreateStockCategory(ctx context.Context, stockCategory *stock.StockCategory) (err error) {
	err = global.GVA_DB.Create(stockCategory).Error
	return err
}

// DeleteStockCategory 删除股票分类记录
// Author [yourname](https://github.com/yourname)
func (stockCategoryService *StockCategoryService)DeleteStockCategory(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&stock.StockCategory{},"id = ?",ID).Error
	return err
}

// DeleteStockCategoryByIds 批量删除股票分类记录
// Author [yourname](https://github.com/yourname)
func (stockCategoryService *StockCategoryService)DeleteStockCategoryByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]stock.StockCategory{},"id in ?",IDs).Error
	return err
}

// UpdateStockCategory 更新股票分类记录
// Author [yourname](https://github.com/yourname)
func (stockCategoryService *StockCategoryService)UpdateStockCategory(ctx context.Context, stockCategory stock.StockCategory) (err error) {
	err = global.GVA_DB.Model(&stock.StockCategory{}).Where("id = ?",stockCategory.ID).Updates(&stockCategory).Error
	return err
}

// GetStockCategory 根据ID获取股票分类记录
// Author [yourname](https://github.com/yourname)
func (stockCategoryService *StockCategoryService)GetStockCategory(ctx context.Context, ID string) (stockCategory stock.StockCategory, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&stockCategory).Error
	return
}
// GetStockCategoryInfoList 分页获取股票分类记录
// Author [yourname](https://github.com/yourname)
func (stockCategoryService *StockCategoryService)GetStockCategoryInfoList(ctx context.Context, info stockReq.StockCategorySearch) (list []stock.StockCategory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&stock.StockCategory{})
    var stockCategorys []stock.StockCategory
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Name != nil && *info.Name != "" {
        db = db.Where("name LIKE ?", "%"+ *info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&stockCategorys).Error
	return  stockCategorys, total, err
}
func (stockCategoryService *StockCategoryService)GetStockCategoryPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
