package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	// 通过go向redis写入数据和读取数据
	// 1. 连接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err =", err)
		return
	}
	defer conn.Close()
	//fmt.Println("redis conn is", conn)
	// 2. 通过go 向redis写入数据 string [key-value]
	_, err = conn.Do("set", "name", "tomjerry")
	if err != nil {
		fmt.Println("set err =", err)
		return
	}
	// 3. 读取数据 string [key-value]
	r, err := redis.String(conn.Do("get", "name")) // 转换成string
	if err != nil {
		fmt.Println("get err =", err)
		return
	}
	fmt.Println("操作ok，返回数据:", r)

}
