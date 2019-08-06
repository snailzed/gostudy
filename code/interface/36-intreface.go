package _interface

import "fmt"

//定义一个接口
type Simpler interface {
	//方法只声明，不需要实现
	Set(int)
	Get() int
}

type Shape struct {
}

func (s *Shape) Set(a int) {

}
func (s *Shape) Get() int {
	return 1
}

func TestSimplerInterface(s Simpler) {
	s.Set(1)
	fmt.Println(s.Get())
}

func main() {
	var s Simpler
	//TestSimplerInterface(s); //错误 不能运行，可以有接口变量，但不能直接用接口变量去调用未实现的方法

	sha := new(Shape)
	s = sha

	s.Get()
	s.Set(1)
	typeAssertion()
}

//类型断言
func typeAssertion() {

	var sim Simpler
	sha := new(Shape)
	sim = sha

	if v, ok := sim.(*Shape); ok {
		fmt.Printf("%T\n", v)
	} else {
		fmt.Printf("noT!\n")
	}
}
