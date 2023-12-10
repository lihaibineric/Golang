package main

import (
	"fmt"
)

/* func swap(a int, b int){
	var temp int
	temp = a
	a = b
	b = temp
} */

func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func main() {
	var a int = 10
	var b int = 20
	fmt.Println("a=", a, "b=", b)
	swap(&a, &b)
	var p *int //一级指针
	p = &a
	fmt.Println("p =", *p)

	var pp **int //二级指针, 这个时候表示的是指针的指针地址
	pp = &p
	fmt.Println("pp=", **pp)
	fmt.Println("a=", a, "b=", b)
}
