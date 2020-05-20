package scheduler

import (
	"ProjectGallery/helpers"
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/zhashkevych/scheduler"
)

func TestPingRedis() error {
	log.Printf("masuk testpingredis")
	conn := helpers.NewPool().Get()
	log.Printf("masuk setelah testpingredis")
	defer conn.Close()

	response, err := conn.Do("PING")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("result: %v", response)
	return err
}

func CacheMostLiked() error {
	//caching most liked project every 1 hour
	log.Printf("masuk sini setiap 10 detik\n")

	return nil
}

func InitScheduler() {
	ctx := context.Background()

	sc := scheduler.NewScheduler()
	sc.Add(ctx, CacheMostLiked(), time.Second*10)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	sc.Stop()
}
