package main

import "fmt"

/**
map 是在 Go 中将值（value）与键（key）关联的内置类型。通过相应的键可以获取到值。

*/
func main() {
	// 通过向 make 函数传入键和值的类型，可以创建 map。make(map[type of key]type of value) 是创建 map 的语法。
	// map 的零值是 nil。如果你想添加元素到 nil map 中，会触发运行时 panic。因此 map 必须使用 make 函数初始化。
	var personSalary map[string]int
	// personSalary["key"] = 1 // panic: assignment to entry in nil map
	if personSalary == nil {
		fmt.Println("map is nil. Going to make one.")
		personSalary = make(map[string]int)
	}

	// 声明的时候初始化 map
	/*personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}*/

	// 添加元素
	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary)

	// 获取元素
	employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalary[employee])

	// 如果获取一个不存在的元素，会发生什么呢？map 会返回该元素类型的零值。
	fmt.Println("Salary of joe is", personSalary["joe"])

	// value, ok := map[key]
	// 上面就是获取 map 中某个 key 是否存在的语法。如果 ok 是 true，表示 key 存在，key 对应的值就是 value ，反之表示 key 不存在。
	newEmp := "joe"
	value, ok := personSalary[newEmp]
	if ok {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp, "not found")
	}

	// 遍历 map 中所有的元素需要用 for range 循环。
	// 有一点很重要，当使用 for range 遍历 map 时，不保证每次执行程序获取的元素顺序相同。
	fmt.Println("All items of a map")
	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}

	// 删除 map 中 key 的语法是 delete(map, key)。这个函数没有返回值。
	fmt.Println("map before deletion", personSalary)
	delete(personSalary, "steve")
	fmt.Println("map after deletion", personSalary)

	// 获取 map 的长度使用 len 函数。
	fmt.Println("length is", len(personSalary))

	// 和 slices 类似，map 也是引用类型。当 map 被赋值为一个新变量的时候，它们指向同一个内部数据结构。
	// 因此，改变其中一个变量，就会影响到另一变量。
	// map 作为函数参数传递时也会发生同样的情况。函数中对 map 的任何修改，对于外部的调用都是可见的。
	fmt.Println("Original person salary", personSalary)
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary)

	// map之间不能使用 == 操作符判断，== 只能用来检查 map 是否为 nil。
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	map2 := map1
	if map1 == map2 { // invalid operation: map1 == map2 (map can only be compared to nil)
	}
}
