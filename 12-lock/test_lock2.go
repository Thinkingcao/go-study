package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	sum2 = 0
	mutex sync.Mutex
    wg sync.WaitGroup
)
func main() {
	//开启100个协程让sum+10
	run()
	doOnce()
}

func run()  {
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go add2()
	}
	//防止提前退出
	//time.Sleep(2 * time.Second)
	wg.Wait()
	fmt.Println("和为:",sum2)
}
func add2() {
	mutex.Lock()
	defer wg.Done()
	defer mutex.Unlock()
	sum2 ++
}
//这是 Go 语言自带的一个示例，虽然启动了 10 个协程来执行 onceBody 函数，但是因为用了 once.Do 方法，所以函数 onceBody 只会被执行一次。也就是说在高并发的情况下，sync.Once 也会保证 onceBody 函数只执行一次。
//
//sync.Once 适用于创建某个对象的单例、只加载一次的资源等只执行一次的场景。
func doOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	//用于等待协程执行完毕
	done := make(chan bool)
	//启动10个协程执行once.Do(onceBody)
	for i := 0; i < 10; i++ {
		go func() {
			//把要执行的函数(方法)作为参数传给once.Do方法即可
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

//10个人赛跑，1个裁判发号施令
func race(){
	cond :=sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)
	for i:=0;i<10; i++ {
		go func(num int) {
			defer  wg.Done()
			fmt.Println(num,"号已经就位")
			cond.L.Lock()
			cond.Wait()//等待发令枪响
			fmt.Println(num,"号开始跑……")
			cond.L.Unlock()
		}(i)
	}
	//等待所有goroutine都进入wait状态
	time.Sleep(2*time.Second)
	go func() {
		defer  wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家准备跑")
		cond.Broadcast()//发令枪响
	}()
	//防止函数提前返回退出
	wg.Wait()
}
//通过 sync.NewCond 函数生成一个 *sync.Cond，用于阻塞和唤醒协程；
//
//然后启动 10 个协程模拟 10 个人，准备就位后调用 cond.Wait() 方法阻塞当前协程等待发令枪响，这里需要注意的是调用 cond.Wait() 方法时要加锁；
//
//time.Sleep 用于等待所有人都进入 wait 阻塞状态，这样裁判才能调用 cond.Broadcast() 发号施令；
//
//裁判准备完毕后，就可以调用 cond.Broadcast() 通知所有人开始跑了。