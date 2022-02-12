package main

import "fmt"

func test(arg interface{}) {
	fmt.Println("arg=", arg)
	val, ok := arg.(string)
	if ok {
		fmt.Println("arg is a string ,arg=", val)
	} else {
		fmt.Println("arg is not string")
	}
}

type Book struct {
	Name string
}

func main() {
	test("forezp")
	fmt.Println("-----")
	test(100)
	book := Book{"go in action"}
	fmt.Println("-----")
	test(book)
}
