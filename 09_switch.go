package main

import "fmt"

func main() {
	finger := 4
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	// case 4: // 重复项, duplicate case 4 in switch
	//	fmt.Println("Another Ring")
	case 5:
		fmt.Println("Pinky")
	}

	// default
	// finger 变量的作用域仅限于这个 switch 内。
	switch finger := 8; finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default: // 默认情况
		fmt.Println("incorrect finger number")
	}

	// 多表达式判断
	switch letter := "i"; letter {
	case "a", "e", "i", "o", "u": // 一个选项多个表达式
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}

	// 无表达式的 switch
	num := 75
	switch { // 表达式被省略了
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("num is greater than 100")
	}

	// Fallthrough 语句
	// 在 Go 中，每执行完一个 case 后，会从 switch 语句中跳出来，不再做后续 case 的判断和执行。
	// 使用 fallthrough 语句可以在已经执行完成的 case 之后，把控制权转移到下一个 case 的执行代码中。
	// fallthrough 语句应该是 case 子句的最后一个语句。
	// 如果它出现在了 case 语句的中间，编译器将会报错：fallthrough statement out of place
	switch num := number(); { // num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
	}
}

func number() int {
	num := 15 * 5
	return num
}
