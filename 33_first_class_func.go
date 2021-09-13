package main

import "fmt"

/** 什么是头等（第一类）函数？
支持头等函数（First Class Function）的编程语言，可以把函数赋值给变量，也可以把函数作为其它函数的参数或者返回值。
Go 语言支持头等函数的机制。
*/
func main() {
	// 匿名函数,由于没有名称，这类函数称为匿名函数（Anonymous Function）。
	a := func() {
		fmt.Println("hello world first class function")
	}
	// 调用
	a()
	fmt.Printf("%T\n", a)
	//要调用一个匿名函数，可以不用赋值给变量。
	func() {
		fmt.Println("hello world first class function")
	}()

	// 就像其它函数一样，还可以向匿名函数传递参数。
	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")

	// 用户自定义的函数类型,赋值了一个符合 add 类型签名的函数。
	var add add = func(a int, b int) int {
		return a + b
	}
	s := add(5, 6)
	fmt.Println("Sum", s)

	/** 高阶函数
	wiki 把高阶函数（Hiher-order Function）定义为：满足下列条件之一的函数：
	接收一个或多个函数作为参数
	返回值是一个函数
	*/
	// 把函数作为参数，传递给其它函数
	f := func(a, b int) int {
		return a + b
	}
	simple(f)

	// 在其它函数中返回函数
	simple := simple2()
	fmt.Println(simple(60, 7))

	/** 闭包
	闭包（Closure）是匿名函数的一个特例。当一个匿名函数所访问的变量定义在函数体的外部时，就称这样的匿名函数为闭包。
	*/
	a1 := appendStr()
	b := appendStr()
	fmt.Println(a1("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a1("Gopher"))
	fmt.Println(b("!"))

	// 头等函数的实际用途
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	stu := []student{s1, s2}
	filterStu := filter(stu, func(s student) bool {
		return s.grade == "B"
	})
	fmt.Println(filterStu)

	// 我们把这种对集合中的每个元素进行操作的函数称为 map 函数。
	arr := []int{5, 6, 7, 8, 9}
	r := iMap(arr, func(n int) int {
		return n * 5
	})
	fmt.Println(r)
}

type add func(a int, b int) int

func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func simple2() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}
