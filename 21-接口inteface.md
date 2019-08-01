# 21、接口interface
>  `interface`定义需要interface关键字，go的接口是非常灵活的；接口通常以`[e]r`、`able`为后缀或者以`I`开头；
go中可以定义一个接口类型的变量`（ var a Namer）`，任何实现了该接口的类型都可以将值赋值给`a`。
接口是一种契约，实现类型必须满足它，它描述了类型的行为，规定类型可以做什么。接口彻底将类型能做什么，以及如何做分离开来，使得相同接口的变量在不同的时刻表现出不同的行为，这就是多态的本质
## 1、定义接口
    
       type Namer interface {
           Method1(param_list) return_type
           Method2(param_list) return_type
           ...
       }
       
       type 接口名 interface {
          //方法，只有声明，没有实现，由别的类型（自定义类型）实现
          方法名(参数列表)(返回参数列表)
       }   
> 接口变量：  var 变量名 接口名 //此时变量名就是接口变量

## 2、接口继承
> 接口是可以继承接口的，新的接口拥有继承的接口所有定义的方法
```go
type ReadWrite interface {
    Read(b Buffer) bool
    Write(b Buffer) bool
}

type Lock interface {
    Lock()
    Unlock()
}

type File interface {
    ReadWrite //匿名字段继承
    Lock //匿名字段继承
    Close()
}
```
`File`接口继承了`ReadWrite`和`Lock`接口，拥有所有接口方法

## 3、类型断言、类型判断
> 接口变量在运行时的实际类型是不确定的，故需要`类型断言`来帮助判断实际类型。在处理未知类型的数据非常有用

        var intf 接口类型名
        v := intf.(T) // intf必须是一个接口变量
        
> `if v,ok := 接口变量.(T)`类型断言
        
        if v,ok := intf.(T);ok {
           //转换可能失败，返回的第二个参数表示是否转换成功，第一个参数是转换为T的值，转换失败则为T类型的默认值
        }
 ```go
//类型断言
func typeAssertion() {

	var sim Simpler
	sha := new(Shape)
	sim = sha

	if v, ok := sim.(*Shape); ok {
		fmt.Printf("%T\n", v)
	} else {
		fmt.Printf("noT!\n")
	}
}

```          
> `type-switch`类型判断 `接口变量.(type)`为获得的类型 `case T类型`；
因为空接口的存在，所以所有变量都可以使用类型断言或者类型判断
```go

//类型判断
switch t := areaIntf.(type) {
case *Square:
	fmt.Printf("Type Square %T with value %v\n", t, t)
case *Circle:
	fmt.Printf("Type Circle %T with value %v\n", t, t)
case nil:
	fmt.Printf("nil value: nothing to check?\n")
default:
	fmt.Printf("Unexpected type %T\n", t)
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}
```   

## 4、空接口        
> `interface{}`，空接口没有任何方法，可以看做所有类型都实现了该接口，对于接口类型的变量，只要实现了该接口，那么就可以将类型实例赋值给该接口变量，所以空接口是万能类型     
 
      var t interface{} //定义一个万能类型

## 5、接口类型方法集

 
       类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
       类型 T 的可调用方法集包含接受者为 T 的所有方法
       类型 T 的可调用方法集不包含接受者为 *T 的方法
       
## 注意事项
- a、go中默认会有空接口`interface {}`,里面没有任何方法，换言之所有类型都实现了空接口，即空接口可以表示任何类型的变量
- b、`接口都是隐式实现的`，多个类型可以实现同一个接口
- c、实现接口的类型可以实现其他方法
- d、只要类型`T`实现了该接口，就可以将`T`类型的实例赋值给该接口变量;可以利用该特性实现多态（传入接口变量，根据实际的类型调用接口方法，同一个类型有不同的表现方式为多态）；
