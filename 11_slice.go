package main

import "fmt"

/**
尽管数组看上去似乎足够灵活，但是它们具有固定长度的限制，不可能增加数组的长度。
这就要用到 切片 了。事实上，在 Go 中，切片比传统数组更常见。
切片是由数组建立的一种方便、灵活且功能强大的包装（Wrapper）。切片本身不拥有任何数据。它们只是对现有数组的引用。
*/
func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] // 使用语法 a[start:end] 创建一个从 a 数组索引 start 开始到 end - 1 结束的切片。
	fmt.Println(b)

	c := []int{6, 7, 8} // 创建一个有 3 个整型元素的数组，并返回一个存储在 c 中的切片引用。
	fmt.Println(c)

	// 切片自己不拥有任何数据。它只是底层数组的一种表示。对切片所做的任何修改都会反映在底层数组中。
	arr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	slice := arr[2:5]
	fmt.Println("array before", arr)
	for i := range slice {
		slice[i]++
	}
	fmt.Println("array after", arr)

	// 当多个切片共用相同的底层数组时，每个切片所做的更改将反映在数组中。
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] // numa[:] 缺少开始和结束值。开始和结束的默认值分别为 0 和 len (numa)
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	/** 切片的长度和容量
	切片的长度是切片中的元素数。切片的容量是从创建切片索引开始的底层数组中元素数。
	*/
	fruitArray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitSlice := fruitArray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitSlice), cap(fruitSlice))
	fruitSlice = fruitSlice[:cap(fruitSlice)] // 重置容量
	fmt.Println("After re-slicing length is", len(fruitSlice), "and capacity is", cap(fruitSlice))

	/**
	使用 make 创建一个切片
	func make（[]T，len，cap）[]T 通过传递类型，长度和容量来创建切片。
	容量是可选参数, 默认值为切片长度。
	make 函数创建一个数组，并返回引用该数组的切片。
	*/
	i := make([]int, 5, 5) // 使用 make 创建切片时默认情况下这些值为零。
	fmt.Println(i)

	/** 追加切片元素
	正如我们已经知道数组的长度是固定的，它的长度不能增加。切片是动态的，使用 append 可以将新元素追加到切片上。
	append 函数的定义是 func append（s[]T，x ... T）[]T。
	x ... T 在函数定义中表示该函数接受参数 x 的个数是可变的。这些类型的函数被称为可变函数。
	有一个问题可能会困扰你。如果切片由数组支持，并且数组本身的长度是固定的，那么切片如何具有动态长度。
	以及内部发生了什么，当新的元素被添加到切片时，会创建一个新的数组。现有数组的元素被复制到这个新数组中，
	并返回这个新数组的新切片引用。现在新切片的容量是旧切片的两倍。
	*/
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // 容量最初是 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) // 容量翻了一番，变成了 6

	// 切片类型的零值为 nil。一个 nil 切片的长度和容量为 0。可以使用 append 函数将值追加到 nil 切片。
	var names []string
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}

	// 也可以使用 ... 运算符将一个切片添加到另一个切片。
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)

	/** 切片的函数传递
	切片包含长度、容量和指向数组第零个元素的指针。当切片传递给函数时，即使它通过值传递，指针变量也将引用相同的底层数组。
	因此，当切片作为参数传递给函数时，函数内所做的更改也会在函数外可见
	*/
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subTactOne(nos)                               // function modifies the slice
	fmt.Println("slice after function call", nos) // modifications are visible outside

	// 类似于数组，切片可以有多个维度。
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}

	/** 内存优化
	切片持有对底层数组的引用。只要切片在内存中，数组就不能被垃圾回收。在内存管理方面，这是需要注意的。
	让我们假设我们有一个非常大的数组，我们只想处理它的一小部分。然后，我们由这个数组创建一个切片，并开始处理切片。
	这里需要重点注意的是，在切片引用时数组仍然存在内存中。
	*/
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}
func subTactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}
func countries() []string {
	/**
	一种解决方法是使用 copy 函数 func copy(dst，src[]T)int 来生成一个切片的副本。
	这样我们可以使用新的切片，原始数组可以被垃圾回收。
	neededCountries := countries[:len(countries)-2 创建一个去掉尾部 2 个元素的切片 countries，
	将 neededCountries 复制到 countriesCpy 同时在函数的下一行返回 countriesCpy。
	现在 countries 数组可以被垃圾回收, 因为 neededCountries 不再被引用。
	*/
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}
