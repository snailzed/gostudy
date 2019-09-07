package cpu

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func Test() {
	ch := make(chan int, 1)
	for {
		select {
		case ret := <-ch:
			fmt.Println(ret)
		default:
		}
	}
}
func main() {
	fd, err := os.Create("./cpu.pprof")
	if err != nil {
		panic(err)
	}
	err = pprof.StartCPUProfile(fd)
	if err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()
	for i := 0; i < 8; i++ {
		go Test()
	}
	time.Sleep(10 * time.Second)
}
