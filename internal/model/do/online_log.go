// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// OnlineLog is the golang structure of table online_log for DAO operations like Where/Data.
type OnlineLog struct {
	g.Meta     `orm:"table:online_log, do:true"`
	Id         interface{} //
	Uuid       interface{} // 用户标识
	Token      interface{} // 用户token
	CreateTime interface{} // 登录时间
	UserName   interface{} // 用户名
	Ip         interface{} // 登录ip
	Explorer   interface{} // 浏览器
	Os         interface{} // 操作系统
	Type       interface{} // 类型1=平台。2=b端，3=C端
}
