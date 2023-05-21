package main

import (
	"fmt"
	"sync"
	"time"
)

// 需求：通知多个goroutine关闭
// 解决：1. 最简单的就是通过建多个channel，分别通知，但是代价太大了；
// 		2. 其实可以通过关闭一个channel，通知所有的goroutine

func worker(cancel chan struct{}, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-cancel:
			fmt.Println(i, "goroutine 已退出...")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println(i, ": hello...")
		}
	}
}

func main() {
	cancel := make(chan struct{})
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(cancel, i+1, &wg)
	}

	time.Sleep(2 * time.Second) // 模拟工作时间
	close(cancel)
	wg.Wait() // 等待最后的退出日志
}
