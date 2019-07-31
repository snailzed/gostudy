package main

import "fmt"

func main() {
	var i int = 6
	ip := *&i          //ip是一个指针
	fmt.Println(i, ip) // 6 6¬

	var j *int
	j = &i

	fmt.Println(*j)
}
