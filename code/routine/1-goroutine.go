package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//go SayHello()
	//go SayHello()
	//time.Sleep(time.Second * 2)
	testGosched()
}

func SayHello() {
	fmt.Println("Hello")
}

func testGosched() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("routine:", i)
		}
	}()

	for i := 0; i < 10; i++ {
		runtime.Gosched()
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
