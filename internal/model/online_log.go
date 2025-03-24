package model

// OnlineParams 用户在线状态写入参数
type OnlineParams struct {
	UserAgent string
	Uuid      string
	Token     string
	Username  string
	Ip        string
	Type      int
}
