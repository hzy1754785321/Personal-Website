package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.MainController{}, "*:Home")
	beego.Router("/handleLogin", &controllers.DataController{}, "post:HandleLogin")
	beego.Router("/HandleRegister", &controllers.DataController{}, "post:HandleRegister")
	beego.Router("/GetSessionUserInfo", &controllers.DataController{}, "post:GetSessionUserInfo")
	beego.Router("/GetSessionPersonal", &controllers.DataController{}, "post:GetSessionPersonal")
	beego.Router("/SavePersonal", &controllers.DataController{}, "post:SavePersonal")
	beego.Router("/personal", &controllers.MainController{}, "*:Personal")
	beego.Router("/personal/userInfo", &controllers.MainController{}, "*:UserInfo")
	beego.Router("/personal/security", &controllers.MainController{}, "*:Security")
	beego.Router("/personal/icon", &controllers.MainController{}, "*:Icon")
}
