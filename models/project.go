package models

import (
	"ProjectGallery/helpers"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"
)

func init() {
	orm.RegisterModel(new(Project))
}

type Project struct {
	Id          int64     `orm:"PK" json:"id" form:"-"`
	Name        string    `json:"name" form:"name"`
	Author      string    `json:"author" form:"author"`
	ProjectPic  string    `json:"project_pic" form:"project_pic"`
	Description string    `json:"description" form:"description"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
}

type ProjectData struct {
	Project       Project `json:"project_data"`
	CompressedPic string  `json:"compressed_image"`
}

type ProjectList struct {
	NumProject int64          `json:"total_project"`
	Data       []*ProjectData `json:"data"`
}

type FilteredProject struct {
	Id          int64     `orm:"PK" json:"id" form:"-"`
	Name        string    `json:"name" form:"name"`
	Author      string    `json:"author" form:"author"`
	ProjectPic  string    `json:"project_pic" form:"project_pic"`
	Description string    `json:"description" form:"description"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	TotalLike   int64     `json:"total_like"`
}

type FilteredProjectData struct {
	Project       FilteredProject `json:"project_data"`
	CompressedPic string          `json:"compressed_image"`
}

type FilteredProjectList struct {
	NumProject int64                  `json:"total_project"`
	Data       []*FilteredProjectData `json:"data"`
}

func AddProject(u Project) (*ProjectData, error) {
	//ORM database
	o := orm.NewOrm()

	//checking author already member
	_, err := GetAccount(u.Author)
	if err != nil {
		return nil, err
	}

	newId, err := o.Insert(&u)
	if err == nil {
		//successfully inserted
		u.Id = newId
	} else {
		err = o.Read(&u)
		if err != nil {
			errMessage := helpers.ErrorMessage(helpers.Post)
			return nil, errMessage
		}
	}
	resp := &ProjectData{}
	resp.Project = u
	log.Printf("inside model project: %v", resp.Project.ProjectPic)
	if u.ProjectPic != "" {
		url := u.ProjectPic[:strings.LastIndexByte(u.ProjectPic, '.')] + "-compressed.png"
		resp.CompressedPic = url
	} else {
		resp.CompressedPic = ""
	}
	return resp, nil
}

func GetProjects(projectName string) *FilteredProjectList {
	o := orm.NewOrm()
	list := &FilteredProjectList{}
	var projects []*FilteredProject
	sql := "SELECT project.id, project.name, project.author, project.project_pic, project.description, project.created_at, (SELECT COUNT(vote.id) FROM vote WHERE vote.vote = 1 AND vote.project_id = project.id) as total_like FROM project WHERE LOWER(name) LIKE '%" + projectName + "%';"

	num, err := o.Raw(sql).QueryRows(&projects)
	if err != nil {
		log.Print("error query: ", err)
		return nil
	}

	var projectData []*FilteredProjectData

	for _, v := range projects {
		u := &FilteredProjectData{}
		u.Project = *v
		if v.ProjectPic != "" {
			url := v.ProjectPic[:strings.LastIndexByte(v.ProjectPic, '.')] + "-compressed.png"
			u.CompressedPic = url
		} else {
			u.CompressedPic = ""
		}
		projectData = append(projectData, u)
	}

	list.Data = projectData
	list.NumProject = num

	return list

}

func GetProjectById(Id int64) (*FilteredProjectData, error) {
	o := orm.NewOrm()
	project := Project{Id: Id}

	err := o.Read(&project)
	log.Printf("model GetProjectByID project: %v", project)
	if err == orm.ErrNoRows {
		errMessage := helpers.ErrorMessage(helpers.CheckProject)
		return nil, errMessage
	} else if err != nil {
		errMessage := helpers.ErrorMessage(helpers.Get)
		return nil, errMessage
	} else {
		u := &FilteredProjectData{}
		temp := FilteredProject{}
		temp.Id = project.Id
		temp.Name = project.Name
		temp.Author = project.Author
		temp.ProjectPic = project.ProjectPic
		temp.Description = project.Description
		temp.CreatedAt = project.CreatedAt

		var total int

		err = o.Raw("SELECT COUNT(id) FROM vote WHERE project_id = ? AND vote = 1", project.Id).QueryRow(&total)
		if err != nil {
			log.Printf("err: %v", err)
			errMessage := helpers.ErrorMessage(helpers.QueryError)
			return nil, errMessage
		}

		temp.TotalLike = int64(total)
		u.Project = temp
		if project.ProjectPic != "" {
			url := project.ProjectPic[:strings.LastIndexByte(project.ProjectPic, '.')] + "-compressed.png"
			u.CompressedPic = url
		} else {
			u.CompressedPic = ""
		}
		return u, nil
	}
}

func UpdateProject(Id int64, uu *Project) (a *ProjectData, err error) {
	o := orm.NewOrm()

	u, err := GetProjectById(Id)
	a = &ProjectData{}
	project := Project{}
	project.Id = u.Project.Id
	project.Name = u.Project.Name
	project.Author = u.Project.Author
	project.ProjectPic = u.Project.ProjectPic
	project.Description = u.Project.Description
	project.CreatedAt = u.Project.CreatedAt

	if err == nil {
		if uu.Author != project.Author {
			return nil, errors.New("not matching author")
		}
		if uu.Description != "" {
			project.Description = uu.Description
		}
		if uu.Name != "" {
			project.Name = uu.Name
		}
		if uu.ProjectPic != "" {
			project.ProjectPic = uu.ProjectPic
			url := uu.ProjectPic[:strings.LastIndexByte(uu.ProjectPic, '.')] + "-compressed.png"
			a.CompressedPic = url
		} else {
			project.ProjectPic = uu.ProjectPic
			a.CompressedPic = ""
		}

		a.Project = project

		log.Printf("Updating project: %v", project)
		// ORM Update
		_, err1 := o.Update(&project)
		log.Print(u, err)

		if err1 == nil {
			//update successful
			return a, nil
		} else {
			errMessage := helpers.ErrorMessage(helpers.Put)
			return nil, errMessage
		}
	} else {
		return nil, err
	}
}

func DeleteProject(Id int64) error {
	o := orm.NewOrm()

	_, err := GetProjectById(Id)
	if err != nil {
		return err
	}

	_, err = o.Delete(&Project{Id: Id})

	log.Print(err)

	if err != nil {
		log.Println("delete Project failed")
		errMessage := helpers.ErrorMessage(helpers.Delete)
		return errMessage
	}
	return nil
}

func FilterMostLikeProject() *FilteredProjectList {
	o := orm.NewOrm()
	list := &FilteredProjectList{}
	var projects []*FilteredProject
	sql := "SELECT project.id, project.name, project.author, project.project_pic, project.description, project.created_at, (SELECT COUNT(vote.id) FROM vote WHERE vote.vote = 1 AND vote.project_id = project.id) as total_like FROM project ORDER BY total_like DESC;"
	num, err := o.Raw(sql).QueryRows(&projects)
	if err != nil {
		log.Print("error query: ", err)
		return nil
	}

	var projectData []*FilteredProjectData

	for _, v := range projects {
		u := &FilteredProjectData{}
		u.Project = *v
		if v.ProjectPic != "" {
			url := v.ProjectPic[:strings.LastIndexByte(v.ProjectPic, '.')] + "-compressed.png"
			u.CompressedPic = url
		} else {
			u.CompressedPic = ""
		}
		projectData = append(projectData, u)
	}

	list.Data = projectData
	list.NumProject = num
	
	return list
}

func GetMostLikeProject() *FilteredProjectList {
	//get from cache
	data := &FilteredProjectList{}
	conn := helpers.NewPool().Get()
	defer conn.Close()

	v, err := redis.Bytes(conn.Do("HGET", "filtered-data", "data"))
	if err != nil {
		log.Printf("Error getting cache: %v\n", err)
	} else {
		if err := json.Unmarshal(v, data); err != nil {
			log.Printf("err GetMostLikeProject: %v", err)
		}
		if data != nil {
			log.Printf("getting from redis\n")
			return data
		}
	}
	//get from DB
	data = FilterMostLikeProject()

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("error marshaling data")
	}

	//set cache
	_, err = conn.Do("HSET", "filtered-data", "data", jsonData)
	if err != nil {
		log.Printf("error setting cache from model: %v", err)
	}

	return data
}
