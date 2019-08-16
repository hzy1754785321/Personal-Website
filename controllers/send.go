package controllers

import (
	"github.com/astaxie/beego"
	m "blog/models"
	"time"
	"os"
	"path"
	"gopkg.in/gomail.v2"
	"strconv"
	"fmt"
	"encoding/json"
	"math/rand"
)

//DataController data
type SendController struct {
	beego.Controller
}

type Sizer interface {
	Size() int64
 }

func SendMail(mailTo []string,subject string, body string ) error {
    mailConn := map[string]string {
        "user": "1754785321@qq.com", 
        "pass": "dudivkb?hzy?xdiwpbggb",   //授权码
        "host": "smtp.qq.com",
        "port": "465",
    }

    port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

    m := gomail.NewMessage()
    m.SetHeader("From","hzy" + "<" + mailConn["user"] + ">") 
    m.SetHeader("To", mailTo...)  //发送给多个用户
    m.SetHeader("Subject", subject)  //设置邮件主题
    m.SetBody("text/html", body)     //设置邮件正文

    d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

    err := d.DialAndSend(m)
    return err
}

func (this *SendController) UploadFiles() {
	f, h, err := this.GetFile("file") 
	if err != nil {
	   this.Data["json"] = m.GeneralResp{Code: 1, Error: "get file fail!"}
	   this.ServeJSON()
	   return
	}
	defer f.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求
	var AllowExtMap map[string]bool = map[string]bool{
	   ".jpg":  true,
	   ".jpeg": true,
	   ".png":  true,
	//    ".gif":  true,
	//    ".csv":  true,
	//    ".docx": true,
	//    ".xlsx": true,
	//    ".xls":  true,
	//    ".doc":  true,
	//    ".pdf":  true,
	//    ".txt":  true,
	//    ".html": true,
	//    ".ppt":  true,
	//    ".pptx": true,
	}
	var Filebytes = 1 << 24 //文件小于16兆
	if _, ok := AllowExtMap[ext]; !ok {
	   this.Data["json"] = m.GeneralResp{Code: 1, Error: "not allowed file format!"}
	   this.ServeJSON()
	   return
	}
 
	if fileSizer, ok := f.(Sizer); ok {
	   fileSize := fileSizer.Size()
	   if fileSize > int64(Filebytes) {
		  this.Data["json"] = m.GeneralResp{Code: 1, Error: "upload file error: file size exceeds 16M!"}
		  this.ServeJSON()
	   } else {
		  uploadDir := "./static/img/upload/" + time.Now().Format("2006/01/02/")
		  err = os.MkdirAll(uploadDir, os.ModePerm)
		  if err != nil {
			 this.Data["json"] = m.GeneralResp{Code: 1, Error: "create upload dir fail:" + err.Error()}
			 this.ServeJSON()
			 return
		}
		 
		fpath := uploadDir + h.Filename
		err = this.SaveToFile("file", fpath)
		if err != nil {
			this.Data["json"] = m.GeneralResp{Code: 1, Error: err.Error()}
			this.ServeJSON()
		}
		this.Data["json"] = m.GeneralResp{Code: 0, Data: fpath[1:len(fpath)]}
		this.ServeJSON()
	   }
	} else {
	   this.Data["json"] = m.GeneralResp{Code: 1, Error: "unable to read file size!"}
	   this.ServeJSON()
	}
 
 }

 func (s *SendController) SendVerifyMail() {
	sessionID := s.GetString("sessionID")
	sessionTemp := s.GetSession(sessionID)
	sessionDat, _ := sessionTemp.(string)
	var session m.Session
	json.Unmarshal([]byte(sessionDat),&session)
	emailAddress := s.GetString("email")
	mailTo := []string {
		emailAddress,
	}
	timeNow := time.Now().Format("2006-01-02 15:04")
	controlType := s.GetString("type")
	if controlType == "bind"{
		if m.CheckRedis(emailAddress){
			s.Data["json"] = map[string]interface{}{"status": false, "msg": "邮箱已绑定"}
			s.ServeJSON()
			return
		}	
		subject := "邮箱绑定"
		cdkey := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
		m.SetRedis("verity" + session.PersonalData.Username, cdkey)
		m.SetKeyExpire("verity" + session.PersonalData.Username,300)  //设置5分钟Key过期
		body := fmt.Sprintf("亲爱的 %s，您好！<br><br>您正于 %s 申请绑定hzy网站的邮箱,本次请求的邮件验证码是：<strong>%s</strong>。 本验证码5分钟内有效，请及时输入。<br>如非本人操作，请忽略该邮件。<br>（请不要回复本邮件）",session.UserInfoData.Nickname,timeNow,cdkey)
		SendMail(mailTo, subject, body)
		s.Data["json"] = map[string]interface{}{"status": true, "msg": "发送绑定邮件成功！"}
	}else if controlType == "unbind"{
		subject := "邮箱解绑"
		cdkey := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
		m.SetRedis("verity" + session.PersonalData.Username, cdkey)
		m.SetKeyExpire("verity" + session.PersonalData.Username,300)  //设置5分钟Key过期
		body := fmt.Sprintf("亲爱的 %s，您好！<br><br>您正于 %s 申请解绑hzy网站的邮箱,本次请求的邮件验证码是：<strong>%s</strong>。 本验证码5分钟内有效，请及时输入。<br>如非本人操作，请忽略该邮件。<br>（请不要回复本邮件）",session.UserInfoData.Nickname,timeNow,cdkey)
		SendMail(mailTo, subject, body)
		s.Data["json"] = map[string]interface{}{"status": true, "msg": "发送解绑邮件成功！"}
	}
	s.ServeJSON()
 }

 func (s *SendController) VerifyMailCode() {
	sessionID := s.GetString("sessionID")
	sessionTemp := s.GetSession(sessionID)
	sessionDat, _ := sessionTemp.(string)
	var session m.Session
	json.Unmarshal([]byte(sessionDat),&session)
	controlType := s.GetString("type")
	if controlType == "bind"{
		if m.CheckRedis("verity" + session.PersonalData.Username){
			cdkey := m.GetRedis("verity" + session.PersonalData.Username)
			cdkeyUser := s.GetString("cdkey")
			if cdkey != cdkeyUser {
				s.Data["json"] = map[string]interface{}{"status": false, "msg": "验证码错误！"}
			}else {
				emailAddress := s.GetString("email")
				m.SetRedis(emailAddress,session.PersonalData.Username)
				session.UserInfoData.Email = emailAddress
				loginjs, _ := json.Marshal(session.UserInfoData)
				m.SetRedis(session.UserInfoData.Username,string(loginjs))
				sessionDats, _ := json.Marshal(session)
				s.SetSession(sessionID,string(sessionDats))
				s.Data["json"] = map[string]interface{}{"status": true, "msg": "绑定成功！"}
			}
		}else{
			s.Data["json"] = map[string]interface{}{"status": false, "msg": "请先发送验证邮件！"}
		}
	}else if controlType == "unbind"{
		if m.CheckRedis("verity" + session.PersonalData.Username){
			cdkey := m.GetRedis("verity" + session.PersonalData.Username)
			cdkeyUser := s.GetString("cdkey")
			if cdkey != cdkeyUser {
				s.Data["json"] = map[string]interface{}{"status": false, "msg": "验证码错误！"}
			}else {
				emailAddress := s.GetString("email")
				m.DeleteRedis(emailAddress)
				session.UserInfoData.Email = ""
				loginjs, _ := json.Marshal(session.UserInfoData)
				m.SetRedis(session.UserInfoData.Username,string(loginjs))
				sessionDats, _ := json.Marshal(session)
				s.SetSession(sessionID,string(sessionDats))
				s.Data["json"] = map[string]interface{}{"status": true, "msg": "解绑成功！"}
			}
		}else{
			s.Data["json"] = map[string]interface{}{"status": false, "msg": "请先发送验证邮件！"}
		}
	}
	s.ServeJSON()
}
