package main

import (
	"fmt"
)

//相当于go中任何类型都继承了interface{}的类型

//interface{}万能类型

func myFunc(arg interface{}) {
	fmt.Println("myfunc is called...")
	fmt.Println(arg)

	//interface{}如何确定类型是什么呢？

	//类型断言机制 用于判断是什么类型
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not a string")
		fmt.Printf("the value is %T\n", arg) //the value is main.Book
	} else {
		fmt.Println("arg is string type,is = ", value)
	}

}

type Book struct {
	auth string
}

func main() {
	book := Book{"golang"}
	myFunc(book) //{golang}
	myFunc(100)  //100
	myFunc("goland")
}
