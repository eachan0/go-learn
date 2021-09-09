package rectangle

import (
	"fmt"
	"math"
)

/**
在 Go 中，任何以大写字母开头的变量或者函数都是被导出的名字。其它包只能访问被导出的函数和变量。
*/
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}

/*
 * 所有包都可以包含一个 init 函数。init 函数不应该有任何返回值类型和参数，在我们的代码中也不能显式地调用它。
 * init 函数可用于执行初始化任务，也可用于在开始执行之前验证程序的正确性。
 */
func init() {
	fmt.Println("rectangle package initialized")
}
