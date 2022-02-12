package main

import "fmt"

type AnimalIF interface {
	eat(s string)
	sleep()
	run()
}

type Cat struct {
	Name string
	Age  int
}

func (this *Cat) eat(s string) {
	fmt.Println("cat eat food:", s)
}

func (this *Cat) sleep() {
	fmt.Println("cat is sleeping")
}

func (this *Cat) run() {
	fmt.Println("cat is running")
}

type Dog struct {
	Name string
}

func (this *Dog) eat(s string) {
	fmt.Println("Dog eat food:", s)
}

func (this *Dog) sleep() {
	fmt.Println("Dog is sleeping")
}

func (this *Dog) run() {
	fmt.Println("Dog is running")
}
func main() {
	var ani AnimalIF
	ani = &Cat{"tom", 3}
	ani.eat("fish")
	ani.sleep()
	ani.run()

	fmt.Println("------")
	ani=&Dog{"aHuang"}
	ani.eat("shit")
	ani.sleep()
	ani.run()
}
