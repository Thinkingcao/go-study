package main

import "fmt"

func main() {
	c := foo1(100, "test")
	fmt.Println("c=", c)
	a, b := foo2(100, 200)
	fmt.Println("a=", a, "b=", b)
	c, d := foo3(300, 400)
	fmt.Println("c=", c, "d=", d)

	e, f := foo4(500, 600)
	fmt.Println("e=", e, "f=", f)

	fmt.Println("foo5=", foo5(1, 2, 3, 4, 5, 6, 7, 8, 9))

}

func foo1(a int, b string) int {
	fmt.Println("a=", a, " b=", b)
	c := 1
	return c
}

func foo2(a int, b int) (int, int) {
	return a, b
}

func foo3(a int, b int) (r1 int, r2 int) {
	r1 = 300
	r2 = 400
	return
}

func foo4(a int, b int) (r1, r2 int) {
	r1 = 500
	r2 = 600
	return
}

func foo5(params ...int) (sum int) {
	sum = 0
	for _, v := range params {
		sum += v
	}
	return
}
