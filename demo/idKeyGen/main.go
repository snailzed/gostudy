package main

import (
	"fmt"
	"github.com/sony/sonyflake"
)

func main() {
	settings := sonyflake.Settings{}
	sf := sonyflake.NewSonyflake(settings)
	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}
	fmt.Printf("sony/sonyflake:%d\n", id)
}
