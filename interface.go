/*
一系列家族定义的接口，每个子类能够重写方法，实现同一个方法有多个接口表现形式
*/
package main

import (
	"fmt"
)

// 本质是一个指针多态
type Animal interface {
	//给出接口包含的多态的函数
	Sleep()
	GetColor() string
	GetType() string
}

// 定义一个具体的类
type Cat struct {
	/*如果继承一个接口interface 那么就不需要直接写出来继承，
	只需要进行实现就可以认为继承了这个接口inerface
	*/
	color string
}

/* 必须要完全重写所有的接口才能认为是多态满足 */
func (this *Cat) Sleep() {
	fmt.Println("cat sleep...")
}

func (this *Cat) GetColor() string {
	fmt.Printf("the cat color is %v\n", this.color)
	return this.color
}

func (this *Cat) GetType() string {
	fmt.Printf("the type is cat\n")
	return "Cat"
}

// 第二个多态的类
type Dog struct {
	//同样需要进行继承这个interface
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("dog sleep...")
}
func (this *Dog) GetColor() string {
	fmt.Printf("the dog color is %v\n", this.color)
	return this.color
}

func (this *Dog) GetType() string {
	fmt.Printf("the type is dog\n")
	return "Dog"
}

func showanimal(animal Animal) {
	animal.Sleep()
	// fmt.Println("color = ",animal.GetColor())
	// fmt.Println("type = ",animal.GetType())
	animal.GetType()
	animal.GetColor()
}

func main() {
	// var animal Animal //接口的数据类型，父类指针
	// animal=&Cat{"green"} //这里相当于是给类型中的变量进行赋值color
	// animal.Sleep()
	// animal.GetType()
	// animal.GetColor()

	// animal=&Dog{"Yellow"}//同样的道理
	// animal.Sleep()
	// animal.GetType()
	// animal.GetColor()
	cat := Cat{"Green"}
	dog := Dog{"Yellow"}

	//通过传递指针和引用来实现多态性
	showanimal(&cat)
	showanimal(&dog)

}
