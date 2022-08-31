package main

import (
	"fmt"
	"time"
)

/*
	要求：
		1. 开启一个writData协程，向管道intChan写入50个整数；
		2. 开启一个readData协程，从管道intChan读取写入的数据；
		3. 注意：
		   1. writData和readData操作的是同一个管道
		   2. 主线程需要等待writData和readData完成才能退出
*/

// 1. write data
func writData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("writData 写入数据：", i)
		//time.Sleep(time.Second)
	}
	close(intChan)
}

// 2. read data
func readData(intChan chan int, exitChan chan bool) {
	for {
		if v, ok := <-intChan; ok {
			fmt.Println("readData 读到数据：", v)
			time.Sleep(time.Second)
		} else {
			break
		}
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	// 3. 创建channel
	intChan := make(chan int, 20)
	exitChan := make(chan bool, 1)
	go writData(intChan)
	go readData(intChan, exitChan)

	for {
		if done := <-exitChan; done {
			break
		}
	}
}
