package main

import "fmt"

func main() {
	var myMap map[string]string
	if myMap == nil {
		fmt.Println("mymap is nil")
	}
	myMap = make(map[string]string)
	myMap["name"] = "forzp"
	myMap["age"] = "10"
	myMap["addres"] = "sz"

	for k, v := range myMap {
		fmt.Println("key=", k, "val=", v)
	}

	fmt.Println("=======")
	myMap2 := make(map[int]string, 10)
	myMap2[1] = "1"
	myMap2[2] = "2"
	for k, v := range myMap2 {
		fmt.Println("key=", k, "val=", v)
	}

	fmt.Println("=======")

	myMap3 := map[string]string{
		"one":   "php",
		"two":   "c++",
		"three": "python",
	}
	printMap(myMap3)
	//新增
	myMap3["four"] = "java"
	printMap(myMap3)
	//删除
	delete(myMap3, "four")
	printMap(myMap3)
	//修改
	myMap3["one"] = "golang"
	printMap(myMap3)
	ChangeValue(myMap3)
	printMap(myMap3)
}

func ChangeValue(cityMap map[string]string) {
	//引用的传递
	cityMap["one"] = "lua"
}

func printMap(myMap map[string]string) {
	fmt.Println("---------")
	//引用的传递
	for key, value := range myMap {
		fmt.Println("key = ", key, "value=", value)
	}
}
