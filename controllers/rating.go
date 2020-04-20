package controllers

import (
	"ProjectGallery/models"
	"encoding/json"
	"log"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about Users
type RatingController struct {
	beego.Controller
}

// @Title Post
// @Description get all Users
// @Success 200 {object} models.Rating
// @router / [post]
func (u *RatingController) Post() {
	var rating models.Rating
	json.Unmarshal(u.Ctx.Input.RequestBody, &rating)
	log.Print(rating)
	newrating, err := models.AddRating(rating)
	if err != nil {
		u.Data["json"] = err
	} else {
		u.Data["json"] = newrating
	}
	u.ServeJSON()
}

// @Title GetProjectRating
// @Description get rating of project
// @Success 200 {object} models.Rating
// @router / [get]
func (u *RatingController) GetProjectRating() {

	rating := u.Ctx.Request.URL.Query()
	log.Print(rating)

	authorID := rating["authorID"][0]
	projectID := ""
	if _, ok := rating["projectID"]; ok {
		projectID = rating["projectID"][0]
	}

	authID, err := strconv.ParseInt(authorID, 10, 64)
	if err != nil {
		u.Data["json"] = err.Error()
	}
	proID, err := strconv.ParseInt(projectID, 10, 64)
	if err != nil {
		u.Data["json"] = err.Error()
	}

	log.Print("authorID ", authID, " projectID", proID)
	if authorID == "" && projectID != "" && err == nil {
		//get average project rating
		avgRating := models.GetAverageRating(proID)
		u.Data["json"] = avgRating
	} else if authorID != "" && projectID != "" && err == nil {
		//get specific rating from authorID to projectID
		rate, _ := models.GetRating(authID, proID)
		u.Data["json"] = rate
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the rating
// @Param	body		body 	models.Rating	true		"body for user content"
// @Success 200 {object} models.Rating
// @Failure 403 body is null
// @router / [put]
func (u *RatingController) Put() {
	var rating models.Rating
	json.Unmarshal(u.Ctx.Input.RequestBody, &rating)

	authorID, projectID := rating.AuthorID, rating.ProjectID

	log.Print("authorID ", authorID, " projectID ", projectID)

	if authorID != 0 && projectID != 0 {
		uu, err := models.UpdateRating(authorID, projectID, &rating)
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
// @Failure 403 parameter is invalid
// @router / [delete]
func (u *RatingController) Delete() {
	rating := u.Ctx.Request.URL.Query()
	log.Print(rating)

	authorID := rating["authorID"][0]
	projectID := ""
	if _, ok := rating["projectID"]; ok {
		projectID = rating["projectID"][0]
	}

	authID, err := strconv.ParseInt(authorID, 10, 64)
	if err != nil {
		u.Data["json"] = err.Error()
	}
	proID, err := strconv.ParseInt(projectID, 10, 64)
	if err != nil {
		u.Data["json"] = err.Error()
	}

	log.Print("authorID ", authorID, " projectID", projectID)
	if authorID != "" && projectID != "" && err == nil {
		models.DeleteRating(authID, proID)
		u.Data["json"] = "delete success!"
	} else {
		u.Data["json"] = err.Error()
	}
	u.ServeJSON()
}
