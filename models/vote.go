package models

import (
	"ProjectGallery/helpers"
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

func AddVote(user string, u Vote) (*Vote, error) {
	//ORM database
	o := orm.NewOrm()

	_, err := GetProjectById(user, u.ProjectId)
	if err != nil {
		return nil, err
	}
	_, err = GetAccount(u.Author)
	if err != nil {
		return nil, err
	}

	_, err = GetVote(u.Author, u.ProjectId)
	errMessage := helpers.ErrorMessage(helpers.CheckVote)

	if err != nil || err == errMessage {
		newId, err := o.Insert(&u)
		if err == nil {
			//successfully inserted
			u.Id = newId
			return &u, nil
		} else {
			err = o.Read(&u)
			if err != nil {
				errMessage := helpers.ErrorMessage(helpers.Post)
				return nil, errMessage
			}
		}
	} else {
		errMessage := helpers.ErrorMessage(helpers.VoteExist)
		return nil, errMessage
	}
	return &u, err
}

func GetVote(author string, projectId int64) (u *Vote, err error) {
	//ORM
	o := orm.NewOrm()
	vote := Vote{Author: author, ProjectId: projectId}
	err = o.Read(&vote, "author", "projectId")
	if err != nil {
		if err == orm.ErrNoRows {
			errMessage := helpers.ErrorMessage(helpers.CheckVote)
			return nil, errMessage
		}
		errMessage := helpers.ErrorMessage(helpers.Get)
		return nil, errMessage
	} else {
		return &vote, nil
	}

}

func GetTotalVote(user string, projectID int64) (error, int64) {

	o := orm.NewOrm()
	_, err := GetProjectById(user, projectID)
	if err != nil {
		return err, -1
	}

	var total int

	err = o.Raw("SELECT COUNT(id) FROM vote WHERE project_id = ? AND vote = 1", projectID).QueryRow(&total)
	if err != nil {
		log.Printf("err: %v", err)
		errMessage := helpers.ErrorMessage(helpers.QueryError)
		return errMessage, -1
	}

	return nil, int64(total)
}

func UpdateVote(uu *Vote) (u *Vote, err error) {
	o := orm.NewOrm()

	u, err = GetVote(uu.Author, uu.ProjectId)

	if err == nil {
		if uu.Vote != u.Vote {
			u.Vote = uu.Vote
		}
		// ORM Update
		_, err1 := o.Update(u)
		log.Print(u, err)

		if err1 == nil {
			//update successful
			return u, nil
		} else {
			errMessage := helpers.ErrorMessage(helpers.Put)
			return nil, errMessage
		}
	} else {
		return nil, err
	}
}

func DeleteVote(author string, projectId int64) error {
	o := orm.NewOrm()

	_, err := GetVote(author, projectId)
	if err != nil {
		return err
	}

	_, err = o.Delete(&Vote{Author: author, ProjectId: projectId}, "author", "projectId")

	if err != nil {
		log.Println("delete Vote failed")
		errMessage := helpers.ErrorMessage(helpers.Delete)
		return errMessage
	}
	return nil
}
