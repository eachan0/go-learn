package main

import (
	"fmt"
	"time"
)

/**什么是 select？
select 语句用于在多个发送/接收信道操作中进行选择。select 语句会一直阻塞，直到发送/接收操作准备就绪。
如果有多个信道操作准备完毕，select 会随机地选取其中之一执行。该语法与 switch 类似，所不同的是，这里的每个 case 语句都是信道操作。
*/
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	// 当 select 由多个 case 准备就绪时，将会随机地选取其中之一去执行。
	// 如果你在 playground 上在线运行的话，它的输出总是一样的，这是由于 playground 不具有随机性所造成的。
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}

	// 在没有 case 准备就绪时，可以执行 select 语句中的默认情况（Default Case）。这通常用于防止 select 语句一直阻塞。
	ch := make(chan string)
	/*go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
		default:
			fmt.Println("no value received")
		}
	}*/

	// 死锁与默认情况
	// 试图读取信道 ch。由于没有 Go 协程向该信道写入数据，因此 select 语句会一直阻塞，导致死锁。
	// 触发运行时 panic: fatal error: all goroutines are asleep - deadlock!
	// 如果存在默认情况，就不会发生死锁，因为在没有其他 case 准备就绪时，会执行默认情况。
	var ch1 chan string
	select {
	case <-ch:
	// 如果 select 只含有值为 nil 的信道，也同样会执行默认情况。
	case v := <-ch1:
		fmt.Println("received value", v)
	default:
		fmt.Println("default case executed")
	}

	// 空 select
	// 我们已经知道，除非有 case 执行，select 语句就会一直阻塞着。在这里，select 语句没有任何 case，因此它会一直阻塞，导致死锁。
	select {}
}

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}
func process1(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}
