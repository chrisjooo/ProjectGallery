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
	ID          int64 `orm:"PK"`
	Name        string
	Rating      int
	AuthorID    int64
	Description string
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time `orm:"auto_now;type(datetime)"`
}

func AddProject(u Project) (*Project, error) {
	log.Print(u)

	//ORM database
	o := orm.NewOrm()

	//checking author already member
	var acc Account
	log.Print("checking account")
	err := o.Raw("SELECT * FROM account WHERE username = ?", u.AuthorID).QueryRow(&acc)
	if err != nil {
		return nil, errors.New("acc is not a member")
	}

	log.Print("entering insertion", u)
	_, err = o.Insert(&u)
	if err == nil {
		//successfully inserted
		log.Print("success")
		return &u, nil
	} else {
		log.Print("error", err)
		return nil, errors.New("error inserting vote")
	}

}

func GetProjects(name string) map[string]*Project {
	o := orm.NewOrm()
	ProjectList := make(map[string]*Project)
	var Projects []*Project

	_, err := o.Raw("SELECT * FROM project WHERE name = %?%", name).QueryRows(&Projects)
	if err != nil {
		return nil
	}

	for _, v := range Projects {
		ProjectList[v.Name] = v
	}

	return ProjectList

}

func GetProjectByID(id int64) (*Project, error) {
	o := orm.NewOrm()
	project := Project{ID: id}

	err := o.Read(&project)
	if err != nil {
		return nil, errors.New("Account not exists")
	} else {
		return &project, nil
	}
}

func UpdateProject(id int64, uu *Project) (a *Project, err error) {
	o := orm.NewOrm()

	u, err := GetProjectByID(id)

	log.Print(*u)

	if err == nil {
		if uu.Description != "" {
			u.Description = uu.Description
		}
		if uu.Name != "" {
			u.Name = uu.Name
		}
		if uu.Rating != 0 {
			u.Rating = uu.Rating
		}
		log.Print("REACHED HERE")
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
		return nil, errors.New("Project Not Exist")
	}
}

func DeleteProject(id int64) {
	o := orm.NewOrm()
	_, err := o.Delete(&Project{ID: id})

	log.Print(err)

	if err != nil {
		log.Fatal("delete Project failed")
	}
}
