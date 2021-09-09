package main

import "fmt"

/**
在面向对象的领域里，接口一般这样定义：接口定义一个对象的行为。
接口只指定了对象应该做什么，至于如何实现这个行为（即实现细节），则由对象本身去确定。

在 Go 语言中，接口就是方法签名（Method Signature）的集合。
当一个类型定义了接口中的所有方法，我们称它实现了该接口。这与面向对象编程（OOP）的说法很类似。
接口指定了一个类型应该具有的方法，并由该类型决定如何实现这些方法。
*/
func main() {
	name := MyString("Sam Anderson")
	var v VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c", v.FindVowels())

	// 接口的实际用途
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)

	// 类型断言
	var s interface{} = 56
	//var s interface{} = "56" // panic: interface conversion: interface {} is string, not int
	var i interface{} = "Steven Paul"
	assert(s)
	assert(i)

	// 类型选择（Type Switch）
	findType("Naveen")
	findType(77)
	findType(89.98)

	// 还可以将一个类型和接口相比较。如果一个类型实现了接口，那么该类型与其实现的接口就可以互相比较。
	findType1("Naveen")
	p := PersonItf{
		name: "Naveen R",
		age:  25,
	}
	findType1(p)
}

// 接口定义
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

// 我们给接受者类型（Receiver Type） MyString 添加了方法 FindVowels() []rune。
//现在，我们称 MyString 实现了 VowelsFinder 接口。这就和其他语言（如 Java）很不同，其他一些语言要求一个类使用 implement 关键字，
// 来显式地声明该类实现了接口。而在 Go 中，并不需要这样。如果一个类型包含了接口中声明的所有方法，那么它就隐式地实现了 Go 接口。
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

// 计算工资接口
type SalaryCalculator interface {
	CalculateSalary() int
}

// 长期员工
type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

// 合同员工
type Contract struct {
	empId    int
	basicpay int
}

// 长期员工工资计算公式
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

// 合同员工工资计算公式
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

// 统计总工资
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		// 我们可以把接口看作内部的一个元组 (type, value)。 type 是接口底层的具体类型（Concrete Type），而 value 是具体类型的值。
		describe(v)
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d\n", expense)
}

// 没有包含方法的接口称为空接口。空接口表示为 interface{}。由于空接口没有方法，因此所有类型都实现了空接口。
func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

/**
类型断言用于提取接口的底层值（Underlying Value）。
在语法 v = i.(T) 中，接口 i 的具体类型是 T，该语法用于获得接口的底层值。不是T类型程序报错
v, ok := i.(T)
如果 i 的具体类型是 T，那么 v 赋值为 i 的底层值，而 ok 赋值为 true。
如果 i 的具体类型不是 T，那么 ok 赋值为 false，v 赋值为 T 类型的零值，此时程序不会报错。
*/
func assert(i interface{}) {
	// 不是 int 类型程序报错
	/*s := i.(int)
	fmt.Println(s)*/
	v, ok := i.(int)
	fmt.Println(v, ok)
}

/**
类型选择用于将接口的具体类型与很多 case 语句所指定的类型进行比较。它与一般的 switch 语句类似。
唯一的区别在于类型选择指定的是类型，而一般的 switch 指定的是值。
类型选择的语法类似于类型断言。类型断言的语法是 i.(T)，而对于类型选择，类型 T 由关键字 type 代替
*/
func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

type DescriberItf interface {
	DescribeItf()
}
type PersonItf struct {
	name string
	age  int
}

func (p PersonItf) DescribeItf() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType1(i interface{}) {
	switch v := i.(type) {
	case DescriberItf:
		v.DescribeItf()
	default:
		fmt.Printf("unknown type\n")
	}
}
