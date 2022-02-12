package main

import "fmt"

type Animal struct {
	Name string
	Age  int
}

func (this *Animal) setName(name string) {
	this.Name = name
}

func (this *Animal) setAge(age int) {
	this.Age = age
}

func (this *Animal) run() {
	fmt.Println("ani is running")
}

type Bird struct {
	Animal
	Color string
}

func (this *Bird) setColor(color string) {
	this.Color = color
}

func (this *Bird) fly() {
	fmt.Println("bird is flying")
}

func main() {
	var ani Animal
	ani.setName("monster")
	ani.setAge(999)
	fmt.Printf("ani = %v\n", ani)
	ani.run()

	fmt.Println("-----")

	bird := Bird{Animal{"cannery", 3}, "red"}
	fmt.Printf("bird %v\n", bird)
	bird.setAge(5)
	fmt.Printf("bird %v\n", bird)
	bird.run()
	bird.fly()
}
