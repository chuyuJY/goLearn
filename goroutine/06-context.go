package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Context: 用来简化处理 多个 Goroutine 之间与请求域的数据、超时和退出等操作

func worker(ctx context.Context, i int, wg *sync.WaitGroup) error {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			// 退出
			fmt.Printf("%v goroutine 已退出...\n", i)
			return ctx.Err()
		default:
			fmt.Printf("%v: hello...\n", i)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(ctx, i+1, &wg)
	}
	wg.Wait()
}
