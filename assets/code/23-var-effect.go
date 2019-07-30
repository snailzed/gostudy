package main

import "fmt"

var a int = 100 //全局变量
var b int = 1   //全局变量
func main() {
	var b int = 11 //局部变量
	fmt.Println("a =", a)
	fmt.Println("b =", b) //就近原则访问变量
	testLocalVar()
}

func testLocalVar() {
	//if else  a1的作用域仅局限于if else代码块
	if a1 := 1; a1 > 2 {
		fmt.Println("true:a=", a1)
	} else {
		fmt.Println("false:a=", a1)
	}
	//fmt.Println("a1=", a1)//报错

	//for代码块 i仅在{}生效
	for i := 1; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}
	//fmt.Printf("i = %d\n", i) //报错 undefined:i

	//switch代码块
	switch j := 1; j == 1 {
	case true:
	case false:
	}
	//fmt.Println(j)//报错
}
