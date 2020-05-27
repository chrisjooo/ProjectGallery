package routers

import (
	"ProjectGallery/controllers"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/astaxie/beego/plugins/cors"
)

func init() {
	log.Printf("routers initialized")

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))

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
	log.Printf("routers done initialized")
}
