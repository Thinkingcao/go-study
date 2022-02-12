package main

import "fmt"

var (
	ga = 100
	gb = "jack"
)

var gc int = 1000
var gd string = "rose"
var ge = "tai"

func main() {
	var a int
	fmt.Println("a=", a)
	fmt.Printf("type of a = %T\n", a)

	var b int = 100
	fmt.Println("b=", b)

	var c = 200
	fmt.Println("c=", c)

	d := 300
	fmt.Println("d=", d)

	var e, f int = 500, 600
	fmt.Println("e=", e, "f=", f)
	var g, h = 700, "alice"
	fmt.Println("g=", g, "h=", h)

	var (
		i = 800
		j = 900
	)
	fmt.Println("i=", i, "j=", j)
	fmt.Println("ga=", ga, "gb=", gb)
	fmt.Println("gc=", gc, "gd=", gd, "ge=", ge)

}
