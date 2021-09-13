package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

/** 什么是错误？
错误表示程序中出现了异常情况。在 Go 中，错误一直是很常见的。错误用内建的 error 类型来表示。
*/
func main() {
	f, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err) // open /test.txt: The system cannot find the file specified.

	} else {
		fmt.Println(f.Name(), "opened successfully")
	}

	// 1. 断言底层结构体类型，使用结构体字段获取更多信息
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
	} else {
		fmt.Println(f.Name(), "opened successfully")
	}

	// 2. 断言底层结构体类型，调用方法获取更多信息
	addr, err := net.LookupHost("golangbot123.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
	} else {
		fmt.Println(addr)
	}

	// 3. 与 error 类型的变量直接比较。
	files, err := filepath.Glob("[")
	if err != nil && err == filepath.ErrBadPattern {
		fmt.Println(err)
		//return
	}
	// 绝不要忽略错误。忽视错误会带来问题。输出看起来就像是没有任何匹配了 glob 模式的文件，但实际上这是因为模式的写法不对。
	fmt.Println("matched files", files)
}
