package main

import "fmt"

func main() {
	res := reverseV2("qwer12341")
	fmt.Println(res)

	res1 := reverseUtf8V2("上海自来水来自海上")
	fmt.Println(res1)

	res2 := isHuiwen("上海自来水来自海上")
	fmt.Println(res2)
}

//对英文字符串逆序
func reverse(str string) string {
	var result []byte
	for i := len(str) - 1; i >= 0; i-- {
		result = append(result, str[i])
	}
	return string(result)
}
func reverseV2(str string) string {
	result := []byte(str)
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}
	return string(result)
}

//对包含中文字符串逆序
func reverseUtf8(str string) string {
	var result []rune
	strRune := []rune(str)
	for i := len(strRune) - 1; i >= 0; i-- {
		result = append(result, strRune[i])
	}
	return string(result)
}
func reverseUtf8V2(str string) string {
	result := []rune(str)
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}
	return string(result)
}

//判断是否字符串是否回文
func isHuiwen(str string) bool {
	result := []rune(str)
	for i := 0; i < len(result)/2; i++ {
		if result[i] != result[len(result)-i-1] {
			return false
		}
	}
	return true
}
