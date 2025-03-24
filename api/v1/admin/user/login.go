package user

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" summary:"登录" tags:"登录" noAuth:"true"`
	UserName string `json:"username" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码"`
}

type LoginRes struct {
	Token string `json:"token" dc:"登录令牌；请求头：Authorization=Bearer token"`
}

type GetAdminInfoReq struct {
	g.Meta `path:"/admin/info" method:"get" summary:"获取用户信息" tags:"登录"`
}

type GetAdminInfoRes struct {
	Token string `json:"token"`
}
