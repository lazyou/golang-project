package user

import (
	"context"
	"demo/internal/dao"

	"demo/api/user/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	_, err = dao.User.Ctx(ctx).WherePri(req.Id).Delete()

	return
}
