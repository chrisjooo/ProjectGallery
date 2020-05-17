package controllers

import (
	"ProjectGallery/models"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about Users
type VoteController struct {
	beego.Controller
}

// @Title Post
// @Description get all Users
// @Success 200 {object} models.Vote
// @router / [post]
func (u *VoteController) Post() {
	var rating models.Vote
	json.Unmarshal(u.Ctx.Input.RequestBody, &rating)
	newrating, err := models.AddVote(rating)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = newrating
	}
	u.ServeJSON()
}

// @Title GetProjectVote
// @Description get rating of project
// @Success 200 {object} models.Vote
// @router /:projectId [get]
func (u *VoteController) GetProjectVote() {
	projectId, err := u.GetInt64(":projectId")

	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		err, totalLike := models.GetTotalVote(projectId)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = totalLike
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the rating
// @Param	body		body 	models.Vote	true		"body for user content"
// @Success 200 {object} models.Vote
// @Failure 403 body is null
// @router / [put]
func (u *VoteController) Put() {
	var rating models.Vote
	json.Unmarshal(u.Ctx.Input.RequestBody, &rating)

	author, projectId := rating.Author, rating.ProjectId
	if author != "" && projectId != 0 {
		uu, err := models.UpdateVote(&rating)
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
// @Param	author		path 	string	true		"The username you want to delete"
// @Param	project_id		path 	string	true		"The projectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 parameter is invalid
// @router / [delete]
func (u *VoteController) Delete() {
	rating := u.Ctx.Request.URL.Query()

	author := rating["author"][0]
	projectId := ""
	if _, ok := rating["projectId"]; ok {
		projectId = rating["projectId"][0]
	}

	proId, err := strconv.ParseInt(projectId, 10, 64)
	if err != nil {
		u.Data["json"] = err.Error()
	}

	if author != "" && projectId != "" && err == nil {
		err = models.DeleteVote(author, proId)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = "delete success!"
		}
	} else {
		u.Data["json"] = err.Error()
	}
	u.ServeJSON()
}
