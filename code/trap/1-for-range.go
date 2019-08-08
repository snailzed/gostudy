package main

import "fmt"

func main() {
	s := "hello world"
	ret := make(map[string]*int)
	for i, char := range s {
		index := i
		fmt.Printf("s[%d]=%c\n", i, char)
		ret[string(char)] = &index //取地址赋值，因为被遍历的变量是将下标和值分别赋值给 i和 char，每次用的都是一个内存地址，所以这里i其实最后的值为 10
	}
	/**
	遍历过程：声明了i和char，将每次遍历的值分别赋值给i和char

	遍历次数  i的值  char的值
	1        0     h
	2        1     e


	*/

	for j, d := range ret {
		fmt.Printf("s[%s]=%d\n", j, *d)
	}

	/**
	s[0]=h
	s[1]=e
	s[2]=l
	s[3]=l
	s[4]=o
	s[5]=
	s[6]=w
	s[7]=o
	s[8]=r
	s[9]=l
	s[10]=d
	s[l]=10
	s[o]=10
	s[ ]=10
	s[w]=10
	s[r]=10
	s[d]=10
	s[h]=10
	s[e]=10
	*/
}
