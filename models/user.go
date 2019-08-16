package models

//UserInfo 用户信息
type UserInfo struct {
	Username string
	Nickname string
	Password string
	LastTime string
	LoginCount int
	LastIP string
	Icon string
	Email string
	Iphone string
}

type PersonalInfo struct{
	Username string
	Nickname string
	Icon string
	Coin int
	Sex string
	Created string
	City string  //市
	Province string //省
	Area string  //区
	Birth string  
	Like []string 
	Sign string
}

// type LoginInfo struct{
// 	LastTime string
// 	Coin int
// 	LastIP string
// 	LastAddress string
// }

// type LoginInfos struct{
// 	Info []LoginInfo
// }

type GeneralResp struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
 }

//  var question = []string{"你的", "str2", "str3", "str4"}