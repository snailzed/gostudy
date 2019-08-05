package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go sendData(ch)
	go getData(ch)
	time.Sleep(time.Second)
}

//向通道发送消息 <-
func sendData(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func getData(ch chan int) {
	for {
		fmt.Println("GETDATA")
		data := <-ch
		fmt.Println(data)
	}
}
