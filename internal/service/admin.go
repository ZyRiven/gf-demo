// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sleep-service/api/v1/admin/user"
	"sleep-service/internal/model/entity"
)

type (
	IAdmin interface {
		GetAdminByUsernamePass(ctx context.Context, req *user.LoginReq) (admin *entity.AdminAccount, err error)
	}
)

var (
	localAdmin IAdmin
)

func Admin() IAdmin {
	if localAdmin == nil {
		panic("implement not found for interface IAdmin, forgot register?")
	}
	return localAdmin
}

func RegisterAdmin(i IAdmin) {
	localAdmin = i
}
