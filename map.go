package main

import (
	"fmt"
)

func main() {
	//第一种声明方式
	var test1 map[string]string
	//在使用map之前需要进行make，make的作用是给map分配数据空间
	test1 = make(map[string]string, 10)
	test1["one"] = "php"
	test1["two"] = "golang"
	test1["three"] = "java"
	fmt.Println(test1) //map[one:php three:java two:golang]

	//第二种声明方式
	test2 := make(map[string]string)
	test2["one"] = "php"
	test2["two"] = "golang"
	test2["three"] = "java"
	fmt.Println(test2) //map[one:php three:java two:golang]

	//第三种声明方式
	test3 := map[string]string{
		"one":   "php",
		"two":   "golang",
		"three": "java", //注意map中的最有一个也要保留逗号
	}
	fmt.Println(test3) //map[one:php three:java two:golang]

	//双层的map 相当于市矩阵
	language := make(map[string]map[string]string)
	language["php"] = make(map[string]string, 2) //设置成长度是2的map
	language["php"]["id"] = "1"
	language["php"]["desc"] = "php是世界上最美的语言"
	language["golang"] = make(map[string]string)
	language["golang"]["id"] = "2"
	language["golang"]["desc"] = "golang 抗并发非常牛"
	//输出值
	/* map[golang:map[desc:golang 抗并发非常牛 id:2] php:map[desc:php是世界上最美的语言 id:1]] */

	fmt.Println(language)
	//增删查改
	val, key := language["php"]
	if key {
		fmt.Printf("%v", val)
	} else {
		fmt.Printf("no")
	}
	language["php"]["id"] = "3"     //修改php子元素的id值
	language["php"]["kfc"] = "vm50" //增加php中元素的值
	delete(language, "php")         //删除php子元素
	fmt.Println(language)

}
