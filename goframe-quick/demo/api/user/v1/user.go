package v1

import (
	"demo/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

// CreateReq 创建接口
type CreateReq struct {
	// tags接口分组(用于生成接口文档)、summary接口描述
	g.Meta `path:"/user" method:"post" tags:"User" summary:"Create user"`
	// 请求表单验证. dc标签是description的缩写
	Name string `v:"required|length:3,10" dc:"user name"`
	Age  uint   `v:"required|between:18,200" dc:"user age"`
}
type CreateRes struct {
	Id int64 `json:"id" dc:"user id"`
}

// DeleteReq 删除接口
type DeleteReq struct {
	g.Meta `path:"/user/{id}" method:"delete" tags:"User" summary:"Delete user"`
	Id     int64 `v:"required" dc:"user id"`
}
type DeleteRes struct{}

// UpdateReq 更新接口
type UpdateReq struct {
	g.Meta `path:"/user/{id}" method:"put" tags:"User" summary:"Update user"`
	Id     int64   `v:"required" dc:"user id"`
	Name   *string `v:"length:3,10" dc:"user name"`
	Age    *uint   `v:"between:18,200" dc:"user age"`
	Status *Status `v:"in:0,1" dc:"user status"`
}
type UpdateRes struct{}

// Status 用户状态类型 -- TODO: 居然单独一个自定义类型!.
type Status int

const (
	StatusOK       Status = 0 // User is OK.
	StatusDisabled Status = 1 // User is disabled.
)

// GetOneReq 查询接口（单个）
type GetOneReq struct {
	g.Meta `path:"/user/{id}" method:"get" tags:"User" summary:"Get one user"`
	Id     int64 `v:"required" dc:"user id"`
}
type GetOneRes struct {
	*entity.User `dc:"user"`
}

// GetListReq 查询接口（列表）
type GetListReq struct {
	g.Meta `path:"/user" method:"get" tags:"User" summary:"Get users"`
	Age    *uint   `v:"between:18,200" dc:"user age"`
	Status *Status `v:"in:0,1" dc:"user status"`
}
type GetListRes struct {
	List []*entity.User `json:"list" dc:"user list"`
}
