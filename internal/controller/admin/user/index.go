package user

import (
	"context"
	"fmt"
	"sleep-service/api/v1/admin/user"
	"sleep-service/internal/dao"
	"sleep-service/internal/model/entity"
	"sleep-service/utility"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

var Login = cLogin{}

type cLogin struct{}

// Login 登录
func (a *cLogin) Login(ctx context.Context, req *user.LoginReq) (res *user.LoginRes, err error) {
	var admin *entity.AdminAccount
	err = dao.AdminAccount.Ctx(ctx).Where(dao.AdminAccount.Columns().Nickname,req.UserName).Scan(&admin)
	if err != nil || admin == nil {
		err = gerror.New("账号不存在")
		return
	}
	pass := utility.Md5(req.Password + admin.Salt)

	if pass != admin.Password {
		err = gerror.New("密码错误")
		return
	}
	//token
	token := middleware.GenerateToken(&middleware.UserClaims{
		ID:             res["id"].(int64),
		Accountid:      res["accountID"].(int64),
		StandardClaims: jwt.StandardClaims{},
	})
	model.DB().Table("admin_account").Where("id", res["id"]).Data(map[string]interface{}{"loginstatus": 1, "lastLoginTime": time.Now().Unix(), "lastLoginIp": gf.GetIp(c)}).Update()
	//登录日志
	model.DB().Table("login_logs").
		Data(map[string]interface{}{"type": 1, "uid": res["id"], "out_in": "in",
			"createtime": time.Now().Unix(), "loginIP": gf.GetIp(c)}).Insert()
	return
}