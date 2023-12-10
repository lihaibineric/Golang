package main

import (
	"fmt"
)

func printArray2(myarray []int) {
	//这是一个引用传递，而不是值传递
	for _, value := range myarray {
		fmt.Println("dynamic array:=", value)
	}
	myarray[0] = 100
}

func pirntArray(myarray [4]int) {
	//传递一个固定数组长度
	//本质上是一个值拷贝的过程
	for index, value := range myarray {
		fmt.Println("index = ", index, "value = ", value)
	}
}

func main() {
	//固定长度的数组
	var myarray1 [10]float32

	/* for i :=0;i<10;i++{

	} */

	myarray2 := [10]int{1, 2, 3, 4}
	myarray3 := [4]int{1, 2, 3, 4}

	for i := 0; i < len(myarray1); i++ {
		fmt.Println(myarray1[i])
	}

	for index, value := range myarray2 {
		fmt.Println("index= ", index, "value= ", value)
	} // 使用关键字range来遍历元素

	fmt.Printf("myarray1 type = %T\n", myarray1)
	fmt.Printf("myarray2 type = %T\n", myarray2)
	//静态数组的值拷贝
	pirntArray(myarray3)

	//动态数组部分*************slice
	myarray := []int{1, 2, 3, 4}          //动态数组，切片
	fmt.Printf("array[] = %T\n", myarray) //输出的长度是空
	printArray2(myarray)
	//此时的动态数组中的值发生了改变
	for _, value := range myarray {
		fmt.Println("dddd=", value)
	}

}
