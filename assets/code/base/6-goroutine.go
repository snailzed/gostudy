package base

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Goroutine!")
	}()
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Goroutine!")
	}()
	time.Sleep(6 * time.Second)

}
