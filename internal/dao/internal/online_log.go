// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OnlineLogDao is the data access object for the table online_log.
type OnlineLogDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  OnlineLogColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// OnlineLogColumns defines and stores column names for the table online_log.
type OnlineLogColumns struct {
	Id         string //
	Uuid       string // 用户标识
	Token      string // 用户token
	CreateTime string // 登录时间
	UserName   string // 用户名
	Ip         string // 登录ip
	Explorer   string // 浏览器
	Os         string // 操作系统
	Type       string // 类型1=平台。2=b端，3=C端
}

// onlineLogColumns holds the columns for the table online_log.
var onlineLogColumns = OnlineLogColumns{
	Id:         "id",
	Uuid:       "uuid",
	Token:      "token",
	CreateTime: "create_time",
	UserName:   "user_name",
	Ip:         "ip",
	Explorer:   "explorer",
	Os:         "os",
	Type:       "type",
}

// NewOnlineLogDao creates and returns a new DAO object for table data access.
func NewOnlineLogDao(handlers ...gdb.ModelHandler) *OnlineLogDao {
	return &OnlineLogDao{
		group:    "default",
		table:    "online_log",
		columns:  onlineLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OnlineLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OnlineLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OnlineLogDao) Columns() OnlineLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OnlineLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OnlineLogDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *OnlineLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
