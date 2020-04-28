package models

import (
	"errors"
	"log"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Vote))
}

type Vote struct {
	Id        int64  `orm:"PK" json:"id"`
	Author    string `json:"author"`
	ProjectId int64  `json:"project_id"`
	Vote      bool   `json:"isLiked"`
}

func AddVote(u Vote) (*Vote, error) {
	log.Print(u)

	_, err := GetProjectById(u.ProjectId)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, errors.New("Project not found")
		}
		return nil, err
	}
	_, err = GetAccount(u.Author)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, errors.New("Account not found")
		}
		return nil, err
	}

	//ORM database
	o := orm.NewOrm()
	temp := Vote{Author: u.Author, ProjectId: u.ProjectId}
	err = o.Read(&temp, "Author", "ProjectId")
	if err != nil && err != orm.ErrNoRows {
		return nil, err
	}

	newId, err := o.Insert(&u)
	if err == nil {
		//successfully inserted
		u.Id = newId
		return &u, nil
	} else {
		err = o.Read(&u)
		if err != nil {
			return nil, err
		}
	}

	return &u, err
}

func GetVote(author string, projectId int64) (u *Vote, err error) {
	//ORM
	o := orm.NewOrm()
	vote := Vote{Author: author, ProjectId: projectId}
	err = o.Read(&vote, "author", "projectId")
	if err != nil {
		return nil, errors.New("Vote not exists")
	} else {
		return &vote, nil
	}

}

func GetTotalVote(projectID int64) int64 {

	o := orm.NewOrm()
	var votes []*Vote

	nums, err := o.Raw("SELECT * FROM vote WHERE projectID = ? AND vote = 1", projectID).QueryRows(&votes)
	if err != nil {
		return 0
	}

	return int64(nums)
}

func UpdateVote(author string, projectId int64, uu *Vote) (u *Vote, err error) {
	o := orm.NewOrm()

	u, err = GetVote(author, projectId)

	log.Print(*u)

	if err == nil {
		if uu.Vote != u.Vote {
			u.Vote = uu.Vote
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
		if err == orm.ErrNoRows {
			return nil, errors.New("Vote Not Exist")
		} else {
			return nil, err
		}
	}
}

func DeleteVote(author string, projectId int64) {
	o := orm.NewOrm()
	_, err := o.Delete(&Vote{Author: author, ProjectId: projectId}, "author", "projectId")

	if err != nil {
		log.Fatal("delete Vote failed")
	}
}
