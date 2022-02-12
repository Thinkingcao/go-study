package main

import "fmt"

func main() {

	//var slice []int
	//slice:= []int {1,2,4}
	//var slice []int =make([]int ,3)
	slice := make([]int, 3)
	if slice == nil {
		fmt.Println("slice是一个空切片")
	} else {
		fmt.Println("slice已经开辟了空间")
		fmt.Println()
	}
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice = append(slice, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice), cap(slice), slice)

	s1 := slice[0:2] // [1, 2]
	fmt.Println(s1)
	//copy 可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 5) //s2 = [0,0,0]
	copy(s2, slice)
	fmt.Println(s2)
}
