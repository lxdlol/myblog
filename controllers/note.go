package controllers

import (
	"MyBlog/models"
	"MyBlog/syserror"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type NoteControllers struct {
	Base
}

func (this *NoteControllers) MustPrepare() {
	this.MustLogin()
	if this.User.Role != 0 {
		this.Abort500(syserror.New("你没有权限添加文章", nil))
	}
}

// @router /new [get]
func (this *NoteControllers) GetNew() {
	this.Data["key"] = uuid.Must(uuid.NewV1()).String()
	this.TplName = "addNote.html"
}

// @router /new/:key [post]
func (this *NoteControllers) PostNew() {
	key := this.Ctx.Input.Param(":key")
	title := this.GetMustString("title", "标题不能为空")
	content := this.GetMustString("content", "文章详情不能为空")
	note, err := models.QuerryByKey(key)
	var n models.Note
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			n = models.Note{
				Key:     key,
				Title:   title,
				Comtent: content,
				UserId:  int(this.User.ID),
			}
			sum, _ := GetSummary(n.Comtent)
			n.Summary = sum
			if e := models.SaveNote(&n); e != nil {
				this.Abort500(syserror.New("提交失败", e))
			}
		} else {
			this.Abort500(syserror.New("保存失败", err))
		}
	} else {
		note.Title = title
		note.Comtent = content
		note.Summary, _ = GetSummary(note.Comtent)
		if e := models.SaveNote(&note); e != nil {
			this.Abort500(syserror.New("提交失败", e))
		}
	}
	conn, _ := redis.Dial("tcp", ":6379")
	defer conn.Close()
	_, e := conn.Do("get", key)
	if e == nil {
		marshal, _ := json.Marshal(note)
		conn.Do("set", key, string(marshal))
	}
	this.SaveJson("提交成功", fmt.Sprintf("/details/%s", key))
}

// @router /edit/:key [get]
func (this *NoteControllers) Edit() {
	key := this.Ctx.Input.Param(":key")
	note, err := models.QuerryByKeyAndUserId(key, int(this.User.ID))
	if err != nil {
		this.Abort500(err)
	}
	this.Data["note"] = note
	this.Data["key"] = key
	this.TplName = "addNote.html"
}

// @router /delete/:key [get]
func (this *NoteControllers) Delete() {
	key := this.Ctx.Input.Param(":key")
	user := this.User
	e := models.DeleteByKeyAndId(key, int(user.ID))
	if e != nil {
		this.Abort500(e)
	}
	conn, _ := redis.Dial("tcp", ":6379")
	defer conn.Close()
	conn.Do("del", key)
	this.Redirect("/", 302)
}

func GetSummary(content string) (summary string, err error) {
	var buffbytes bytes.Buffer
	if _, e := buffbytes.Write([]byte(content)); e != nil {
		return "", e
	}
	document, e := goquery.NewDocumentFromReader(&buffbytes)
	if e != nil {
		return "", e
	}
	text := document.Find("body").Text()
	if len(text) > 400 {
		text = text[:400]
	}
	return text, nil
}
