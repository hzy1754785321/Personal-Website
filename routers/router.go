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
}
