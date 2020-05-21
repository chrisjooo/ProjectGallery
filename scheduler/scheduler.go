package scheduler

import (
	"ProjectGallery/helpers"
	"ProjectGallery/models"
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/zhashkevych/scheduler"
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

func CacheMostLiked(ctx context.Context) {
	//caching most liked project every 1 hour
	ctx, _ = context.WithTimeout(ctx, time.Second*30)
	conn := helpers.NewPool().Get()
	defer conn.Close()
	_, err := conn.Do("FLUSHALL")
	if err != nil {
		log.Printf("error flushing: %v\n", err)
	}
	projectList := models.FilterMostLikeProject()
	_, err = conn.Do("HSET", "filtered-data", "data", projectList)
	if err != nil {
		log.Printf("Error setting cache: %v", err)
	}

	select {
	case <-ctx.Done():
		log.Println("Scheduler timeout-30second-")
		return
	default:
	}
}

func InitScheduler() {
	ctx := context.Background()

	sc := scheduler.NewScheduler()
	sc.Add(ctx, CacheMostLiked, time.Minute*1)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	sc.Stop()
}
