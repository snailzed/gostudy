package main

import "fmt"

func main() {
	var a int8 // -128~ 127
	a = 18
	a = -1
	a = 127
	a = -128
	fmt.Println(a)

	var b int
	b = 1288888889999999
	fmt.Println(b)

	var c float32 = 1
	fmt.Println(c)

	b = int(a)

	fmt.Println(b)
}
