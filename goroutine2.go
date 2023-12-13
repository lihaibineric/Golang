package main

import (
	"fmt"
	"time"
	"runtime"
)

func main(){

	//用go创建承载一个形参为空，返回值也为空的一个函数
	//匿名函数
	go func() {
		defer  fmt.Println("A.defer")
		//匿名函数
		func(){
			defer fmt.Println("B.defer")

			//如果想退出当前的goroutine
			runtime.Goexit()
			/*  输出的值
				B.defer
				A.defer
			*/

			fmt.Println("B")
		}()
		//后面加一个小括号表示调用匿名函数

		fmt.Println("A")
	} ()

	//对于goroutine的形式没办法拿到返回的bool值 无法拿到当前的值
	go func(a int ,b int)bool{
		fmt.Println("a = ",a, ",b = ", b)
		return true
	}(10,20)

	//死循环
	for{
		time.Sleep(1*time.Second)
	}
}