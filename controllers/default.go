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

func (c *MainController) Personal()  {

	c.TplName= "personal/personal.html"
}

func (c *MainController) UserInfo()  {

	c.TplName= "personal/userInfo.html"
}

func (c *MainController) Security()  {

	c.TplName= "personal/security.html"
}

func (c *MainController) Icon()  {

	c.TplName= "personal/icon.html"
}

