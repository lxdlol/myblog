package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["MyBlog/controllers:CommentController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:CommentController"],
		beego.ControllerComments{
			Method: "GetCount",
			Router: `/comment/count`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MyBlog/controllers:CommentController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:CommentController"],
		beego.ControllerComments{
			Method: "PostNewComment",
			Router: `/comment/new/?:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MyBlog/controllers:CommentController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:CommentController"],
		beego.ControllerComments{
			Method: "GetQuery",
			Router: `/comment/query`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MyBlog/controllers:IndexControllers"] = append(beego.GlobalControllerRouter["MyBlog/controllers:IndexControllers"],
		beego.ControllerComments{
			Method: "GetIndex",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MyBlog/controllers:IndexControllers"] = append(beego.GlobalControllerRouter["MyBlog/controllers:IndexControllers"],
		beego.ControllerComments{
			Method: "GetAbout",
			Router: `/about`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MyBlog/controllers:IndexControllers"] = append(beego.GlobalControllerRouter["MyBlog/controllers:IndexControllers"],
		beego.ControllerComments{
			Method: "GetComment",
			Router: `/comment/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MyBlog/controllers:IndexControllers"] = append(beego.GlobalControllerRouter["MyBlog/controllers:IndexControllers"],
		beego.ControllerComments{
			Method: "GetDetails",
			Router: `/details/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MyBlog/controllers:IndexControllers"] = append(beego.GlobalControllerRouter["MyBlog/controllers:IndexControllers"],
		beego.ControllerComments{
			Method: "GetLogin",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MyBlog/controllers:IndexControllers"] = append(beego.GlobalControllerRouter["MyBlog/controllers:IndexControllers"],
		beego.ControllerComments{
			Method: "GetMessage",
			Router: `/message`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MyBlog/controllers:NoteControllers"] = append(beego.GlobalControllerRouter["MyBlog/controllers:NoteControllers"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MyBlog/controllers:NoteControllers"] = append(beego.GlobalControllerRouter["MyBlog/controllers:NoteControllers"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/edit/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MyBlog/controllers:NoteControllers"] = append(beego.GlobalControllerRouter["MyBlog/controllers:NoteControllers"],
		beego.ControllerComments{
			Method: "GetNew",
			Router: `/new`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MyBlog/controllers:NoteControllers"] = append(beego.GlobalControllerRouter["MyBlog/controllers:NoteControllers"],
		beego.ControllerComments{
			Method: "PostNew",
			Router: `/new/:key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MyBlog/controllers:UserController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:UserController"],
		beego.ControllerComments{
			Method: "AddUser",
			Router: `/addUser`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MyBlog/controllers:UserController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAddUser",
			Router: `/addUser`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MyBlog/controllers:UserController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:UserController"],
		beego.ControllerComments{
			Method: "PostLogin",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["MyBlog/controllers:UserController"] = append(beego.GlobalControllerRouter["MyBlog/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
