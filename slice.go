package main

import (
	"fmt"
)

func main() {
	/*
		声明切片之后的长度是3，同时初始化的值是1，2，3
	*/
	slice1 := []int{1, 2, 3}
	//%v表示的是打印出全部的表示信息
	fmt.Printf("len = %d, slice =%v\n", len(slice1), slice1)
	//直接对没有赋予空间的位置修改会出现位置越界
	//slice1[3]=999
	slice1 = make([]int, 4)
	//但是这个时候就只会重新赋予空间并复制0
	fmt.Printf("len = %d, slice =%v\n", len(slice1), slice1)

	/*
		声明slice是切片,但是没有分配空间
	*/
	var slice2 []int
	fmt.Printf("len = %d, slice =%v\n", len(slice2), slice2)

	//slice2[0]=2 //直接赋予数值会出现错误，越界
	slice2 = make([]int, 3)
	//开辟空间，但是默认值都是0
	fmt.Printf("len = %d, slice =%v\n", len(slice2), slice2)
	slice2[0] = 1000
	fmt.Printf("len = %d, slice =%v\n", len(slice2), slice2)
	//此时就修改成功

	/*
		声明slice类型，同时分配空间
	*/
	var slice3 = make([]int, 5)
	fmt.Printf("len = %d, slice =%v\n", len(slice3), slice3)

	/*
		判断一个slice切片是不是为0
	*/
	if slice1 == nil {
		fmt.Println("slice1 is null")
	} else {
		fmt.Println("slice1 is not null")
	}

	/*
		声明一个切片
	*/
	var slice4 = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice =%v\n", len(slice4), cap(slice4), slice4)
	//追加一个元素 这个1是元素的值
	slice4 = append(slice4, 1)
	fmt.Printf("len = %d, cap = %d, slice =%v\n", len(slice4), cap(slice4), slice4)
	//可以追加元素
	slice4 = append(slice4, 2)
	fmt.Printf("len = %d, cap = %d, slice =%v\n", len(slice4), cap(slice4), slice4)
	/* 如果这个时候超了cap的空间之后呢
	结果就是会以cap相同大小继续开辟相同的可使用空间大小
	*/
	slice4 = append(slice4, 1999)
	fmt.Printf("len = %d, cap = %d, slice =%v\n", len(slice4), cap(slice4), slice4)

	/*
		如果声明一个切片的时候不显式声明cap的值，那么cap和len的值相同
	*/
	var slice5 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice =%v\n", len(slice5), cap(slice5), slice5)

	/*
		切片的截取
	*/
	var slice6 = make([]int, 3)
	slice6[0] = 1
	slice6[2] = 33
	fmt.Printf("len = %d, cap = %d, slice =%v\n", len(slice6), cap(slice6), slice6)
	//截取的过程中是左闭右开
	s1 := slice6[0:2]
	// s1:=slice6[:]//表示截取全部
	fmt.Println(s1)
	//注意此时的s1的地址和slice6的位置一样了

	//copy
	var s2 = make([]int, 3)
	copy(s2, slice6)
	fmt.Println(s2) // 只会从头开始截取s2长度的，如果长了那么就是0，
	//否则就是选择slice6中相同长度的元素，从左到右边

}
