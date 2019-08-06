package _interface

import "fmt"

//接口
type A interface {
	sayHi()
}

type B interface {
	sayHello()
}

//自定义类型
type tinyint uint

//T类型接收者
func (t tinyint) sayHi() {
	fmt.Println("Hi")
}

//*T类型接收者
func (t *tinyint) sayHello() {
	fmt.Println("Hello")
}

func testA(b B) {
	b.sayHello()
}

//测试接口类型的方法集
func main() {
	var t tinyint
	//t := new(tinyint)
	t.sayHi()
	//t.sayHello()
	testA(t)
}
