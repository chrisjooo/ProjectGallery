package models

import (
	"ProjectGallery/helpers"
	"log"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Account))
}

type Account struct {
	Id          int64     `orm:"PK" json:"id" form:"-"`
	Username    string    `orm:"unique" json:"username" form:"username"`
	Password    string    `json:"password" form:"password"`
	FullName    string    `json:"fullname" form:"fullname"`
	Email       string    `json:"email" form:"email"`
	ProfilePic  string    `json:"profile_pic" form:"profile_pic"`
	Description string    `json:"description" form:"description"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
}

type AccountData struct {
	Account       Account `json:"account_data"`
	CompressedPic string  `json:"compressed_image"`
}

type AccList struct {
	NumAcc int64          `json:"total_account"`
	Data   []*AccountData `json:"data"`
}

func AddAccount(u Account) (*Account, error) {
	password := helpers.HashAndSalt([]byte(u.Password))
	u.Password = password
	//ORM database
	o := orm.NewOrm()

	//check username
	acc := Account{Username: u.Username}
	err := o.Read(&acc, "Username")
	if err == nil || err != orm.ErrNoRows {
		log.Print(err)
		errMessage := helpers.ErrorMessage(helpers.AccountExist)
		return nil, errMessage
	}

	newId, err := o.Insert(&u)
	if err == nil {
		//successfully inserted
		u.Id = newId
		return &u, nil
	} else {
		errMessage := helpers.ErrorMessage(helpers.Post)
		return nil, errMessage
	}
}

func GetAccount(username string) (u *AccountData, err error) {

	//ORM
	o := orm.NewOrm()
	acc := Account{Username: username}
	err = o.Read(&acc, "Username")
	if err != nil {
		if err == orm.ErrNoRows {
			errMessage := helpers.ErrorMessage(helpers.CheckAccount)
			return nil, errMessage
		}
		log.Print("read account error: ", err)
		errMessage := helpers.ErrorMessage(helpers.Get)
		return nil, errMessage
	} else {
		u = &AccountData{}
		u.Account = acc
		if acc.ProfilePic != "" {
			url := acc.ProfilePic[:strings.LastIndexByte(acc.ProfilePic, '.')] + "-compressed.png"
			u.CompressedPic = url
		} else {
			u.CompressedPic = ""
		}
		return u, nil
	}

}

func GetAllAccounts() *AccList {

	o := orm.NewOrm()
	list := &AccList{}
	var account []*Account
	o.QueryTable(new(Account)).All(&account)
	var accountData []*AccountData

	for _, v := range account {
		u := &AccountData{}
		u.Account = *v
		if v.ProfilePic != "" {
			url := v.ProfilePic[:strings.LastIndexByte(v.ProfilePic, '.')] + "-compressed.png"
			u.CompressedPic = url
		} else {
			u.CompressedPic = ""
		}
		accountData = append(accountData, u)
	}

	list.Data = accountData
	list.NumAcc = int64(len(account))

	return list

}

func UpdateAccount(username string, uu *Account) (u *AccountData, err error) {
	o := orm.NewOrm()

	u, err = GetAccount(username)
	acc := Account{}
	acc = u.Account
	u = &AccountData{}

	if err == nil {
		if uu.Email != "" {
			acc.Email = uu.Email
		}
		if uu.Description != "" {
			acc.Description = uu.Description
		}
		if uu.FullName != "" {
			acc.FullName = uu.FullName
		}
		if uu.Password != "" {
			password := helpers.HashAndSalt([]byte(uu.Password))
			acc.Password = password
		}
		if uu.ProfilePic != "" {
			acc.ProfilePic = uu.ProfilePic
			url := uu.ProfilePic[:strings.LastIndexByte(uu.ProfilePic, '.')] + "-compressed.png"
			u.CompressedPic = url
		} else {
			acc.ProfilePic = uu.ProfilePic
			u.CompressedPic = ""
		}
		// ORM Update
		_, err1 := o.Update(&acc)
		log.Print(u, err)

		if err1 == nil {
			//update successful
			u.Account = acc
			return u, nil
		} else {
			errMessage := helpers.ErrorMessage(helpers.Put)
			return nil, errMessage
		}
	} else {
		return nil, err
	}
}

func DeleteAccount(username string) error {
	o := orm.NewOrm()

	_, err := GetAccount(username)
	if err != nil {
		return err
	}

	_, err = o.Delete(&Account{Username: username}, "Username")

	if err != nil {
		log.Println("delete Account failed")
		errMessage := helpers.ErrorMessage(helpers.Delete)
		return errMessage
	}
	return nil
}

func Login(username, password string) (*helpers.TokenDetails, error) {
	o := orm.NewOrm()
	acc := Account{Username: username}

	err := o.Read(&acc, "username")

	if err != nil {
		errMessage := helpers.ErrorMessage(helpers.AccountLogin)
		return nil, errMessage
	}

	check := helpers.ComparePassword(acc.Password, []byte(password))
	if !check {
		errMessage := helpers.ErrorMessage(helpers.AccountLogin)
		return nil, errMessage
	}

	//JWT TOKEN

	token, err := helpers.CreateToken(acc.Username)
	if err != nil {
		errMessage := helpers.ErrorMessage(helpers.JWTLogin)
		return nil, errMessage
	}
	return token, nil

}
