package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// bool
	boolA := true
	boolB := false
	fmt.Println("bool_a:", boolA, "bool_b:", boolB)
	boolC := boolA && boolB
	fmt.Println("c:", boolC)
	boolD := boolA || boolB
	fmt.Println("d:", boolD)

	/** 有符号整型
	int8：表示8位有符号整型
	大小：8位
	范围：-128～127

	int16：表示16位有符号整型
	大小：16位
	范围：-32768～32767

	int32：表示32位有符号整型
	大小：32位
	范围：-2147483648～2147483647

	int64：表示64位有符号整型
	大小：64位
	范围：-9223372036854775808～9223372036854775807

	int：根据不同的底层平台（Underlying Platform），表示 32 或 64 位整型。除非对整型的大小有特定的需求，否则你通常应该使用 int 表示整型。
	大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
	范围：在 32 位系统下是 -2147483648～2147483647，而在 64 位系统是 -9223372036854775808～9223372036854775807。*/

	/*无符号整型
	uint8：表示 8 位无符号整型
	大小：8 位
	范围：0～255

	uint16：表示 16 位无符号整型
	大小：16 位
	范围：0～65535

	uint32：表示 32 位无符号整型
	大小：32 位
	范围：0～4294967295

	uint64：表示 64 位无符号整型
	大小：64 位
	范围：0～18446744073709551615

	uint：根据不同的底层平台，表示 32 或 64 位无符号整型。
	大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
	范围：在 32 位系统下是 0～4294967295，而在 64 位系统是 0～18446744073709551615。*/

	var intA = 89
	intB := 95
	fmt.Println("value of int_a is", intA, "and int_b is", intB)
	// unsafe 包应该小心使用，因为使用 unsafe 包可能会带来可移植性问题。
	fmt.Printf("type of int_a is %T, size of int_a is %d\n", intA, unsafe.Sizeof(intA)) // int_a 的类型和大小
	fmt.Printf("type of int_b is %T, size of int_b is %d\n", intB, unsafe.Sizeof(intB)) // int_b 的类型和大小

	/*浮点型
	float32：32 位浮点数
	float64：64 位浮点数
	*/

	fA, fB := 5.67, 8.97
	fmt.Printf("type of f_a %T f_b %T\n", fA, fB)
	fSum := fA + fB
	fDiff := fA - fB
	fmt.Println("float_sum", fSum, "float_diff", fDiff)
	no1, no2 := 56, 89
	fmt.Println("float_sum", no1+no2, "float_diff", no1-no2)

	/* 复数类型
	complex64：实部和虚部都是 float32 类型的的复数。
	complex128：实部和虚部都是 float64 类型的的复数。
	*/
	c1 := complex(5, 7) // 内建函数 complex 用于创建一个包含实部和虚部的复数。
	c2 := 8 + 27i       // 还可以使用简短语法来创建复数：
	cAdd := c1 + c2
	fmt.Println("complex_sum:", cAdd)
	cMul := c1 * c2
	fmt.Println("complex_product:", cMul)

	/*其他数字类型
	byte 是 uint8 的别名。
	rune 是 int32 的别名。
	*/

	/*string类型
	在 Golang 中，字符串是字节的集合。
	*/
	first := "Naveen"
	last := "Ramanathan"
	fullname := first + " " + last
	fmt.Println("My name is", fullname)

	/**类型转换**/
	i := 55   // int
	j := 67.8 // float64
	// sum := i + j // 不允许 int + float64
	sum := i + int(j)
	fmt.Println(sum)

	k := 10
	// var l float64 = k // 若没有显式转换，该语句会报错
	var l = float64(k)
	fmt.Println("l=", l)
}
