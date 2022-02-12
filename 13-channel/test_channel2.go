package main

import "fmt"

func main() {
	channel:=make(chan int,2)
	go func() {
		defer fmt.Println("goroutine1 end...")
		fmt.Println("goroutine1 is running...")
		fmt.Println("goroutine1 send data{888} to chan...")
		channel <- 777 //将666 发送给c
	}()

	go func() {
		defer fmt.Println("goroutine2 end...")
		fmt.Println("goroutine2 is running...")
		fmt.Println("goroutine2 send data{888} to chan...")
		channel <- 888 //将666 发送给c
	}()

	for i := 0; i < 2; i++ {
		num := <-channel //从c中接收数据，并赋值给num
		fmt.Println("num = ", num)
	}

	fmt.Println("main 结束")
}