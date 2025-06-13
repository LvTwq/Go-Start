package main

import (
	"fmt"
	"unsafe"
)

/*
常量学习
*/
func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	// 多重赋值
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d\n", area)
	fmt.Println("=============================")
	fmt.Println(a, b, c)
	fmt.Println("=============================")
	fmt.Println(a1, b1, c1)
	fmt.Println("=============================")
	const (
		a2 = iota //0
		b2        //1
		c2        //2
		d2 = "ha" //独立值，iota += 1
		e2        //"ha"   iota += 1
		f2 = 100  //iota +=1
		g2        //100  iota +=1
		h2 = iota //7,恢复计数
		i2        //8
	)
	fmt.Println(a2, b2, c2, d2, e2, f2, g2, h2, i2)
}

// 常量还可以用作枚举：
const (
	Unknown = 0
	Female  = 1
	Male    = 2

	a1 = "abc"
	b1 = len(a1)
	// 一个变量或类型在内存中占用的字节数
	c1 = unsafe.Sizeof(a1)
)
