package main

import "fmt"

func main() {
	var d1 DescriberItf2
	p1 := PersonItf2{"Sam", 25}
	// 在讨论方法的时候就已经提到过，使用值接受者声明的方法，既可以用值来调用，也能用指针调用。
	// 不管是一个值，还是一个可以解引用的指针，调用这样的方法都是合法的。
	d1 = p1
	d1.DescribeItf2()
	p2 := PersonItf2{"James", 32}
	d1 = &p2
	d1.DescribeItf2()

	var d2 DescriberItf2
	a := AddressItf2{"Washington", "USA"}

	/* 如果下面一行取消注释会导致编译错误：
	   cannot use a (type Address) as type Describer
	   in assignment: Address does not implement
	   Describer (Describe method has pointer
	   receiver)
	*/
	// 其原因是：对于使用指针接受者的方法，用一个指针或者一个可取得地址的值来调用都是合法的。
	// 但接口中存储的具体值（Concrete Value）并不能取到地址，因此在第 45 行，对于编译器无法自动获取 a 的地址，于是程序报错。
	//d2 = a

	d2 = &a // 这是合法的
	// 因为 AddressItf2 类型的指针实现了 DescriberItf2 接口
	d2.DescribeItf2()

	// 实现多个接口
	e := EmployeeItf2{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var s SalaryCalculatorItf2 = e
	s.DisplaySalary()
	var l LeaveCalculatorItf2 = e
	fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())

	// 接口的嵌套
	// 尽管 Go 语言没有提供继承机制，但可以通过嵌套其他的接口，创建一个新接口。
	var empOp EmployeeOperations = e
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())

	// 接口的零值是 nil。对于值为 nil 的接口，其底层值（Underlying Value）和具体类型（Concrete Type）都为 nil。
	var d3 DescriberItf2
	if d3 == nil {
		fmt.Printf("d1 is nil and has type %T value %v\n", d3, d3)
	}
	// 对于值为 nil 的接口，由于没有底层值和具体类型，当我们试图调用它的方法时，程序会产生 panic 异常。
	d3.DescribeItf2()
}

type DescriberItf2 interface {
	DescribeItf2()
}
type PersonItf2 struct {
	name string
	age  int
}

func (p PersonItf2) DescribeItf2() { // 使用值接受者实现
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type AddressItf2 struct {
	state   string
	country string
}

func (a *AddressItf2) DescribeItf2() { // 使用指针接受者实现
	fmt.Printf("State %s Country %s\n", a.state, a.country)
}

type SalaryCalculatorItf2 interface {
	DisplaySalary()
}

type LeaveCalculatorItf2 interface {
	CalculateLeavesLeft() int
}

type EmployeeItf2 struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e EmployeeItf2) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d\n", e.firstName, e.lastName, e.basicPay+e.pf)
}

func (e EmployeeItf2) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

type EmployeeOperations interface {
	SalaryCalculatorItf2
	LeaveCalculatorItf2
}
