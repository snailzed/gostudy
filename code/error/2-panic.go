package main

import "fmt"

func main() {
	t1()
	t2()
	t3()
	//testReceover(5)
}

func t1() {
	fmt.Println("t1")
}

func t2() {
	defer func() {
		fmt.Println("defer")
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("t2")
}

func t3() {
	fmt.Println("t3")
}

var arr [5]int

func testReceover(i int) {
	//recover只能在defer语句修饰的函数使用
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(arr[i]) //数组越界会panic  panic: runtime error: index out of range
}
