package models

import (
	"ProjectGallery/helpers"
	"errors"
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Account))
}

type Account struct {
	Id          int64     `orm:"PK" json:"id"`
	Username    string    `orm:"unique" json:"username"`
	Password    string    `json:"password"`
	FullName    string    `json:"fullname"`
	Email       string    `json:"email"`
	Description string    `json:"description"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
}

type AccList struct {
	NumAcc int64      `json:"total_account"`
	Data   []*Account `json:"data"`
}

func AddAccount(u Account) (*Account, error) {
	log.Print(u)
	log.Println("creating account")
	password := helpers.HashAndSalt([]byte(u.Password))
	u.Password = password
	//ORM database
	o := orm.NewOrm()

	//check username
	acc := Account{Username: u.Username}
	err := o.Read(&acc, "Username")
	if err == nil || err != orm.ErrNoRows {
		log.Print(err)
		return nil, errors.New("username already exist")
	}
	log.Print("sampai sini")

	newId, err := o.Insert(&u)
	if err == nil {
		//successfully inserted
		u.Id = newId
		log.Print("new: ", u)
		return &u, nil
	} else {
		log.Print("error here")
		return nil, errors.New("error inserting account")
	}

	return &u, err
}

func GetAccount(username string) (u *Account, err error) {

	//ORM
	o := orm.NewOrm()
	acc := Account{Username: username}
	err = o.Read(&acc, "Username")
	if err != nil {
		log.Print("read account error: ", err)
		return nil, errors.New("Account not exists")
	} else {
		return &acc, nil
	}

}

func GetAllAccounts() *AccList {

	o := orm.NewOrm()
	list := &AccList{}
	var account []*Account
	o.QueryTable(new(Account)).All(&account)

	list.Data = account
	list.NumAcc = int64(len(account))

	return list

}

func UpdateAccount(username string, uu *Account) (a *Account, err error) {
	o := orm.NewOrm()

	u, err := GetAccount(username)

	log.Print(*u)

	if err == nil {
		if uu.Email != "" {
			u.Email = uu.Email
		}
		if uu.Description != "" {
			u.Description = uu.Description
		}
		if uu.FullName != "" {
			u.FullName = uu.FullName
		}
		if uu.Password != "" {
			password := helpers.HashAndSalt([]byte(uu.Password))
			u.Password = password
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
		return nil, errors.New("Account Not Exist")
	}
}

func DeleteAccount(username string) {
	o := orm.NewOrm()
	_, err := o.Delete(&Account{Username: username}, "Username")

	if err != nil {
		log.Fatal("delete Account failed")
	}
}

func Login(username, password string) (bool, error) {
	o := orm.NewOrm()
	acc := Account{Username: username}

	err := o.Read(&acc, "username")

	if err != nil {
		return false, errors.New("InvalId username or password")
	}

	check := helpers.ComparePassword(acc.Password, []byte(password))
	if check {
		return true, nil
	} else {
		return false, errors.New("InvalId username or password")
	}
}

// func Logout(username string) (error) {

// }
