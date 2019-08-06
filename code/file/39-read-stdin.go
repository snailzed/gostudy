package main

import (
	"bufio"
	"fmt"
	"os"
)

//读取标准输入输出
func main() {
	//1、读取标准输入
	//testScanf()
	//testScan()
	//testScanln()
	//testSscan()
	testBufioread()
}

//scanf决定输入的格式
func testScanf() {
	var (
		a int
		b string
		c float32
	)
	fmt.Scanf("%d %s %f", &a, &b, &c)
	fmt.Println(a, b, c)
}

//Scan以换行符和空格作为分隔符
func testScan() {
	var (
		a int
		b string
		c float32
	)
	fmt.Scan(&a, &b, &c)
	//fmt.Scan(&a)
	//fmt.Scan(&b)
	//fmt.Scan(&c)
	fmt.Println(a, b, c)
}

//一行输入，以空格作为分隔符
func testScanln() {
	var (
		a int
		b string
		c float32
	)
	fmt.Scanln(&a, &b, &c)
	fmt.Println(a, b, c)
}

func testSscan() {
	var (
		a int
		b string
		c float32
	)
	s := "1234 qwer 5.565"
	fmt.Sscan(s, &a, &b, &c)
	fmt.Println(a, b, c)
}

//使用bufio读取的字符串会包括分隔符
func testBufioread() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	fmt.Println(input)
}
