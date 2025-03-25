package user

import (
	"context"
	"fmt"
	"sleep-service/api/v1/admin/user"
)

var Admin = cAdmin{}

type cAdmin struct{}

func (a *cAdmin) GetAdminInfo(ctx context.Context, req *user.GetAdminInfoReq) (res *user.GetAdminInfoRes, err error) {
	fmt.Println("info")
	return
}
