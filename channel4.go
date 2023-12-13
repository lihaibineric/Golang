package main

import (
	"fmt"
)

func main(){
	c := make(chan int) //channel要通过make来触发，如果是nil channel 会发生阻塞

	go func(){
		defer fmt.Println("sub finished ...")
		for i := 0;i<5;i++{
			c <- i
			// close(c)
			//如果向已经关闭的channel发送数据就会出现panic的错误
		}
		// close 关闭一个channel
		close(c)
		//如果去掉关闭channel这个开关，那么会出现死锁的情况，就是main的进程数据都在等待塞数据
	}()

	
	//利用range进行简写
	//利用range不断迭代从channel中操作数据
	for data := range c {
		fmt.Println(data)
	}
	/* for{
		//如果ok为ture 那么channel没有关闭, 如果为false 那么已经被关闭
		if data, ok := <-c; ok{
			fmt.Println(data)
		}else{
			break
		}
	} */

	fmt.Println("Main finished ...")
}

/* 输出的值go run channel4.go
0
1
2
3
4
Main finished ...
sub finished ... */