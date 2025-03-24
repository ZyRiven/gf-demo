// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sleep-service/internal/model"
)

type (
	IOnlineLog interface {
		Invoke(ctx context.Context, params *model.OnlineParams)
		SaveOnline(ctx context.Context, params *model.OnlineParams)
	}
)

var (
	localOnlineLog IOnlineLog
)

func OnlineLog() IOnlineLog {
	if localOnlineLog == nil {
		panic("implement not found for interface IOnlineLog, forgot register?")
	}
	return localOnlineLog
}

func RegisterOnlineLog(i IOnlineLog) {
	localOnlineLog = i
}
