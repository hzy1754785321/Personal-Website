package models

import "time"

type User struct {
	Username string
	Nickname string
	Password string
//	Email string
	LoginCount int
	LastTime time.Time
	LastIp string
//	State int8
	Created time.Time
//	Updated time.Time
}

