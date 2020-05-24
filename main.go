package main

import (
	_ "ProjectGallery/routers"
	"ProjectGallery/scheduler"
	"log"

	"github.com/astaxie/beego"

	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	dbUser := beego.AppConfig.String("mysqluser")
	dbPwd := beego.AppConfig.String("mysqlpass")
	dbName := beego.AppConfig.String("mysqldb")
	dbUrls := beego.AppConfig.String("mysqlurls")
	dbString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True&charset=utf8mb4", dbUser, dbPwd, dbUrls, dbName)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbString)
	name := "default"
	force := false
	verbose := true

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	err1 := scheduler.TestPingRedis()
	log.Printf("%v", err1)
	scheduler.InitScheduler()
	log.Printf("already init scheduler time to init beego\n")
	beego.Run()

}
