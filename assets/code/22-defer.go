package main

import "fmt"

func main() {
	//testDefer()
	testDefer1()
}

//多个defer执行顺序：栈，先进后出，跟析构函数一样
func testDefer() {
	defer fmt.Println(1111)
	defer fmt.Println(8888)
	fmt.Println(2222)
	fmt.Println(3333)
	return
}

//defer定义时，如果使用了变量，则变量的值为执行到该行defer的值
func testDefer1() {
	//执行顺序： testErr(0) -> fmt.Println("END")
	i := 1
	defer fmt.Println(i)     // 1，不是等于100
	defer fmt.Println("END") //还是会执行
	defer testErr(0)         //报错 integer divide by zero
	i = 100
	fmt.Println(i)
}

func testErr(a int) int {
	return 1 / a
}
