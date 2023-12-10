package main

import (
	"fmt"
)

// type myint int  //本质上是声明一种数据类型

type Book struct {
	title string
	auth  string
}

func changeBook(book Book) {
	//这是一种值拷贝的改变，不能从地址层进行改变
	book.auth = "li"
}
func changeBook2(book *Book) {
	book.auth = "li"
}

func main() {
	var book1 Book //定义一个结构体
	book1.title = "golang"
	book1.auth = "zhang"
	//输出结构体中的内容
	fmt.Println(book1) //{golang zhang}
	changeBook(book1)
	fmt.Println(book1) //{golang zhang}
	changeBook2(&book1)
	fmt.Println(book1) //{golang li}

}
