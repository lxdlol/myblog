package controllers

import (
	"MyBlog/models"
	"MyBlog/syserror"
	"github.com/astaxie/beego"
)

type IndexControllers struct {
	Base
}

// @router / [get]
func (this *IndexControllers) GetIndex() {
	//查询总页数
	page, e := this.GetInt("page", 1)
	if e != nil && page < 1 {
		this.Abort500(e)
	}
	//每页最小显示5篇文章
	var limit = 5
	title := this.GetString("title")
	//通过page查询当前页面文章
	note, err := models.QuerrByPageTitle(page, limit, title)
	if err != nil {
		this.Abort500(err)
	}
	this.Data["note"] = note
	//查询总页数，有没有下一页
	count, err2 := models.QuerryCount(title)
	if err2 != nil {
		this.Abort500(err2)
	}
	allpage := count / limit
	if count%limit != 0 {
		allpage += 1
	}
	this.Data["allpage"] = allpage
	this.Data["page"] = page
	this.Data["title"] = title
	this.TplName = "index.html"
}

// @router /about [get]
func (this *IndexControllers) GetAbout() {
	this.TplName = "about.html"
}

// @router /message [get]
func (this *IndexControllers) GetMessage() {
	this.TplName = "message.html"
}

// @router /login [get]
func (this *IndexControllers) GetLogin() {
	this.TplName = "login.html"
}

// @router /details/:key [get]
func (this *IndexControllers) GetDetails() {
	key := this.Ctx.Input.Param(":key")
	note, err := models.QuerryByKey(key)
	if err != nil {
		this.Abort500(syserror.New("文章不存在1", err))
	}
	comments, e := models.QuerryByNotekey(key)
	if e != nil {
		this.Abort500(syserror.New("文章不存在44", e))
	}
	this.Data["note"] = note
	this.Data["comments"] = comments
	this.TplName = "details.html"
}

// @router /comment/:key [get]
func (this *IndexControllers) GetComment() {
	this.MustLogin()
	key := this.Ctx.Input.Param(":key")
	note, err := models.QuerryByKey(key)
	if err != nil {
		beego.Info(err)
		this.Abort500(syserror.New("ok", err))
	}
	this.Data["note"] = note
	this.TplName = "comment.html"
}
