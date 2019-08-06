package function

import "fmt"

func main() {
	testVoid()
	testSingleReturnV1()
	testSingleReturnV2()

	testChangeableParam(1, 23, 4)
	testFuncType()
}

//1、无参无返回值
func testVoid() {
	fmt.Println("testVoid:无参无返回值")
}

//2、单一返回值
func testSingleReturnV1() int {
	return 1
}

//2.1、单一返回值
func testSingleReturnV2() (a int) {
	a = 1
	return
}

//3、多返回值
func testMultiReturnV1(a, b int) (int, int) {
	return a + b, a - b
}
func testMultiReturnV2(a, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}
func testMultiReturnV3(a, b int) (sum, sub int) { //返回值如果和后面的返回类型相同则类型可省略不写
	sum = a + b
	sub = a - b
	return
}

//4、多形参列表
func testMultiParamV1(a string, b string) string {
	return a + b
}

//多形参列表 左边的参数如没写明类型，则类型继承下一个最近的有声明类型的形参
func testMultiParamV2(a, b string) string {
	return a + b
}

//5、可变参数
func testChangeableParam(args ...int) {
	fmt.Printf("%T", args)
	fmt.Println(args)
}

//统计字符串字符个数,转成[]rune
func cal(str string) {

}

// 6、函数类型
func testFuncType() {
	a := testMultiParamV2
	fmt.Println(a("a", "b"))
}
