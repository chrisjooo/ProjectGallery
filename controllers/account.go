package controllers

import (
	"ProjectGallery/helpers"
	"ProjectGallery/models"
	"ProjectGallery/validations"
	"encoding/json"
	"strings"

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
	validationErr := validations.AccountValidation(&account)
	if validationErr == nil {
		newAcc, err := models.AddAccount(account)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = newAcc
		}
	} else {
		u.Data["json"] = validationErr.Error()
	}

	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.AccList
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

		account := models.Account{}

		u.ParseForm(&account)

		validationErr := validations.AccountValidation(&account)
		if validationErr != nil {
			u.Data["json"] = validationErr.Error()
		} else {
			file, header, err := u.GetFile("profile_pic") // where <<this>> is the controller and <<file>> the id of your form field
			if file != nil {
				// get the filename
				fileName := header.Filename
				url := "./static/images/accounts/"

				fileType := fileName[strings.IndexByte(fileName, '.'):]
				newFileName := url + username + fileType
				err = u.SaveToFile("profile_pic", newFileName)
				if err != nil {
					u.Data["json"] = err.Error()
				} else {
					//helper function
					err = helpers.CompressToPNG(newFileName)
					if err != nil {
						u.Data["json"] = err.Error()
					} else {
						account.ProfilePic = newFileName

						uu, err1 := models.UpdateAccount(username, &account)
						if err1 != nil {
							u.Data["json"] = err1.Error()
						} else {
							u.Data["json"] = uu
						}
					}
				}
			} else {
				uu, err := models.UpdateAccount(username, &account)
				if err != nil {
					u.Data["json"] = err.Error()
				} else {
					u.Data["json"] = uu
				}
			}
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
	err := models.DeleteAccount(username)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = "delete success!"
	}
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *AccountController) Login() {
	var account models.Account
	json.Unmarshal(u.Ctx.Input.RequestBody, &account)
	check, err := models.Login(account.Username, account.Password)

	if check {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = err.Error()
	}
	u.ServeJSON()
}

// // @Title logout
// // @Description Logs out current logged in user session
// // @Success 200 {string} logout success
// // @router /logout [get]
// func (u *AccountController) Logout() {
// 	u.Data["json"] = "logout success"
// 	u.ServeJSON()
// }
