package main

import "fmt"

func main() {
	if true { // if 和 { 需要在同一行
		fmt.Println(true)
	}

	if a := 1; a == 1 { //支持一个语句初始化，中间需要用分号隔开
		fmt.Println("a == 1")
	}
}
