package main

import "fmt"

// go build -gcflags '-m -l' 01-type.go
// go build -gcflags '-m -m -l' 01-type.go	// 更详细

// 变量类型不确定
func main() {
	a := 123
	fmt.Println(a)
}
