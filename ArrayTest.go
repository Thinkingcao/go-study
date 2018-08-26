package main

import (
	"fmt"
)

func main()  {
	a:=[...]int{1,2,3,4,5}
	fmt.Println(a)

	b:=[...]int{19:1}
	fmt.Println(b)



	//冒泡排序
	c:=[...]int{55,2,45,3,77}
	fmt.Println(c)
	num:=len(c)
	for i:=0; i<num;i++  {
		for j:=i+1;j<num ;j++  {
			if c[i]<c[j]{
				temp:=c[i]
				c[i]=c[j]
				c[j]=temp
			}
		}
	}
	fmt.Println(c)

	for i:=0;i<3 ;i++  {
		v:=1
		fmt.Println(&v)
	}
}
