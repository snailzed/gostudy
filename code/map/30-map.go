package _map

import "fmt"

func main() {
	//定义map  map[键类型]值类型
	var m = map[int]string{}
	m[1] = "string"

	m1 := map[int]string{1: "HI"}
	m1[2] = "LOO"

	//使用make创建map  make:array、slice、map
	m2 := make(map[int]string, 10) //值不为nil

	var m3 map[int][]int // 定义未初始化的值为nil
	if m3 == nil {
		fmt.Println("HHHH")
	}

	//判断是否存在该键，ok为true存在，false不存在
	val, ok := m2[1]
	if ok {
		fmt.Printf("key[%d] is exists in m2,val = \n", 1, val)
	} else {
		fmt.Printf("key[%d] is not  exists in m2\n", 1)
	}

	delete(m2, 111111) //key不存在不会产生错误
	fmt.Printf("m=%v\n", m)
	fmt.Printf("m1=%v\n", m1)
	fmt.Printf("m2=%v\n", m2)

	testCreateMap()
	testMap()
}

func testMap() {
	m := map[int]int{10: 1, 2: 2, 3: 3, 4: 4}
	//遍历map
	for key, value := range m {
		fmt.Printf("key[%d]=val[%d]\n", key, value)
	}
	//遍历只获取key
	for key := range m {
		fmt.Printf("key=[%d]\n", key)
	}
}

func testCreateMap() {
	var m = map[struct {
		a int
		b int
	}]int{struct {
		a int
		b int
	}{a: 1, b: 1}: 1}

	fmt.Printf("%v", m)
}
