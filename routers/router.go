package routers

import (
	"MyBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(&controllers.IndexControllers{})
	beego.Include(&controllers.UserController{})
	beego.AddNamespace(beego.NewNamespace("note", beego.NSInclude(&controllers.NoteControllers{})))
	beego.Include(&controllers.CommentController{})
	//beego.Include(&controllers.PraiseControllers{})
	beego.Router("/praise/:type/:key", &controllers.PraiseControllers{}, "post:Praise")
}
