package onlineLog

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/mssola/user_agent"
	"sleep-service/internal/dao"
	"sleep-service/internal/model"
	"sleep-service/internal/model/do"
	"sleep-service/internal/model/entity"
	"sleep-service/internal/service"
	"sleep-service/utility/liberr"
	"time"
)

type sOnlineLog struct {
	Pool *grpool.Pool
}

func init() {
	service.RegisterOnlineLog(New())
}

func New() *sOnlineLog {
	return &sOnlineLog{
		Pool: grpool.New(100),
	}
}

func (s *sOnlineLog) Invoke(ctx context.Context, params *model.OnlineParams) {
	err := s.Pool.Add(ctx, func(ctx context.Context) {
		//写入数据
		s.SaveOnline(ctx, params)
	})
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

func (s *sOnlineLog) SaveOnline(ctx context.Context, params *model.OnlineParams) {
	err := g.Try(ctx, func(ctx context.Context) {
		ua := user_agent.New(params.UserAgent)
		browser, _ := ua.Browser()
		os := ua.OS()
		var (
			info *entity.OnlineLog
			data = &do.OnlineLog{
				Type:       params.Type,
				Uuid:       params.Uuid,
				Token:      params.Token,
				CreateTime: time.Now().Unix(),
				UserName:   params.Username,
				Ip:         params.Ip,
				Explorer:   browser,
				Os:         os,
			}
		)
		//查询是否已存在当前用户
		err := dao.OnlineLog.Ctx(ctx).Fields(dao.OnlineLog.Columns().Id).
			Where(dao.OnlineLog.Columns().Token, data.Token).
			Scan(&info)
		liberr.ErrIsNil(ctx, err)
		//若已存在则更新
		if info != nil {
			_, err = dao.OnlineLog.Ctx(ctx).
				Where(dao.OnlineLog.Columns().Id, info.Id).
				FieldsEx(dao.OnlineLog.Columns().Id).Update(data)
			liberr.ErrIsNil(ctx, err)
		} else { //否则新增
			_, err = dao.OnlineLog.Ctx(ctx).
				FieldsEx(dao.OnlineLog.Columns().Id).Insert(data)
			liberr.ErrIsNil(ctx, err)
		}
	})
	if err != nil {
		g.Log().Error(ctx, err)
	}
}
