package main

import (
	"fmt"
	"time"
)

func  main(){
	c := make(chan int ,3) //带有缓冲的channel

	//查看当前的channel中的缓冲的容量大小
	//len表示的是元素数量 cap表示的是缓冲的容量大小
	fmt.Println("len(c) = ", len(c), ", cap(c)", cap(c))

	//定义一个子goroutine
	go func(){
		defer fmt.Println("sub goroutine结束")
		
		// for i := 0; i<3; i++{
		for i := 0; i<4; i++{// 这种情况就会出现阻塞，因为超过了缓冲的容量
			c <- i
			fmt.Println("sub goroutine 正在运行: len(c) = ", len(c), ", cap(c) =  ", cap(c)," 发送的元素是= ",i)
		}
	}()

	time.Sleep(1*time.Second)

	//有缓存的话就不会出现阻塞的情况

	for i := 0;i < 3;i++{
		num := <-c //从c中接受数据并赋值给num
		fmt.Println("num = ", num)
	}
	fmt.Println("main 结束")
}

/* 
	输出的内容如下：
	len(c) =  0 , cap(c) 3
	sub goroutine 正在运行: len(c) =  1 , cap(c) =   3  发送的元素是=  0
	sub goroutine 正在运行: len(c) =  2 , cap(c) =   3  发送的元素是=  1
	sub goroutine 正在运行: len(c) =  3 , cap(c) =   3  发送的元素是=  2
	
	####sub goroutine结束
	
	num =  0
	num =  1
	num =  2
	main 结束
 */