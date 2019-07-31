package main

import "fmt"

func main() {
	testSliceWriteStyle()

	//定义切片
	var arr = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	s := arr[:]
	s1 := arr[0:3] //s1 s 都指向arr，对s1 s修改都会修改arr
	fmt.Println("before arr =", arr)
	fmt.Println("before s =", s)
	testSlice(s)
	testSlice(s1)
	fmt.Println("after arr = ", arr)
	fmt.Println("after s = ", s)
}

func testSliceWriteStyle() {
	//定义切片
	var arr = [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	//以下4中写法相同，切片元素与数组相同，并且都指向arr
	s1 := arr[:]
	s2 := arr[0:]
	s3 := arr[:len(arr)]
	s4 := arr[0:len(arr)]

	fmt.Printf("slice s1 = %v\n", s1)
	fmt.Printf("slice s2 = %v\n", s2)
	fmt.Printf("slice s3 = %v\n", s3)
	fmt.Printf("slice s4 = %v\n", s4)
}
func testSlice(a []int) {
	a[0] *= 10
	a[1] *= 10
	a[2] *= 10
}
