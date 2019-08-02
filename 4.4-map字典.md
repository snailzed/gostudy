# 19、map字典【类似php的关联数组】

## 1、定义
> `key`可以是任意可以用 == 或者 != 操作符比较的类型，如int、float、string，只包含内建类型的struct；
`value`可以是任意类型的值,如`[]int` `*[]int`

    var map1 map[keytype]valuetype
    var map1 map[string]int
    
    var map2 = make(map[keytype]valuetype)
    
```go
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

```    
## 2、初始化
>定义后手动初始化，可以在定义时进行初始化
   
       //定义并初始化
       var m = map[int]string{1:"HI"}
      
       //定义后，手动赋值
       var m1 = map[int]string
          m1[1] = "string"    
 
## 3、map操作
- a、检测是否存在某个`key`，`ok`语法

    
    //isPresent为true说明存在，为false则不存在，此时val1为nil
    val1, isPresent = map1[key1]
    
     //判断是否存在该键，ok为true存在，false不存在
     val, ok := m2[1]
     if ok {
     	fmt.Printf("key[%d] is exists in m2,val = \n", 1, val)
     } else {
     	fmt.Printf("key[%d] is not  exists in m2\n", 1)
     }
     
     if _,ok := m2[key];ok {
        fmt.Println("exists")
     } else{
        fmt.Println("not exists")
     }

- b、删除某个`key`

      delete(map,key) //key不存在不会产生错误

## 4、`for range`遍历map
 > for range遍历map，第一个返回值为key，第二个返回值为key对应的value,value是一个值拷贝，修改value不会修改map；map是无序的。
 
       for key, value := range map1 {
       	...
       }

> 如果只想获取`key`

       //遍历只获取key
       	for key := range m {
       		fmt.Printf("key=[%d]\n", key)
       	}            
         
## map注意事项

- a、未初始化的map值为`nil`，使用`make`创建的map`非nil` 
- b、`map`是引用类型，作为参数传递是引用传递 
- c、map的性能比数组和切片慢，性能考虑则尽量使用切片来解决
- d、map是可变长的，长度可变   
- e、map只有len，没有`cap容量`
- f、map是无序的，map排序可以先将`key`取出再利用`sort包`排序