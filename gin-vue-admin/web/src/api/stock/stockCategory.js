import service from '@/utils/request'
// @Tags StockCategory
// @Summary 创建股票分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.StockCategory true "创建股票分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /stockCategory/createStockCategory [post]
export const createStockCategory = (data) => {
  return service({
    url: '/stockCategory/createStockCategory',
    method: 'post',
    data
  })
}

// @Tags StockCategory
// @Summary 删除股票分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.StockCategory true "删除股票分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stockCategory/deleteStockCategory [delete]
export const deleteStockCategory = (params) => {
  return service({
    url: '/stockCategory/deleteStockCategory',
    method: 'delete',
    params
  })
}

// @Tags StockCategory
// @Summary 批量删除股票分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除股票分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stockCategory/deleteStockCategory [delete]
export const deleteStockCategoryByIds = (params) => {
  return service({
    url: '/stockCategory/deleteStockCategoryByIds',
    method: 'delete',
    params
  })
}

// @Tags StockCategory
// @Summary 更新股票分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.StockCategory true "更新股票分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /stockCategory/updateStockCategory [put]
export const updateStockCategory = (data) => {
  return service({
    url: '/stockCategory/updateStockCategory',
    method: 'put',
    data
  })
}

// @Tags StockCategory
// @Summary 用id查询股票分类
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.StockCategory true "用id查询股票分类"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /stockCategory/findStockCategory [get]
export const findStockCategory = (params) => {
  return service({
    url: '/stockCategory/findStockCategory',
    method: 'get',
    params
  })
}

// @Tags StockCategory
// @Summary 分页获取股票分类列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取股票分类列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stockCategory/getStockCategoryList [get]
export const getStockCategoryList = (params) => {
  return service({
    url: '/stockCategory/getStockCategoryList',
    method: 'get',
    params
  })
}

// @Tags StockCategory
// @Summary 不需要鉴权的股票分类接口
// @Accept application/json
// @Produce application/json
// @Param data query stockReq.StockCategorySearch true "分页获取股票分类列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /stockCategory/getStockCategoryPublic [get]
export const getStockCategoryPublic = () => {
  return service({
    url: '/stockCategory/getStockCategoryPublic',
    method: 'get',
  })
}
