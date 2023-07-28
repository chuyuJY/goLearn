package main

import (
	"log"
	"time"

	"github.com/Jeffail/tunny"
)

func main() {
	// 第一个参数是协程池的大小(poolSize)，第二个参数是协程运行的函数(worker)
	pool := tunny.NewFunc(3, func(i interface{}) interface{} {
		log.Println(i)
		time.Sleep(time.Second)
		return nil
	})
	defer pool.Close() // 关闭协程池

	for i := 0; i < 10; i++ {
		go pool.Process(i) // 将参数 i 传递给协程池定义好的 worker 处理
	}
	time.Sleep(time.Second * 4)
}
