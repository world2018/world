package routers

import (
	"github.com/astaxie/beego"
	"world/controllers/Admin"
	"world/controllers/Home"
)

func init() {
	//---------------------------------- 前台路由 --------------------------------------
	beego.Router("/", &Home.IndexController{}, "*:Index")
	//---------------------------------- 后台路由 --------------------------------------

	nsadmin := beego.NewNamespace("admin",
		//基础操作
		beego.NSRouter("/", &Admin.IndexController{}, "*:Index"),
		beego.NSRouter("/login", &Admin.LoginController{}, "*:Index"),
		beego.NSRouter("/logout", &Admin.LoginController{}, "get:Logout"),
		beego.NSRouter("/editPwd", &Admin.IndexController{}, "get,post:EditPwd"),
		//系统管理
		beego.NSRouter("/sys/admin/?:splat", &Admin.AdminController{}, "get:Index"),
		beego.NSRouter("/sys/admin/add", &Admin.AdminController{}, "get,post:Add"),
		beego.NSRouter("/sys/admin/edit/:id", &Admin.AdminController{}, "get,post:Edit"),
		beego.NSRouter("/sys/admin/del", &Admin.AdminController{}, "post:Del"),
	)
	beego.AddNamespace(nsadmin)
}