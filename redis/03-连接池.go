package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// 定义一个全局pool
var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) { // 初始化连接的代码，连接哪个ip，端口
			return redis.Dial("tcp", "localhost:6379")
		},
		TestOnBorrow:    nil,
		MaxIdle:         8,   // 最大空闲连接数
		MaxActive:       0,   // 表示和数据库的最大连接数，0表示没有限制
		IdleTimeout:     100, // 最大空闲时间
		Wait:            false,
		MaxConnLifetime: 0,
	}
}

func operatePool() {
	// 先从一个pool取出一个连接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("set", "name", "tom")
	if err != nil {
		fmt.Println("set err =", err)
		return
	}

	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("get err =", err)
		return
	}
	fmt.Println("get r =", r)

	// 关闭pool
	pool.Close()
	conn = pool.Get()
	// 此时取不出了，因为pool已关闭
}

func main() {
	operatePool()
}
