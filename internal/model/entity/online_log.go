// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// OnlineLog is the golang structure for table online_log.
type OnlineLog struct {
	Id         uint64 `json:"id"         orm:"id"          description:""`
	Uuid       string `json:"uuid"       orm:"uuid"        description:"用户标识"`
	Token      string `json:"token"      orm:"token"       description:"用户token"`
	CreateTime int64  `json:"createTime" orm:"create_time" description:"登录时间"`
	UserName   string `json:"userName"   orm:"user_name"   description:"用户名"`
	Ip         string `json:"ip"         orm:"ip"          description:"登录ip"`
	Explorer   string `json:"explorer"   orm:"explorer"    description:"浏览器"`
	Os         string `json:"os"         orm:"os"          description:"操作系统"`
	Type       int    `json:"type"       orm:"type"        description:"类型1=平台。2=b端，3=C端"`
}
