package controllers

import (
	"github.com/astaxie/beego"
	m "blog/models"
	"time"
	"encoding/json"
	"fmt"
	"crypto/md5"
	"io"
	"math/rand"
)

//DataController data
type DataController struct {
	beego.Controller
}

//HandleLogin 处理登录
func (c *DataController) HandleLogin() {
	username, password := c.GetString("login"), c.GetString("pwd")
	if m.CheckRedis(username){
		var user m.UserInfo
		js := m.GetRedis(username)
		json.Unmarshal([]byte(js),&user)
		if user.Password == password{
			user.LastTime = time.Now().Format("2006-01-02 15:04:05")
			if user.LoginCount > 1{
				msg :=  fmt.Sprintf("登录成功！<br /><br />欢迎回来,%s",user.Nickname)
				c.Data["json"] = map[string]interface{}{"status": true, "msg": msg}
			}else{
				c.Data["json"] = map[string]interface{}{"status": true, "msg": "登录成功！<br /><br />欢迎来到hzy的个人网站"}
			}
			user.LoginCount++
			jsDat, _ := json.Marshal(user)
			m.SetRedis(username,string(jsDat))
			var person m.PersonalInfo
			perjs := m.GetRedis("personal" + username)
			json.Unmarshal([]byte(perjs),&person)
			md := md5.New()
    		io.WriteString(md, username)
			sessionID := fmt.Sprintf("%x", md.Sum(nil))
			var session m.Session
			session.UserInfoData = user
			session.PersonalData = person
			sessionDat, _ := json.Marshal(session)
			c.SetSession(sessionID,string(sessionDat))
			c.Ctx.SetCookie("sessionID",sessionID)
		}else{
			c.Data["json"] = map[string]interface{}{"status": false, "msg": "密码错误"}
		}
	}else{
		c.Data["json"] = map[string]interface{}{"status": false, "msg": "用户名不存在"}
	}
	c.ServeJSON()
}


//HandleRegister 处理注册
func (c *DataController) HandleRegister() {
	username, password := c.GetString("username"),  c.GetString("pwd") 
	nickname := c.GetString("nickname")
	if m.CheckRedis(username){
		c.Data["json"] = map[string]interface{}{"status": false, "msg": "用户名已存在"}
	}else{
		var user m.UserInfo
		var person m.PersonalInfo
		user.Username = username
		user.Nickname = nickname
		user.Password = password
		user.Icon = "/static/img/pic.jpg"
		user.LoginCount = 1
		nowTime := time.Now().Format("2006-01-02 15:04:05")
		user.LastTime = nowTime
		js, _ := json.Marshal(user)
		m.SetRedis(username,string(js))
		person.Username = username
		person.Nickname = nickname
		person.Created = nowTime
		coinAdd := rand.Intn(10)
		person.Coin += coinAdd
		perjs, _ := json.Marshal(person)
		m.SetRedis("personal" + username,string(perjs))
		c.Data["json"] = map[string]interface{}{"status": true, "msg": "注册成功"}
	}
	c.ServeJSON()
}

//GetSessionUserInfo  获得session
func (c *DataController) GetSessionUserInfo() {
	sessionID := c.GetString("sessionID")
	sessionTemp := c.GetSession(sessionID)
	sessionDat, _ := sessionTemp.(string)
	var user m.UserInfo
	var session m.Session
	json.Unmarshal([]byte(sessionDat),&session)
	user = session.UserInfoData
	c.Data["json"] = map[string]interface{}{"Username": user.Username , "Nickname": user.Nickname,"LastTime": user.LastTime,"Icon":user.Icon}
	c.ServeJSON()
}

func (c *DataController) GetSessionPersonal() {
	sessionID := c.GetString("sessionID")
	sessionTemp := c.GetSession(sessionID)
	sessionDat, _ := sessionTemp.(string)
	var session m.Session
	var person m.PersonalInfo
	json.Unmarshal([]byte(sessionDat),&session)
	person = session.PersonalData
	perjs, _ := json.Marshal(person)
	per := make(map[string]interface{})
	json.Unmarshal(perjs, &per)
	c.Data["json"] = per
	c.ServeJSON()
}

func (c *DataController) SavePersonal() {
	var person m.PersonalInfo
	sessionID := c.GetString("sessionID")
	sessionTemp := c.GetSession(sessionID)
	sessionDat, _ := sessionTemp.(string)
	var session m.Session
	json.Unmarshal([]byte(sessionDat),&session)
	person = session.PersonalData
	person.Nickname, person.Birth , person.Sex =   c.GetString("NickName") , c.GetString("Birth"), c.GetString("Sex")
	if c.GetString("Province") != "" {
		person.City , person.Province , person.Area  =  c.GetString("City"),  c.GetString("Province"), c.GetString("Area")
	}
	perjs, _ := json.Marshal(person)
	m.SetRedis("personal" + person.Username,string(perjs))
	session.UserInfoData = session.UserInfoData
	session.PersonalData = person
	sessionDats, _ := json.Marshal(session)
	c.SetSession(sessionID,string(sessionDats))
	c.Data["json"] = map[string]interface{}{"status": false, "msg": "修改成功！"}
	c.ServeJSON()
}

func (c *DataController) SaveIcon() {
	sessionID := c.GetString("sessionID")
	icon := c.GetString("Icon")
	sessionTemp := c.GetSession(sessionID)
	sessionDat, _ := sessionTemp.(string)
	var session m.Session
	json.Unmarshal([]byte(sessionDat),&session)
	session.UserInfoData.Icon = icon
	session.PersonalData.Icon = icon	
	perjs, _ := json.Marshal(session.PersonalData)
	m.SetRedis("personal" + session.PersonalData.Username,string(perjs))
	loginjs, _ := json.Marshal(session.UserInfoData)
	m.SetRedis(session.UserInfoData.Username,string(loginjs))
	sessionDats, _ := json.Marshal(session)
	c.SetSession(sessionID,string(sessionDats))
	c.Data["json"] = map[string]interface{}{"status": true, "msg": "修改成功！","path":icon}
	c.ServeJSON()
}

func (c *DataController) ChangeSecurity() {
	sessionID := c.GetString("sessionID")
	sessionTemp := c.GetSession(sessionID)
	sessionDat, _ := sessionTemp.(string)
	var session m.Session
	json.Unmarshal([]byte(sessionDat),&session)
	securityType := c.GetString("securityType")
	switch securityType{
	case "0": session.UserInfoData.Password = c.GetString("Password")
	case "1": session.UserInfoData.Email = c.GetString("Email")
	case "2": session.UserInfoData.Iphone = c.GetString("iphone")
	}
	loginjs, _ := json.Marshal(session.UserInfoData)
	m.SetRedis(session.UserInfoData.Username,string(loginjs))
	sessionDats, _ := json.Marshal(session)
	c.SetSession(sessionID,string(sessionDats))
	c.Data["json"] = map[string]interface{}{"status": true, "msg": "修改成功！"}
	c.ServeJSON()
}

//getClientIp 获取用户IP地址
func (p *DataController) getClientIp() string {
	var ip = p.Ctx.Input.IP()
	fmt.Printf(ip)	
	// s := strings.Split(p.Ctx.Request.RemoteAddr, ":")
	return ip	
}

func(c *DataController) CheckPassword() {
	sessionID := c.GetString("sessionID")
	sessionTemp := c.GetSession(sessionID)
	sessionDat, _ := sessionTemp.(string)
	var session m.Session
	json.Unmarshal([]byte(sessionDat),&session)
	password := c.GetString("pwd") 
	if password != session.UserInfoData.Password {
		c.Data["json"] = map[string]interface{}{"status": false, "msg": "密码错误"}
	}else{
		c.Data["json"] = map[string]interface{}{"status": true, "msg": "密码正确"}
	}
	c.ServeJSON()
}
