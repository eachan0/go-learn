package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
	"time"
)

var pool *redis.Pool

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redisConn(), nil },
	}
}

func closePool() {
	err := pool.Close()
	if err != nil {
		return
	}
	log.Println("redis connect pool close")
}

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	pool = newPool()
}

func main() {
	key := "go:test:key"
	// 直连方式
	// conn := redisConn()
	// defer closeConn(conn)
	// 使用连接池
	conn := pool.Get()
	defer closePool()
	_, _ = conn.Do("SET", key, 3)
	reply, _ := redis.String(conn.Do("get", key))
	fmt.Println(key + " =" + reply)
	_, _ = conn.Do("SET", "anotherkey", "will expire in a minute", "EX", 600)
	_, _ = conn.Do("EXPIRE", key, 600)
	intReply, err := redis.Int(conn.Do("TTL", key))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(key + " TTL:" + strconv.Itoa(intReply))
	}
	intReply, err = redis.Int(conn.Do("TTL", "anotherkey"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("anotherkey TTL:" + strconv.Itoa(intReply))
	}
	reply, _ = redis.String(conn.Do("GET", "anotherkey"))
	fmt.Println("anotherkey=" + reply)
	defer conn.Close()
}

func redisConn() redis.Conn {
	//c, err := redis.Dial("tcp", "127.0.0.1:6379")
	c, err := redis.DialURL("redis://127.0.0.1:6379", redis.DialPassword("123456"), redis.DialDatabase(7))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("redis connect inited")
	return c
}

func closeConn(conn redis.Conn) {
	err := conn.Close()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("redis connect close")
}
