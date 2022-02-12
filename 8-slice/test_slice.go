package main

import "fmt"

func main() {

	var slice1 = []int{1, 2, 3}
	printSlice(slice1)
	fmt.Println("--------")
	for i, v := range slice1 {
		fmt.Println("index=", i, "val=", v)
	}
}

func printSlice(slice []int) {
	//引用传递
	for i, v := range slice {
		fmt.Println("index=", i, "val=", v)
	}
	slice[0] = 100
}
