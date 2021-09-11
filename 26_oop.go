package main

import "study/oop/employee"

func main() {
	//e := employee.Employee{FirstName: "Sam", LastName: "Adolf", TotalLeaves: 30, LeavesTaken: 20}

	/**可以看到，使用 Employee 创建的零值变量没有什么用。它没有合法的姓名，也没有合理的休假细节。
	在像 Java 这样的 OOP 语言中，是使用构造器来解决这种问题的。一个合法的对象必须使用参数化的构造器来创建。
	Go 并不支持构造器。如果某类型的零值不可用，需要程序员来隐藏该类型，避免从其他包直接访问。
	程序员应该提供一种名为 NewT(parameters) 的 函数，按照要求来初始化 T 类型的变量。
	按照 Go 的惯例，应该把创建 T 类型变量的函数命名为 NewT(parameters)。这就类似于构造器了。
	如果一个包只含有一种类型，按照 Go 的惯例，应该把函数命名为 New(parameters)， 而不是 NewT(parameters)。
	*/
	//var e employee.Employee
	//e.LeavesRemaining()

	/**
	把 Employee 结构体的首字母改为小写 e，也就是将 type Employee struct 改为了 type employee struct。
	通过这种方法，我们把 employee 结构体变为了不可引用的，防止其他包对它的访问。
	除非有特殊需求，否则也要隐藏所有不可引用的结构体的所有字段，这是 Go 的最佳实践。
	由于我们不会在外部包需要 employee 的字段，因此我们也让这些字段无法引用。
	*/
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()

	// 虽然 Go 不支持类，但结构体能够很好地取代类，而以 New(parameters) 签名的方法可以替代构造器。
}
