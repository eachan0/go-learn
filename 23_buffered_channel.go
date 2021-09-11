package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/** 什么是缓冲信道？
在上一教程里，我们讨论的主要是无缓冲信道。我们在信道的教程里详细讨论了，无缓冲信道的发送和接收过程是阻塞的。
我们还可以创建一个有缓冲（Buffer）的信道。只在缓冲已满的情况，才会阻塞向缓冲信道（Buffered Channel）发送数据。
同样，只有在缓冲为空的时候，才会阻塞从缓冲信道接收数据。
通过向 make 函数再传递一个表示容量的参数（指定缓冲的大小），可以创建缓冲信道。
ch := make(chan type, capacity)
要让一个信道有缓冲，上面语法中的 capacity 应该大于 0。无缓冲信道的容量默认为 0，因此我们在上一教程创建信道时，省略了容量参数。

缓冲信道的重要应用之一就是实现工作池。
一般而言，工作池就是一组等待任务分配的线程。一旦完成了所分配的任务，这些线程可继续等待任务的分配。
*/
func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func main() {
	ch0 := make(chan string, 2)
	ch0 <- "naveen"
	ch0 <- "paul"
	fmt.Println(<-ch0)
	fmt.Println(<-ch0)

	/*ch := make(chan int, 2)
	go write(ch)
	time.Sleep(time.Second * 2)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}*/

	/** 死锁
	我们向容量为 2 的缓冲信道写入 3 个字符串。当在程序控制到达第 3 次写入时（第 11 行），由于它超出了信道的容量，因此这次写入发生了阻塞。
	现在想要这次写操作能够进行下去，必须要有其它协程来读取这个信道的数据。但在本例中，并没有并发协程来读取这个信道，因此这里会发生死锁（deadlock）
	*/
	ch0 <- "naveen"
	ch0 <- "paul"
	// ch0 <- "steve"
	fmt.Println(<-ch0)
	fmt.Println(<-ch0)

	/** 长度 vs 容量
	缓冲信道的容量是指信道可以存储的值的数量。我们在使用 make 函数创建缓冲信道的时候会指定容量大小。
	缓冲信道的长度是指信道中当前排队的元素个数。
	*/
	ch0 <- "naveen"
	ch0 <- "paul"
	fmt.Println("capacity is", cap(ch0))
	fmt.Println("length is", len(ch0))
	fmt.Println("read value", <-ch0)
	fmt.Println("new length is", len(ch0))

	// 工作池
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

type Job struct {
	id, randomNo int
}
type Result struct {
	job Job
	sum int
}

func digits23(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

// 工作协程（Worker Goroutine）会监听缓冲信道 jobs 里更新的作业。一旦工作协程完成了作业，其结果会写入缓冲信道 results。
// worker 函数接收了一个 WaitGroup 类型的 wg 作为参数，当所有的 jobs 完成的时候，调用了 Done() 方法。
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits23(job.randomNo)}
		results <- output
	}
	wg.Done()
}

/** 在创建 Go 协程之前，它调用了 wg.Add(1) 方法，于是 WaitGroup 计数器递增。
接下来，我们创建工作协程，并向 worker 函数传递 wg 的地址。创建了需要的工作协程后，函数调用 wg.Wait()，等待所有的 Go 协程执行完毕。
所有协程完成执行之后，函数会关闭 results 信道。因为所有协程都已经执行完毕，于是不再需要向 results 信道写入数据了。
*/
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomNo := rand.Intn(999)
		job := Job{i, randomNo}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomNo, result.sum)
	}
	done <- true
}
