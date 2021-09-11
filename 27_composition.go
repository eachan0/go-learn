package main

import "fmt"

/** 组合取代继承
Go 不支持继承，但它支持组合（Composition）。组合一般定义为“合并在一起”。汽车就是一个关于组合的例子：一辆汽车由车轮、引擎和其他各种部件组合在一起。
*/
func main() {
	author1 := author27{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}
	post1 := post27{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	post2 := post27{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	post3 := post27{
		"Concurrency",
		"Go is a concurrent language and not a parallel one",
		author1,
	}
	w := website{
		posts: []post27{post1, post2, post3},
	}
	w.contents()
}

type author27 struct {
	firstName string
	lastName  string
	bio       string
}

func (a author27) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post27 struct {
	title   string
	content string
	author27
}

func (p post27) details() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	fmt.Println("Author: ", p.author27.fullName())
	fmt.Println("Bio: ", p.author27.bio)
}

type website struct {
	// 结构体不能嵌套一个匿名切片。我们需要一个字段名。
	posts []post27
}

func (w website) contents() {
	fmt.Println("Contents of Website")
	for _, v := range w.posts {
		v.details()
		fmt.Println()
	}
}
