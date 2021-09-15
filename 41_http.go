package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// 基本的HTTP/HTTPS请求 Get、Head、Post和PostForm函数发出HTTP/HTTPS请求。
	// GET请求示例
	server()
	get()
}

func get() {
	resp, err := http.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	// 程序在使用完response后必须关闭回复的主体。
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read from resp.Body failed,err:", err)
		return
	}
	fmt.Print(string(body))
}

func server() {
	var addr = ":8080"
	mux := &http.ServeMux{}
	//根路径
	mux.HandleFunc("/", myHandler)
	//邮件夹
	mux.HandleFunc("/list", myHandler)
	//登陆界面
	mux.HandleFunc("/login", myHandler)
	s := &http.Server{
		Addr:           addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("start on %s\n", addr)
}
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there!\n")
}
