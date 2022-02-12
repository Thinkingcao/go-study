package main

import "fmt"

func main() {
	channel := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("send data to channel,data=", i)
			channel <- i
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("get data from channel,data=", data)
	}

	fmt.Println("main finish")

}
