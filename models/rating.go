package models

import (
	"errors"
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Rating))
}

type Rating struct {
	ID        int64 `orm:"PK"`
	AuthorID  int64
	ProjectID int64
	Rating    int64
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func AddRating(u Rating) (*Rating, error) {
	log.Print(u)

	project, err := GetProjectByID(u.ProjectID)
	if err != nil {
		return nil, errors.New("Project not found")
	}
	_, err = GetAccountByID(u.AuthorID)
	if err != nil {
		return nil, errors.New("Account not found")
	}
	if u.Rating == 0 {
		return nil, errors.New("Rating not inserted")
	}
	//ORM database
	o := orm.NewOrm()
	_, err = o.Insert(&u)
	if err == nil {
		//successfully inserted
		return &u, nil
	} else {
		err = o.Read(&u)
		if err != nil {
			return nil, err
		}
	}
	//Update project rating
	project.Rating = GetAverageRating(project.ID)
	_, err = UpdateProject(project.ID, project)
	if err != nil {
		return nil, errors.New("failed to update rating")
	}

	return &u, err
}

func GetRating(authorID, projectID int64) (u *Rating, err error) {
	//ORM
	o := orm.NewOrm()
	rating := Rating{AuthorID: authorID, ProjectID: projectID}
	err = o.Read(&rating, "authorID", "projectID")
	if err != nil {
		return nil, errors.New("Rating not exists")
	} else {
		return &rating, nil
	}

}

func GetAverageRating(projectID int64) int {

	o := orm.NewOrm()
	avg := int64(0)
	var ratings []*Rating

	nums, err := o.Raw("SELECT * FROM rating WHERE projectID = ?", projectID).QueryRows(&ratings)
	if err != nil {
		return 0
	}

	for _, v := range ratings {
		avg += v.Rating
	}

	return int(avg / nums)

}

func UpdateRating(authorID, projectID int64, uu *Rating) (a *Rating, err error) {
	o := orm.NewOrm()

	u, err := GetRating(authorID, projectID)

	log.Print(*u)

	if err == nil {
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
		return nil, errors.New("Rating Not Exist")
	}
}

func DeleteRating(authorID, projectID int64) {
	o := orm.NewOrm()
	_, err := o.Delete(&Rating{AuthorID: authorID, ProjectID: projectID}, "authorID", "projectID")

	if err != nil {
		log.Fatal("delete Account failed")
	}
}
