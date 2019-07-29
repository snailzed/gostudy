package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	a = "1121"
	fmt.Println(a)

	//原样输出
	a = `12121阿达撒多所;大所多`
	fmt.Println(a)

	//长度
	alen := len(a)

	//是否包含某个子串
	res := strings.Contains(a, "12")

	fmt.Println(alen, res) //30 true
	//拼接
	a += "哈哈哈哈啊哈"
	fmt.Println(a) //12121阿达撒多所;大所多哈哈哈哈啊哈

	a = fmt.Sprintf("%s%s", a, a)
	fmt.Println(a) //12121阿达撒多所;大所多哈哈哈哈啊哈12121阿达撒多所;大所多哈哈哈哈啊哈

	//分割
	res1 := strings.Split(a, ";")
	fmt.Println(res1) //[12121阿达撒多所 大所多哈哈哈哈啊哈12121阿达撒多所 大所多哈哈哈哈啊哈]

	//前后缀判断
	res2 := strings.HasPrefix(a, "12")
	res3 := strings.HasSuffix(a, "哈哈哈哈啊哈")
	fmt.Println(res2, res3) //true true

	//判断子串第一次出现位置
	res4 := strings.Index(a, "212")
	res5 := strings.LastIndex(a, "2128")
	fmt.Println(res4, res5) //1 -1

	//join操作
	var b []string = []string{"12", "34", "56"}
	res6 := strings.Join(b, "")
	fmt.Println(res6) //123456
}
