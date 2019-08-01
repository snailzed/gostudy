package main

import "fmt"

type Animal struct {
	id   int
	name string
	age  int
}

type Cat struct {
	leg    byte
	Animal //结构体类型匿名字段，继承Animal结构体，拥有Animal结构体所有字段
	int    //非结构体类型， Cat.int来使用该字段，意味着只有结构体同时只能拥有一种类型的匿名字段
}

func main() {

	cat := new(Cat)
	//id name age来自Animal
	cat.id = 1
	cat.name = "myCat"
	cat.age = 1

	cat.leg = 'w'
	//cat.int调用匿名字段
	cat.int = 50

	fmt.Println("cat=", cat)

	//顺序初始化
	cat2 := Cat{60, Animal{2, "45", 12}, 1}
	fmt.Println("cat2=", cat2)

	//指定字段初始化
	cat3 := Cat{leg: 80, int: 69, Animal: Animal{id: 4}}
	fmt.Println("cat3=", cat3)
}
