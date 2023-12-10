package main

import (
	"fmt"
)

func fool(a string, b int) int {
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	c := 100
	return c
}

func fool2(a string, b int) (int, int) {
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	return b, 1000
}

func fool3(a string, b int) (r1 int, r2 int) {
	fmt.Println("--------fool3--------")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	r1 = 89
	r2 = 98

	return r1, r2
}

func main() {
	c := fool("33", 100)
	fmt.Println("c= ", c)

	re1, re2 := fool2("abs", 1009)
	fmt.Println("re1=", re1, "re2=", re2)

	ret1, ret2 := fool3("hhh", 898)
	fmt.Println("ret1=", ret1, "ret2=", ret2)

	retu1, retu2 := fool4("helo", 798)
	fmt.Println("retu1=", retu1, "retu2=", retu2)

	fmt.Println("----55----")
	fr1, fr2 := fu5(55, "32")
	fmt.Println("fu5-re1:", fr1, "fu3-re2:", fr2)
	fmt.Println("xxx", fool("xx", 90))
}
func fool4(a string, b int) (r1, r2 int) {
	fmt.Println("----fool4-----")
	r1 = 1010
	r2 = 2030
	return r1, r2
}

func fu5(a int, b string) (r1 int, r2 string) {
	fmt.Println("---func5----")
	// r1=90
	// r2=10.1
	return r1, r2
}
