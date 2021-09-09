package main

import (
	"fmt"
	"unicode/utf8"
)

/*Go 语言中的字符串是一个字节切片。把内容放在双引号""之间。
Go 中的字符串是兼容 Unicode 编码的，并且使用 UTF-8 进行编码。
*/
func main() {
	name := "Hello World"
	printBytes(name)
	printChars(name)
	name = "Señor"
	printBytes(name)
	/**
	我们尝试输出 Señor 的字符，但却输出了错误的 S e Ã ± o r。 为什么程序分割 Hello World 时表现完美，但分割 Señor 就出现了错误呢？
	这是因为 ñ 的 Unicode 代码点（Code Point）是 U+00F1。它的 UTF-8 编码占用了 c3 和 b1 两个字节。
	而我们打印字符时，却假定每个字符的编码只会占用一个字节，这是错误的。
	在 UTF-8 编码中，一个代码点可能会占用超过一个字节的空间。那么我们该怎么办呢？rune 能帮我们解决这个难题。
	*/
	printCharsBak(name)

	/**
	rune 是 Go 语言的内建类型，它也是 int32 的别称。在 Go 语言中，rune 表示一个代码点。
	代码点无论占用多少个字节，都可以用一个 rune 来表示
	*/
	printChars(name)

	// 上面的程序是一种遍历字符串的好方法，但是 Go 给我们提供了一种更简单的方法来做到这一点：使用 for range 循环。
	printCharsAndBytes(name)

	// 用字节切片构造字符串
	// byteSlice 包含字符串 Café 用 UTF-8 编码后的 16 进制字节。程序输出结果是 Café。
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)
	// 如果把 16 进制换成对应的 10 进制值会怎么样呢？
	byteSlice = []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str = string(byteSlice)
	fmt.Println(str) // 输出结果也是Café

	// 用 rune 切片构造字符串
	// runeSlice 包含字符串 Señor的 16 进制的 Unicode 代码点。这个程序将会输出Señor。
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str = string(runeSlice)
	fmt.Println(str)

	/*字符串的长度
	utf8 package 包中的 func RuneCountInString(s string) (n int) 方法用来获取字符串的长度。
	这个方法传入一个字符串参数然后返回字符串中的 rune 的数量。
	*/
	length(str)

	/**
	Go 中的字符串是不可变的。一旦一个字符串被创建，那么它将无法被修改。
	为了修改字符串，可以把字符串转化为一个 rune 切片。然后这个切片可以进行任何想要的改变，然后再转化为一个字符串。
	*/
	fmt.Println(mutate(str))
}

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		// %x 格式限定符用于指定 16 进制编码。
		fmt.Printf("%x ", s[i])
	}
	fmt.Printf("\n")
}

func printChars(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		// %c 格式限定符用于打印字符串的字符。
		fmt.Printf("%c ", runes[i])
	}
	fmt.Printf("\n")
}

func printCharsBak(s string) {
	for i := 0; i < len(s); i++ {
		// %c 格式限定符用于打印字符串的字符。
		fmt.Printf("%c ", s[i])
	}
	fmt.Printf("\n")
}

func printCharsAndBytes(s string) {
	for index, runes := range s {
		fmt.Printf("%c starts at byte %d\n", runes, index)
	}
}

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func mutate(s string) string {
	// 由于字符串是不可变的，因此这个操作是非法的。所以程序抛出了一个错误 cannot assign to s[0]。
	// s[0] = 'a' // any valid unicode character within single quote is a rune
	runes := []rune(s)
	runes[0] = 'a'
	return string(runes)
}
