package main

import (
	"fmt"
	"sync"
)

/** 临界区
在学习 Mutex 之前，我们需要理解并发编程中临界区（Critical Section）的概念。
当程序并发地运行时，多个 Go 协程不应该同时访问那些修改共享资源的代码。这些修改共享资源的代码称为临界区。
根据上下文切换的不同情形，x 的最终值不确定。这种不太理想的情况称为竞态条件（Race Condition），其程序的输出是由协程的执行顺序决定的。
果在任意时刻只允许一个 Go 协程访问临界区，那么就可以避免竞态条件。而使用 Mutex 可以达到这个目的。
Mutex
Mutex 用于提供一种加锁机制（Locking Mechanism），可确保在某时刻只有一个协程在临界区运行，以防止出现竞态条件。
Mutex 可以在 sync 包内找到。Mutex 定义了两个方法：Lock 和 Unlock。所有在 Lock 和 Unlock 之间的代码，都只能由一个 Go 协程执行，于是就可以避免竞态条件。
有一个 Go 协程已经持有了锁（Lock），当其他协程试图获得该锁时，这些协程会被阻塞，直到 Mutex 解除锁定为止。
*/
var x = 0
var y = 0

func main() {
	var w sync.WaitGroup
	var m = sync.Mutex{}
	for i := 0; i < 1000; i++ {
		w.Add(1)
		// 传递 Mutex 的地址很重要。如果传递的是 Mutex 的值，而非地址，那么每个协程都会得到 Mutex 的一份拷贝，竞态条件还是会发生。
		go incrementByMutex(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)

	// 使用信道处理竞态条件
	// 们创建了容量为 1 的缓冲信道，该缓冲信道用于保证只有一个协程访问增加 x 的临界区。
	// 具体的实现方法是在 x 增加之前，传入 true 给缓冲信道。由于缓冲信道的容量为 1，所以任何其他协程试图写入该信道时，都会发生阻塞，
	// 直到 x 增加后，信道的值才会被读取（第 10 行）。实际上这就保证了只允许一个协程访问临界区。
	var w1 sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w1.Add(1)
		go incrementByChannel(&w1, ch)
	}
	w1.Wait()
	fmt.Println("final value of y", y)

	/** Mutex vs 信道
	通过使用 Mutex 和信道，我们已经解决了竞态条件的问题。那么我们该选择使用哪一个？答案取决于你想要解决的问题。
	如果你想要解决的问题更适用于 Mutex，那么就用 Mutex。如果需要使用 Mutex，无须犹豫。而如果该问题更适用于信道，那就使用信道。:)
	由于信道是 Go 语言很酷的特性，大多数 Go 新手处理每个并发问题时，使用的都是信道。这是不对的。Go 给了你选择 Mutex 和信道的余地，选择其中之一都可以是正确的。
	总体说来，当 Go 协程需要与其他协程通信时，可以使用信道。而当只允许一个协程访问临界区时，可以使用 Mutex。
	就我们上面解决的问题而言，我更倾向于使用 Mutex，因为该问题并不需要协程间的通信。所以 Mutex 是很自然的选择。
	我的建议是去选择针对问题的工具，而别让问题去将就工具。
	*/
}

func incrementByMutex(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}
func incrementByChannel(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	y = y + 1
	<-ch
	wg.Done()
}
