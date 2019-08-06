package function //必须

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

//函数也是一种数据类型， 通过type给一个函数类型起名
//FuncType它是一个函数类型
type FuncType func(int, int) int //没有函数名字，没有{}

func main() {
	var result int
	result = Add(1, 1) //传统调用方式
	fmt.Println("result = ", result)

	//声明一个函数类型的变量，变量名叫fTest
	var fTest FuncType
	fTest = Add            //是变量就可以赋值
	result = fTest(10, 20) //等价于Add(10, 20)
	fmt.Println("result2 = ", result)

	fTest = Minus
	result = fTest(10, 5) //等价于Minus(10, 5)
	fmt.Println("result3 = ", result)

	//函数作为参数传递
	fmt.Println(cal(1, 2, Add))
	fmt.Println(cal(1, 2, Minus))
}

//函数作为参数传递
func cal(a, b int, op FuncType) int {
	return op(a, b)
}

//多形参列表 左边的参数如没写明类型，则类型继承下一个最近的有声明类型的形参
func testMultiParamV2(a, b string) string {
	return a + b
}
func testFuncType() {
	a := testMultiParamV2
	fmt.Println(a("a", "b"))
}
