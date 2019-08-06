package main

import "fmt"

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go tel(ch, quit)
	for {
		select {
		case num := <-ch:
			fmt.Println(num)
		case <-quit:
			fmt.Println("QUIT")
			return
		}
	}
}

func tel(ch chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	quit <- true
}
