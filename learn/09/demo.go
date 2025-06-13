package main

import "fmt"

/*
结构体定义需要使用 type 和 struct 语句
*/
func main() {
	/*
		创建对象
	*/
	// 创建一个新结构体
	fmt.Println(Books{"Go Programming", "www.w3resource.com", "Go Programming", 6495407})
	// 也可以使用key:value结构
	fmt.Println(Books{title: "Go Programming", author: "www.w3resource.com", subject: "Go Programming", bookId: 6495407})
	// 忽略的字段为0或空
	fmt.Println(Books{title: "Go Programming", author: "www.w3resource.com"})

	/*
		访问结构体字段
	*/
	var Book1 Books /* 声明 Book1 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.bookId = 6495407

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.bookId)

	/*
		结构体作为函数参数
	*/
	/* 打印 Book1 信息 */
	printBook(Book1)

	/*
		结构体指针
	*/
	printBook1(&Book1)
}

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.bookId)
}
func printBook1(book *Books) {
	fmt.Printf("printBook1 Book title : %s\n", book.title)
	fmt.Printf("printBook1 Book author : %s\n", book.author)
	fmt.Printf("printBook1 Book subject : %s\n", book.subject)
	fmt.Printf("printBook1 Book book_id : %d\n", book.bookId)
}
