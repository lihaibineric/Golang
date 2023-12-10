package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() { //为什么这个地方返回类型不是*
	fmt.Println("user is called ..")
	fmt.Printf("%v\n", this)
}

func DoFiledAndMethod(input interface{}) {
	//获取输入的类型type
	inputType := reflect.TypeOf(input)
	fmt.Println("input type is: ", inputType.Name())
	//output: input type is:  User

	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("input value is: ", inputValue)
	//output: input value is:  {1 eric 19}

	//通过type获取其中的字段
	/*
		1、获取interface中的reflect的type, 通过type得到numfield，进行遍历
		2、得到每个filed，就是数据类型
		3、通过filed中有一个interface()方法得到对应的value
	*/
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		// value:=inputType.Field{i}.interface()
		value := inputValue.Field(i).Interface()
		// fmt.Println(field) //{Id  int  0 [0] false}、{Name  string  8 [1] false}
		/*

			每个field表示的就是一行的元素内容，其中Name表示了这一行的变量名， Type表示的是这一行的类型
			通过value单独存在field里面，根据索引值寻找通过Interface()调用
		*/
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)

	}

	/* 如何遍历选择方法呢
	也是根据inputType来进行划分
	*/
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

func reflectNum(arg interface{}) {
	fmt.Println("type:", reflect.TypeOf(arg))
	fmt.Println("value:", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 1.2345
	reflectNum(num)
	user := User{1, "eric", 19}
	DoFiledAndMethod(user)
}
