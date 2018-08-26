package main

import "fmt"

type A2 struct {
	name string
}

func main() {
	a:=A2{}
	a.Print()
}

func (a2 A2)Print()  {
	fmt.Println("A")
}