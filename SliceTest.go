package main

import "fmt"

func main()  {

	s1:=make([]int ,3,10)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	fmt.Println(s1)

	s2:=make([]int,3,6)
	fmt.Printf("%p\n",s2)
	s2=append(s2,1,2,4)
	fmt.Printf("%v %p\n",s2,s2)
	s2=append(s2,1,2,4)
	fmt.Printf("%v %p\n",s2,s2)

}
