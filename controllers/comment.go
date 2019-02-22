package controllers

import (
	"MyBlog/models"
	"MyBlog/syserror"
	"github.com/satori/go.uuid"
)

type CommentController struct {
	Base
}

// @router /comment/new/?:key [post]
func (this *CommentController) PostNewComment() {
	this.MustLogin()
	noteKey := this.Ctx.Input.Param(":key")
	key := uuid.Must(uuid.NewV1()).String()
	comment := this.GetMustString("content", "请输入内容")
	c := models.Comment{
		Key:     key,
		NoteKey: noteKey,
		User:    this.User,
		UserId:  int(this.User.ID),
		Content: comment,
	}
	if e := models.SaveComment(&c); e != nil {
		this.Abort500(syserror.New("提交到数据库发生错误", e))
	}
	if noteKey == "" {
		this.SaveJsonok("提交成功", H{"data": c})
	} else {
		this.SaveJson("提交失败", "")
	}

}

// @router /comment/count [get]
func (this *CommentController) GetCount() {
	count, err := models.QuerryCountByNoteKey("")
	if err != nil {
		this.Abort500(syserror.New("查询留言数量失败", err))
	}
	this.SaveJsonok("查询成功", H{"count": count})
}

// @router /comment/query [get]
func (this *CommentController) GetQuery() {
	pageon, e := this.GetInt("pageon", 1)
	if e != nil || pageon < 1 {
		pageon = 1
	}
	limit, err := this.GetInt("limit", 5)
	if err != nil {
		limit = 10
	}
	comments, err2 := models.QuerrByPageCommentByNoteKey(pageon, limit)
	if err2 != nil {
		this.Abort500(syserror.New("查询留言内容失败", err2))
	}

	this.SaveJsonok("查询成功", H{"data": comments})
}
