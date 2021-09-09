package main

import "fmt"

// for 是 Go 语言唯一的循环语句。Go 语言中并没有其他语言比如 C 语言中的 while 和 do while 循环

/* 语法
for initialisation; condition; post {
}
*/
// 这三个组成部分，即初始化，条件和 post 都是可选的。
func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}
	fmt.Printf("\n")
	// break 语句用于在完成正常执行之前突然终止 for 循环，之后程序将会在 for 循环下一行代码开始执行。
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop\n")

	// continue 语句用来跳出 for 循环中当前循环。在 continue 语句后的所有的 for 循环语句都不会在本次循环中执行。
	// 循环体会在一下次循环中继续执行。
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")

	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

	count := 1
	// 无限循环
	for {
		fmt.Println("Hello World")
		count++
		if count > 10 {
			break
		}
	}
}
