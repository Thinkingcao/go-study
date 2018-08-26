package main

import "fmt"

func main() {
	A1()
	B1()
	C1()
}

func A1()  {
	fmt.Println("func a1")
}

func B1()  {
	defer func() {

		if err:=recover();err!=nil{
			fmt.Println("recovery in b1")
		}
 	}()

	panic("paninc in b1")
}

func C1()  {

	fmt.Println("func c1")
}
