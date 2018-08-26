package main

import "fmt"

func main() {

	//指针传递
	a := 1
	A(&a)
	fmt.Println(a)

	//匿名函数
	b := func() {
		fmt.Println("anouyos func")

	}
	b();
	//闭包
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func A(a *int) {
	*a = 2
	fmt.Println(*a)

}
