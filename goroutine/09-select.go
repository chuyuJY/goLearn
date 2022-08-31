package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
	// 1. 传统方法遍历管道时，如果不关闭管道将造成deadlock
	// 但实际中，很难确定何时关闭管道
	// 2. 使用select解决
	for {
		time.Sleep(time.Second)
		select {
		case v := <-intChan: // 注意，此处如果intChan一直没有关闭，也不会一直阻塞在这里，会继续往下走
			fmt.Println("从intChan中读取数据：", v)
		case v := <-stringChan:
			fmt.Println("从stringChan中读取数据：", v)
		default:
			fmt.Println("哪个管道都读不到数据了...")
			return
		}
	}
}
