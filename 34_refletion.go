package main

import (
	"fmt"
	"reflect"
)

/** 什么是反射？
反射就是程序能够在运行时检查变量和值，求出它们的类型。
在 Go 语言中，reflect 实现了运行时反射。reflect 包会帮助识别 interface{} 变量的底层具体类型和具体值。
reflect.Type 表示 interface{} 的具体类型，而 reflect.Value 表示它的具体值。
reflect.TypeOf() 和 reflect.ValueOf() 两个函数可以分别返回 reflect.Type 和 reflect.Value。这两种类型是我们创建查询生成器的基础。
Type 表示 interface{} 的实际类型（在这里是 main.Order)，而 Kind 表示该类型的特定类别（在这里是 struct）。
NumField() 和 Field() 方法
NumField() 方法返回结构体中字段的数量，而 Field(i int) 方法返回字段 i 的 reflect.Value。
*/
func main() {
	// Int 和 String 可以帮助我们分别取出 reflect.Value 作为 int64 和 string。
	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)
	b := "Naveen"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)

	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
	i := 90
	createQuery(i)
}

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) {
	// NumField 方法只能在结构体上使用
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
		/**
		反射是 Go 语言中非常强大和高级的概念，我们应该小心谨慎地使用它。使用反射编写清晰和可维护的代码是十分困难的。
		你应该尽可能避免使用它，只在必须用到它时，才使用反射。
		*/
	}
	fmt.Println("unsupported type")
}
