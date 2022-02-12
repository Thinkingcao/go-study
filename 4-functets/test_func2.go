package main

import "fmt"


func main() {

	//匿名函数：
	sum2 := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum2(1, 2))

	//闭包这都得益于匿名函数闭包的能力，让我们自定义的 colsure 函数，可以返回一个匿名函数，并且持有外部函数 colsure 的变量 i。因而在 main 函数中，每调用一次 cl()，i 的值就会加 1。
	//
	//小提示：在 Go 语言中，函数也是一种类型，它也可以被用来声明函数类型的变量、参数或者作为另一个函数的返回值类型。
	cl:=colsure()
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())
}

func colsure() func() int {
	i:=0
	return func() int {
		i++
		return i
	}
}