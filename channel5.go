package main

import (
	"fmt"
)

func fibonacii(c , quit chan int){
	x,y:=1,1
	for {
		select{// select 能处理多个channel的状态
		case c <- x: //本质上就是判断会不会返回ok
			//如果c可写，则该case就会执行
			x = y
			y = x + y
		case <-quit ://如果这个时候quit可读，说明塞进了一个0
			fmt.Println("quit")
			return
		}
	}

}
func main(){
	c := make(chan int)
	quit := make(chan int)

	go func(){
		for i:= 0; i<6; i++{
			fmt.Println(<-c)
			//从c中读数据，如果没有那么就会阻塞
		}

		quit <- 0 //表示退出就会向quit中发送数字0
	}()

	//main go
	fibonacii(c,quit) //相当于main的goroutine 管道在做选择
}