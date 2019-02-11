package demo

import "fmt"

func PrintVar() {
	var age int // 未赋值
	fmt.Println("age 未赋值：", age)
	age = 12 // 赋值
	fmt.Println("age 赋值：", age)
	var age2 = 10 // 省略类型
	fmt.Println("age2 初始化赋值：", age2)

	// 批量赋值
	var width, height = 100, 50
	fmt.Println("width is", width, "height is", height)
	// 声明不同类型变量
	var (
		name = "navven"
		age3 = 20
	)
	fmt.Println("name is", name, "age3 is", age3)

	// 简短声明
	// 所有变量都必须有初始值
	name1, age4 := "naveen", 21
	fmt.Println("name1 is", name1, "age4 is", age4)

}
