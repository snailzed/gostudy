package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	i := 10
	for i < 20 {
		fmt.Println(i)
		i++
	}

	for i < 30 {
		fmt.Println(i)
		i++
	}
}
