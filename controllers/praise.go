package controllers

import (
	"MyBlog/models"
	"MyBlog/syserror"
	"github.com/astaxie/beego"
	"github.com/pkg/errors"
)

type PraiseControllers struct {
	Base
}

func (this *PraiseControllers) MustPrepare() {
	this.MustLogin()
}

func (this *PraiseControllers) Praise() {
	ttype := this.Ctx.Input.Param(":type")
	key := this.Ctx.Input.Param(":key")
	var table string
	switch ttype {
	case "comment":
		table = "comments"
	case "note":
		table = "notes"
	default:
		this.Abort500(errors.New("未知错误"))
	}
	pcnt, err := models.UpdatePraise(table, key, int(this.User.ID))
	if err != nil {
		if e, ok := err.(syserror.HasPriseError); ok {
			this.Abort500(e)
		}
		this.Abort500(errors.New("未知错误"))
	}
	beego.Info(1111)
	this.SaveJsonok("点赞成功", H{"praise": pcnt})
}
