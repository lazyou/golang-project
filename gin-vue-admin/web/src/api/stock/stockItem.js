import service from '@/utils/request'
// @Tags StockItem
// @Summary 创建股票
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.StockItem true "创建股票"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /stockItem/createStockItem [post]
export const createStockItem = (data) => {
  return service({
    url: '/stockItem/createStockItem',
    method: 'post',
    data
  })
}

// @Tags StockItem
// @Summary 删除股票
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.StockItem true "删除股票"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stockItem/deleteStockItem [delete]
export const deleteStockItem = (params) => {
  return service({
    url: '/stockItem/deleteStockItem',
    method: 'delete',
    params
  })
}

// @Tags StockItem
// @Summary 批量删除股票
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除股票"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stockItem/deleteStockItem [delete]
export const deleteStockItemByIds = (params) => {
  return service({
    url: '/stockItem/deleteStockItemByIds',
    method: 'delete',
    params
  })
}

// @Tags StockItem
// @Summary 更新股票
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.StockItem true "更新股票"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /stockItem/updateStockItem [put]
export const updateStockItem = (data) => {
  return service({
    url: '/stockItem/updateStockItem',
    method: 'put',
    data
  })
}

// @Tags StockItem
// @Summary 用id查询股票
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.StockItem true "用id查询股票"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /stockItem/findStockItem [get]
export const findStockItem = (params) => {
  return service({
    url: '/stockItem/findStockItem',
    method: 'get',
    params
  })
}

// @Tags StockItem
// @Summary 分页获取股票列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取股票列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stockItem/getStockItemList [get]
export const getStockItemList = (params) => {
  return service({
    url: '/stockItem/getStockItemList',
    method: 'get',
    params
  })
}
// @Tags StockItem
// @Summary 获取数据源
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /stockItem/findStockItemDataSource [get]
export const getStockItemDataSource = () => {
  return service({
    url: '/stockItem/getStockItemDataSource',
    method: 'get',
  })
}

// @Tags StockItem
// @Summary 不需要鉴权的股票接口
// @Accept application/json
// @Produce application/json
// @Param data query stockReq.StockItemSearch true "分页获取股票列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /stockItem/getStockItemPublic [get]
export const getStockItemPublic = () => {
  return service({
    url: '/stockItem/getStockItemPublic',
    method: 'get',
  })
}
