package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	//第一种连接方法
	conn := redisConn()

	defer closeConn(conn)
}

func redisConn() redis.Conn {
	//c, err := redis.Dial("tcp", "127.0.0.1:6379")
	c, err := redis.DialURL("redis://192.168.1.161:6379", redis.DialPassword("ttfs.4090"), redis.DialDatabase(7))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("redis connect inited")
	return c
}

func closeConn(conn redis.Conn) {
	err := conn.Close()
	log.Println("redis connect c")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("redis connect close")
}
