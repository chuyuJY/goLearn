package main

import "fmt"

/*
	3. struct{} 用作仅包含方法的结构体
	在部分场景下，结构体只包含方法，不包含任何的字段。
	无论是 int 还是 bool 都会浪费额外的内存，因此呢，这种情况下，声明为空结构体是最合适的。
*/

type Door struct{}

func (d Door) Open() {
	fmt.Println("Open the door")
}

func (d Door) Close() {
	fmt.Println("Close the door")
}
