package models

import (
	"ProjectGallery/helpers"
	"encoding/json"
	"errors"
	"log"
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

type ProjectList struct {
	NumProject int64      `json:"total_project"`
	Data       []*Project `json:"data"`
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

type FilteredProjectList struct {
	NumProject int64              `json:"total_project"`
	Data       []*FilteredProject `json:"data"`
}

func AddProject(u Project) (*Project, error) {
	//ORM database
	o := orm.NewOrm()

	//checking author already member
	_, err := GetAccount(u.Author)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, errors.New("account is not a member")
		}
		log.Println("error: ", err)
		return nil, err
	}

	newId, err := o.Insert(&u)
	if err == nil {
		//successfully inserted
		u.Id = newId
	} else {
		err = o.Read(&u)
		if err != nil {
			return nil, errors.New("failed insertion")
		}
	}
	return &u, nil
}

func GetProjects(projectName string) *ProjectList {
	o := orm.NewOrm()
	list := &ProjectList{}
	var projects []*Project
	sql := "SELECT * FROM project WHERE LOWER(name) LIKE '%" + projectName + "%'"
	num, err := o.Raw(sql).QueryRows(&projects)
	if err != nil {
		log.Print("error query: ", err)
		return nil
	}
	list.Data = projects
	list.NumProject = num

	return list

}

func GetProjectById(Id int64) (*Project, error) {
	o := orm.NewOrm()
	project := Project{Id: Id}

	err := o.Read(&project)
	if err == orm.ErrNoRows {
		return nil, errors.New("Project not exists")
	} else if err != nil {
		return nil, err
	} else {
		return &project, nil
	}
}

func UpdateProject(Id int64, uu *Project) (a *Project, err error) {
	o := orm.NewOrm()

	u, err := GetProjectById(Id)

	if err == nil {
		if uu.Author != u.Author {
			return nil, errors.New("not matching author")
		}
		if uu.Description != "" {
			u.Description = uu.Description
		}
		if uu.Name != "" {
			u.Name = uu.Name
		}
		if uu.ProjectPic != "" {
			u.ProjectPic = uu.ProjectPic
		}
		// ORM Update
		_, err1 := o.Update(u)
		log.Print(u, err)

		if err1 == nil {
			//update successful
			return u, nil
		} else {
			return nil, err1
		}
	} else {
		if err == orm.ErrNoRows {
			return nil, errors.New("Project Not Exist")
		}
		return nil, err
	}
}

func DeleteProject(Id int64) error {
	o := orm.NewOrm()

	_, err := GetProjectById(Id)
	if err != nil {
		if err == orm.ErrNoRows {
			return errors.New("Project Not Exist")
		}
		return err
	}

	_, err = o.Delete(&Project{Id: Id})

	log.Print(err)

	if err != nil {
		log.Fatal("delete Project failed")
		return err
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
	list.Data = projects
	list.NumProject = num

	log.Printf("%v\n", list.Data)
	log.Printf("%v\n", list.NumProject)

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
		err = json.Unmarshal(v, data)
		return data
	}
	//get from DB
	data = FilterMostLikeProject()
	//set cache
	_, err = conn.Do("HSET", "filtered-data", &data)
	if err != nil {
		log.Printf("error setting cache from model: %v", err)
	}

	return data
}
