package main

import "fmt"

func main() {

	var a  =make(map[int]string)
	a[0]="ok";
	for k,v:=range a {
		fmt.Println(k,v)
	}

	for _,v:=range a{
		fmt.Println(v)
	}

	m:=map[int]string{1:"a",2:"b"}
	fmt.Println(m)
}
