package main

import (
	"fmt"
)

// 如果这个类如果是大写，那么其他的Package也能够访问
type Human struct {
	Name string //其中变量名称表示大写说明是能够对外界可见的public
	Age  int
	//如果是小写那么就是private的属性
}

func (this *Human) GetName() {
	fmt.Println(this.Name)
}
func (this *Human) SetName(newname string) {
	//只有是引用地址的传递的时候才是能够修改的
	this.Name = newname
}
func (this Human) SetName1(newname string) {
	//默认是值传递
	this.Name = newname
}

type Superman struct {
	Human //表示Superman继承了Human，同时这里没有C++中的公有保护等其他类型的继承
	//在子类中重新增加变量
	Level int
}

// 对于父类方法进行重写
func (this *Superman) GetName() {
	fmt.Println(this.Name)
	fmt.Println(this.Level)
}

// 子类中的新方法
func (this *Superman) LevelUp() {
	fmt.Println("level up")
	this.Level = this.Level + 1
	fmt.Println(this)
}

func main() {
	human := Human{Name: "zhang", Age: 99}
	human.SetName1("li")
	fmt.Println(human) //{zhang 99}
	human.SetName("li")
	fmt.Println(human) //{li 99}
	human.GetName()
	fmt.Println("-------------")
	superman := Superman{Human{"li4", 18}, 99}
	//父类方法的重写
	superman.GetName()
	//子类新方法
	superman.LevelUp()
	/* level up
	&{{li4 18} 100}	 */
	//父类方法原封不动
	superman.SetName("wang5")
	fmt.Println(superman) //{{wang5 18} 100}

	fmt.Println("+++++++++++++")
	//第二种继承类对象的声明
	var super Superman
	super.Name = "zhangmazi"
	super.Level = 100
	super.Age = 19
	fmt.Println(super)
}
