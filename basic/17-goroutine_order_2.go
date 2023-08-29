package main

import (
	"fmt"
)

var (
	catChan  = make(chan int)
	fishChan = make(chan int)
	dogChan  = make(chan int)
	allChan  = make(chan struct{})
)

func cat() {
	v := <-catChan
	fmt.Printf("%v cat\n", v)
	fishChan <- v
}

func fish() {
	v := <-fishChan
	fmt.Printf("%v fish\n", v)
	dogChan <- v
}

func dog() {
	v := <-dogChan
	fmt.Printf("%v dog\n", v)
	if v == 10 {
		allChan <- struct{}{}
		return
	}
	catChan <- v + 1
}

func main() {
	for i := 0; i < 10; i++ {
		go cat()
		go fish()
		go dog()
	}
	catChan <- 1
	<-allChan
}
