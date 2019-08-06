package base

import "fmt"

func main() {

	//指定类型
	var a bool //默认值 false
	var a1 bool = true
	var a2 bool = false

	//不指定类型
	var b1 = true
	var b2 = false

	//自动推导类型
	c1 := true
	c2 := false

	fmt.Println(a, a1, a2, b1, b2, c1, c2) //false true false true false true false
	fmt.Print("%t", a)
}
