package main

import "fmt"

func main() {

	var arr1 [10]int
	for i := 0; i < len(arr1); i++ {
		fmt.Print(" ", arr1[i])
	}

	arr2 := [5]int{1, 2, 3, 4, 5}
	for index, v := range arr2 {
		fmt.Println("index=", index, "val=", v)
	}
	fmt.Printf("arr2 types = %T\n", arr2)
	fmt.Println("-------------")
	printArr(arr2)

	fmt.Println("---------")

	for index, v := range arr2 {
		fmt.Println("index=", index, "val=", v)
	}
}

func printArr(arr [5]int) {
	//值拷贝，不是引用拷贝
	for index, value := range arr {
		fmt.Println("index = ", index, ", value = ", value)
	}
	arr[0]=111

}
