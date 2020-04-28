package models

import (
	"errors"
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Project))
}

type Project struct {
	Id          int64     `orm:"PK" json:"id"`
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
}

type ProjectList struct {
	NumProject int64      `json:"total_project"`
	Data       []*Project `json:"data"`
}

func AddProject(u Project) (*Project, error) {
	log.Print(u)
	//ORM database
	o := orm.NewOrm()

	//checking author already member
	log.Print("checking account")
	_, err := GetAccount(u.Author)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, errors.New("account is not a member")
		}
		log.Println("error :", err)
		return nil, err
	}

	log.Print("entering insertion", u)
	newId, err := o.Insert(&u)
	if err == nil {
		//successfully inserted
		log.Print("success")
		u.Id = newId
		return &u, nil
	} else {
		log.Print("error", err)
		return nil, errors.New("error inserting project")
	}

}

func GetProjects(projectName string) *ProjectList {
	o := orm.NewOrm()
	list := &ProjectList{}
	var projects []*Project
	sql := "SELECT * FROM project WHERE LOWER(name) LIKE '%" + projectName + "%'"
	log.Print("query: ", sql)
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
		return nil, errors.New("Account not exists")
	} else if err != nil {
		return nil, err
	} else {
		return &project, nil
	}
}

func UpdateProject(Id int64, uu *Project) (a *Project, err error) {
	o := orm.NewOrm()

	u, err := GetProjectById(Id)

	log.Print(*u)

	if err == nil {
		if uu.Description != "" {
			u.Description = uu.Description
		}
		if uu.Name != "" {
			u.Name = uu.Name
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

func DeleteProject(Id int64) {
	o := orm.NewOrm()
	_, err := o.Delete(&Project{Id: Id})

	log.Print(err)

	if err != nil {
		log.Fatal("delete Project failed")
	}
}
