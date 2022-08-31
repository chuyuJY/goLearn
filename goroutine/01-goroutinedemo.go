package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
	1. 在主线程开启一个goroutine，每一秒输出一个"hello, world"
	2. 在主线程也每一秒输出一个"hello, world"
	3. 要求：主线程和goroutine同时执行
*/

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() hello, world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	// 1. 在主线程开启一个goroutine
	go test()
	// 2. 主线程的输出
	for i := 0; i < 10; i++ {
		fmt.Println("main() hello, world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}
