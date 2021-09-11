package main

import (
	"fmt"
	"sync"
)

/** 什么是 defer？
defer 语句的用途是：含有 defer 语句的函数，会在该函数将要返回之前，调用另一个函数。
*/
func main() {
	nums := []int{78, 109, 2, 563, 300}
	largest(nums)

	// defer 不仅限于函数的调用，调用方法也是合法的。
	p := person29{
		firstName: "John",
		lastName:  "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ")

	// 在 Go 语言中，并非在调用延迟函数的时候才确定实参，而是当执行 defer 语句的时候，就会对延迟函数的实参进行求值。
	a := 5
	/** defer 栈
	当一个函数内多次调用 defer 时，Go 会把 defer 调用放入到一个栈中，随后按照后进先出（Last In First Out, LIFO）的顺序执行
	*/
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)

	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
		//如果你仔细观察，会发现 wg.Done() 只在 area 函数返回的时候才会调用。
		//wg.Done() 应该在 area 将要返回之前调用，并且与代码流的路径（Path）无关，
		//因此我们可以只调用一次 defer，来有效地替换掉 wg.Done() 的多次调用。
		wg.Add(1)
		go v.areaWithDefer(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

type person29 struct {
	firstName string
	lastName  string
}

func (p person29) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		wg.Done()
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		wg.Done()
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
	wg.Done()
}

func (r rect) areaWithDefer(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
}
