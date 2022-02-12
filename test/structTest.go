package main

import "fmt"

type Person struct {
	name string
	age int
}

func main() {

	p:=Person{}
	p.age=10
	p.name="forezp"
	fmt.Println(p)

	b:= struct {
		Name string
		Age int
	}{
		Name:"bill",
		Age:15,
	}

	fmt.Println(b)

	c:=P{"1",1}
	fmt.Println(c)
}

type P struct {
	string
	int
}


