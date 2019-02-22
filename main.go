package main

import (
	"MyBlog/models"
	_ "MyBlog/routers"
	"encoding/gob"
	"github.com/astaxie/beego"
	_"MyBlog/models"
	"strings"
)

func main() {
	initTemplete()
	initSession()
	beego.Run()
}

func initTemplete(){
	beego.AddFuncMap("equal", func(x,y string)bool {
		x= strings.Trim(x, "/")
		y= strings.Trim(y, "/")
		return strings.Compare(x,y)==0
	})
	beego.AddFuncMap("add", func(x,y int)int {
		return x+y
	})
}

func initSession(){
	gob.Register(models.User{})
	beego.BConfig.WebConfig.Session.SessionOn=true
	beego.BConfig.WebConfig.Session.SessionName="lxdbolg"
	beego.BConfig.WebConfig.Session.SessionProvider="file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}