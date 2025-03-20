// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AdminAccount is the golang structure for table admin_account.
type AdminAccount struct {
	Id            uint   `json:"id"            orm:"id"            description:""`
	Uid           int    `json:"uid"           orm:"uid"           description:"添加用户"`
	AccountID     int    `json:"accountID"     orm:"accountID"     description:"账号id"`
	DeptId        int    `json:"deptId"        orm:"dept_id"       description:"部门id"`
	Username      string `json:"username"      orm:"username"      description:"用户账号"`
	Password      string `json:"password"      orm:"password"      description:"密码"`
	Salt          string `json:"salt"          orm:"salt"          description:"密码盐"`
	Name          string `json:"name"          orm:"name"          description:"姓名"`
	Nickname      string `json:"nickname"      orm:"nickname"      description:"昵称"`
	Avatar        string `json:"avatar"        orm:"avatar"        description:"头像"`
	Mobile        string `json:"mobile"        orm:"mobile"        description:"手机号码"`
	Email         string `json:"email"         orm:"email"         description:"邮箱"`
	LastLoginIp   string `json:"lastLoginIp"   orm:"lastLoginIp"   description:"最后登录IP"`
	LastLoginTime int    `json:"lastLoginTime" orm:"lastLoginTime" description:"最后登录时间"`
	Status        int    `json:"status"        orm:"status"        description:"状态0=正常，1=禁用"`
	Validtime     int    `json:"validtime"     orm:"validtime"     description:"账号有效时间0=无限"`
	Createtime    int    `json:"createtime"    orm:"createtime"    description:"创建时间"`
	Updatetime    int    `json:"updatetime"    orm:"updatetime"    description:"修改时间"`
	Address       string `json:"address"       orm:"address"       description:"地址"`
	City          string `json:"city"          orm:"city"          description:"城市"`
	Remark        string `json:"remark"        orm:"remark"        description:"描述"`
	Company       string `json:"company"       orm:"company"       description:"公司名称"`
	Province      string `json:"province"      orm:"province"      description:"省份"`
	Area          string `json:"area"          orm:"area"          description:"地区"`
}
