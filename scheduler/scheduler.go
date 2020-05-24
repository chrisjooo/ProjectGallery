package scheduler

import (
	"ProjectGallery/helpers"
	"ProjectGallery/models"
	"encoding/json"
	"log"

	"gopkg.in/robfig/cron.v2"
)

func TestPingRedis() error {
	conn := helpers.NewPool().Get()
	defer conn.Close()

	response, err := conn.Do("PING")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("result: %v", response)
	return err
}

func CacheMostLiked() {
	conn := helpers.NewPool().Get()
	defer conn.Close()
	projectList := models.FilterMostLikeProject()

	jsonData, err := json.Marshal(projectList)
	if err != nil {
		log.Printf("error marshaling data")
	}

	_, err = conn.Do("HSET", "filtered-data", "data", jsonData)
	if err != nil {
		log.Printf("Error setting cache: %v", err)
	} else {
		log.Printf("Success inserting every 1 minute")
	}
}

func InitScheduler() {
	c := cron.New()
	c.AddFunc("@every 0h1m0s", func() {
		log.Println("Cron Jobs every 1 minute")
		CacheMostLiked()
	})
	c.Start()
}
