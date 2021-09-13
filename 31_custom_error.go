package main

import (
	"fmt"
	"math"
)

/** 使用 New 函数创建自定义错误
创建自定义错误最简单的方法是使用 errors 包中的 New 函数。
*/
func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
	}

	area, err = circleArea1(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			if err.radiusNegative() {
				fmt.Println("半径小于0")
				return
			}
			fmt.Printf("Radius %0.2f is less than zero", err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		//return 0, errors.New("Area calculation failed, radius is less than zero")
		/** 使用 Errorf 给错误添加更多信息
		fmt 包中的 Errorf 函数了。Errorf 函数会根据格式说明符，规定错误的格式，并返回一个符合该错误的字符串。
		*/
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}

/**
使用结构体类型和字段提供错误的更多信息
错误还可以用实现了 error 接口的结构体来表示。这种方式可以更加灵活地处理错误。
错误类型的命名约定是名称以 Error 结尾。
*/
type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func circleArea1(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

// 使用结构体类型的方法来提供错误的更多信息
func (e *areaError) radiusNegative() bool {
	return e.radius < 0
}
