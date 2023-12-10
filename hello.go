package main //程序的包 就是这个.go中包含的main函数名字

// import "fmt"//导入的包
// import "time"

import (
	"fmt"
	"time"
)

// main函数
func main() { // 左花括号必须同行
	fmt.Println("hello go!")
	//暂停1s
	time.Sleep(1 * time.Second)
}
