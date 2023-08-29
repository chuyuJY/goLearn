package main

import (
	"context"
	"time"
)

func f1(ch chan struct{}) {
	time.Sleep(time.Second * 1)
	ch <- struct{}{}
}

func f2(ch chan struct{}) {
	time.Sleep(time.Second * 3)
	ch <- struct{}{}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	ch := make(chan struct{}, 1)
	go func() {
		go f1(ch)
		select {
		case <-ctx.Done():
			println("f1 timeout")
		case <-ch:
			println("f1 done")
		}
	}()
	go func() {
		go f2(ch)
		select {
		case <-ctx.Done():
			println("f2 timeout")
		case <-ch:
			println("f2 done")
		}
	}()
	time.Sleep(time.Second * 5)
}
