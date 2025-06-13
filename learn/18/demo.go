package main

import (
	"errors"
	"fmt"
)

// 定义一个 DivideError 结构
type DivideError struct {
	dividee int
	divider int
}

// 实现 `error` 接口
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

var ErrNotFound = errors.New("not found")

func findItem(id int) error {
	return fmt.Errorf("database error: %w", ErrNotFound)
}

func main() {
	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

	err := findItem(1)
	// 检查某个错误是否是特定错误或由该错误包装而成。
	if errors.Is(err, ErrNotFound) {
		fmt.Println("item not found")
	} else {
		fmt.Println("unknown error:", err)
	}

	err1 := getError()
	var myErr *MyError
	//将错误转换为特定类型以便进一步处理
	if errors.As(err1, &myErr) {
		fmt.Printf("Custom error - Code: %d, Msg: %s\n", myErr.Code, myErr.Msg)
	}

	fmt.Println("Starting program...")
	safeFunction()
	fmt.Println("Program continued after panic")

	moreDefer()
}

type MyError struct {
	Code int
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Code: %d, Msg: %s", e.Code, e.Msg)
}

func getError() error {
	return &MyError{Code: 404, Msg: "Not Found"}
}

/*
Go 的 panic 用于处理不可恢复的错误，recover 用于从 panic 中恢复。
panic:导致程序崩溃并输出堆栈信息，常用于程序无法继续运行的情况
recover:捕获 panic，避免程序崩溃
*/
func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("something went wrong")
}

/*
多个 defer 语句的执行顺序依照后进先出的原则，输出如下：
函数自身代码
Three defer
Second defer
First defer
*/
func moreDefer() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Three defer")
	fmt.Println("函数自身代码")
}
