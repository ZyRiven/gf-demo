package model

import "sleep-service/internal/model/entity"

type Context struct {
	User *ContextUser // User in context.
}

type ContextUser struct {
	*entity.AdminAccount
}
