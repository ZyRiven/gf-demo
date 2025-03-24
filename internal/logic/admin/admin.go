package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"sleep-service/api/v1/admin/user"
	"sleep-service/internal/dao"
	"sleep-service/internal/model/entity"
	"sleep-service/internal/service"
	"sleep-service/utility/libUtils"
	"sleep-service/utility/liberr"
)

type sAdmin struct{}

func init() {
	service.RegisterAdmin(New())
}

func New() *sAdmin {
	return &sAdmin{}
}

func (s *sAdmin) GetAdminByUsernamePass(ctx context.Context, req *user.LoginReq) (admin *entity.AdminAccount, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.AdminAccount.Ctx(ctx).Where(dao.AdminAccount.Columns().Username, req.UserName).Scan(&admin)
		liberr.ErrIsNil(ctx, err)
		liberr.ValueIsNil(admin, "暂无此人")
		//验证密码
		if libUtils.EncryptPassword(req.Password, admin.Salt) != admin.Password {
			liberr.ErrIsNil(ctx, gerror.New("账号或密码错误"))
		}
		//账号状态
		if admin.Status == 0 {
			liberr.ErrIsNil(ctx, gerror.New("账号已被冻结"))
		}
	})
	return
}
