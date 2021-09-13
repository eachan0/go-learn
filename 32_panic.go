package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

/** 什么是 panic？
在 Go 语言中，程序中一般是使用错误来处理异常情况。对于程序中出现的大部分异常情况，错误就已经够用了。
但在有些情况，当程序发生异常时，无法继续运行。在这种情况下，我们会使用 panic 来终止程序。
当函数发生 panic 时，它会终止运行，在执行完所有的延迟函数后，程序控制返回到该函数的调用方。
这样的过程会一直持续下去，直到当前协程的所有函数都返回退出，然后程序会打印出 panic 信息，接着打印出堆栈跟踪（Stack Trace），最后程序终止
可以认为 panic 和 recover 与其他语言中的 try-catch-finally 语句类似，只不过一般我们很少使用 panic 和 recover。
而当我们使用了 panic 和 recover 时，也会比 try-catch-finally 更加优雅，代码更加整洁。
*/
func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")

	// panic，recover 和 Go 协程
	a()
	fmt.Println("normally returned from main")

	// 运行时 panic
	// 运行时错误（如数组越界）也会导致 panic。这等价于调用了内置函数 panic，其参数由接口类型 runtime.Error 给出。
	arr()
	fmt.Println("normally returned from arr")
}

func fullName(firstName *string, lastName *string) {
	defer fmt.Println("deferred call in fullName")
	defer recoverName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

/**
recover 是一个内建函数，用于重新获得 panic 协程的控制。
recover 函数的标签如下所示：
func recover() interface{}
只有在延迟函数的内部，调用 recover 才有用。在延迟函数内调用 recover，可以取到 panic 的错误信息，
并且停止 panic 续发事件（Panicking Sequence），程序运行恢复正常。如果在延迟函数的外部调用 recover，就不能停止 panic 续发事件。
*/
func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func a() {
	defer recovery()
	fmt.Println("Inside A")
	// 只有在相同的 Go 协程中调用 recover 才管用。recover 不能恢复一个不同协程的 panic。
	b()
	// go b() // 无法对 panic 生效
	time.Sleep(1 * time.Second)
}

func b() {
	fmt.Println("Inside B")
	panic("oh! B panicked")
}

func arr() {
	// 恢复一个运行时 panic
	defer r()
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		// 当我们恢复 panic 时，我们就释放了它的堆栈跟踪。实际上，在上述程序里，恢复 panic 之后，我们就失去了堆栈跟踪。
		// 有办法可以打印出堆栈跟踪，就是使用 Debug 包中的 PrintStack 函数。
		debug.PrintStack()
	}
}
