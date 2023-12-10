package lib1

import (
	"fmt"
)

// 当前package中的接口函数
func Lib1test() {
	fmt.Println("lib1test()")
}

func init() {
	fmt.Println("lib1.init()...")
}
