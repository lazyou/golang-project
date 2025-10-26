package stock

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/stock"
	stockReq "github.com/flipped-aurora/gin-vue-admin/server/model/stock/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StockCategoryApi struct{}

// CreateStockCategory 创建股票分类
// @Tags StockCategory
// @Summary 创建股票分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body stock.StockCategory true "创建股票分类"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /stockCategory/createStockCategory [post]
func (stockCategoryApi *StockCategoryApi) CreateStockCategory(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var stockCategory stock.StockCategory
	err := c.ShouldBindJSON(&stockCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = stockCategoryService.CreateStockCategory(ctx, &stockCategory)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteStockCategory 删除股票分类
// @Tags StockCategory
// @Summary 删除股票分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body stock.StockCategory true "删除股票分类"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /stockCategory/deleteStockCategory [delete]
func (stockCategoryApi *StockCategoryApi) DeleteStockCategory(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := stockCategoryService.DeleteStockCategory(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteStockCategoryByIds 批量删除股票分类
// @Tags StockCategory
// @Summary 批量删除股票分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /stockCategory/deleteStockCategoryByIds [delete]
func (stockCategoryApi *StockCategoryApi) DeleteStockCategoryByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := stockCategoryService.DeleteStockCategoryByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateStockCategory 更新股票分类
// @Tags StockCategory
// @Summary 更新股票分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body stock.StockCategory true "更新股票分类"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /stockCategory/updateStockCategory [put]
func (stockCategoryApi *StockCategoryApi) UpdateStockCategory(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var stockCategory stock.StockCategory
	err := c.ShouldBindJSON(&stockCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = stockCategoryService.UpdateStockCategory(ctx, stockCategory)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindStockCategory 用id查询股票分类
// @Tags StockCategory
// @Summary 用id查询股票分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询股票分类"
// @Success 200 {object} response.Response{data=stock.StockCategory,msg=string} "查询成功"
// @Router /stockCategory/findStockCategory [get]
func (stockCategoryApi *StockCategoryApi) FindStockCategory(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	restockCategory, err := stockCategoryService.GetStockCategory(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(restockCategory, c)
}

// GetStockCategoryList 分页获取股票分类列表
// @Tags StockCategory
// @Summary 分页获取股票分类列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query stockReq.StockCategorySearch true "分页获取股票分类列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /stockCategory/getStockCategoryList [get]
func (stockCategoryApi *StockCategoryApi) GetStockCategoryList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo stockReq.StockCategorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := stockCategoryService.GetStockCategoryInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetStockCategoryPublic 不需要鉴权的股票分类接口
// @Tags StockCategory
// @Summary 不需要鉴权的股票分类接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /stockCategory/getStockCategoryPublic [get]
func (stockCategoryApi *StockCategoryApi) GetStockCategoryPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	stockCategoryService.GetStockCategoryPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的股票分类接口信息",
	}, "获取成功", c)
}
