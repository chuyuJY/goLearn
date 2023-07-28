package main

import "sync"

// 1. 利用 sync.Once 只关闭一次 channel
// 比直接关闭显得礼貌了一些，但并不优雅，因为并不能完全有效地避免数据竞争

type MyChannel struct {
	C    chan int
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan int)}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}
