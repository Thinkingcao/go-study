package main

import "fmt"

const ga int = 200
const (
	SHEN_ZHENG = iota + 1
	GUANGZHOU
	FOSHAN
)

const (
	q, w = iota + 1, iota + 2
	e, r
	t, y
)

func main() {

	const a = 100
	fmt.Println("a=", a)
	fmt.Println("ga=", ga)
	fmt.Println("shenzhen=", SHEN_ZHENG)
	fmt.Println("guangzhou=", GUANGZHOU)
	fmt.Println("foshan=", FOSHAN)

	fmt.Println("q=", q)
	fmt.Println("w=", w)
	fmt.Println("e=", e)
	fmt.Println("r=", r)
	fmt.Println("t=", t)
	fmt.Println("y=", y)

}
