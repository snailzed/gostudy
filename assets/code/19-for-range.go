package main

import (
	"fmt"
	"time"
)

func main() {
	//1、遍历字符串
	a := "qwertyuuip"
	for i, char := range a {
		fmt.Printf("a[%d]=%c\n", i, char)
	}

	//2、遍历数组和切片
	b := []string{"HI", "你好", "你你好"}
	for i, str := range b {
		fmt.Printf("b[%d]=%s\n", i, str)
	}

	//3、遍历map，map是无序的，所以每次遍历出来的结果可能不一致
	cMap := map[int]string{0: "HI", 5: "HII", 10: "Hello"}
	for i, d := range cMap {
		fmt.Printf("cMap[%d]=%s\n", i, d)
	}

	//4、遍历管道数据
	channel := time.Tick(2 * time.Second)
	for d := range channel {
		fmt.Println("Rev channel data:", d)
	}
}
