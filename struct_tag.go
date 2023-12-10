package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	/*
		通过以下形式给结构体中的变量添加标签
		作用： 其他包在调用这个当前包的时候对于某个属性的一个说明，指示某个包在具体使用中的作用
	*/
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findtag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagstring := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info:", tagstring, "doc", tagdoc)
	}
}

func main() {
	var re resume
	findtag(&re)
}
