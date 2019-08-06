package main

import (
	"fmt"
	"time"
)

func main() {
	timeout()
	return
	tick := time.Tick(time.Second)
	boom := time.After(5 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func timeout() {
	ret := make(chan string)

	go func() {
		ret <- "DONE!"
	}()
	for {
		select {

		//select每次都是类似采用轮询机制，判断每个通道是否可用，下面这个定时器每次
		case <-time.After(5 * time.Second):
			fmt.Println("timeout")
			return
		case r := <-ret:
			fmt.Println("ret=", r)
		}
	}
}
