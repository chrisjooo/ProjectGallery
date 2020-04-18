package controllers

import (
	"ProjectGallery/models"
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
)

// Operations about Users
type AccountController struct {
	beego.Controller
}

// @Title Post
// @Description get all Users
// @Success 200 {object} models.Account
// @router / [post]
func (u *AccountController) Post() {
	var account models.Account
	json.Unmarshal(u.Ctx.Input.RequestBody, &account)
	log.Print(account)
	newAcc, err := models.AddAccount(account)
	if err != nil {
		u.Data["json"] = err
	} else {
		u.Data["json"] = newAcc
	}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.Account
// @router / [get]
func (u *AccountController) GetAll() {
	accounts := models.GetAllAccounts()
	u.Data["json"] = accounts
	u.ServeJSON()
}

// @Title GetByUsername
// @Description get user by username
// @Param	username		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Account
// @Failure 403 :username is empty
// @router /:username [get]
func (u *AccountController) GetByUsername() {
	username := u.GetString(":username")
	if username != "" {
		account, err := models.GetAccount(username)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = account
		}
	}
	u.ServeJSON()
}

// @Title GetByID
// @Description get user by username
// @Param	id		path 	int	true		"The key for staticblock"
// @Success 200 {object} models.Account
// @Failure 403 :id is empty
// @router /:id [get]
func (u *AccountController) GetByID() {
	id := u.GetInt64(":id")
	if id != 0 {
		account, err := models.GetAccountByID(id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = account
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the account
// @Param	username		path 	string	true		"The username you want to update"
// @Param	body		body 	models.Account	true		"body for user content"
// @Success 200 {object} models.Account
// @Failure 403 :username is null
// @router /:username [put]
func (u *AccountController) Put() {
	username := u.GetString(":username")
	if username != "" {
		var account models.Account
		json.Unmarshal(u.Ctx.Input.RequestBody, &account)
		uu, err := models.UpdateAccount(username, &account)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	username		path 	string	true		"The username you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 username is empty
// @router /:username [delete]
func (u *AccountController) Delete() {
	username := u.GetString(":username")
	models.DeleteAccount(username)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// // @Title Login
// // @Description Logs user into the system
// // @Param	username		query 	string	true		"The username for login"
// // @Param	password		query 	string	true		"The password for login"
// // @Success 200 {string} login success
// // @Failure 403 user not exist
// // @router /login [get]
// func (u *AccountController) Login() {
// 	username := u.GetString("username")
// 	password := u.GetString("password")
// 	if models.Login(username, password) {
// 		u.Data["json"] = "login success"
// 	} else {
// 		u.Data["json"] = "user not exist"
// 	}
// 	u.ServeJSON()
// }

// // @Title logout
// // @Description Logs out current logged in user session
// // @Success 200 {string} logout success
// // @router /logout [get]
// func (u *AccountController) Logout() {
// 	u.Data["json"] = "logout success"
// 	u.ServeJSON()
// }