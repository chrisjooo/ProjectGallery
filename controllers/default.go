package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// func (c *MainController) Get() {
// 	c.Data["Website"] = "beego.me"
// 	c.Data["Email"] = "astaxie@gmail.com"
// 	c.TplName = "index.tpl"
// }

// @Title Ping
// @Description get all Users
// @Success 200 {string} login success
// @router /ping [get]
func (u *MainController) Ping() {
	u.Data["json"] = "Pong!"
	u.ServeJSON()
}