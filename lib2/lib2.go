package lib2

import (
	"fmt"
)

// 当前package中的接口函数
func Lib2test() {
	fmt.Println("lib2test()")
}

func init() {
	fmt.Println("lib2.init()...")
}
