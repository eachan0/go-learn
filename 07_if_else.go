package main

import "fmt"

// if 是条件语句。if 语句的语法是
/*
if condition {
}
or
if condition {
} else if condition {
} else {
}
*/
// if 还有另外一种形式，它包含一个 statement 可选语句部分，该组件在条件判断之前运行
// else 语句应该在 if 语句的大括号 } 之后的同一行中。如果不是，编译器会不通过。原因是 Go 语言的分号是自动插入。
func main() {

	num := 10
	if num%2 == 0 { //checks if number is even
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}

	// num1 在 if 语句中进行初始化，num1 只能从 if 和 else 中访问。
	// 也就是说 num1 的范围仅限于 if else 代码块。
	// 如果我们试图从其他外部的 if 或者 else 访问 num1,编译器会不通过。
	if num1 := 11; num1%2 == 0 { //checks if number is even
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}

	if num = 99; num <= 50 {
		fmt.Println("number is less than or equal to 50")
	} else if num >= 51 && num <= 100 {
		fmt.Println("number is between 51 and 100")
	} else {
		fmt.Println("number is greater than 100")
	}
}
