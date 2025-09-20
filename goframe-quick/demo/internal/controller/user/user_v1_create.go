package user

import (
	"context"
	"demo/api/user/v1"
	"demo/internal/dao"
	"demo/internal/model/do"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	// 调用生成的 dao 操作数据, 调用生成的 do 结构体赋值数据(字段为any类型, 避免了其它类型零值可能不写入数据的问题)
	// do 结构体用在【创建数据】的情况!
	insertId, err := dao.User.Ctx(ctx).Data(do.User{
		Name:   req.Name,
		Status: v1.StatusOK,
		Age:    req.Age,
	}).InsertAndGetId()

	if err != nil {
		return nil, err
	}

	res = &v1.CreateRes{
		Id: insertId,
	}

	return
}
