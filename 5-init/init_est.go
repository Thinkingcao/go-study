package main

import (
	"fmt"
	mylib1 "go-study/5-init/lib1"
	_ "go-study/5-init/lib2"
)
func main() {
	mylib1.Test()
	//lib2.Test()
}

func init() {
	fmt.Println("main init")
}
