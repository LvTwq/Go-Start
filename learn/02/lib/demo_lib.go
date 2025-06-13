package lib

import (
	"fmt"
	"time"
)

func Hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

/*
首字母大写，其他包也能访问
*/
func SayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("hello world")
		time.Sleep(100 * time.Millisecond)
	}
}
