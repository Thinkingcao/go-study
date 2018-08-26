package main

import "fmt"

type human struct {
	Sex int
}

type teacher struct {
	human
	Name string
	Age int
}

type student struct {
	human
	Name string
	Age int
}


func main() {
 a:=teacher{Name:"ss",Age:1,human:human{Sex:1}}
 fmt.Println(a)
}
