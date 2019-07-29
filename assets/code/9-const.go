package main

import "fmt"

func main() {
	// const 常量名 [类型] = 初始值
	const a int = 1
	fmt.Printf("a = %d\n", a) //a = 1

	//常量定义之后不能修改
	//a = 2
	//fmt.Print(a) // cannot assign to a

	//优雅的写法
	const (
		a1 int    = 2
		a2        //默认会继承上一个定义的常量的类型和值 a2=2
		b1 string = "hi"
		b2        //默认会继承上一个定义的常量的类型和值 a2=3
		b3        //默认会继承上一个定义的常量的类型和值 a2=3
	)
	fmt.Printf("a1 = %d a2=%d b1=%s b2=%s b3=%s\n", a1, a2, b1, b2, b3) //a1 = 2 a2=2 b1=hi b2=hi

	// iota使用
	const (
		c  = iota //0
		c1 = iota //1
		c2 = iota //2
		c3        //3 等价于 c3 = iota，此时iota=3，自动递增
	)
	fmt.Println(c, c1, c2, c3) //0 1 2 3

	const (
		d = iota
		d1
		d2
		d3
	)
	fmt.Println(d, d1, d2, d3) //0 1 2 3

	//左移运算符 << ：二进制数左移右位补0 左移n位 = 乘以2的n次方
	//右移运算符 << ：二进制数右移左位补1 右移n位 = 除以2的n次方
	const (
		e  = 1 << iota //0 1左移0位
		e1             // e1 = 1 << iota 1左移1位 = 2  0001=> 0010 = 2
		e2             // e2 = 1 << iota 1左移2位 = 4  0001=> 0100 = 4
	)
	fmt.Println(e, e1, e2) //1 2 4

	//iota默认值等于0，逐行递增，跟每行是否有`iota`无关
	const (
		f  = iota //0
		f1 = 2    //2
		f2 = iota //2
	)
	fmt.Println(f, f1, f2) // 0 2 2

	const (
		A = iota //0
		B        //1
		C        //2
		D = 8    //8
		E        //8
		F = iota //5
		G        //6
	)
	fmt.Println(A, B, C, D, E, F, G)
}
