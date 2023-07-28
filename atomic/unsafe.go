package main

import (
	"fmt"
	"unsafe"
)

func main() {
	bytes := []byte("hello")
	fmt.Println(bytes)          // 输出 [104 101 108 108 111]
	p := unsafe.Pointer(&bytes) // 强制转换成unsafe.Pointer，编译器不会报错
	str := *(*string)(p)        // 然后强制转换成string类型的指针，再将这个指针的值当做string类型取出来
	fmt.Println(str)            // 输出 "hello"
}
