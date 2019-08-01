package main

import "fmt"

type student struct {
	name string
	age  int
}
type rectangle struct {
	width int
	long  int
}

func main() {
	var s student   //普通结构体实例
	var s1 *student //结构体指针

	s = student{name: "snailzed", age: 24}
	s1 = &student{name: "oyxp", age: 24} //指定字段初始化，未初始化的字段则为该类型的默认值
	s2 := student{"oyxp1", 1}            //顺序初始化
	s3 := new(student)                   //使用new获取实例
	s3 = &student{}
	s3.age = 456

	testStruct(s2)

	fmt.Println(s, s.name, s.age)
	fmt.Println(s1, s1.name, s1.age)
	fmt.Println(s2, s2.name, s2.age)
	fmt.Println(s3, s3.name, s3.age)

	r := rectangle{width: 100, long: 100}
	Area(r)

	fmt.Println(FileInstance(1, "test"))
}

//普通结构体变量作为参数是值传递，指针变量为引用传递
func testStruct(s student) {
	s.age = 999
	s.name = "12131qweqweqw"
}

func Area(r rectangle) {
	fmt.Println(r.long * r.width)
}

//名字改为小写的
type file struct {
	fd   int
	name string
}

func FileInstance(fd int, name string) *file {
	return &file{fd: fd, name: name}
}
