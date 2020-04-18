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
	ID          int64  `orm:"PK"`
	Username    string `orm:"unique"`
	Password    string
	FullName    string
	Email       string
	Description string
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
}

func AddAccount(u Account) (*Account, error) {
	log.Print(u)
	password := helpers.HashAndSalt([]byte(u.Password))
	u.Password = password
	//ORM database
	o := orm.NewOrm()
	_, err := o.Insert(&u)
	if err == nil {
		//successfully inserted
		return &u, nil
	} else {
		err = o.Read(&u)
		if err != nil {
			return nil, err
		}
	}

	return &u, err
}

func GetAccount(username string) (u *Account, err error) {

	//ORM
	o := orm.NewOrm()
	acc := Account{Username: username}
	err = o.Read(&acc, "Username")
	if err != nil {
		return nil, errors.New("Account not exists")
	} else {
		return &acc, nil
	}

}

func GetAccountByID(id int64) (u *Account, err error) {
	o := orm.NewOrm()
	acc := Account{ID: id}
	err = o.Read(&acc)
	if err != nil {
		return nil, errors.New("Account not exists")
	} else {
		return &acc, nil
	}
}

func GetAllAccounts() map[string]*Account {

	o := orm.NewOrm()
	AccList := make(map[string]*Account)
	var Accounts []*Account
	o.QueryTable(new(Account)).All(&Accounts)

	for _, v := range Accounts {
		AccList[v.Username] = v
	}

	return AccList

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

// func Login(username, password string) (bool, error) {
// 	o := orm.NewOrm()
// 	acc := Account{Username: username, Password: password}
// 	err := o.Read(&acc, "username", "password")
// 	if err != nil {
// 		return true, nil
// 	} else {
// 		return false, errors.New("Invalid username or password")
// 	}
// }

// func Logout(username string) (error) {

// }
