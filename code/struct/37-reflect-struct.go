package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name string "it is name"
	Age  int    "it is age"
	Desc string "it is desc"
}

func (a Animal) SetName(name string) {
	a.Name = name
}
func main() {
	GetMethod()
	return
	a := new(Animal)
	v := reflect.ValueOf(a) //必须传入指针，才能修改值
	//Elem获取指向的内存空间，Field(index):设置指定的字段  SetXXX:只能设置可导出的字段
	v.Elem().Field(0).SetString("Snail")
	v.Elem().FieldByName("Age").SetInt(43)

	fmt.Println(v, v.Type().Elem().Field(0).Tag)
}

func GetMethod() {
	a := new(Animal)
	v := reflect.ValueOf(a)

	for i := 0; i < v.NumMethod(); i++ {
		m := v.Type().Method(i)
		fmt.Println(m.Name, m.Type)
	}
}
