package main

import (
	"fmt"
	"study/structs/computer"
)

/**
结构体是用户定义的类型，表示若干个字段（Field）的集合。有时应该把数据整合在一起，而不是让这些数据没有联系。这种情况下可以使用结构体。
*/
func main() {
	// 创建了一个命名的结构体 Employee
	// 通过指定每个字段名的值，我们定义了结构体变量 emp1。字段名的顺序不一定要与声明结构体类型时的顺序相同。
	// 在这里，我们改变了 lastName 的位置，将其移到了末尾。这样做也不会有任何的问题。
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}
	// 定义 emp2 时我们省略了字段名。在这种情况下，就需要保证字段名的顺序与声明结构体时的顺序相同。
	emp2 := Employee{"Thomas", "Paul", 29, 800}
	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	// 创建匿名结构体
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	fmt.Println("Employee 3", emp3)

	// 当定义好的结构体并没有被显式地初始化时，该结构体的字段将默认赋为零值。
	var emp4 Employee
	fmt.Println("Employee 4", emp4)

	// 当然还可以为某些字段指定初始值，而忽略其他字段。这样，忽略的字段名会赋值为零值。
	emp5 := Employee{
		firstName: "John",
		lastName:  "Paul",
	}
	fmt.Println("Employee 5", emp5)

	// 点号操作符 . 用于访问结构体的字段。
	emp6 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d\n", emp6.salary)

	// 还可以创建零值的 struct，以后再给各个字段赋值。
	var emp7 Employee
	emp7.firstName = "Jack"
	emp7.lastName = "Adams"
	fmt.Println("Employee 7:", emp7)

	// 还可以创建指向结构体的指针。
	// Go 语言允许我们在访问字段时，可以使用 emp8.field 来代替显式的解引用 (*emp8).field
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", (*emp8).firstName)
	fmt.Println("Last Name:", emp8.lastName)
	fmt.Println("Age:", (*emp8).age)
	fmt.Printf("Salary: $%d\n", emp8.salary)

	// 创建一个匿名字段的结构体
	ap := AnonymousPerson{"Naveen", 50}
	fmt.Println(ap)
	// 虽然匿名字段没有名称，但其实匿名字段的名称就默认为它的类型。Go 默认这些字段名是它们各自的类型
	var ap1 AnonymousPerson
	ap1.string = "naveen"
	ap1.int = 50
	fmt.Println(ap1)

	// 嵌套结构体
	var p Person
	p.name = "Naveen"
	p.age = 50
	p.address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)

	var p1 Person1
	p1.name = "Naveen"
	p1.age = 50
	p1.Address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p1.name)
	fmt.Println("Age:", p1.age)
	fmt.Println("City:", p1.city)   // city is promoted field (提升字段)
	fmt.Println("State:", p1.state) // state is promoted field (提升字段)

	// 导出结构体和字段
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	// spec.model // spec.model undefined (cannot refer to unexported field or method model)
	fmt.Println("Spec:", spec)

	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	/** 结构体相等性（Structs Equality）
	结构体是值类型。如果它的每一个字段都是可比较的，则该结构体也是可比较的。
	如果两个结构体变量的对应字段相等，则这两个变量也是相等的。
	*/
	name3 := name{firstName: "Steve", lastName: "Jobs"}
	name4 := name{}
	name4.firstName = "Steve"
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}
	// 如果结构体包含不可比较的字段，则结构体变量也不可比较。
	image1 := image{data: map[int]int{
		0: 155,
	}}
	image2 := image{data: map[int]int{
		0: 155,
	}}
	/*if image1 == image2 { // invalid operation: image1 == image2 (struct containing map[int]int cannot be compared)
		fmt.Println("image1 and image2 are equal")
	}*/
	fmt.Println(image1)
	fmt.Println(image2)
}

// Employee 这种成为命名的结构体（Named Structure）。
// 我们创建了名为 Employee 的新类型，而它可以用于创建 Employee 类型的结构体变量。
type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

// 声明结构体时也可以不用声明一个新类型，这样的结构体类型称为 匿名结构体（Anonymous Structure）。
var _ struct {
	// 通过把相同类型的字段声明在同一行，结构体可以变得更加紧凑
	firstName, lastName string
	age, salary         int
}

// AnonymousPerson :当我们创建结构体时，字段可以只有类型，而没有字段名。这样的字段称为匿名字段（Anonymous Field）
type AnonymousPerson struct {
	string
	int
}

type Address struct {
	city, state string
}

// Person 嵌套结构体
type Person struct {
	name string
	age  int
	// 结构体的字段有可能也是一个结构体。这样的结构体称为嵌套结构体。
	address Address
}

// Person1 结构体中有匿名的结构体类型字段，则该匿名结构体里的字段就称为提升字段。
// 这是因为提升字段就像是属于外部结构体一样，可以用外部结构体直接访问。
type Person1 struct {
	name string
	age  int
	Address
}
type name struct {
	firstName string
	lastName  string
}

type image struct {
	data map[int]int
}
