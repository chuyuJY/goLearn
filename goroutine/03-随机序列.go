package main

import "fmt"

// select 会随机选择一个 channel， 以此特性实现随机序列生成器

func main() {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()

	for v := range ch {
		fmt.Print(v)
	}
}
