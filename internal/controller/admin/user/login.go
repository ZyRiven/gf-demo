package user

import (
	"context"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"sleep-service/api/v1/admin/user"
	"sleep-service/internal/consts"
	"sleep-service/internal/model"
	"sleep-service/internal/model/entity"
	"sleep-service/internal/service"
	"sleep-service/utility/libUtils"

	"github.com/gogf/gf/v2/util/gconv"
)

var Login = cLogin{}

type cLogin struct{}

// Login 登录
func (a *cLogin) Login(ctx context.Context, req *user.LoginReq) (res *user.LoginRes, err error) {
	var (
		admin *entity.AdminAccount
		token string
	)
	ip := libUtils.GetClientIp(ctx)
	userAgent := libUtils.GetUserAgent(ctx)
	admin, err = service.Admin().GetAdminByUsernamePass(ctx, req)
	if err != nil {
		return
	}

	key := gconv.String(admin.Id) + "-" + gmd5.MustEncryptString(admin.Name) + gmd5.MustEncryptString(admin.Password)
	if g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool() {
		key = gconv.String(admin.Id) + "-" + gmd5.MustEncryptString(admin.Name) + gmd5.MustEncryptString(admin.Password+ip+userAgent)
	}
	token, err = service.GfToken().GenerateToken(ctx, key, admin)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("登录失败")
		return
	}

	//用户在线状态保存
	service.OnlineLog().Invoke(gctx.New(), &model.OnlineParams{
		Type:      consts.AdminOnlineType,
		UserAgent: userAgent,
		Uuid:      gmd5.MustEncrypt(token),
		Token:     token,
		Username:  admin.Name,
		Ip:        ip,
	})
	return
}
