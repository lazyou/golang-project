
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type StockItemSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      StockCategoryId  *int `json:"stockCategoryId" form:"stockCategoryId"` 
      Name  *string `json:"name" form:"name"` 
      Code  *string `json:"code" form:"code"` 
      CompanyName  *string `json:"companyName" form:"companyName"` 
    request.PageInfo
}
