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
				&controllers.MainController{},
			),
		),
		beego.NSNamespace("/project",
			beego.NSInclude(
				&controllers.MainController{},
			),
		),
		beego.NSNamespace("/rating",
			beego.NSInclude(
				&controllers.MainController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
