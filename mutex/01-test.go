package main

import (
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	go func() {
		mu.Lock()
		time.Sleep(10 * time.Second)
		mu.Unlock()
	}()
	time.Sleep(time.Second)
	mu.Unlock()
	select {}
}
