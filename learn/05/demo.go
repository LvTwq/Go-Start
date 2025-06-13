package main

import "fmt"

/*
运算符	描述				实例
&		返回变量存储地址	&a; 将给出变量的实际地址。
*		指针变量。		*a; 是一个指针变量
*/
func main() {
	var a int = 4
	// 表示 ptr 是一个指向 int 类型的指针
	var ptr *int
	// 'ptr' 包含了 'a' 变量的地址
	ptr = &a

	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)

	fmt.Printf("a 的值为  %d\n", a)
	fmt.Printf("a 的值为  %d\n", ptr)
	// value 是 4，即 ptr 指向的值
	value := *ptr
	fmt.Printf("*ptr 为 %d\n", value)
}
