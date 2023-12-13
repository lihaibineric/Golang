package main

import (
	"fmt"
)

func main(){
	//首先定义一个channel
	c := make(chan int) //定义一个整型 没有缓存的定义方式

	//channel作用于两个goroutine之间的通信

	//一个main的goroutine 新建一个goroutine
	go func(){
		defer fmt.Println("goroutine结束")
		fmt.Println("goroutine 正在运行...")

		c <- 666 //将666发送给channel c
	}()

	num := <-c //从c中接受数据，并赋值给num
	//如果接受到了那么就会打出来，实现了两个goroutine之间的通信
	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束...")

}