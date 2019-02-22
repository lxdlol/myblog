package controllers

import (
	"MyBlog/models"
	"MyBlog/syserror"
	"github.com/astaxie/beego"
)

const SESSION = "SESSION_USER_KEY"

type Base struct {
	beego.Controller
	User    models.User
	Islogin bool
}
type MustPrepare interface {
	MustPrepare()
}

func (this *Base) Prepare() {
	this.Data["Path"] = this.Ctx.Request.RequestURI
	this.Islogin = false
	u, ok := this.GetSession(SESSION).(models.User)
	if ok {
		this.Islogin = true
		this.User = u
		this.Data["User"] = this.User
	}
	this.Data["Islogin"] = this.Islogin
	if prepare, ok := this.AppController.(MustPrepare); ok {
		prepare.MustPrepare()
	}

}
func (this *Base) Abort500(err error) {
	this.Data["error"] = err
	this.Abort("500")
}
func (this *Base) GetMustString(key, msg string) string {
	getString := this.GetString(key)
	if len(getString) == 0 {
		this.Abort500(syserror.New(msg, nil))
	}
	return getString
}
func (this *Base) MustLogin() {
	if !this.Islogin {
		this.Abort500(syserror.NoUser{})
	}
}
func (this *Base) SaveJson(msg, url string) {
	this.Data["json"] = map[string]interface{}{
		"code":   0,
		"action": url,
		"msg":    msg,
	}
	this.ServeJSON()
}

type H map[string]interface{}

func (this *Base) SaveJsonok(msg string, h H) {
	h["code"] = 0
	h["msg"] = msg
	this.Data["json"] = h
	this.ServeJSON()
}
