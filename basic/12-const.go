package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c) // 0, 1, 2
}
