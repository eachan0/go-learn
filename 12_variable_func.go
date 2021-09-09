package main

import "fmt"

/**
可变参数函数是一种参数个数可变的函数。
如果函数最后一个参数被记作 ...T ，这时函数可以接受任意个 T 类型参数作为最后一个参数。
请注意只有函数的最后一个参数才允许是可变的。
可变参数函数的工作原理是把可变参数转换为一个新的切片。
*/
func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)

	nums := []int{89, 90, 95}
	// find(89, nums) // cannot use nums (type []int) as type int in argument to find
	/**
	nums ...int 意味它可以接受 int 类型的可变参数。
	nums 作为可变参数传入 find 函数。前面我们知道，这些可变参数参数会被转换为 int 类型切片然后在传入 find 函数中。
	但是在这里 nums 已经是一个 int 类型切片，编译器试图在 nums 基础上再创建一个切片,像下面这样
	find(89, []int{nums})
	这里之所以会失败是因为 nums 是一个 []int类型 而不是 int类型。
	有一个可以直接将切片传入可变参数函数的语法糖，你可以在在切片后加上 ... 后缀。
	如果这样做，切片将直接传入函数，不再创建新的切片
	*/
	find(89, nums...)

	/**
	使用了 ... ，welcome 切片本身会作为参数直接传入，不需要再创建一个新的切片。这样参数 welcome 将作为参数传入 changeStr 函数
	在 changeStr 函数中，切片的第一个元素被替换成 Go
	*/
	welcome := []string{"hello", "world"}
	changeStr(welcome...)
	fmt.Println(welcome)
}

/**
在下面程序中 func find(num int, nums ...int) 中的 nums 可接受任意数量的参数。
在 find 函数中，参数 nums 相当于一个整型切片。
*/
func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func changeStr(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}
