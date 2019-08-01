package models

//UserInfo 用户信息
type UserInfo struct {
	Username string
	Nickname string
	Password string
//	Email string
	LoginCount int
	LastTime string
	LastIP string
//	State int8
	Created string
//	Updated time.Time
}


var userData string = `{
	"Username": "",
	"Nickname": "",
	"Password": "",
	"LoginCount": 0,
	"LastTime": "",
	"LastIP": "",
	"Created": "",
}`

