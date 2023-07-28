package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 实现：各线程只初始化一次instance

type singleton struct {
}

// 1. 原子操作相比互斥锁要更节省资源，因此通过原子检测标志位状态来降低互斥锁的使用
/*

var (
	instance    *singleton
	initialized uint32
	mu          sync.Mutex
)

func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}
	return instance
}
*/

// 2. 将上述代码通用部分提取，成为atomic.Once
//type Once struct {
//	mu   sync.Mutex
//	done uint32
//}
//
//func (o *Once) Do(f func()) {
//	if atomic.LoadUint32(&o.done) == 1 {
//		return
//	}
//
//	o.mu.Lock()
//	defer o.mu.Unlock()
//
//	if o.done == 0 {
//		defer atomic.StoreUint32(&o.done, 1)
//		f()
//	}
//}
//
//var (
//	instance *singleton
//	once     Once
//)
//
//func Instance() *singleton {
//	once.Do(func() {
//		instance = &singleton{}
//	})
//	return instance
//}

// 3. 使用 sync.Once 来保证代码只执行一次
func doOnce() {
	var once sync.Once
	onecBody := func() {
		fmt.Println("Only once")
	}
	// 等待协程执行完毕
	done := make(chan int)
	for i := 0; i < 10; i++ {
		id := i
		go func() {
			fmt.Println("当前协程id:", id)
			once.Do(onecBody)
			done <- id
		}()
	}

	for i := 0; i < 10; i++ {
		fmt.Println("协程id:", <-done, " 已执行完毕...")
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	doOnce()
}
