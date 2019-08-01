# 20、struct结构体
> 属于复合类型，结构体里面的数据称为`字段`，类似于OOP语言中的类，每一个结构体都有自己的属性和方法

## 1、定义结构体

         type identifier struct {
             field1 type1
             field2 type2
             ...
         }
          
          type 结构体名 struct{
              字段名 字段类型
              ... 
          }
 

## 2、获取实例和初始化                  
> 结构体初始化跟普通变量初始化一致，包括`顺序初始化`（必须和结构体字段顺序一致，直接填入值即可）和`指定字段初始化`(字段:值)



      var  变量名 结构体名   //普通结构体实例
      var  变量名 *结构体名  //结构体指针实例
      var  变量名 = 结构体名{ field:value }  //未初始化的值为默认值
      var  变量名 = *结构体名{ v1, v2,v3 } //未初始化的值为默认值,顺序初始化
      
      var s = struct1{}    //返回普通结构体
      var s1 = *struct1{}   //返回指针类型
      var s2 = new(struct) //返回指针类型
       *T{} 等价于 new(T)
      
```go
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

	fmt.Println(s, s.name, s.age)
	fmt.Println(s1, s1.name, s1.age)
	fmt.Println(s2, s2.name, s2.age)
	fmt.Println(s3, s3.name, s3.age)
}
//普通结构体变量作为参数是值传递，指针变量为引用传递
func testStruct(s student) {
	s.age = 999
	s.name = "12131qweqweqw"
}

func Area(r rectangle) {
	fmt.Println(r.long * r.width)
}


```      
      
       
## 3、结构体操作
>获取字段值使用`.`选择器,`.`既可以获取普通结构体的字段值，也可以获取结构体指针的字段值

- a、结构体工厂
> 根据可见性规则，将结构体名写成小写，然后专门提供一个函数用来获取结构体实例
      
 ```go
//名字改为小写的
type file struct {
	fd   int
	name string
}

func FileInstance(fd int, name string) *file {
	return &file{fd: fd, name: name}
}
```       
- b、使用`reflect`包获取结构体中`tag`标签
> 结构体中的`tag`标签可以使用`reflect`反射获取
   
```go

package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	name string "it is name"
	age  int    "it is age"
	desc string "it is desc"
}

func main() {
	a := new(Animal)
	atype := reflect.TypeOf(*a)
	fmt.Println("field 0 type = ", atype.Field(0).Tag)
	fmt.Println("field 1 type = ", atype.Field(1).Tag)
	fmt.Println("field 2 type = ", atype.Field(2).Tag)
	/**
	field 0 type =  it is name
	field 1 type =  it is age
	field 2 type =  it is desc
	*/
}
``` 
- c、匿名字段以及初始化
> 匿名字段：没有显式的名字，只有数据类型；
 
匿名字段可以是一个结构体类型，如果是结构体类型，可以看成是当前结构体继承了该结构体，拥有了该结构体的所有成员和方法。
也可以是其他类型`T`，可以通过`结构体名.T`来操作该字段。可以顺序初始化和指定字段初始化。
                       
```go
package main

import "fmt"

type Animal struct {
	id   int
	name string
	age  int
}

type Cat struct {
	leg byte
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
```
- d、同名字段规则
> 在结构体中，同名字段使用规则（就近原则）：如果在本作用域找到该成员则使用该成员，否则从继承的结构体中查找成员；如果同一作用域中出现同名字段，使用时则会报错。
 
 
        type A struct {
           a int
           b int
        }    
        
        type B struct {
           a int
        }
        
        type C struct {
           A
           B
           b int
        }
        
        c := C{},操作c.a时会报错，因为无法确定是c.A.a还是c.B.a （同一作用域同名字段会报错）；
        操作c.b则实际操作的是C本身的b，不是c.A.b,就近原则
       
       
## 4、结构体方法
> 有接收者的函数为方法（给某个类型绑定一个函数）

      func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... } 
      
      //调用：recv.methodName(parameter_list) ,其中recv为receiver_type实例

其中接收者类型`receiver_type`不能为指针类型（不能是一个直接定义的指针变量，var i *int），但是它可以是任何其他允许类型的指针;可以是结构体类型、函数类型、(int、string、bool、数组等值类型)的别名类型

## 5、指针或值作为接收者
> 指针作为接收者，方法调用时为引用传递，可以直接修改接收者的内容，不消耗内存；而值作为接收者，方法调用时为值拷贝，会消耗内存；
方法调用时，接受者既可以为指针也可以为值，Go底层会自动解引用（rev <=> *rev，指针和值会自动转换）

```go
package main

import (
	"fmt"
)

type B struct {
	thing int
}

func (b *B) change() { b.thing = 1 }

func (b B) write() string { return fmt.Sprint(b) }

func main() {
	var b1 B // b1是值
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) // b2是指针
	b2.change()
	fmt.Println(b2.write())
}

/* 输出：
{1}
{1}
*/

```  
## 6、方法继承
> 结构体中存在结构体匿名字段，那么当前结构体则会自动继承匿名结构体的所有成员以及方法
```go

package main

import "fmt"

type Person struct {
	name string //名字
	sex  byte   //性别, 字符类型
	age  int    //年龄
}

//Person类型，实现了一个方法
func (tmp *Person) PrintInfo() {
	fmt.Printf("name=%s, sex=%c, age=%d\n", tmp.name, tmp.sex, tmp.age)
}

//有个学生，继承Person字段，成员和方法都继承了
type Student struct {
	Person //匿名字段
	id     int
	addr   string
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 666, "bj"}
	s.PrintInfo()
}
```                
## 结构体注意事项
- a、结构体在内存中是连续块形式存在，性能很强
- b、`普通结构体作为参数传递为值传递，指针类型结构体为引用传递`
- c、`make`函数只能串创建内置引用类型（`slice、map、channel`），`new`函数用来创建`值类型和用户自定义类型`
- d、在一个结构体中对于每一种数据类型`只能有一个匿名字段`
- e、匿名字段为结构体，则当前结构体继承了该结构体的所有成员
- f、结构体和其方法可以不在同一个文件内，但必须在同一个包，所以int、string等这些类型上不能直接定义方法，因为不在同一个包，但可以通过别名方式定义
- g、不支持方法重载，支持方法重写
               