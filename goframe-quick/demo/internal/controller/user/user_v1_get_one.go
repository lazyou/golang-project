package user

import (
	"context"
	"demo/internal/dao"

	"demo/api/user/v1"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	// res 【命名返回值】初始化是 nil, &res.User也是nil. 所以这里才需要【字面量再次初始化】!
	res = &v1.GetOneRes{}
	err = dao.User.Ctx(ctx).WherePri(req.Id).Scan(&res.User)
	return
}
