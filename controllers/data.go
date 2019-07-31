package controllers

import (
	"github.com/astaxie/beego"
	m"blog/models"
)

//DataController data
type DataController struct {
	beego.Controller
}

//HandleLogin 处理登录
func (c *DataController) HandleLogin() {
	username, password := c.GetString("login"), c.GetString("pwd")
	if username == "admin" && password == "123456" {
	//	c.SetSession("username", username)
		c.Data["json"] = map[string]interface{}{"status": true, "msg": "登录成功"}
	} else {
		c.Data["json"] = map[string]interface{}{"status": false, "msg": "登录失败"}
	}
	c.ServeJSON()
}

//HandleRegister 处理注册
func (c *DataController) HandleRegister() {
	username, password := c.GetString("username"), c.GetString("pwd")
	if CheckRedis(username){
		c.Data["json"] = map[string]interface{}{"status": false, "msg": "用户名已存在"}
	}else{
		
		
	}
	
}