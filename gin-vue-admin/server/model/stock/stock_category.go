// 自动生成模板StockCategory
package stock

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 股票分类 结构体  StockCategory
type StockCategory struct {
	global.GVA_MODEL
	Name   *string `json:"name" form:"name" gorm:"uniqueIndex;comment:名字;column:name;size:191;"` //名字
	Sort   *int32  `json:"sort" form:"sort" gorm:"comment:排序(越小越靠前);column:sort;"`               //排序(越小越靠前)
	Remark *string `json:"remark" form:"remark" gorm:"comment:备注;column:remark;"`                //备注
}

// TableName 股票分类 StockCategory自定义表名 stock_category
func (StockCategory) TableName() string {
	return "stock_category"
}
