package main

import (
	"fmt"
	"math"
)

func main() {
	const a = 5 // 允许
	a = 89      // 不允许重新赋值

	fmt.Println("Hello, playground")
	var b = math.Sqrt(4)   // 允许
	const c = math.Sqrt(4) // 不允许,常量的值会在编译的时候确定。因为函数调用发生在运行时，所以不能将函数的返回值赋值给常量。

	fmt.Printf("a=%d,b=%f,c=%f", a, b, c)

	/* 字符串常量
	双引号中的任何值都是 Go 中的字符串常量。例如像 Hello World 或 Sam 等字符串在 Go 中都是常量。他们是无类型的。
	无类型的常量有一个与它们相关联的默认类型，并且当且仅当一行代码需要时才提供它。
	在声明中 var name = "Sam" ， name 需要一个类型，它从字符串常量 Sam 的默认类型中获取。
	*/
	var name = "Sam"
	fmt.Printf("type %T value %v", name, name)
	// 创建一个有类型常量
	const typedHello string = "Hello World"

	var defaultName = "Sam" // 允许
	type myString string
	var customName myString = "Sam" // 允许
	customName = defaultName        // 不允许

	/* 布尔常量
	布尔常量和字符串常量没有什么不同。他们是两个无类型的常量 true 和 false。字符串常量的规则适用于布尔常量，*/
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst       // 允许
	var customBool myBool = trueConst // 允许
	defaultBool = customBool          // 不允许

	/* 数字常量
	常量 a 是没有类型的，它的值是 5,a 的语法是通用的（它可以代表一个浮点数、整数甚至是一个没有虚部的复数），因此可以将其分配给任何兼容的类型。
	这些常量的默认类型可以被认为是根据上下文在运行中生成的。
	var intVar int = a 要求 a 是 int，所以它变成一个 int 常量。
	var complex64Var complex64 = a 要求 a 是 complex64，因此它变成一个复数类型。
	*/

	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	/* 数字表达式
	5.9 在语法中是浮点型，8 是整型，5.9/8 是允许的，因为两个都是数字常量。
	除法的结果是 0.7375 是一个浮点型，所以 t 的类型是浮点型
	*/
	var t = 5.9 / 8
	fmt.Printf("a's type %T value %v", t, t)
}
