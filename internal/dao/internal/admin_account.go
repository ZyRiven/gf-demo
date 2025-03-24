// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminAccountDao is the data access object for the table admin_account.
type AdminAccountDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  AdminAccountColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// AdminAccountColumns defines and stores column names for the table admin_account.
type AdminAccountColumns struct {
	Id         string //
	Uid        string // 添加用户
	AccountID  string // 账号id
	DeptId     string // 部门id
	Username   string // 用户账号
	Password   string // 密码
	Salt       string // 密码盐
	Name       string // 姓名
	Avatar     string // 头像
	Mobile     string // 手机号码
	Email      string // 邮箱
	Status     string // 状态0=正常，1=禁用
	Validtime  string // 账号有效时间0=无限
	Createtime string // 创建时间
	Updatetime string // 修改时间
	Address    string // 地址
	City       string // 城市
	Remark     string // 描述
	Company    string // 公司名称
	Province   string // 省份
	Area       string // 地区
}

// adminAccountColumns holds the columns for the table admin_account.
var adminAccountColumns = AdminAccountColumns{
	Id:         "id",
	Uid:        "uid",
	AccountID:  "accountID",
	DeptId:     "dept_id",
	Username:   "username",
	Password:   "password",
	Salt:       "salt",
	Name:       "name",
	Avatar:     "avatar",
	Mobile:     "mobile",
	Email:      "email",
	Status:     "status",
	Validtime:  "validtime",
	Createtime: "createtime",
	Updatetime: "updatetime",
	Address:    "address",
	City:       "city",
	Remark:     "remark",
	Company:    "company",
	Province:   "province",
	Area:       "area",
}

// NewAdminAccountDao creates and returns a new DAO object for table data access.
func NewAdminAccountDao(handlers ...gdb.ModelHandler) *AdminAccountDao {
	return &AdminAccountDao{
		group:    "default",
		table:    "admin_account",
		columns:  adminAccountColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AdminAccountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AdminAccountDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AdminAccountDao) Columns() AdminAccountColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AdminAccountDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AdminAccountDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AdminAccountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
