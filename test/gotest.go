package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	mu.Lock()
	func(){
		fmt.Println("hello")
		mu.Unlock()
	}()
	mu.Lock()
	fmt.Println("world")

}