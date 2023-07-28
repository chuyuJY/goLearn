package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func watchDog(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "已收到停止指令, 马上停止")
			return
		default:
			fmt.Println(name, "正在监控...")
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		watchDog(ctx, "dahuang")
	}()
	go func() {
		defer wg.Done()
		watchDog(ctx, "dabai")
	}()
	time.Sleep(5 * time.Second) // 先让监控狗监控5秒
	cancel()                    // 通知多个 goroutine 退出
	wg.Wait()
}
