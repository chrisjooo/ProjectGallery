package controllers

import (
	"ProjectGallery/helpers"
	"ProjectGallery/models"
	"strconv"

	"strings"

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
	project := models.Project{}
	u.ParseForm(&project)

	uu, err := models.AddProject(project)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		file, header, err := u.GetFile("project_pic")
		if file != nil {
			// get the filename
			fileName := header.Filename
			url := "./static/images/projects/"
			fileType := fileName[strings.IndexByte(fileName, '.'):]
			newFileName := url + strconv.FormatInt(uu.Id, 10) + fileType
			project.ProjectPic = newFileName
			err = u.SaveToFile("project_pic", newFileName)
			if err != nil {
				u.Data["json"] = err.Error()
			} else {
				err = helpers.CompressToPNG(newFileName)
				if err != nil {
					u.Data["json"] = err.Error()
				} else {
					//update
					uu, err = models.UpdateProject(uu.Id, &project)
					if err != nil {
						u.Data["json"] = err.Error()
					} else {
						u.Data["json"] = uu
					}
				}
			}
		}
	}
	u.ServeJSON()
}

// @Title GetProjectsByName
// @Description get all projects with specific name
// @Success 200 {object} models.Project
// @router /:name [get]
func (u *ProjectController) GetProjectsByName() {
	name := u.GetString(":name")
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
	id, err := u.GetInt64(":id")
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		if id != 0 {
			project, err := models.GetProjectById(id)
			if err != nil {
				u.Data["json"] = err.Error()
			} else {
				u.Data["json"] = project
			}
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
	id, err := u.GetInt64(":id")
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		if id != 0 {
			project := models.Project{}
			u.ParseForm(&project)

			file, header, err := u.GetFile("project_pic") // where <<this>> is the controller and <<file>> the id of your form field
			if file != nil {
				// get the filename
				fileName := header.Filename
				url := "./static/images/projects/"
				fileType := fileName[strings.IndexByte(fileName, '.'):]
				newFileName := url + strconv.FormatInt(id, 10) + fileType

				err = u.SaveToFile("project_pic", newFileName)
				if err != nil {
					u.Data["json"] = err.Error()
				} else {
					err = helpers.CompressToPNG(newFileName)
					if err != nil {
						u.Data["json"] = err.Error()
					} else {
						project.ProjectPic = newFileName
						uu, err1 := models.UpdateProject(id, &project)
						if err1 != nil {
							u.Data["json"] = err1.Error()
						} else {
							u.Data["json"] = uu
						}
					}
				}
			} else {
				uu, err := models.UpdateProject(id, &project)
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
// @Failure 403 id is empty
// @router /:id [delete]
func (u *ProjectController) Delete() {
	id, err := u.GetInt64(":id")
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		if id != 0 {
			err = models.DeleteProject(id)
			if err != nil {
				u.Data["json"] = err.Error()
			} else {
				u.Data["json"] = "delete success!"
			}
			u.ServeJSON()
		}
	}

}
