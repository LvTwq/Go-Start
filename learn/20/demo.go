package main

import (
	"fmt"
	"sync"
)

// 共享的资源
var (
	sum   int
	mutex sync.Mutex
)

func main() {

	//开启100个协程让sum+10
	/*	for i := 0; i < 100; i++ {
			go add(10)
		}

		//防止提前退出
		time.Sleep(2 * time.Second)
		fmt.Println("和为:", sum)*/

	run()
}

func add(i int) {
	mutex.Lock()
	defer mutex.Unlock()
	// 可能同时会有多个协程交叉执行
	sum += i
}

func run() {
	var wg sync.WaitGroup
	//因为要监控110个协程，所以设置计数器为110
	wg.Add(110)
	for i := 0; i < 100; i++ {
		go func() {
			//计数器减1
			defer wg.Done()
			add(10)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			//计数器-1
			defer wg.Done()
			fmt.Println("sum=", readSum())
		}()
	}

	//一直等待，只要计数器值为0
	wg.Wait()

	doOnce()
}

var rwMutex sync.RWMutex

func readSum() int {
	//只获取读锁
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	b := sum

	return b

}

func doOnce() {
	/*
		保证代码只执行一次
		sync.Once 适用于创建某个对象的单例、只加载一次的资源等只执行一次的场景
	*/
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}

	// 用于等待协程执行完毕
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
