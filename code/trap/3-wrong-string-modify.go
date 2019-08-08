package main

import "fmt"

func main() {

	str := "I am a phper."
	//修改字符串
	//str[0]= 'U' //错误写法，字符串不可变，必须先转成字节或字符数组

	strByte := []byte(str)
	strByte[0] = 'U'
	str = string(strByte)
	fmt.Println(str)
}
