package main

import "fmt"

/*
变量学习
*/
func main() {
	var a string = "hello"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	/*
		相当于
		var intVal int
		intVal =1
	*/
	intVal := 1
	fmt.Println(intVal)

	// 相当于 var f string = "Runoob"
	f := "Runoob"
	fmt.Println(f)

	// 这种不带声明格式的只能在函数体中出现
	g, h := 123, "hello"
	fmt.Println(x, y, a1, b1, c, f, g, h)
}

// 这种因式分解关键字的写法一般用于声明全局变量
var x, y int
var (
	a1 int
	b1 bool
)
