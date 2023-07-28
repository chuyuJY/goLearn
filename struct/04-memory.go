package main

import (
	"fmt"
	"unsafe"
)

type Args struct {
	num1 int32
	num2 int16
	num3 int8
}

type Flag struct {
	num1 int16
	num2 int32
	num3 int8
}

func main() {
	fmt.Println(unsafe.Sizeof(Args{}), unsafe.Alignof(Args{})) // 8 byte   4 byte
	fmt.Println(unsafe.Sizeof(Flag{}), unsafe.Alignof(Flag{})) // 12 byte  4 byte
}
