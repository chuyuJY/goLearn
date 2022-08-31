package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 1. 返回机器的cpu个数
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)
	// 2. 设置CPU数目
	runtime.GOMAXPROCS(cpuNum)
}
