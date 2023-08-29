package main

import (
	"fmt"
	"time"
)

var (
	ch1 = make(chan byte)
	ch2 = make(chan int)
)

func letter() {
	for i := 0; i < 10; i++ {
		v := <-ch1
		fmt.Println(string(v))
		ch2 <- int(v - 'a' + 1)
	}

}

func number() {
	for i := 0; i < 10; i++ {
		v := <-ch2
		fmt.Println(v)
		if v == 10 {
			return
		}
		ch1 <- byte(v + 'a')
	}
}

func main() {
	go letter()
	go number()
	ch1 <- 'a'
	time.Sleep(time.Second * 2)
}
