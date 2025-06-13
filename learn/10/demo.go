package main

import "fmt"

/*
Go 语言切片是对数组的抽象。

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")
与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
*/
func main() {
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	printSlice(numbers)
	if numbers == nil {
		fmt.Printf("切片是空的")
	}

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	fmt.Println("=============================")

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

	/* 同时添加多个元素 */
	number3 = append(number3, 2, 3, 4)
	printSlice(number3)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers4 := make([]int, len(numbers), cap(numbers)*2)
	var numbers5 []int
	/* 拷贝 numbers4 的内容到 numbers5 */
	copy(numbers4, numbers5)
	printSlice(numbers4)
	printSlice(numbers5)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
