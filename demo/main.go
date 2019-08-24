package main

import (
	"fmt"
	"gostudy/demo/logger"
)

func main1() {
	log := logger.NewFileLogger(logger.LOG_INFO, "./runtime/info.log", "./runtime/warn.log", "./runtime/error.log")
	log.Debug("it is debug log")
	log.Info("it is info log")
	log.Warn("it is warn log")
	log.Error("it is error log")
	log.Fatal("it is fatal log")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}
