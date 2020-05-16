package controllers

import (
	"ProjectGallery/models"
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
	project := models.Project{}
	u.ParseForm(&project)

	log.Printf("project: %v", project)

	file, header, err := u.GetFile("project_pic") // where <<this>> is the controller and <<file>> the id of your form field
	log.Printf("\nGoing through err: %v", err)
	if file != nil {
		// get the filename
		fileName := header.Filename
		log.Printf("\nfilename: %v", fileName)
		url := "./static/images/projects/" + fileName
		project.ProjectPic = url
		uu, err := models.AddProject(project)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			err = u.SaveToFile("project_pic", url)
			if err != nil {
				u.Data["json"] = err.Error()
			} else {
				u.Data["json"] = uu
			}
		}
	} else {
		uu, err := models.AddProject(project)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
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

			log.Printf("project: %v\n", project)

			file, header, err := u.GetFile("project_pic") // where <<this>> is the controller and <<file>> the id of your form field
			log.Printf("\nGoing through err: %v\n\n header: %v", err, header)
			if file != nil {
				// get the filename
				fileName := header.Filename
				log.Printf("\nfilename: %v", fileName)
				url := "./static/images/projects/" + fileName
				project.ProjectPic = url
				uu, err := models.UpdateProject(id, &project)
				if err != nil {
					u.Data["json"] = err.Error()
				} else {
					err = u.SaveToFile("project_pic", url)
					if err != nil {
						u.Data["json"] = err.Error()
					} else {
						u.Data["json"] = uu
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
			models.DeleteProject(id)
			u.Data["json"] = "delete success!"
			u.ServeJSON()
		}
	}

}
