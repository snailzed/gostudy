package array

import "fmt"

func main() {
	//定义数组
	var arr [2]int
	arr1 := new([2]int)

	//初始化
	arr[0] = 1
	arr[1] = 2

	arr1[0] = 11
	arr1[1] = 22

	fmt.Printf("arr=%v,arr[0]=%d\n", arr, arr[0])     //arr=[1 2],arr[0]=1
	fmt.Printf("arr1=%v,arr1[0]=%d\n", arr1, arr1[0]) //arr1=&[11 22],arr1[0]=11

	//定义并初始化
	var a1 = [2]int{1, 2}       //指明数组长度
	var a2 = [...]int{1, 2}     //自动计算数组长度
	var a3 = [5]int{0: 1, 3: 8} // key value指明

	fmt.Println(a1) //[1 2]
	fmt.Println(a2) //[1 2]
	fmt.Println(a3) //[1 0 0 8 0]

	//值类型，值拷贝
	var testArr = [2]int{1, 2}
	fmt.Println("testArr1:", testArr)
	testArray(testArr)
	fmt.Println("testArr2:", testArr)
}

func testArray(a [2]int) {
	a[1] = 3
	fmt.Println("testArrIN:", a)
}
