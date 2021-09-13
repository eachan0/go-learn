package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gobuffalo/packr"
	"io/ioutil"
	"log"
	"os"
)

/**
将整个文件读取到内存
	使用绝对文件路径
	使用命令行标记来传递文件路径
	将文件绑定在二进制文件中

分块读取文件

逐行读取文件
*/
func main() {
	// 将整个文件读取到内存是最基本的文件操作之一。这需要使用 ioutil 包中的 ReadFile 函数。
	/**有三种方法可以解决路径问题。
	使用绝对文件路径,文件必须放在程序指定的路径中，否则就会出错。
	使用命令行标记来传递文件路径
	将文件绑定在二进制文件中
	*/

	// 使用命令行标记来传递文件路径
	// 使用 flag 包，我们可以从输入的命令行获取到文件路径，接着读取文件内容。
	// 首先我们来看看 flag 包是如何工作的。flag 包有一个名为 String 的函数。该函数接收三个参数。
	// 第一个参数是标记名，第二个是默认值，第三个是标记的简短描述。
	fptr := flag.String("path", "test.txt", "file path to read from")
	// 在程序访问 flag 之前，必须先调用 flag.Parse()。
	flag.Parse()
	fmt.Println("value of fpath is", *fptr)

	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
	} else {
		fmt.Println("Contents of file:", string(data))
	}

	/** 3. 将文件绑定在二进制文件中
	虽然从命令行获取文件路径的方法很好，但还有一种更好的解决方法。如果我们能够将文本文件捆绑在二进制文件，岂不是很棒？这就是我们下面要做的事情。
	有很多包可以帮助我们实现。我们会使用 packr，因为它很简单，并且我在项目中使用它时，没有出现任何问题。
	第一步就是安装 packr 包。
	在命令提示符中输入下面命令，安装 packr 包。
	go get -u github.com/gobuffalo/packr
	packr 会把静态文件（例如 .txt 文件）转换为 .go 文件，接下来，.go 文件会直接嵌入到二进制文件中。
	packer 非常智能，在开发过程中，可以从磁盘而非二进制文件中获取静态文件。在开发过程中，当仅仅静态文件变化时，可以不必重新编译。
	*/
	box := packr.NewBox("static")
	data1, _ := box.FindString("test1.txt")
	fmt.Println("Contents of file:", data1)

	// 分块读取文件
	// 在前面的章节，我们学习了如何把整个文件读取到内存。当文件非常大时，尤其在 RAM 存储量不足的情况下，把整个文件都读入内存是没有意义的。
	// 更好的方法是分块读取文件。这可以使用 bufio 包来完成。
	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	} else {
		r := bufio.NewReader(f)
		b := make([]byte, 3)
		for {
			_, err := r.Read(b)
			if err != nil {
				fmt.Println("Error reading file:", err)
				break
			}
			fmt.Println(string(b))
		}
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	// 逐行读取文件
	// 本节我们讨论如何使用 Go 逐行读取文件。这可以使用 bufio 来实现。
	//逐行读取文件涉及到以下步骤。
	//1.打开文件；
	//2.在文件上新建一个 scanner；
	//3.扫描文件并且逐行读取。
	f, err = os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
