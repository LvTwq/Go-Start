package main

import (
	"fmt"
	"strconv"
)

/*
数值类型转换
*/
func main() {

	// 数值类型转换
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("平均值是 %f\n", mean)

	// 字符串类型转换
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("字符串转换失败")
	} else {
		fmt.Printf("字符串 %s 转换为数字 %d\n", str, num)
	}

	num2 := 123
	str2 := strconv.Itoa(num2)
	fmt.Printf("数字 %d 转换为字符串 %s\n", num2, str2)

	//类型断言
	var i interface{} = "hello"
	str, ok := i.(string)
	if ok {
		fmt.Printf("i is string, value is %s\n", str)
	} else {
		fmt.Printf("i is not string\n")
	}

	// 类型转换
	// 创建一个 StringWriter 实例并赋值给 Writer 接口变量
	var w Writer = &StringWriter{}
	// 将 Writer 接口类型转换为 StringWriter 类型
	sw := w.(*StringWriter)
	// 修改 StringWriter 的字段
	sw.str = "hello"
	fmt.Println(sw.str)
}

type Writer interface {
	Write([]byte) (int, error)
}

type StringWriter struct {
	str string
}

func (sw *StringWriter) Write(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}
