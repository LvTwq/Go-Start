package main

import "fmt"

func main() {
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	// 简短声明语法
	numbers1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers1)

	// 如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度：
	balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)

	// 如果设置了数组的长度，我们还可以通过指定下标来初始化元素，将索引为 1 和 3 的元素初始化
	balance1 := [5]float32{1: 2.0, 3: 7.0}
	fmt.Println(balance1)
	// 修改
	balance1[4] = 50.0
	fmt.Println(balance1)

	fmt.Println("=============================")

	/* n 是一个长度为 10 的数组 */
	var n [10]int
	var i, j int
	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}
