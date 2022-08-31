package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello, world", i)
	}
}

func testError() {
	// 使用defer + recover 捕获异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("testError 发生异常：", err)
		}
	}()

	// 异常报错
	var myMap map[int]int
	myMap[0] = 0
}

func main() {
	go sayHello()
	go testError()

	// 主线程
	for i := 0; i < 10; i++ {
		fmt.Println("main() hello", i)
		time.Sleep(time.Second)
	}
}
