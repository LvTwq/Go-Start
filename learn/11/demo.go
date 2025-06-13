package main

import "fmt"

/*
Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。

for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环
*/

// 声明一个包含 2 的幂次方的切片
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// 遍历简单的切片
	for i, v := range pow {
		fmt.Printf("2 的 %d 次方是 %d\n", i, v)
	}

	// 迭代字符串时，返回每个字符的索引和 Unicode 代码点（rune）
	for i, c := range "hello" {
		fmt.Printf("hello 字符串的第 %d 个字符是 %q\n", i, c)
	}

	// for 循环的 range 格式可以省略 key 和 value
	map1 := make(map[int]float32)

	map1[1] = 1.1
	map1[2] = 2.2
	// 遍历 map1，读取 key 和 value
	for k, v := range map1 {
		fmt.Printf("map1[%d] = %f\n", k, v)
	}

	// 遍历 map1，只读取 key
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}

	// 遍历 map1，只读取 value
	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}

	// range 遍历从通道接收的值，直到通道关闭。
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	for v := range ch {
		fmt.Printf("ch value is: %d\n", v)
	}

	// 在遍历时可以使用 _ 来忽略索引或值。
	nums := []int{2, 3, 4}
	// 忽略索引
	for _, num := range nums {
		fmt.Printf("num is: %d\n", num)
	}
	// 忽略值
	for i := range nums {
		fmt.Printf("index is: %d\n", i)
	}
}
