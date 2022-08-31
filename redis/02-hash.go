package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// 操作hash

func operateHash() {
	// 1. 连接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err =", err)
		return
	}
	defer conn.Close()
	//fmt.Println("redis conn is", conn)
	// 2. 通过go 向redis写入数据 hset [key-field-value]
	/*	// 单个添加
		_, err = conn.Do("hset", "user1", "name", "tomjerry")
		if err != nil {
			fmt.Println("hset err =", err)
			return
		}
		_, err = conn.Do("hset", "user1", "age", "20")
		if err != nil {
			fmt.Println("hset err =", err)
			return
		}*/
	// 批量添加
	_, err = conn.Do("hmset", "user2", "name", "john", "age", "18")
	if err != nil {
		fmt.Println("hmset err =", err)
		return
	}
	// 3. 读取数据 string [key-value]
	/*	// 单个读取
		r1, err := redis.String(conn.Do("hget", "user1", "name")) // 转换成string
		if err != nil {
			fmt.Println("hget err =", err)
			return
		}
		r2, err := redis.Int(conn.Do("hget", "user1", "age")) // 转换成string
		if err != nil {
			fmt.Println("hget err =", err)
			return
		}
		fmt.Printf("操作ok，返回数据: %v: %v\n", r1, r2)*/

	// 批量读取
	r, err := redis.Strings(conn.Do("hmget", "user2", "name", "age")) // 转换成string
	if err != nil {
		fmt.Println("hmget err =", err)
		return
	}
	for i, v := range r {
		fmt.Printf("%v: %v\n", i, v)
	}

}
func main() {
	operateHash()
}
