package user

import (
	"context"
	"demo/internal/dao"
	"demo/internal/model/do"

	"demo/api/user/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res = &v1.GetListRes{}
	// do 结构体用在【查询条件】的情况!
	err = dao.User.Ctx(ctx).
		Where(do.User{
			Age:    req.Age,
			Status: req.Status,
		}).Scan(&res.List)

	return
}
