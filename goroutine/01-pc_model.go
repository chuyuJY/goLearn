package main

import (
	"fmt"
	"time"
)

// 最常见的就是生产者消费者模型
// 生产者
func producer(factor int, out chan<- int) {
	i := 0
	for {
		out <- i * factor
		i++
	}
}

// 消费者
func consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 64)
	go producer(3, ch) // 生成 3 的倍数的序列
	go producer(5, ch) // 生成 5 的倍数的序列
	go consumer(ch)    // 消费生成的队列
	// 运行一定时间后退出
	time.Sleep(3 * time.Second)
}
