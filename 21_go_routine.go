package main

import (
	"fmt"
	"time"
)

/**
Go 协程是与其他函数或方法一起并发运行的函数或方法。Go 协程可以看作是轻量级线程。
与线程相比，创建一个 Go 协程的成本很小。因此在 Go 应用中，常常会看到有数以千计的 Go 协程并发地运行。
Go 协程相比于线程的优势
1. 相比线程而言，Go 协程的成本极低。堆栈大小只有若干 kb，并且可以根据应用的需求进行增减。而线程必须指定堆栈的大小，其堆栈是固定不变的。
2. Go 协程会复用（Multiplex）数量更少的 OS 线程。即使程序有数以千计的 Go 协程，也可能只有一个线程。
如果该线程中的某一 Go 协程发生了阻塞（比如说等待用户输入），那么系统会再创建一个 OS 线程，并把其余 Go 协程都移动到这个新的 OS 线程。
所有这一切都在运行时进行，作为程序员，我们没有直接面临这些复杂的细节，而是有一个简洁的 API 来处理并发。
3. Go 协程使用信道（Channel）来进行通信。信道用于防止多个协程访问共享内存时发生竞态条件（Race Condition）。
信道可以看作是 Go 协程之间通信的管道。
*/

func hello() {
	fmt.Println("Hello world goroutine")
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	// 如何启动一个 Go 协程？
	// 调用函数或者方法时，在前面加上关键字 go，可以让一个新的 Go 协程并发地运行。
	// go hello() 启动了一个新的 Go 协程。现在 hello() 函数与 main() 函数会并发地执行。
	// 主函数会运行在一个特有的 Go 协程上，它称为 Go 主协程（Main Goroutine）。
	go hello()
	fmt.Println("main function")

	/*运行一下程序，你会很惊讶！
	该程序只会输出文本 main function。我们启动的 Go 协程究竟出现了什么问题？要理解这一切，我们需要理解两个 Go 协程的主要性质。
	1. 启动一个新的协程时，协程的调用会立即返回。与函数不同，程序控制不会去等待 Go 协程执行完毕。
	在调用 Go 协程之后，程序控制会立即返回到代码的下一行，忽略该协程的任何返回值。
	2. 如果希望运行其他 Go 协程，Go 主协程必须继续运行着。如果 Go 主协程终止，则程序终止，于是其他 Go 协程也不会继续运行。
	*/
	// 在 Go 主协程中使用休眠，以便等待其他协程执行完毕，这种方法只是用于理解 Go 协程如何工作的技巧。
	// 信道可用于在其他协程结束执行之前，阻塞 Go 主协程。
	time.Sleep(1 * time.Second)

	// 启动多个 Go 协程
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}
