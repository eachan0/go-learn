package main

import "fmt"

/**
数组是同一类型元素的集合。例如，整数集合 5,8,9,79,76 形成一个数组。
Go 语言中不允许混合不同类型的元素，例如包含字符串和整数的数组。
（译者注：当然，如果是 interface{} 类型数组，可以包含任意类型）
*/
func main() {
	// 声明了一个长度为 3 的整型数组。数组中的所有元素都被自动赋值为数组类型的零值。
	var a [3]int
	fmt.Println(a)

	// 数组元素赋值
	a[0] = 12 // 数组的索引从 0 开始到 length - 1 结束。
	a[1] = 78
	a[2] = 50

	//  简略声明 来创建相同的数组。
	// a := [3]int{12, 78, 50}
	fmt.Println(a)

	// 在简略声明中，不需要将数组中所有的元素赋值。
	b := [3]int{12}
	fmt.Println(b)

	// 甚至可以忽略声明数组的长度，并用 ... 代替，让编译器为你自动计算长度
	c := [...]int{12, 78, 50}
	fmt.Println(c)

	// 数组的大小是类型的一部分。因此 [5]int 和 [25]int 是不同类型。数组不能调整大小，
	//d := [3]int{5, 78, 8}
	//var e [5]int
	//e = d              // 不允许 since [3]int and [5]int are distinct types
	//fmt.Println(e);

	/**
	数组是值类型
	Go 中的数组是值类型而不是引用类型。这意味着当数组赋值给一个新的变量时，该变量会得到一个原始数组的一个副本。
	如果对新变量进行更改，则不会影响原始数组。
	*/
	f := [...]string{"USA", "China", "India", "Germany", "France"}
	g := f // a copy of a is assigned to b
	g[0] = "Singapore"
	fmt.Println("a is ", f)
	fmt.Println("b is ", g)

	// 同样，当数组作为参数传递给函数时，它们是按值传递，而原始数组保持不变。
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)

	/**
	数组的长度
	通过将数组作为参数传递给 len 函数，可以得到数组的长度。
	*/
	h := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of a is", len(h))

	// 迭代数组
	// for 循环遍历数组中的元素，从索引 0 到 length of the array - 1。
	for l, i := len(h), 0; i < l; i++ {
		fmt.Printf("%d th element of a is %.2f\n", i, h[i])
	}
	// Go 提供了一种更好、更简洁的方法，通过使用 for 循环的 range 方法来遍历数组。range 返回索引和该索引处的值。
	sum := float64(0)
	for i, v := range h {
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("sum of all elements of a", sum)

	// 如果你只需要值并希望忽略索引，则可以通过用 _ 空白标识符替换索引来执行，同样值也可以被忽略。
	sum = 0
	for _, v := range h { // ignores index
		sum += v
	}
	fmt.Println("sum of all elements of a", sum)

	// 多维数组
	x := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, // 逗号是必需的。这是因为根据 Go 语言的规则自动插入分号。
	}
	printArray(x)
	var y [3][2]string
	y[0][0] = "apple" // 通过每个索引一个一个添加。这是另一种初始化二维数组的方法。
	y[0][1] = "samsung"
	y[1][0] = "microsoft"
	y[1][1] = "google"
	y[2][0] = "AT&T"
	y[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printArray(y)
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}

func printArray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}
