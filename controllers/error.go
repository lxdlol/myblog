package controllers

import (
	"MyBlog/syserror"
	"github.com/astaxie/beego"
)

type ErrorController struct {
	Base
}

//如果请求是ajax请求，返回jason{code:,msg:,ression:error}数据hahh
func (c *ErrorController) Error404() {
	c.TplName = "error/404.html"
	if c.IsAjax() {
		c.jerror(syserror.Error404{})
	}
}

func (c *ErrorController) Error500() {
	c.TplName = "error/500.html"
	err, ok := c.Data["error"].(error)
	if !ok {
		err = syserror.New("未知错误", nil)
	}
	sysError, ok := err.(syserror.SysError)
	if !ok {
		sysError = syserror.New(err.Error(), nil)
	}
	if sysError.Ression() != nil {
		beego.Info(sysError.Error(), sysError.Ression())
	}
	beego.Info(sysError.Error())
	beego.Info(c.IsAjax)
	if c.IsAjax() {

		c.jerror(sysError)
	} else {

		c.Data["comtent"] = sysError.Error()
	}

}

func (c *ErrorController) jerror(err syserror.SysError) {
	c.Ctx.Output.Status = 200
	c.Data["json"] = map[string]interface{}{
		"code": err.Code(),
		"msg":  err.Error(),
	}
	c.ServeJSON()
}
