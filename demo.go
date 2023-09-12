package main

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func process() {
	for {
		fmt.Println("--------------->")

		time.Sleep(2 * time.Second)
	}
}

func main() {
	//var s = fmt.Sprintf("hello %s", "taoyf")
	//fmt.Println(s)

	startT := time.Now()

	zap.S().Info("#start ----> ")

	go process()

	//go speed()

	//接收终止信号 Signal表示操作系统信号
	quit := make(chan os.Signal)

	//接收control+c
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-quit
	zap.S().Infof("#-----------> done, serve %v", time.Since(startT))
}
