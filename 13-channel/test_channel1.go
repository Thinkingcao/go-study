package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		defer fmt.Println("goroutine end...")
		fmt.Println("goroutine is running...")
		fmt.Println("goroutine send data{888} to chan...")
		channel <- 888 //将666 发送给c
	}()

	num := <-channel
	fmt.Println("main func get data from chan ，data=", num)
	fmt.Println("main func end...")
}
