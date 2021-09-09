package main

import "fmt"

/**
指针是一种存储变量内存地址（Memory Address）的变量。
指针变量的类型为 *T，该指针指向一个T类型的变量。
*/
func main() {
	b := 255
	// & 操作符用于获取变量的地址。
	var a *int
	// 指针的零值是 nil。
	fmt.Println("a is", a)
	if a == nil {
		a = &b
		fmt.Println("a after initialization is", a)
	}
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	// 指针的解引用可以获取指针所指向的变量的值。将 a 解引用的语法是 *a。
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("new value of b is", b)

	// 向函数传递指针参数
	fmt.Println("value of a before function call is", b)
	changeVal(a)
	fmt.Println("value of a after function call is", b)

	/**
	这种方式向函数传递一个数组指针参数，并在函数内修改数组。尽管它是有效的，但却不是 Go 语言惯用的实现方式。我们最好使用切片来处理。
	*/
	arr := [3]int{89, 90, 91}
	fmt.Println(arr)
	modifyArrByPointer(&arr)
	fmt.Println(arr)
	// 不要向函数传递数组的指针，而应该使用切片,代码更加简洁，也更符合 Go 语言的习惯。
	modifyArrBySlice(arr[:])
	fmt.Println(arr)

	// Go 并不支持其他语言（例如 C）中的指针运算。
	x := [...]int{109, 110, 111}
	y := &x
	// y++ // invalid operation: y++ (non-numeric type *[3]int)
	fmt.Println(y)
}

func changeVal(val *int) {
	*val = 55
}

func modifyArrByPointer(arr *[3]int) {
	(*arr)[0] = 90
	// arr[x] 是 (*arr)[x] 的简写形式
	arr[0] = 91
}
func modifyArrBySlice(sls []int) {
	sls[0] = 90
}
