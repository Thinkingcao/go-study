package main

import (
	"fmt"
	"time"
)

var sum = 0
func main() {
	//开启100个协程让sum+10
	for i := 0; i < 10000; i++ {
		go add()
	}
	//防止提前退出
	time.Sleep(2 * time.Second)
	fmt.Println("和为:",sum)
}
func add() {
	sum ++
}