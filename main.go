package main

import (
	_ "ProjectGallery/routers"
	"ProjectGallery/scheduler"
	"log"

	"github.com/astaxie/beego"

	// "github.com/joho/godotenv"

	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// var Pool *redis.Pool

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	log.Printf("\ninitialize redis\n")

	// redisUrls := beego.AppConfig.String("redisurls")
	// redisString := fmt.Sprintf("%s:6379", redisUrls)

	// Pool = &redis.Pool{
	// 	MaxIdle:     10,
	// 	IdleTimeout: 240 * time.Second,
	// 	Dial: func() (redis.Conn, error) {
	// 		return redis.Dial("tcp", "localhost:6379")
	// 	},
	// }

	// conn, err := redis.Dial("tcp", "localhost:6379")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer conn.Close()

	// s, err := redis.String(conn.Do("PING"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("masuk ping: %s", s)

	// log.Printf("test ping redis\n")
	// err := scheduler.TestPingRedis()
	// log.Printf("error pinging: %v", err)

	err1 := scheduler.TestPingRedis()
	log.Printf("%v", err1)

	dbUser := beego.AppConfig.String("mysqluser")
	dbPwd := beego.AppConfig.String("mysqlpass")
	dbName := beego.AppConfig.String("mysqldb")
	dbUrls := beego.AppConfig.String("mysqlurls")
	dbString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True&charset=utf8mb4", dbUser, dbPwd, dbUrls, dbName)

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
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	beego.Run()
}
