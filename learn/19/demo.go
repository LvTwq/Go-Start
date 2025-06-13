package main

import (
	"fmt"
	"time"
)

/*
Go 语言中没有线程的概念，只有协程，也称为 goroutine。相比线程来说，协程更加轻量
如果启动了多个 goroutine，它们之间通过 channel（通道）通信
*/
func main1() {
	go sayHello()
	for i := 0; i < 5; i++ {
		fmt.Println("main")
		time.Sleep(100 * time.Millisecond)
	}

	s := []int{7, 2, 8, -9, 4, 0}
	/*
				声明一个通道，chan 是一个关键字，表示是 channel 类型。后面的 int 表示 channel 里的数据是 int 类型
			一个 chan 的操作只有两种：发送和接收
		1. 接收：获取 chan 中的值，操作符为 <- chan
		2. 发送：向 chan 发送值，操作符为 chan <-
	*/
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// 从通道 c 中接收
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	fmt.Println("===============")

	//定义了一个可以存储整数类型的带缓冲通道，缓冲区大小为2
	ch := make(chan int, 2)
	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据，而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Go 遍历通道与关闭通道
	c1 := make(chan int, 10)
	go fibonacci(cap(c1), c)
	/*
		  range 函数遍历每个从通道接收到的数据，c1发完10个数据之后，就关闭了通道
		所以这里，range函数接收到10个数据之后，就结束了
		如果c1通道不关闭，那么range函数就不会结束，从而在接受第11个数据的时候就阻塞了
	*/
	for i := range c1 {
		fmt.Println(i)
	}

	c2 := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c2)
		}
		quit <- 0
	}()
	fibonacci1(c2, quit)
}

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("hello world")
		time.Sleep(100 * time.Millisecond)
	}
}

/*
ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据

	// 并把值赋给 v
*/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// 把 sum 发送到通道 c
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

/*
select 语句使得一个 goroutine 可以等待多个通信操作。
select 会阻塞，直到其中的某个 case 可以继续执行
*/

func fibonacci1(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

/*
在 Go 语言中，提倡通过通信来共享内存，而不是通过共享内存来通信，其实就是提倡通过 channel 发送接收消息的方式进行数据传递，而不是通过修改同一个变量。
所以在数据流动、传递的场景中要优先使用 channel，它是并发安全的，性能也不错。
*/
func main() {
	ch := make(chan string)
	go func() {
		fmt.Println("飞雪无情")
		ch <- "goroutine 完成"
	}()
	fmt.Println("我是 main goroutine")
	v := <-ch
	fmt.Println("接收到的chan中的值为：", v)

	fmt.Println("===============")
	//声明三个存放结果的channel
	firstCh := make(chan string)
	secondCh := make(chan string)
	threeCh := make(chan string)

	//同时开启3个goroutine下载
	go func() {
		firstCh <- downloadFile("firstCh")
	}()

	go func() {
		secondCh <- downloadFile("secondCh")
	}()

	go func() {
		threeCh <- downloadFile("threeCh")
	}()

	//开始select多路复用，哪个channel能获取到值，
	//就说明哪个最先下载好，就用哪个。
	select {
	case filePath := <-firstCh:
		fmt.Println(filePath)
	case filePath := <-secondCh:
		fmt.Println(filePath)
	case filePath := <-threeCh:
		fmt.Println(filePath)
	}
}

func downloadFile(chanName string) string {

	//模拟下载文件,可以自己随机time.Sleep点时间试试
	time.Sleep(time.Second)
	return chanName + ":filePath"
}
