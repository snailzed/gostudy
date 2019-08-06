package slice

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

	testSliceCap()
	testSliceLen()
	testCopyAndAppend()
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

func testSliceCap() {
	s := []int{}
	fmt.Printf("len(s)=%d,cap(s)=%d\n", len(s), cap(s)) // len=0 cap=0
	//cap会根据长度增加，一般是长度的2倍
	// 1 1 | 2 2 | 3 4 | 4 4 | 5 8 | 6 8 | 7 8 | 8 8| 9 16 |10 16
	for i := 0; i < 10; i++ {
		s = append(s, i)
		fmt.Printf("len(s)=%d,cap(s)=%d\n", len(s), cap(s))
	}
	/**
	len(s)=0,cap(s)=0
	len(s)=1,cap(s)=1
	len(s)=2,cap(s)=2
	len(s)=3,cap(s)=4
	len(s)=4,cap(s)=4
	len(s)=5,cap(s)=8
	len(s)=6,cap(s)=8
	len(s)=7,cap(s)=8
	len(s)=8,cap(s)=8
	len(s)=9,cap(s)=16
	len(s)=10,cap(s)=16
	*/
}

func testSliceLen() {
	var s = []int{1, 2}
	s = s[0:1]
	fmt.Println(len(s))
}

func testCopyAndAppend() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)
	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}
