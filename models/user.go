package models

import("time")

//UserInfo w
type UserInfo struct {
	Username string
	Nickname string
	Password string
//	Email string
	LoginCount int
	LastTime time.Time
	LastIP string
//	State int8
	Created time.Time
//	Updated time.Time
}

