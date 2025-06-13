package main

import (
	"fmt"
	"math"
)

/*
接口
*/
func main() {
	c := Circle{5}
	var s Shape = c // 接口变量可以存储实现了接口的类型
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())

	printValue(42)          // int
	printValue("hello")     // string
	printValue(3.14)        // float64
	printValue([]int{1, 2}) // slice

	// 类型断言
	var i interface{} = "hello"
	str := i.(string) // 类型断言
	fmt.Println(str)  // 输出：hello

	// 类型选择（type switch）
	printType(42)
	printType("hello")
	printType(3.14)
	printType([]int{1, 2, 3})

	var rw ReaderWriter = File{}
	fmt.Println(rw.Read())
	rw.Write("hello, go")
}

// 定义接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义一个结构体
type Circle struct {
	Radius float64
}

// Circle 实现 Shape 接口
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 空接口 interface{} 是 Go 的特殊接口，表示所有类型的超集
func printValue(val interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", val, val)
}

func printType(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case float64:
		fmt.Println("Float:", v)
	default:
		fmt.Println("Unknown type")
	}
}

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type ReaderWriter interface {
	Reader
	Writer
}

type File struct {
}

func (f File) Read() string {
	return "Reading data"
}

func (f File) Write(data string) {
	fmt.Println("Writing data:", data)
}
