package main

import "fmt"

func Print[T any](v T) {
	fmt.Printf("%T:%[1]v\n", v)
}

type MyInt int

func main() {
	Print("1")
	Print(MyInt(2))
}
