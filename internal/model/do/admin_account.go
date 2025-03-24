// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AdminAccount is the golang structure of table admin_account for DAO operations like Where/Data.
type AdminAccount struct {
	g.Meta     `orm:"table:admin_account, do:true"`
	Id         interface{} //
	Uid        interface{} // 添加用户
	AccountID  interface{} // 账号id
	DeptId     interface{} // 部门id
	Username   interface{} // 用户账号
	Password   interface{} // 密码
	Salt       interface{} // 密码盐
	Name       interface{} // 姓名
	Avatar     interface{} // 头像
	Mobile     interface{} // 手机号码
	Email      interface{} // 邮箱
	Status     interface{} // 状态0=正常，1=禁用
	Validtime  interface{} // 账号有效时间0=无限
	Createtime interface{} // 创建时间
	Updatetime interface{} // 修改时间
	Address    interface{} // 地址
	City       interface{} // 城市
	Remark     interface{} // 描述
	Company    interface{} // 公司名称
	Province   interface{} // 省份
	Area       interface{} // 地区
}
