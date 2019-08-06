package main

import "fmt"

func main() {
	ch1 := make(chan int)
	go pump(ch1)                           // pump hangs
	fmt.Println(len(ch1), cap(ch1), <-ch1) // prints only 0
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
