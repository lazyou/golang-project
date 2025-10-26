// 自动生成模板StockItem
package stock

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 股票 结构体  StockItem
type StockItem struct {
	global.GVA_MODEL
	StockCategoryId *int32     `json:"stockCategoryId" form:"stockCategoryId" gorm:"comment:关联分类;column:stock_category_id;" binding:"required"` //关联分类
	Sort            *int32     `json:"sort" form:"sort" gorm:"comment:排序(越小越靠前);column:sort;"`                                                  //排序(越小越靠前)
	Name            *string    `json:"name" form:"name" gorm:"comment:股票名字;column:name;size:191;" binding:"required"`                           //股票名字
	Code            *string    `json:"code" form:"code" gorm:"uniqueIndex;comment:股票代码;column:code;size:50;" binding:"required"`                //股票代码
	CompanyName     *string    `json:"companyName" form:"companyName" gorm:"comment:公司名称;column:company_name;size:191;"`                        //公司名称
	CompanyProfile  *string    `json:"companyProfile" form:"companyProfile" gorm:"comment:公司简介;column:company_profile;"`                        //公司简介
	Industry        *string    `json:"industry" form:"industry" gorm:"comment:所属行业;column:industry;size:191;"`                                  //所属行业
	ListingDate     *time.Time `json:"listingDate" form:"listingDate" gorm:"comment:上市日期;column:listing_date;"`                                 //上市日期
	Remark          *string    `json:"remark" form:"remark" gorm:"comment:备注;column:remark;size:191;type:text;"`                                //备注
}

// TableName 股票 StockItem自定义表名 stock
func (StockItem) TableName() string {
	return "stock"
}
