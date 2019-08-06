package main

import "fmt"

func main() {
	ch := make(chan int, 1) //缓冲通道
	fmt.Println(len(ch), cap(ch))
	ch <- 1
	fmt.Println(<-ch)
}
