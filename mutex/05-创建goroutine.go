package main

import (
	"fmt"
	"sync"
)

var a string
var mu sync.Mutex
var he sync.Mutex

func f() {
	he.Lock()
	fmt.Println(a)
	mu.Unlock()
}

func hello() {
	// 保证顺序执行
	a = "hello, world"
	he.Unlock()
	go f()
	mu.Lock()
}

func main() {
	mu.Lock()
	he.Lock()
	hello()
}
