package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"
)

const host = "http://design.hmcec.com/"

var path string
var work = false
var author = false

func main() {
	// 定义几个变量，用于接收命令行的参数值
	var workStr string
	var authorStr string
	// &user 就是接收命令行中输入 -u 后面的参数值，其他同理
	flag.StringVar(&path, "d", "D:\\go_images\\", "保存路径,默认为")
	flag.StringVar(&workStr, "w", "true", "作品图片,默认获取")
	flag.StringVar(&authorStr, "a", "true", "作者图片，默认获取")
	// 解析命令行参数写入注册的flag里
	flag.Parse()
	work, _ = strconv.ParseBool(workStr)
	author, _ = strconv.ParseBool(authorStr)
	// 初始化文件夹
	initPath(path)
	// 获取列表页
	hrefs := home()
	done := make(chan string)
	for _, href := range hrefs {
		go item(href, done)
	}
	count := 0
	for v := range done {
		count++
		fmt.Println("[" + v + "]照片获取完")
		if count >= 30 {
			close(done)
			break
		}
	}
}

func initPath(path string) {
	path += "\\"
	s, err := os.Stat(path)
	if err != nil {
		err = os.Mkdir(path, os.ModePerm)
	} else {
		if !s.IsDir() {
			fmt.Println("指定路径不是一个目录")
			syscall.Exit(0)
		}
	}
}

func home() []string {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "http://design.hmcec.com/ranking.asp", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36 Edg/93.0.961.38")
	request.Header.Add("Host", "design.hmcec.com")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Add("Referer", "http://design.hmcec.com/")
	response, _ := client.Do(request)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	dom, _ := goquery.NewDocumentFromReader(strings.NewReader(string(body)))

	hrefs := make([]string, 30, 30)
	dom.Find("tbody > tr > td > a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			hrefs[i] = host + href
		}
	})
	return hrefs
}

func item(url string, done chan string) {
	client := &http.Client{}
	response, _ := client.Get(url)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	dom, _ := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	var name string
	dom.Find("div.panel-body").Each(func(i int, s *goquery.Selection) {
		index := s.Find("p:nth-child(1) > span").Text()
		workImage, exist1 := s.Find("p:nth-child(2) > img").Attr("src")
		authorImage, exist2 := s.Find("p:nth-child(4) > img").Attr("src")
		nameNode := s.Find("p:nth-child(6)").Text()
		addressNode := s.Find("p:nth-child(7)").Text()
		if exist1 && exist2 && index != "" && nameNode != "" && addressNode != "" {
			name = removeSpecialCharacters(nameNode)
			address := removeSpecialCharacters(addressNode)
			if work {
				getImage(host+workImage, index+"_"+name+"_"+address+"_work.jpg")
			}
			if author {
				getImage(host+authorImage, index+"_"+name+"_"+address+"_author.jpg")
			}
		}
	})
	done <- name
}

func getImage(url, filename string) {
	res, _ := http.Get(url)
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReader(res.Body)
	file, err := os.Create(path + filename)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)
	_, err = io.Copy(writer, reader)
	if err != nil {
		panic(err)
	}
}

func removeSpecialCharacters(old string) string {
	temp := strings.ReplaceAll(old, "作者：", "")
	temp2 := strings.ReplaceAll(temp, "\t", "")
	temp3 := strings.ReplaceAll(temp2, "\n", "")
	return strings.ReplaceAll(temp3, " ", "")
}
