package controllers

import (
	"ProjectGallery/helpers"
	"ProjectGallery/models"
	"ProjectGallery/validations"
	"errors"
	"log"
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
	tokenAuth, err := helpers.ExtractTokenMetadata(u.Ctx)
	if err != nil {
		errCode := helpers.ErrorCode(err.Error())
		u.Ctx.ResponseWriter.WriteHeader(errCode)
		u.Data["json"] = err.Error()
		u.ServeJSON()
		return
	}
	err = helpers.FetchAuth(tokenAuth)
	if err != nil {
		errCode := helpers.ErrorCode(err.Error())
		u.Ctx.ResponseWriter.WriteHeader(errCode)
		u.Data["json"] = err.Error()
		u.ServeJSON()
		return
	}
	project := models.Project{}
	u.ParseForm(&project)

	if project.Author != tokenAuth.Username {
		err = errors.New("Unauthorized")
		errCode := helpers.ErrorCode(err.Error())
		u.Ctx.ResponseWriter.WriteHeader(errCode)
		u.Data["json"] = err.Error()
		u.ServeJSON()
		return
	}

	validationErr := validations.ProjectValidation(&project)
	if validationErr != nil {
		errCode := helpers.ErrorCode(validationErr.Error())
		u.Ctx.ResponseWriter.WriteHeader(errCode)
		u.Data["json"] = validationErr.Error()
	} else {
		uu, err := models.AddProject(project)
		if err != nil {
			errCode := helpers.ErrorCode(err.Error())
			u.Ctx.ResponseWriter.WriteHeader(errCode)
			u.Data["json"] = err.Error()
		} else {
			file, header, err := u.GetFile("project_pic")
			if file != nil {
				// get the filename
				fileName := header.Filename
				url := "./static/images/projects/"
				fileType := fileName[strings.IndexByte(fileName, '.'):]
				newFileName := url + strconv.FormatInt(uu.Project.Id, 10) + fileType
				domain := beego.AppConfig.String("domain")
				log.Printf("domain: %v\n", domain)
				urlFileName := domain + "/static/images/accounts/" + strconv.FormatInt(uu.Project.Id, 10) + fileType
				project.ProjectPic = urlFileName
				err = u.SaveToFile("project_pic", newFileName)
				if err != nil {
					errCode := helpers.ErrorCode(err.Error())
					u.Ctx.ResponseWriter.WriteHeader(errCode)
					u.Data["json"] = err.Error()
				} else {
					err = helpers.CompressToPNG(newFileName)
					if err != nil {
						errCode := helpers.ErrorCode(err.Error())
						u.Ctx.ResponseWriter.WriteHeader(errCode)
						u.Data["json"] = err.Error()
					} else {
						//update
						log.Printf("Project: %v\n", project)
						uu, err = models.UpdateProject(uu.Project.Id, &project)
						if err != nil {
							errCode := helpers.ErrorCode(err.Error())
							u.Ctx.ResponseWriter.WriteHeader(errCode)
							u.Data["json"] = err.Error()
						} else {
							log.Printf("response: %v", uu)
							u.Data["json"] = uu
						}
					}
				}
			} else {
				u.Data["json"] = uu
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
		errCode := helpers.ErrorCode(err.Error())
		u.Ctx.ResponseWriter.WriteHeader(errCode)
		u.Data["json"] = err.Error()
	} else {
		if id != 0 {
			project, err := models.GetProjectById(id)
			if err != nil {
				errCode := helpers.ErrorCode(err.Error())
				u.Ctx.ResponseWriter.WriteHeader(errCode)
				u.Data["json"] = err.Error()
			} else {
				u.Data["json"] = project
			}
		}
	}
	u.ServeJSON()
}

// @Title GetLikeProjects
// @Description get all projects filtered by likes
// @Success 200 {object} models.Project
// @router /filter/like [get]
func (u *ProjectController) GetLikeProjects() {
	projects := models.GetMostLikeProject()
	u.Data["json"] = projects

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
	tokenAuth, err := helpers.ExtractTokenMetadata(u.Ctx)
	if err != nil {
		errCode := helpers.ErrorCode(err.Error())
		u.Ctx.ResponseWriter.WriteHeader(errCode)
		u.Data["json"] = err.Error()
		u.ServeJSON()
		return
	}
	err = helpers.FetchAuth(tokenAuth)
	if err != nil {
		errCode := helpers.ErrorCode(err.Error())
		u.Ctx.ResponseWriter.WriteHeader(errCode)
		u.Data["json"] = err.Error()
		u.ServeJSON()
		return
	}
	id, err := u.GetInt64(":id")
	if err != nil {
		errCode := helpers.ErrorCode(err.Error())
		u.Ctx.ResponseWriter.WriteHeader(errCode)
		u.Data["json"] = err.Error()
	} else {
		if id != 0 {
			project := models.Project{}
			u.ParseForm(&project)
			log.Printf("controller update: %v", project)

			if project.Author != tokenAuth.Username {
				err = errors.New("Unauthorized")
				errCode := helpers.ErrorCode(err.Error())
				u.Ctx.ResponseWriter.WriteHeader(errCode)
				u.Data["json"] = err.Error()
				u.ServeJSON()
				return
			}

			file, header, err := u.GetFile("project_pic") // where <<this>> is the controller and <<file>> the id of your form field
			if file != nil {
				// get the filename
				fileName := header.Filename
				url := "./static/images/projects/"
				fileType := fileName[strings.IndexByte(fileName, '.'):]
				newFileName := url + strconv.FormatInt(id, 10) + fileType

				err = u.SaveToFile("project_pic", newFileName)
				if err != nil {
					errCode := helpers.ErrorCode(err.Error())
					u.Ctx.ResponseWriter.WriteHeader(errCode)
					u.Data["json"] = err.Error()
				} else {
					err = helpers.CompressToPNG(newFileName)
					if err != nil {
						errCode := helpers.ErrorCode(err.Error())
						u.Ctx.ResponseWriter.WriteHeader(errCode)
						u.Data["json"] = err.Error()
					} else {
						domain := beego.AppConfig.String("domain")
						log.Printf("domain: %v\n", domain)
						urlFileName := domain + "/static/images/accounts/" + strconv.FormatInt(id, 10) + fileType
						project.ProjectPic = urlFileName
						uu, err1 := models.UpdateProject(id, &project)
						if err1 != nil {
							errCode := helpers.ErrorCode(err1.Error())
							u.Ctx.ResponseWriter.WriteHeader(errCode)
							u.Data["json"] = err1.Error()
						} else {
							u.Data["json"] = uu
						}
					}
				}
			} else {
				log.Printf("harusnya masuk sini: %v", project)
				uu, err := models.UpdateProject(id, &project)
				if err != nil {
					errCode := helpers.ErrorCode(err.Error())
					u.Ctx.ResponseWriter.WriteHeader(errCode)
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
	tokenAuth, err := helpers.ExtractTokenMetadata(u.Ctx)
	if err != nil {
		errCode := helpers.ErrorCode(err.Error())
		u.Ctx.ResponseWriter.WriteHeader(errCode)
		u.Data["json"] = err.Error()
		u.ServeJSON()
		return
	}
	err = helpers.FetchAuth(tokenAuth)
	if err != nil {
		errCode := helpers.ErrorCode(err.Error())
		u.Ctx.ResponseWriter.WriteHeader(errCode)
		u.Data["json"] = err.Error()
		u.ServeJSON()
		return
	}
	id, err := u.GetInt64(":id")
	if err != nil {
		errCode := helpers.ErrorCode(err.Error())
		u.Ctx.ResponseWriter.WriteHeader(errCode)
		u.Data["json"] = err.Error()
	} else {
		if id != 0 {
			project, err := models.GetProjectById(id)
			if project.Project.Author != tokenAuth.Username {
				err1 := errors.New("Unauthorized")
				errCode := helpers.ErrorCode(err1.Error())
				u.Ctx.ResponseWriter.WriteHeader(errCode)
				u.Data["json"] = err1.Error()
				u.ServeJSON()
				return
			}
			err = models.DeleteProject(id)
			if err != nil {
				errCode := helpers.ErrorCode(err.Error())
				u.Ctx.ResponseWriter.WriteHeader(errCode)
				u.Data["json"] = err.Error()
			} else {
				u.Data["json"] = "delete success!"
			}
			u.ServeJSON()
		}
	}

}
