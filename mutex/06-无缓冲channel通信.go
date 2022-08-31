package main

import "fmt"

/*
	用channel保持同步。对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成之前。
*/

var done = make(chan int)
var msg string

func aGoroutine() {
	msg = "hello, world"
	close(done) // 关闭的chan 会返回给接收chan 0
}

func main() {
	go aGoroutine()
	fmt.Println(<-done) // 接受阻塞，保证顺序一致性
	fmt.Println(msg)
}
