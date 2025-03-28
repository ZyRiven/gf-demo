package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast-token/gftoken"
)

type IGfToken interface {
	// GenerateToken 生成token
	GenerateToken(ctx context.Context, key string, data interface{}) (keys string, err error)
	// Middleware 绑定group
	Middleware(group *ghttp.RouterGroup) error
	// ParseToken 解析token (只验证格式并不验证过期)
	ParseToken(r *ghttp.Request) (*gftoken.CustomClaims, error)
	IsLogin(r *ghttp.Request) (b bool, failed *gftoken.AuthFailed)
	GetRequestToken(r *ghttp.Request) (token string)
	RemoveToken(ctx context.Context, token string) (err error)
	GetTokenData(ctx context.Context, token string) (tData *gftoken.TokenData, key string, err error)
	IsNotExpired(token string) (*gftoken.CustomClaims, int)
}

var gt IGfToken

func RegisterGToken(gtk IGfToken) {
	gt = gtk
}

func GfToken() IGfToken {
	if gt == nil {
		panic("implement not found for interface IGfToken, forgot register?")
	}
	return gt
}
