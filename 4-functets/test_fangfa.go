package main

import "fmt"

//不同于函数的方法
//在 Go 语言中，方法和函数是两个概念，但又非常相似，不同点在于方法必须要有一个接收者，这个接收者是一个类型，这样方法就和这个类型绑定在一起，称为这个类型的方法。
//
//在下面的示例中，type Age uint 表示定义一个新类型 Age，该类型等价于 uint，可以理解为类型 uint 的重命名。其中 type 是 Go 语言关键字，表示定义一个类型，在结构体和接口的课程中我会详细介绍。
//和函数不同，定义方法时会在关键字 func 和方法名 String 之间加一个接收者 (age Age) ，接收者使用小括号包围。

type Age uint
func (age Age) String(){
	fmt.Println("the age is",age)
}

func (age *Age) Modify(){
	*age = Age(30)
}

func main() {
	age:=Age(2)
	age.String()

	age2:=Age(25)
	//方法赋值给变量，方法表达式
	sm:=Age.String
	//通过变量，要传一个接收者进行调用也就是age
	sm(age2)
	//我们知道，方法 String 其实是没有参数的，但是通过方法表达式赋值给变量 sm 后，在调用的时候，必须要传一个接收者，这样 sm 才知道怎么调用。
	//
	//小提示：不管方法是否有参数，通过方法表达式调用，第一个参数必须是接收者，然后才是方法自身的参数。

	age3:=Age(26)
	age3.String()
	age3.Modify()
	age3.String()
	//示例中调用指针接收者方法的时候，使用的是一个值类型的变量，并不是一个指针类型，其实这里使用指针变量调用也是可以的，如下面的代码所示：
	//
	//复制代码
	//(&age).Modify()
	//这就是 Go 语言编译器帮我们自动做的事情：
	//
	//如果使用一个值类型变量调用指针类型接收者的方法，Go 语言编译器会自动帮我们取指针调用，以满足指针接收者的要求。
	//
	//同样的原理，如果使用一个指针类型变量调用值类型接收者的方法，Go 语言编译器会自动帮我们解引用调用，以满足值类型接收者的要求。
	defer func() {
		if p:=recover();p!=nil{
			fmt.Println(p)
		}
	}()
	panic("1111")
}