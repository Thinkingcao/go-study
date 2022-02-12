package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	pp:=NewPerson("飞雪无情",20)
	fmt.Println(*pp)
}

func NewPerson(name string,age int) *person{
	p:=new(person)
	p.name = name
	p.age = age
	return p
}