package models

// import (
// "github.com/astaxie/beego/session"
// )

// var GlobalSessions *session.Manager

//Session session信息
type Session struct {
	UserInfoData  UserInfo  `json:"UserInfo"` //用户信息
}