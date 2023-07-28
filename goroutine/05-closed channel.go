package main

import "fmt"

/*
	channel的close和for-range
*/

func testClose() {
	intChan := make(chan int, 3)
	intChan <- 2
	// 1. 关闭一个channel
	close(intChan)
	// 2. 向关闭的channel中写入数据
	//intChan <- 3 // 会报错 panic: send on closed channel
	// 3. 从关闭的channel读取数据
	num, ok := <-intChan
	fmt.Println(num, ok) // 2 true
	num, ok = <-intChan
	fmt.Println(num, ok) // 0 false
}

func testFR() {
	intChan := make(chan int, 100)
	for i := 1; i <= 100; i++ {
		intChan <- i
	}

	// 如果遍历管道，但不关闭管道，将会报错。因为还会傻乎乎的等着管道写数据，造成死锁
	close(intChan)
	// 1. 遍历管道
	for v := range intChan {
		fmt.Println("v = ", v)
	}
}

func main() {
	testClose()
	//testFR()
}
