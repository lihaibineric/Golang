package main

import (
	"fmt"
)

// const 来定义常量枚举类型
const (
	//关键字 iota
	/* BEIJING=0
	SHANGHAI=1
	SHENGZHENG=2 */
	BEIJING = 10 * iota // 默认为0
	SHANGHAI
	HANGZHOU
)
const (
	//这个iota的值表示的是和行数有关的数值，因此计算的时候用行来表示
	a, b = iota + 1, iota + 2 //每次新开一个关于iota的计算 那么后续全部会跟着这个计算方式下去
	c, d
	e, f
	g, h = iota * 2, iota * 3
	//iota只能在const之中使用
	i, k
)

func main() {
	//常量 只读属性
	const length int = 10
	fmt.Println("length =", length)
	fmt.Println("bj = ", BEIJING)
	fmt.Println("sh = ", SHANGHAI)
	fmt.Println("hz = ", HANGZHOU)
	fmt.Println("a=", a, "b=", b, "c=", c, "d=", d, "e=", e, "f=", f, "g=", g, "h=", h, "i=", i, "k=", k)
}
