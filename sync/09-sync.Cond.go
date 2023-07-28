package main

import (
	"fmt"
	"sync"
	"time"
)

// 10个人赛跑，1个裁判发号施令
func race() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()
			fmt.Println(id, "号已准备就绪...")
			cond.L.Lock()
			cond.Wait() // 等待发令枪响
			fmt.Println(id, "号开跑!!!")
			cond.L.Unlock()
		}(i)
	}
	//等待所有goroutine都进入wait状态
	time.Sleep(5 * time.Second)

	go func() {
		defer wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家准备跑")
		cond.Broadcast() // 发令枪响
	}()
	//防止函数提前返回退出
	wg.Wait()
}

func main() {
	race()
}
