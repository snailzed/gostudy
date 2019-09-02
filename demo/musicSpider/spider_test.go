package musicSpider

import (
	"fmt"
	"testing"
	"time"
)

func TestSpiderJayMusic(t *testing.T) {
	ch := make(chan int, 3)
	lenCh := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go do(i, ch, lenCh)
	}

	var i int
	for {
		select {
		case <-lenCh:
			i++
			if i >= 10 {
				fmt.Println("Download finished.")
				return
			}
		}
	}
}

func do(i int, ch chan int, lenCh chan int) {
	ch <- 1
	lenCh <- 1
	time.Sleep(2 * time.Second)
	fmt.Printf("[%d]do do do....\n", i)
	<-ch
}
