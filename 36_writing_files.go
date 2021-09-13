package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

// (word1[\s\S]*word2) | (word2[\s\S]*word1)
func main() {
	// 创建文件
	f, err := os.Create("test_write.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 将字符串写入文件
	// 最常见的写文件就是将字符串写入文件。这个写起来非常的简单。这个包含以下几个阶段。
	// 1.创建文件
	// 2.将字符串写入文件
	_, err = f.WriteString("Hello World\n")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	// 将字节写入文件和写入字符串非常的类似。我们将使用 Write 方法将字节写入到文件。
	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 10}
	_, err = f.Write(d2)
	// 将字符串一行一行的写入文件
	d := []string{"Welcome to the world of Go1.", "Go is a compiled language.",
		"It is easy to learn Go."}
	for _, v := range d {
		_, err = fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 追加到文件
	// 这个文件将以追加和写的方式打开。这些标志将通过 Open 方法实现。当文件以追加的方式打开，我们添加新的行到文件里。
	f, err = os.OpenFile("test_write.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")

	// 并发写文件
	// 当多个 goroutines 同时（并发）写文件时，我们会遇到竞争条件(race condition)。因此，当发生同步写的时候需要一个 channel 作为一致写入的条件。
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done)
	go func() {
		wg.Wait()
		close(data)
	}()
	df := <-done
	if df == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}
}

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consume(data chan int, done chan bool) {
	f, err := os.Create("concurrent")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}
