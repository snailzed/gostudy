# 17、 切片[例子](./assets/code/29-slice.go)
>数组一个连续片段的引用（该数组我们称之为相关数组，通常是匿名的；属于`引用类型`，可以使用`make`关键字创建。

## 1、定义切片
     
       var identifier []type //type为数据类型,不用写长度

## 2、切片初始化

      var slice1 []type = arr1[start:end]
      var slice2 = []type{} //底层会创建一个数组并返回一个切片
      
## 3、切片详解
> 切片是从数组而来的，写法详解
       
       arr[start:end]: 左闭右开
         start:切片开始的下标，包含arr[start]，不写默认为0
         end：切边结束的下标，不包含arr[end],不写默认为数组长度即len(arr)
       
       切片长度： len(slice)
       切片容量： cap(slice),容量会根据长度自动扩容，每次扩容是原先容量的2倍
               cap(slice) >= len(slice)
        
       arr[:]、arr[0:]、arr[:len(arr)]、arr[0:len(arr)]  :切片和数组元素相同,切片指向数组   
                
## 4、切片作为参数
> 切片是对数组的引用，所以是`引用传递`;对切片进行增删改都会影响数组。

## 5、切片在内存中的组织方式

有 3 个域的结构体：`指向相关数组的指针`，`切片长度`以及`切片容量`

## 6、切片复制与追加
>追加：`append(dst []type,src []type)`: 将`src`切片元素复制到`dst`切片  
复制：`copy([]type,ele...)`：将`ele`追加到切片中 
```go

func testCopyAndAppend() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)****
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}
```        
       
## 切片注意事项
- 1 切片是`引用类型`，作为参数传递时属于`引用传递`
- 2 切片的`长度是可变`的，而数组是不可变的
- 3 切片实际是对数组片段（全集或子集）的引用，`修改切片会直接影响数组`
- 4 绝对不要用指针指向 slice。切片本身已经是一个引用类型，所以它本身就是一个指针!!
- 5 可以使用`for` `for range`遍历切片
- 6 切片是指向数组，如果指向数组的切片一直存在，那么数组就不会释放