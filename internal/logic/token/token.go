package token

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tiger1103/gfast-token/gftoken"
	"sleep-service/internal/model"
	"sleep-service/internal/service"
	"sleep-service/utility/liberr"
)

type sToken struct {
	*gftoken.GfToken
}

func init() {
	service.RegisterGToken(New())
}

func New() *sToken {
	var (
		ctx = gctx.New()
		opt *model.TokenOptions
		err = g.Cfg().MustGet(ctx, "gfToken").Struct(&opt)
		fun gftoken.OptionFunc
	)
	liberr.ErrIsNil(ctx, err)
	if opt.CacheModel == "redis" {
		fun = gftoken.WithGRedis()
	} else {
		fun = gftoken.WithGCache()
	}
	return &sToken{
		GfToken: gftoken.NewGfToken(
			gftoken.WithCacheKey(opt.CacheKey),
			gftoken.WithTimeout(opt.Timeout),
			gftoken.WithMaxRefresh(opt.MaxRefresh),
			gftoken.WithMultiLogin(opt.MultiLogin),
			fun,
		),
	}
}
