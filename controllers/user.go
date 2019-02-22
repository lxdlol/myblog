package controllers

import (
	"MyBlog/models"
	"MyBlog/syserror"
	"strings"
)

type UserController struct {
	Base
}

// @router /login [post]
func (this *UserController) PostLogin() {
	//获取email
	email := this.GetMustString("email", "邮箱不能为空")
	//获取密码
	pwd := this.GetMustString("password","密码不能为空")
	//数据库查询
	user, err := models.QuerryByEmail(email)
	if err!=nil {
		this.Abort500(syserror.New("用户不存在",nil))
	}
	if user.Pwd!=pwd {
		this.Abort500(syserror.New("密码错误",nil))
	}
	//登陆成功，将session保存
	this.SetSession(SESSION,user)
	this.Data["json"]=map[string]interface{}{
		"code": 0,
		"action":"/",
	}
	this.ServeJSON()
}
// @router /logout [get]
func(this *UserController)Logout(){
	this.MustLogin()
	this.DelSession(SESSION)
	this.Redirect("/",302)
}


// @router /addUser [post]
func(this *UserController)AddUser(){
	name:= this.GetMustString("name", "邮箱不能为空")
	email := this.GetMustString("email", "邮箱不能为空")
	pwd := this.GetMustString("pwd", "邮箱不能为空")
	pwd1 := this.GetMustString("pwd1", "邮箱不能为空")
	if strings.Compare(pwd,pwd1)!=0{
		this.Abort500(syserror.New("两次输入的密码不一致",nil))
	}
	user, err := models.QuerryByEmail(email)
	if err==nil&&strings.Compare(email,user.Email)==0{
		this.Abort500(syserror.New("此邮箱已注册",nil))
	}
	if user2, e := models.QuerryByName(name);e==nil&&strings.Compare(user2.Name,name)==0{
		this.Abort500(syserror.New("用户名已注册",nil))
	}
	if err:=models.SaveUser(models.User{
		Name:name,
		Email:email,
		Pwd :  pwd,
		Avatar : "/static/images/info-img.png",
		Role :1,
	});err!=nil{
		this.Abort500(syserror.New("注册失败",err))
	}
	this.Data["json"]=map[string]interface{}{
		"code": 0,
		"action":"/",
	}
	this.ServeJSON()
}
// @router /addUser [get]
func(this *UserController)GetAddUser(){
	this.TplName="addUser.html"
}




