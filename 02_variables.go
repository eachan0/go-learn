package main

import "fmt"

func main() {
	var age int // 变量声明,默认0
	fmt.Println("my age is", age)
	age = 29 // 赋值
	fmt.Println("my age is", age)
	age = 54 // 赋值
	fmt.Println("my new age is", age)

	var age1 = 29 //自动类型推断,必须有初始值
	age2 := 23    // 上面的简短式

	fmt.Printf("age1=%d,age2=%d\n", age1, age2)

	// 声明多个变量,使用默认值
	// var width, height int
	// fmt.Printf("width=%d,height=%d", width, height)
	// 带初始值,自动推断
	var width, height = 200, 100
	fmt.Printf("width=%d,height=%d\n", width, height)
	// 多个变量赋值
	width, height = 400, 500
	fmt.Printf("width=%d,height=%d\n", width, height)

	var (
		name    = "naveen"
		age3    = 29
		height1 int
	)
	// name, age3, height1 := "naveen", 29, 0
	fmt.Println("my name is", name, ", age is", age3, "and height is", height1)
	// 简短式赋值
	name, age3, height1 = "tom", 30, 40
	fmt.Println("my name is", name, ", age is", age3, "and height is", height1)

	a, b := 20, 30 // 声明变量a和b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b已经声明，但c尚未声明
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // 给已经声明的变量b和c赋新值
	fmt.Println("changed b is", b, "c is", c)
}
