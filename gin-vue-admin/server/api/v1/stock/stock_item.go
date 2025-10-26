package stock

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/stock"
    stockReq "github.com/flipped-aurora/gin-vue-admin/server/model/stock/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type StockItemApi struct {}



// CreateStockItem 创建股票
// @Tags StockItem
// @Summary 创建股票
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body stock.StockItem true "创建股票"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /stockItem/createStockItem [post]
func (stockItemApi *StockItemApi) CreateStockItem(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var stockItem stock.StockItem
	err := c.ShouldBindJSON(&stockItem)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = stockItemService.CreateStockItem(ctx,&stockItem)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteStockItem 删除股票
// @Tags StockItem
// @Summary 删除股票
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body stock.StockItem true "删除股票"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /stockItem/deleteStockItem [delete]
func (stockItemApi *StockItemApi) DeleteStockItem(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := stockItemService.DeleteStockItem(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteStockItemByIds 批量删除股票
// @Tags StockItem
// @Summary 批量删除股票
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /stockItem/deleteStockItemByIds [delete]
func (stockItemApi *StockItemApi) DeleteStockItemByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := stockItemService.DeleteStockItemByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateStockItem 更新股票
// @Tags StockItem
// @Summary 更新股票
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body stock.StockItem true "更新股票"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /stockItem/updateStockItem [put]
func (stockItemApi *StockItemApi) UpdateStockItem(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var stockItem stock.StockItem
	err := c.ShouldBindJSON(&stockItem)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = stockItemService.UpdateStockItem(ctx,stockItem)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindStockItem 用id查询股票
// @Tags StockItem
// @Summary 用id查询股票
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询股票"
// @Success 200 {object} response.Response{data=stock.StockItem,msg=string} "查询成功"
// @Router /stockItem/findStockItem [get]
func (stockItemApi *StockItemApi) FindStockItem(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	restockItem, err := stockItemService.GetStockItem(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(restockItem, c)
}
// GetStockItemList 分页获取股票列表
// @Tags StockItem
// @Summary 分页获取股票列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query stockReq.StockItemSearch true "分页获取股票列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /stockItem/getStockItemList [get]
func (stockItemApi *StockItemApi) GetStockItemList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo stockReq.StockItemSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := stockItemService.GetStockItemInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}
// GetStockItemDataSource 获取StockItem的数据源
// @Tags StockItem
// @Summary 获取StockItem的数据源
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /stockItem/getStockItemDataSource [get]
func (stockItemApi *StockItemApi) GetStockItemDataSource(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口为获取数据源定义的数据
    dataSource, err := stockItemService.GetStockItemDataSource(ctx)
    if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
   		response.FailWithMessage("查询失败:" + err.Error(), c)
   		return
    }
   response.OkWithData(dataSource, c)
}

// GetStockItemPublic 不需要鉴权的股票接口
// @Tags StockItem
// @Summary 不需要鉴权的股票接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /stockItem/getStockItemPublic [get]
func (stockItemApi *StockItemApi) GetStockItemPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    stockItemService.GetStockItemPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的股票接口信息",
    }, "获取成功", c)
}
