package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego/session"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

var globalSessions *session.Manager

func init() {
//	beego.BConfig.WebConfig.Session.SessionOn = true
	// sessionConfig := &session.ManagerConfig{
	// 	CookieName:"sessionID", 
	// 	EnableSetCookie: true, 
	// 	Gclifetime:3600,
	// 	Maxlifetime: 3600, 
	// 	Secure: false,
	// 	CookieLifeTime: 3600,
	// 	ProviderConfig: "",
	// 	}
	// 	globalSessions, _ = session.NewManager("redis",sessionConfig)
	// 	go globalSessions.GC()
}


func main() {
	
	beego.Run()
}
