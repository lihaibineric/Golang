package main

import (
	"fmt"
	"time"
)

//子goroutine
func  newTask(){
	i:=1
	for{
		i++
		fmt.Printf("new Goroutine: i =%d\n",i)
		time.Sleep(1*time.Second)
	}
}
//主goroutine
func main(){
	//创建一个goroutine去执行newTask()流程
	go newTask()
	//从而实现并发执行
	i :=0
	for {
		i++
		fmt.Printf("main Goroutine: i=%d\n",i)
		time.Sleep(1*time.Second)
	}

	// 如果主进程结束，那么子goroutine也会结束
	// fmt.Println("main goroutine exit")
}