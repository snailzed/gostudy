package main

import "fmt"

func main() {
	test()
	main1()
	var num int
	fmt.Printf("请按下楼层：")
	fmt.Scan(&num)
	switch num { //switch后面写的是变量本身
	case 1:
		fmt.Println("按下的是1楼")
		//break //go语言保留了break关键字，跳出switch语言， 不写，默认就包含
		//fallthrough //不跳出switch语句，后面的无条件执行
	case 2:
		fmt.Println("按下的是2楼")
		//break
		//fallthrough
	case 3:
		fmt.Println("按下的是3楼")
		//break
		//fallthrough
	case 4:
		fmt.Println("按下的是4楼")
		//break
		//fallthrough
	default:
		fmt.Println("按下的是xxx楼")
	}
}
func test() {
	switch a := 8; a {
	case 1, 2, 3, 4:
		fmt.Println(" 1 <= a <=4 ")
	case 5, 6, 7, 8:
		fmt.Println(" 5 <= a <=8 ")
	}
}
func main1() {
	switch i := 1; i > 1 {
	case true:
		fmt.Println(true)
	case false:
		fmt.Println(false)
	}
}
func main2() {
	//支持一个初始化语句， 初始化语句和变量本身， 以分号分隔
	switch num := 4; num { //switch后面写的是变量本身
	case 1:
		fmt.Println("按下的是1楼")

	case 2:
		fmt.Println("按下的是2楼")

	case 3, 4, 5:
		fmt.Println("按下的是yyy楼")

	case 6:
		fmt.Println("按下的是4楼")

	default:
		fmt.Println("按下的是xxx楼")
	}

	score := 85
	switch { //可以没有条件
	case score > 90: //case后面可以放条件
		fmt.Println("优秀")
	case score > 80: //case后面可以放条件
		fmt.Println("良好")
	case score > 70: //case后面可以放条件
		fmt.Println("一般")
	default:
		fmt.Println("其它")
	}
}
