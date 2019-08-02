package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

/**
编写一个程序，从键盘读取输入。当用户输入 'S' 的时候表示输入结束，这时程序输出 3 个数字：
i) 输入的字符的个数，包括空格，但不包括 '\r' 和 '\n'
ii) 输入的单词的个数
iii) 输入的行数
*/
var nrchars, nrwords, nrlines int

func main() {
	calculator()
	return
	testwordCount()
	nrchars, nrwords, nrlines = 0, 0, 0
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input, type S to stop: ")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
		}
		if input == "S\n" { // Windows, on Linux it is "S\n"
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters: %d\n", nrchars)
			fmt.Printf("Number of words: %d\n", nrwords)
			fmt.Printf("Number of lines: %d\n", nrlines)
			os.Exit(0)
		}
		Counters(input)
	}
}

func Counters(input string) {
	nrchars += len(input) - 2 // -2 for \r\n
	nrwords += len(strings.Fields(input))
	nrlines++
}

func testwordCount() {
	s := "123 asd qwe asd qwe 123 q  "
	fmt.Println(len(strings.Fields(s)))
}

/**
编写一个简单的逆波兰式计算器，它接受用户输入的整型数（最大值 999999）和运算符 +、-、*、/。
输入的格式为：number1 ENTER number2 ENTER operator ENTER --> 显示结果
当用户输入字符 'q' 时，程序结束。请使用您在练习11.13中开发的 stack 包。
*/
func calculator() {
	fmt.Println("输入的整型数（最大值 999999）和运算符 +、-、*、/，按q退出")
	reader := bufio.NewReader(os.Stdin)
	s := new(Stack)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("imput error")
			return
		}
		switch input {
		case "q\n":
			fmt.Println("QUIT")
			return
		case "+\n":
			n1 := s.pop()
			n2 := s.pop()
			fmt.Printf("%d%s%d=%d\n", n2, "+", n1, n2+n1)
		case "-\n":
			n1 := s.pop()
			n2 := s.pop()
			fmt.Printf("%d%s%d=%d\n", n2, "-", n1, n2-n1)
		case "*\n":
			n1 := s.pop()
			n2 := s.pop()
			fmt.Printf("%d%s%d=%d\n", n2, "*", n1, n2*n1)
		case "/\n":
			n1 := s.pop()
			n2 := s.pop()
			fmt.Printf("%d%s%d=%d\n", n2, "/", n1, n2/n1)
		default:
			n, err := strconv.Atoi(input[0 : len(input)-1])
			if err != nil {
				fmt.Println("param error,input a number plz.err = ", err)
			}
			s.push(n)
		}

	}
}
