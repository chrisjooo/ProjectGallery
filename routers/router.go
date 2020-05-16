package routers

import (
	"ProjectGallery/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/test",
			beego.NSInclude(
				&controllers.MainController{},
			),
		),
		beego.NSNamespace("/account",
			beego.NSInclude(
				&controllers.AccountController{},
			),
		),
		beego.NSNamespace("/project",
			beego.NSInclude(
				&controllers.ProjectController{},
			),
		),
		beego.NSNamespace("/vote",
			beego.NSInclude(
				&controllers.VoteController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
