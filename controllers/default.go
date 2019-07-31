package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type MainController struct {
	beego.Controller
	controllerName string
}


func (p *MainController) Prepare()  {
	controllerName, _ := p.GetControllerAndAction()
	p.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
}

func (c *MainController) Get() {
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "login.html"
}


func (c *MainController) Home()  {
//	c.list()
	c.TplName= "home.html"
}

//获取用户IP地址
func (p *MainController) getClientIp() string {
	s := strings.Split(p.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

