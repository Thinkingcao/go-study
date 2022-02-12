package main

import (
	"fmt"
	"time"
)

func main() {

	go task()
	fmt.Println("main func run")
	go func() {
		fmt.Println("run in niming func ")
	}()
	time.Sleep(2 * time.Second)
}

func task() {
	for i := 0; i < 10; i++ {
		fmt.Println("task index=", i)
	}
}
