package main

import (
	"fmt"
	"log"
	// 导入了包，却不在代码中使用它，这在 Go 中是非法的。当这么做时，编译器是会报错的。
	// 其原因是为了避免导入过多未使用的包，从而导致编译时间显著增加。
	// 遇到这种情况就可以使用空白标识符 _。
	_ "net"
	"study/geometry/rectangle"
)

/*
 * 1. 包级别变量
 */
var rectLen, rectWidth float64 = -6, 7

/*
*2. init 函数会检查长和宽是否大于0
 */
func init() {
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {
	fmt.Println("Geometrical shape properties")
	/*Area function of rectangle package used*/
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	/*Diagonal function of rectangle package used*/
	fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))
}
