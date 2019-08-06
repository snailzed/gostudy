package function

import "fmt"

func main() {
	fmt.Println(testDefer2()) //2
}

//defer是在函数return之后，返回之前执行的
func testDefer2() (i int) {
	var j int = 1
	defer func() {
		i++
	}()
	return j
}
