package main

import "fmt"

/**
函数是一块执行特定任务的代码。一个函数是在输入源基础上，通过执行一系列的算法，生成预期的输出。
函数声明通用语法:
func functionname(parametername type) returntype {
    // 函数体（具体实现的功能）
}
函数的声明以关键词 func 开始，后面紧跟自定义的函数名 functionname (函数名)。
函数的参数列表定义在 () 之间，返回值的类型则定义在之后的 returntype (返回值类型)处。
声明一个参数的语法采用 (parameter1 type, parameter2 type)即(参数1 参数1的类型,参数2 参数2的类型)的形式指定。
之后包含在 {} 之间的代码，就是函数体。函数中的参数列表和返回值并非是必须的
调用函数的语法为 functionname(parameters)
*/
func main() {

	price, no := 90, 6 // 定义 price 和 no,默认类型为 int
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice) // 打印到控制台上

	// 多返回值
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f\n", area, perimeter)

	// 命名返回值
	area, perimeter = rectProps1(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f\n", area, perimeter)

	/*空白符
	_在Go中被用作空白符，可以用作表示任何类型的任何值。
	继续以 rectProps 函数为例，该函数计算的是面积和周长。
	假使我们只需要计算面积，而并不关心周长的计算结果，该怎么调用这个函数呢？这时，空白符_就上场了。
	*/
	area, _ = rectProps(10.8, 5.6) // 返回值周长被丢弃
	fmt.Printf("Area %f \n", area)
}

// 表示这个函数不需要输入参数，且没有返回值
func functionName() {
}

// 如果有连续若干个参数，它们的类型一致，那么我们无须一一罗列，只需在最后一个参数后添加该类型
//func calculateBill(price int, no int) int {
func calculateBill(price, no int) int {
	return price * no // 商品总价 = 商品单价 * 数量, 返回总价
}

/*多返回值
Go语言支持一个函数可以有多个返回值。
写一个以矩形的长和宽为输入参数，计算并返回矩形面积和周长的函数。
矩形的面积是长度和宽度的乘积, 周长是长度和宽度之和的两倍。*/
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

/*命名返回值
从函数中可以返回一个命名值。一旦命名了返回值，可以认为这些值在函数第一行就被声明为变量了。
函数中的 return 语句没有显式返回任何值。
由于 area 和 perimeter 在函数声明中指定为返回值, 因此当遇到 return 语句时, 它们将自动从函数返回。*/
func rectProps1(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // 不需要明确指定返回值，默认返回 area, perimeter 的值
}
