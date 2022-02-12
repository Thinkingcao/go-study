package main

import (
	"fmt"
)

func main() {

	var a1 [3]int
	fmt.Println(a1)
	var a2 = [...]int{12, 12, 12, 1}
	a2[0]=123
	fmt.Println(a2[0])
	var a3 = []int{1: 1, 3: 3}
	fmt.Println(a2)
	fmt.Println(a3)
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	b := []int{19: 1}
	fmt.Println(b)

	fmt.Println(a)

	//冒泡排序
	c := [...]int{55, 2, 45, 3, 77}
	fmt.Println(c)
	num := len(c)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if c[i] < c[j] {
				temp := c[i]
				c[i] = c[j]
				c[j] = temp
			}
		}
	}
	fmt.Println(c)

	for i := 0; i < 3; i++ {
		v := 1
		fmt.Println(&v)
	}
}
