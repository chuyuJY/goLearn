package main

import (
	"fmt"
	"sync"
)

/*
	摘要：在线程之间进行数据同步：顺序一致性的内存模型
		1. 同步原语
		2. 互斥锁
*/

func main() {
	/*
		done := make(chan int)

		go func() {
			fmt.Println("hello, world")
			done <- 1
		}()
		<-done
	*/
	var mu sync.Mutex
	mu.Lock()
	go func() {
		fmt.Println("hello, world")
		mu.Unlock()
	}()
	mu.Lock() // 主程序阻塞在此处
}
