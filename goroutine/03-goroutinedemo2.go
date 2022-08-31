package main

import (
	"fmt"
	"sync"
	"time"
)

// 需求：要计算1-20的各个数的阶乘，并且把各个数的阶乘放入到map中，最后显示出来。要求使用goroutine。

/*
	思路：
	1. 编写一个函数，计算阶乘，并放入map
	2. 启动多个协程，统计结果放入map（全局共用）
	3. 所有的协程都对同一个map进行操作，会造成资源竞争，因此需要上锁
*/

var (
	myMap = map[int]int{}
	// 3. 全局互斥锁
	lock sync.Mutex
)

func calculate(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	// 加锁
	lock.Lock()
	myMap[n] = res
	// 解锁
	lock.Unlock()
}

func main() {
	// 起了20个协程
	for i := 1; i <= 20; i++ {
		go calculate(i)
	}
	// 此处不能马上遍历，因为不一定执行完毕
	time.Sleep(time.Second)
	for num, res := range myMap {
		fmt.Println(num, res)
	}
}
