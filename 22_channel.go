package main

import (
	"fmt"
	"time"
)

/**
什么是信道？
信道可以想像成 Go 协程之间通信的管道。如同管道中的水会从一端流到另一端，通过使用信道，数据也可以从一端发送，在另一端接收。

信道的声明
所有信道都关联了一个类型。信道只能运输这种类型的数据，而运输其他类型的数据都是非法的。
chan T 表示 T 类型的信道。
信道的零值为 nil。信道的零值没有什么用，应该像对 map 和切片所做的那样，用 make 来定义信道。
*/
func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	}
	//  简短声明通常也是一种定义信道的简洁有效的方法。
	// a := make(chan int)

	// 通过信道进行发送和接收
	// data := <-a // 读取信道 a
	// a <- data   // 写入信道 a

	/**
	信道旁的箭头方向指定了是发送数据还是接收数据。
	箭头对于 a 来说是向外指的，因此我们读取了信道 a 的值，并把该值存储到变量 data。
	在第二行，箭头指向了 a，因此我们在把数据写入信道 a。
	发送与接收默认是阻塞的
	发送与接收默认是阻塞的。这是什么意思？当把数据发送到信道时，程序控制会在发送数据的语句处发生阻塞，直到有其它 Go 协程从信道读取到数据，才会解除阻塞。与此类似，当读取信道的数据时，如果没有其它的协程把数据写入到这个信道，那么读取过程就会一直阻塞着。
	信道的这种特性能够帮助 Go 协程之间进行高效的通信，不需要用到其他编程语言常见的显式锁或条件变量。
	*/
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go helloChannel(done)
	<-done
	fmt.Println("Main received data")

	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)

	/**
	使用信道需要考虑的一个重点是死锁。当 Go 协程给一个信道发送数据时，照理说会有其他 Go 协程来接收数据。
	如果没有的话，程序就会在运行时触发 panic，形成死锁。
	同理，当有 Go 协程等着从一个信道接收数据时，我们期望其他的 Go 协程会向该信道写入数据，要不然程序就会触发 panic。
	*/
	// chd := make(chan int)
	// all goroutines are asleep - deadlock!
	// chd <- 5

	// 单向信道
	// 创建唯送（Send Only）信道 sendch。chan<- int 定义了唯送信道，因为箭头指向了 chan。
	sendch := make(chan<- int)
	go sendData(sendch)
	// fmt.Println(<-sendch) // invalid operation: <-sendch (receive from send-only type chan<- int)

	// 一个不能读取数据的唯送信道究竟有什么意义呢
	// 这就需要用到信道转换（Channel Conversion）了。把一个双向信道转换成唯送信道或者唯收（Receive Only）信道都是行得通的，但是反过来就不行。
	cha1 := make(chan int)
	go sendData(cha1)
	fmt.Println(<-cha1)

	/** 关闭信道和使用 for range 遍历信道
	数据发送方可以关闭信道，通知接收方这个信道不再有数据发送过来。
	当从信道接收数据时，接收方可以多用一个变量来检查信道是否已经关闭。
	v, ok := <- ch
	上面的语句里，如果成功接收信道所发送的数据，那么 ok 等于 true。而如果 ok 等于 false，说明我们试图读取一个关闭的通道。
	从关闭的信道读取到的值会是该信道类型的零值。例如，当信道是一个 int 类型的信道时，那么从关闭的信道读取的值将会是 0。
	*/
	ch := make(chan int)
	go producer(ch)
	/*for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}*/
	//for range 循环用于在一个信道关闭之前，从信道接收数据。
	for v := range ch {
		fmt.Println("Received ", v)
	}
}

func helloChannel(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(1 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func sendData(sendch chan<- int) {
	sendch <- 10
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}
