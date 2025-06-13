package main

import (
	"errors"
	"fmt"
)

/*
Go 语言的错误处理采用显式返回错误的方式，而非传统的异常处理机制
*/
func main() {
	err := errors.New("111111111")
	fmt.Println(err)

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	_, err1 := divide1(10, 0)
	if err1 != nil {
		fmt.Println(err1)
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

// 通过定义自定义类型，可以扩展 error 接口
type DivideError struct {
	Dividend int
	Divisor  int
}

func (e *DivideError) Error() string {
	return fmt.Sprintf("cannot divide %d by %d", e.Dividend, e.Divisor)
}

func divide1(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivideError{Dividend: a, Divisor: b}
	}
	return a / b, nil
}
