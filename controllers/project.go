package controllers

import (
	"ProjectGallery/models"
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
)

// Operations about Users
type ProjectController struct {
	beego.Controller
}

// @Title Post
// @Description get all Users
// @Success 200 {object} models.Project
// @router / [post]
func (u *ProjectController) Post() {
	var project models.Project
	json.Unmarshal(u.Ctx.Input.RequestBody, &project)
	log.Print(project)
	newProject, err := models.AddProject(project)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = newProject
	}
	u.ServeJSON()
}

// @Title GetProjectsByName
// @Description get all projects with specific name
// @Success 200 {object} models.Project
// @router /:name [get]
func (u *ProjectController) GetProjectsByName() {
	name := u.GetString(":name")
	log.Print("\nproject name : ", name, "\n")
	if name != "" {
		projects := models.GetProjects(name)
		u.Data["json"] = projects
	}
	u.ServeJSON()
}

// @Title GetById
// @Description get user by username
// @Param	id		path 	int	true		"The key for staticblock"
// @Success 200 {object} models.Project
// @Failure 403 :id is empty
// @router /id/:id [get]
func (u *ProjectController) GetById() {
	id, _ := u.GetInt64(":id")
	if id != 0 {
		project, err := models.GetProjectById(id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = project
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the Project
// @Param	id		path 	int	true		"The id project you want to update"
// @Param	body		body 	models.Project	true		"body for user content"
// @Success 200 {object} models.Project
// @Failure 403 :id is null
// @router /:id [put]
func (u *ProjectController) Put() {
	id, _ := u.GetInt64(":id")
	if id != 0 {
		var project models.Project
		json.Unmarshal(u.Ctx.Input.RequestBody, &project)
		uu, err := models.UpdateProject(id, &project)
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
// @Failure 403 id is empty
// @router /:id [delete]
func (u *ProjectController) Delete() {
	id, _ := u.GetInt64(":id")
	if id != 0 {
		models.DeleteProject(id)
		u.Data["json"] = "delete success!"
		u.ServeJSON()
	}

}
