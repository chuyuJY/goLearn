package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 1. 用 mutex 实现计数器
func mutexAdd() {
	var a int32 = 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	start := time.Now()
	for i := 0; i < 100000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			a += 1
			mu.Unlock()
		}()
	}
	wg.Wait()
	timeSpends := time.Now().Sub(start).Nanoseconds()
	fmt.Printf("use mutex a is %d, spend time: %v\n", a, timeSpends)
}

// 2. 用 atomic 实现计数器
func AtomicAdd() {
	var a int32 = 0
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&a, 1)
		}()
	}
	wg.Wait()
	timeSpends := time.Now().Sub(start).Nanoseconds()
	fmt.Printf("use atomic a is %d, spend time: %v\n", atomic.LoadInt32(&a), timeSpends)
}
