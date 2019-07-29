package main

import "fmt"

func main() {
	var c string = "qwer1234中"
	d1 := []rune(c)
	for _, d := range c {
		fmt.Printf("%c\n", d)
	}
	fmt.Println(len(c), len(d1))
}
