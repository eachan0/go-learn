package main

import (
	"fmt"
	"math"
)

/**
方法其实就是一个函数，在 func 这个关键字和方法名中间加入了一个特殊的接收器类型。
接收器可以是结构体类型或者是非结构体类型。接收器是可以在方法的内部访问的。
func (t Type) methodName(parameter list) {
}
上面的代码片段创建了一个接收器类型为 Type 的方法 methodName。
*/

// 在非结构体上的方法
// 到目前为止，我们只在结构体类型上定义方法。也可以在非结构体类型上定义方法， 但是有一个问题。
// 为了在一个类型上定义一个方法，方法的接收器类型定义和方法的定义应该在同一个包中。
// 我们尝试把一个 add 方法添加到内置的类型 int。
// 这是不允许的，因为 add 方法的定义和 int 类型的定义不在同一个包中。
// 该程序会抛出编译错误 cannot define new methods on non-local type int。
/*
func (a int) add(b int) {
}
*/

func main() {
	emp1 := EmployeeMethod{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	// 调用 Employee 类型的 displaySalary() 方法
	emp1.displaySalary()
	/**
	调用 displaySalary函数，Employee 结构体被当做参数传递给它。这个程序也产生完全相同的输出
	为什么我们已经有函数了还需要方法呢？
	Go 不是纯粹的面向对象编程语言，而且Go不支持类。因此，基于类型的方法是一种实现和类相似行为的途径。
	相同的名字的方法可以定义在不同的类型上，而相同名字的函数是不被允许的。
	*/
	displaySalary(emp1)
	// 假设我们有一个 Rectangle 和 Circle 结构体。可以在 Square 和 Circle 上分别定义一个 Area 方法。
	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of circle %f", c.Area())

	/* 指针接收器与值接收器
	到目前为止，我们只看到了使用值接收器的方法。还可以创建使用指针接收器的方法。
	值接收器和指针接收器之间的区别在于，在指针接收器的方法内部的改变对于调用者是可见的，然而值接收器的情况不是这样的。
	*/
	e := EmployeeMethod{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("\nEmployee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)
	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	(&e).changeAge(51)
	// Go语言让我们可以直接使用 e.changeAge(52)。e.changeAge(52) 会自动被Go语言解释为 (&e).changeAge(52)
	e.changeAge(52)
	fmt.Printf("\nEmployee age after change: %d\n", e.age)
	/** 那么什么时候使用指针接收器，什么时候使用值接收器？
	一般来说，指针接收器可以使用在：对方法内部的接收器所做的改变应该对调用者可见时。
	指针接收器也可以被使用在如下场景：当拷贝一个结构体的代价过于昂贵时。考虑下一个结构体有很多的字段。
	在方法内使用这个结构体做为值接收器需要拷贝整个结构体，这是很昂贵的。
	在这种情况下使用指针接收器，结构体不会被拷贝，只会传递一个指针到方法内部使用。
	在其他的所有情况，值接收器都可以被使用。
	*/

	/* 匿名字段的方法
	属于结构体的匿名字段的方法可以被直接调用，就好像这些方法是属于定义了匿名字段的结构体一样。
	*/
	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			city:  "Los Angeles",
			state: "California",
		},
	}
	p.fullAddress() //访问 address 结构体的 fullAddress 方法

	/* 在方法中使用值接收器 与 在函数中使用值参数
	当一个函数有一个值参数，它只能接受一个值参数。
	当一个方法有一个值接收器，它可以接受值接收器和指针接收器。
	*/
	re := Rectangle{
		length: 10,
		width:  5,
	}
	Area(re)
	re.Area()
	rep := &re
	// Area(rep)  //  cannot use rep (type *Rectangle) as type Rectangle in argument to Area
	rep.Area() //通过指针调用值接收器

	/* 在方法中使用指针接收器 与 在函数中使用指针参数
	和值参数相类似，函数使用指针参数只接受指针，而使用指针接收器的方法可以使用值接收器和指针接收器。
	*/
	perimeter(rep)
	// perimeter(re) // cannot use re (type Rectangle) as type *Rectangle in argument to perimeter
	re.perimeter()
	rep.perimeter()

	// 在非结构体上的方法
	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}

type EmployeeMethod struct {
	name        string
	salary, age int
	currency    string
}

/*
使用值接收器的方法。
*/
func (e EmployeeMethod) changeName(newName string) {
	e.name = newName
}

/*
使用指针接收器的方法。
*/
func (e *EmployeeMethod) changeAge(newAge int) {
	e.age = newAge
}

/*
  displaySalary() 方法将 Employee 做为接收器类型
*/
func (e EmployeeMethod) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

/*
displaySalary()方法被转化为一个函数，把 Employee 当做参数传入。
*/
func displaySalary(e EmployeeMethod) {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

/**
方法
*/
func (r Rectangle) Area() (area int) {
	area = r.length * r.width
	fmt.Printf("Area Method result: %d\n", area)
	return
}

/**
函数
*/
func Area(r Rectangle) {
	fmt.Printf("Area Function result: %d\n", r.length*r.width)
}

func perimeter(r *Rectangle) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *Rectangle) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
}

// 为内置类型 int 创建一个类型别名，
type myInt int

// 然后创建一个以该类型别名为接收器的方法。
func (a myInt) add(b myInt) myInt {
	return a + b
}
