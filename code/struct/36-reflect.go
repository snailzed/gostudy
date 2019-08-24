package main

import (
	"fmt"
	"reflect"
)

func main() {
	//使用reflect.TypeOf 获取类型
	var a int = 1
	//test_typeof(a)
	test_valueof(&a)
	fmt.Println(a)
}

func test_typeof(a interface{}) {
	//返回Type类型结构体
	fmt.Println(reflect.TypeOf(a))

	//使用Kind()
	b := reflect.TypeOf(a).Kind()
	fmt.Println(b)
	switch b {
	case reflect.Int:
		fmt.Println("int")
	}
}

func test_valueof(a interface{}) {
	//返回Value类型的结构体
	v := reflect.ValueOf(a)
	v.Type() //等价于 reflect.TypeOf()

	switch v.Kind() {
	case reflect.Int:
		v.SetInt(42)
		fmt.Println("int")
	case reflect.Float32:
		fmt.Println("Float32")
	case reflect.Ptr:
		v.Elem().SetInt(2423) //传入的是地址，必须先调用Elem()
		fmt.Println("Ptr")
	}
}
