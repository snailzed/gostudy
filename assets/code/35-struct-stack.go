package main

import "fmt"

type Stack struct {
	data  [4]int
	index int
}

func (s *Stack) push(d int) {
	if s.index >= 4 {
		panic("stack overflow")
	}
	s.data[s.index] = d
	s.index++
}

func (s *Stack) pop() int {
	if s.index <= 0 {
		panic("no ele")
	}
	s.index--
	ret := s.data[s.index]
	s.data[s.index] = 0
	return ret

}

//实现栈（stack）数据结构：
func main() {
	stack := new(Stack)
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	fmt.Println(stack)

	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	stack.push(4)
	stack.push(4)
	fmt.Println(stack)
}
