package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg:=sync.WaitGroup{}
	wg.Add(10)
	for i:=0;i<10 ;i++  {
		go Go(&wg,i)
	}
	wg.Wait()
}

func Go(wg *sync.WaitGroup,index int)  {
	a:=1
	for i:=0;i<1000000 ;i++  {
		a+=i
	}
	fmt.Println(index,a)
	wg.Done()
}
