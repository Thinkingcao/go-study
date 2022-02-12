package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (user User) Hello(name string) {
	fmt.Println("hello world")
}

func main() {
	u := User{1, "ok", 12}
	Info(u)

	x:=123
	v:=reflect.ValueOf(&x)
	v.Elem().SetInt(99)
	fmt.Println(x)

	v2:=reflect.ValueOf(u)
	mv:=v2.MethodByName("Hello")
	args:=[]reflect.Value{reflect.ValueOf("jon")}
	mv.Call(args)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Print("type:", t.Name())

	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("xxx")
		return
	}

	v := reflect.ValueOf(o)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Println(f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}
