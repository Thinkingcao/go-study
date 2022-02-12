package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `info:"名字"`
	Age  int    `info:"年龄"`
}

func (this User) Call() {
	fmt.Println("name=", this.Name, "age=", this.Age)
}

func main() {
	test(111)
	user := User{"forezp", 11}
	test(user)
	testFiledAndMethod(user)
	findTag(user)
}

func findTag(arg interface{}) {
	t := reflect.TypeOf(arg).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		fmt.Println("tag info=", tagInfo)
	}
}

func testFiledAndMethod(arg interface{}) {

	argType := reflect.TypeOf(arg)
	argValue := reflect.ValueOf(arg)
	fmt.Println("type=", argType, "val=", argValue)

	fmt.Println("开始遍历字段")
	//遍历字段
	for i := 0; i < argType.NumField(); i++ {
		field := argType.Field(i)
		value := argValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
	fmt.Println("开始遍历方法")
	for i := 0; i < argType.NumMethod(); i++ {
		m := argType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

func test(arg interface{}) {
	fmt.Println("type=", reflect.TypeOf(arg))
	fmt.Println("value=", reflect.ValueOf(arg))
	fmt.Println("------")
}
