package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func tickerJob() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for i := 0; i < 3; i++ {
		fmt.Println("Tickerジョブ:", time.Now())
		<-ticker.C
	}
}

func cronJob() {
	fmt.Println("Cronジョブ実行:", time.Now())
}

func main() {
	c := cron.New()
	c.AddFunc("@every 5s", cronJob)
	c.Start()

	go tickerJob() // Tickerジョブを別ゴルーチンで動かす

	time.Sleep(30 * time.Second)
	c.Stop()
}
