package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego/session"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	"math/rand"
	"time"
	"os/exec"
	"os"
	"path/filepath"
)

var globalSessions *session.Manager

func init() {
	//	beego.BConfig.WebConfig.Session.SessionOn = true
	// sessionConfig := &session.ManagerConfig{
	// 	Cookie	Name:"sessionID", 
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
	rand.Seed(time.Now().UnixNano())
	initConfig()
	beego.Run()
}

func initConfig() {

rootPath := GetAPPRootPath()

beego.SetViewsPath(rootPath +"/views")

beego.LoadAppConfig("ini", rootPath+"/conf/app.conf")

beego.SetStaticPath("static", rootPath+"/static")

}

func GetAPPRootPath() string {

file, err := exec.LookPath(os.Args[0])

if err != nil {

return ""

  }

p, err := filepath.Abs(file)

if err != nil {

return ""

  }

return filepath.Dir(p)

}