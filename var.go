/*
四种变量的声明方式
*/
package main

import (
	"fmt"
)

// 声明全局变量
var A int = 100

// B:="1sa"
func main() {
	//方法1：声明一个变量 x 类型 默认的值为0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a %T\n", a)

	var b = "abc"
	fmt.Println("b =", b)
	fmt.Printf("b type is %T\n", b)

	//3:自动匹配
	var c = 100.11
	fmt.Println("c =", c)
	fmt.Printf("c type is %T\n", c)

	//省去var关键字
	e := 100
	fmt.Println("e=", e)
	fmt.Printf("type e =%T\n", e)

	f := 100.1
	fmt.Println("f=", f)
	fmt.Printf("type f =%T\n", f)

	//输出全局变量
	fmt.Println("A=", A)
	fmt.Printf("type A =%T\n", A)

	//自动推理的无法用于全局变量
	// fmt.Println("B=",B)
	// fmt.Printf("type B =%T\n",B)

	var xx, yy int = 100, 200
	fmt.Println("xx=", xx, "yy=", yy)

	var xxx, yyy = 100, "sds"
	fmt.Println("xxx=", xxx, "yyy=", yyy)

	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv=", vv, "jj=", jj)
}
