package main

import "fmt"

func main() {
	var a int         //整形
	var b string      //字符串
	var c float32     //浮点型
	var d float64     //浮点型
	var e rune        //字符
	var f bool        //布尔
	var g struct{}    //结构体
	var h interface{} //接口类型，万能类型
	var i byte        //字节

	k := 1 //自动推导类型 int

	var (
		a1 int
		b1 string
	)

	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
	fmt.Printf("c = %v\n", c)
	fmt.Printf("d = %v\n", d)
	fmt.Printf("e = %v，%c\n", e, e)
	fmt.Printf("f = %v\n", f)
	fmt.Printf("g = %v\n", g)
	fmt.Printf("h = %v\n", h)
	fmt.Printf("i = %v\n", i)
	fmt.Printf("k = %v\n", k)
	fmt.Printf("a1 = %v\n", a1)
	fmt.Printf("b1 = %v\n", b1)

	//多重赋值
	A, B, C := 1, 2, 3
	fmt.Println(A, B, C)

	//匿名变量,使用 _ 表示
	A, _ = B, A //将B的值赋值给A，A的值赋值给匿名变量，在函数多个返回值时经常用到
	fmt.Println(A, B)
}
