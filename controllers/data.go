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
	c.Data["json"] = map[string]interface{}{"Username": user.Username , "Nickname": user.Nickname,"LastTime": user.LastTime}
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
	// c.Data["json"] = map[string]interface{}{
	// 	 "Username": person.Username ,
	// 	 "Nickname": person.Nickname,
	// 	 "Coin": person.Coin,
	// 	 "Sex" : person.Sex,
	// 	 "Created" : person.Created,
	// 	 "City" : person.City,
	// 	 "Province" : person.Province,
	// 	 "Sex" : person.Sex,
	// 	 "Sex" : person.Sex,
	// 	 "Sex" : person.Sex,
	// }
	c.ServeJSON()
}

//getClientIp 获取用户IP地址
func (p *DataController) getClientIp() string {
	var ip = p.Ctx.Input.IP()
	fmt.Printf(ip)
	// s := strings.Split(p.Ctx.Request.RemoteAddr, ":")
	return ip	
}
