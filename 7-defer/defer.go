package main

import "fmt"

func main() {
	//test()
	fmt.Println("result=", returnAndDeffer())
}

func deferFun() {
	fmt.Println("defer func invoke")
}

func returnFunc() int {

	fmt.Println("return fun invoke")
	return 1
}

func returnAndDeffer() int {
	defer deferFun()
	return returnFunc()
}
func test() {
	defer fmt.Println("defer one execute")
	defer fmt.Println("defer two execute")

	fmt.Println("execute test")
}
