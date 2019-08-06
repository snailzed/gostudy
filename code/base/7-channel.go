package base

import (
	"fmt"
	"time"
)

//实现协程与主协程的通信
func main() {
	channel := make(chan int, 1)

	go func() {
		fmt.Println("Goroutine!")
		time.Sleep(2 * time.Second)
		channel <- 1
	}()

	res := <-channel
	fmt.Println(res)
}
