package main

import (
	"fmt"
)

func main() {
	suck(pump())
	//time.Sleep(1e9) //主协程退出，子协程也退出
	for {

	}
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
