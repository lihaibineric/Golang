package main

import (
	"fmt"
)

func func1() {
	fmt.Println("111")
}

func func2() {
	fmt.Println("222")
}

func func3() {
	fmt.Println("333")
}

func deferFunc() int {
	fmt.Println("defer func called....")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called....")
	return 0
}

func ReturnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	//写入defer关键字

	//表示的是某个函数或者某个流程结束之前需要触发的 生命周期完全结束才会调用defer
	defer fmt.Println("main end 1") //相当于C++中的析构函数 or final
	defer fmt.Println("main end 2")
	/* 此时会先调用2 再调用1
	本质上就是检索到defer之后，就放进一个栈中，先进后出
	*/
	fmt.Println("hello go 1")
	fmt.Println("hello go 2")

	//defer调用的逻辑是一种压栈输出的方式
	defer func1()
	defer func2()
	defer func3()

	//defer 和return在一个函数中说明return的调用要比defer更前；defer永远是最后的
	ReturnAndDefer()
}
