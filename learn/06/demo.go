package main

import (
	"fmt"
	"math"
)

/*
函数学习

	func function_name( [parameter list] ) [return_types] {
	   函数体
	}
*/
func main() {
	var a int = 100
	var b int = 200
	var ret int

	/* 调用函数并显示函数返回值 */
	ret = max(a, b)

	fmt.Printf("最大值是 : %d\n", ret)

	a1, b1 := swap("hello", "world")
	fmt.Println(a1, b1)

	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

	fmt.Println("=======================")
	// 将匿名函数作为参数传递给其他函数
	calculate := func(operation func(int, int) int, x, y int) int {
		return operation(x, y)
	}

	// 定义一个匿名函数并将其赋值给变量add
	add := func(a, b int) int {
		return a + b
	}

	sum := calculate(add, 2, 8)
	fmt.Println("2 + 8 =", sum)

	// 也可以直接在函数调用中定义匿名函数
	difference := calculate(func(a, b int) int {
		return a - b
	}, 10, 4)
	fmt.Println("10 - 4 =", difference)

	fmt.Println("=======================")

	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的半径为: ", c1.getArea())

	cl := closure()
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())

	age := Age(25)
	age.String()
	age.Modify()
	age.String()
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// Go 函数可以返回多个值
func swap(x, y string) (string, string) {
	return y, x
}

/*
要自定义一个结构体，需要使用 type+struct 关键字组合
*/
type Circle struct {
	radius float64
}

/*
在 Go 语言中，方法和函数是两个概念，但又非常相似，不同点在于方法必须要有一个接收者，这个接收者是一个类型，这样方法就和这个类型绑定在一起，称为这个类型的方法。
该 method 属于 Circle 类型对象中的方法
*/
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

/*
func closure()：定义了一个名为 closure 的函数。
func() int：表示该函数返回一个函数，这个返回的函数没有参数并返回一个整数。
*/
func closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type Age uint

/*
值类型接收者
*/
func (age Age) String() {
	fmt.Println("the age is", age)
}

/*
指针类型接收者
定义的方法的接收者类型是指针，所以我们对指针的修改是有效的
可以简单地理解为，值接收者使用的是值的副本来调用方法，而指针接收者使用实际的值来调用方法
*/
func (age *Age) Modify() {
	*age = Age(30)
}

type person struct {
	name string
	age  uint
	addr address
}

type address struct {
	province string
	city     string
}

func (p person) String() string {
	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

type WalkRun interface {
	Walk()
	Run()
}

func (p *person) Walk() {
	fmt.Printf("%s能走\n", p.name)
}

func (p *person) Run() {
	fmt.Printf("%s能跑\n", p.name)
}
