package main

import "fmt"

type User struct {
	Name    string
	Age     int
	Address string
}

func main() {
	var user = User{"forezp", 10, "sz"}
	fmt.Printf("user = %v\n", user)
	user.Age = 30
	fmt.Printf("user = %v\n", user)
	changeUser(user)
	fmt.Printf("user = %v\n", user)
	changeUser2(&user)
	fmt.Printf("user = %v\n", user)
}

func changeUser(user User) {
	//值的传递
	user.Address = "bj"
}

func changeUser2(user *User) {
	//值的传递
	user.Address = "bj"
}
