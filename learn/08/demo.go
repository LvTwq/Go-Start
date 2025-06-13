package main

import "fmt"

/*
指针：指向一个值的内存地址
*/
func main() {

	var a int = 10 /* 声明实际变量 */
	fmt.Printf("变量的地址：%x\n", &a)

	var ip *int /* 声明指针变量 */
	ip = &a     /* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址：%x\n", ip)
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}
