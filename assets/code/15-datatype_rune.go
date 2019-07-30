package main

import "fmt"

func main() {
	var c string = "qwer1234中"
	d1 := []rune(c)
	for _, d := range c {
		fmt.Printf("%c\n", d)
	}
	fmt.Println(len(c), len(d1))
	testComplex()
}

//复数类型
func testComplex() {
	var t complex128 //声明
	t = 2.1 + 3.14i  //赋值
	fmt.Println("t = ", t)

	//自动推导类型
	t2 := 3.3 + 4.4i
	fmt.Printf("t2 type is %T\n", t2)

	//通过内建函数，取实部和虚部
	fmt.Println("real(t2) = ", real(t2), ", imag(t2) = ", imag(t2))
}
