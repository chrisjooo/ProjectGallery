package main

import (
	_ "ProjectGallery/routers"

	"github.com/astaxie/beego"

	"github.com/joho/godotenv"

	"fmt"
	"log"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := beego.AppConfig.String("mysqluser")
	dbPwd := beego.AppConfig.String("mysqlpass")
	dbName := beego.AppConfig.String("mysqldb")
	dbUrls := beego.AppConfig.String("mysqlurls")
	dbString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", dbUser, dbPwd, dbUrls, dbName)

	// Register Driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// Register default database
	orm.RegisterDataBase("default", "mysql", dbString)

	// autosync
	// db alias
	name := "default"

	// drop table and re-create
	force := false

	// print log
	verbose := true

	// error
	err = orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	beego.Run()
}
