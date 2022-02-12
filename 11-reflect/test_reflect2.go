package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string
	Age int
}

func (p person) Print(prefix string){
	fmt.Printf("%s:Name is %s,Age is %d\n",prefix,p.Name,p.Age)
}

func main() {
	p:=person{Name: "飞雪无情",Age: 20}
	pv:=reflect.ValueOf(p)
	//反射调用person的Print方法
	mPrint:=pv.MethodByName("Print")
	args:=[]reflect.Value{reflect.ValueOf("登录")}
	mPrint.Call(args)
}