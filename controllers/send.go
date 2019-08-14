package controllers

import (
	"github.com/astaxie/beego"
	models "blog/models"
	"time"
	"os"
	"path"
)

//DataController data
type SendController struct {
	beego.Controller
}

type Sizer interface {
	Size() int64
 }
 
//  func (s *SendController) UploadFilesApi() {
// 	f, h, err := s.GetFile("file") 
// 	if err != nil {
// 		s.Data["json"] = models.GeneralResp{Code: 1, Error: "get file fail!"}
// 		s.ServeJSON()
// 		return
// 	}
// 	defer f.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
// 	var tmp = `{"src": "D://test.jpg"}`
// 	s.Data["json"] = map[string]interface{}{"code": 0, "msg": "", "data": tmp}
// 	s.ServeJSON()
//  }


func (this *SendController) UploadFiles() {
	f, h, err := this.GetFile("file") 
	if err != nil {
	   this.Data["json"] = models.GeneralResp{Code: 1, Error: "get file fail!"}
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
	   this.Data["json"] = models.GeneralResp{Code: 1, Error: "not allowed file format!"}
	   this.ServeJSON()
	   return
	}
 
	if fileSizer, ok := f.(Sizer); ok {
	   fileSize := fileSizer.Size()
	   if fileSize > int64(Filebytes) {
		  this.Data["json"] = models.GeneralResp{Code: 1, Error: "upload file error: file size exceeds 16M!"}
		  this.ServeJSON()
	   } else {
		  uploadDir := "./static/img/upload/" + time.Now().Format("2006/01/02/")
		  err = os.MkdirAll(uploadDir, os.ModePerm)
		  if err != nil {
			 this.Data["json"] = models.GeneralResp{Code: 1, Error: "create upload dir fail:" + err.Error()}
			 this.ServeJSON()
			 return
		  }
		 
		  fpath := uploadDir + h.Filename
		  err = this.SaveToFile("file", fpath)
		  if err != nil {
			 this.Data["json"] = models.GeneralResp{Code: 1, Error: err.Error()}
			 this.ServeJSON()
		  }
		  this.Data["json"] = models.GeneralResp{Code: 0, Data: fpath[1:len(fpath)]}
		  this.ServeJSON()
	   }
	} else {
	   this.Data["json"] = models.GeneralResp{Code: 1, Error: "unable to read file size!"}
	   this.ServeJSON()
	}
 
 }
