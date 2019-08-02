# 22、文件操作

## 1、获取用户终端输入
使用`fmt`包的`Scan`、`Scanf`和`Scanln`以及`Sscan`函数。标准输入：`os.Stdin`

        //从终端读取
        fmt.Scanf(format string, a ...interface{}) (n int, err error)：格式化输入，format决定如何读取数据，格式化格式与格式化输出一致，遇到换行符结束
        fmt.Scan(a ...interface{}) (n int, err error) ：从终端获取用户输入,以换行符和空格作为分隔符
        fmt.Scanln(a ...interface{}) (n int, err error) ： 从终端获取用户输入，以空格分割并依次放入到后续的参数内，碰到换行符结束
        
        //从字符串读取，规则和从终端获取一致
        fmt.Sscan()
        fmt.Sscanf()
        fmt.Sscanln()

       //使用 bufio包读取标准输入
       inputReader := bufio.NewReader(os.Stdin) //获取reader指针
       input,_ = inputReader.ReadString('\n') //input为读取的字符串，\n为换行符
       
       
```go
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
```        
