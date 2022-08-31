package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 1; i <= 1000000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool, i int) {
	for num := range intChan {
		isPrime := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primeChan <- num
		}
	}
	// 所有的数字均读取完毕
	//time.Sleep(time.Millisecond * 10)
	fmt.Println("goruotine", i, "已完成...")
	exitChan <- true
}

func normal() {
	for num := 1; num <= 1000000; num++ {
		isPrime := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println("prime:", num)
		}
	}
}

func main() {
	intChan := make(chan int, 20000)
	primeChan := make(chan int, 20000)
	exitChan := make(chan bool, 6)
	// 记录起始时间
	start := time.Now().Unix()
	end := time.Now().Unix()
	// 1. 开启写入数字的goroutine
	go putNum(intChan)
	// 2. 开启四个goroutine计算素数
	for i := 1; i <= 6; i++ {
		go primeNum(intChan, primeChan, exitChan, i)
	}
	// 3. 开启一个goroutine监测上述goroutine是否已全部完成
	// 当取出4个true，证明此时所有goroutine已完成
	go func() {
		for i := 0; i < 6; i++ {
			<-exitChan
		}
		end = time.Now().Unix()
		close(primeChan)
		close(exitChan)
	}()

	// 4. 主线程打印
	//for res := range primeChan {
	//	fmt.Println("prime:", res)
	//}
	// 测试时间专用，避免打印的时间
	for {
		if _, ok := <-primeChan; ok {
		} else {
			break
		}
	}
	fmt.Println("goroutine方式 花费时间：", end-start)

	// normal 方式花费时间
	//start = time.Now().Unix()
	//normal()
	//end = time.Now().Unix()
	//fmt.Println("normal方式 花费时间：", end-start)
}
