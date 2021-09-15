package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	p := &Product{}
	p.Name = "Xiao mi 6"
	p.IsOnSale = true
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 1
	o := &Order{
		Num:     2,
		Total:   2499.00 * 2,
		Product: *p,
	}
	data, _ := json.Marshal(o)

	fmt.Println(string(data))

	var jsonStr = `{"num":2,"total":"4998","product":{"name":"Xiao mi 6","product_id":"1","number":"10000","price":"2499","is_on_sale":"true"}}`
	o0 := &Order{}
	err := json.Unmarshal([]byte(jsonStr), o0)
	fmt.Println(err)
	fmt.Println(*o0)

	// Golang自带的JSON解析功能非常强悍
	/** 实现以下接口
	type Marshaler interface {
	    MarshalJSON() ([]byte, error)
	}
	type Unmarshaler interface {
	    UnmarshalJSON([]byte) error
	}
	*/
	u := &User{"zyc199777@gmail.com"}
	data, err = json.Marshal(u)
	fmt.Println(string(data))
	fmt.Println(err)
	err = json.Unmarshal([]byte("{\"Email\":\"zyc199777@outlook.com\"}"), u)
	fmt.Println(u)
	fmt.Println(err)
}

// Product 商品信息
type Product struct {
	// 何为Tag，tag就是标签，给结构体的每个字段打上一个标签，标签冒号前是类型，后面是标签名
	// omitempty，tag里面加上omitempy，可以在序列化的时候忽略0值或者空值
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,string"` // 表示不进行序列化
	Number    int     `json:"number,string"`
	Price     float64 `json:"price,omitempty,string"`
	IsOnSale  bool    `json:"is_on_sale,string"`
}

// 有些时候，我们在序列化或者反序列化的时候，可能结构体类型和需要的类型不一致，这个时候可以指定,支持string,number和boolean
type Order struct {
	Num     int     `json:"num"`
	Total   float64 `json:"total,string"`
	Product Product `json:"product,string"`
}

type User struct {
	Email string
}

/*func (u *User) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	_, err := buf.WriteString(u.Email)
	return buf.Bytes(), err
}*/

func (u *User) UnmarshalJSON(data []byte) error {
	fmt.Println("UnmarshalJSON")
	// 这里简单演示一下，简单判断即可
	if !bytes.ContainsAny(data, "@") {
		return fmt.Errorf("mail format error")
	}
	u.Email = string(data)
	return nil
}
