package main

import "fmt"

func main() {

	//i是一个局部变量
	var i int
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(i) //0
}
