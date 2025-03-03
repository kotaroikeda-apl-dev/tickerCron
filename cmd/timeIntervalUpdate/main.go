package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

var tickerInterval = 5 * time.Second

func tickerJob() {
	for {
		fmt.Println("Tickerジョブ実行:", time.Now())
		time.Sleep(tickerInterval)
	}
}

func changeInterval() {
	fmt.Println("Tickerの間隔を 2秒 に変更")
	tickerInterval = 2 * time.Second
}

func main() {
	c := cron.New()
	c.AddFunc("@every 15s", changeInterval)
	c.Start()

	go tickerJob() // Tickerをゴルーチンで実行

	time.Sleep(40 * time.Second)
	c.Stop()
}
