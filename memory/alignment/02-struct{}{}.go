package main

import (
	"fmt"
	"unsafe"
)

/*
	2. 空 struct{} 的对齐
	空 struct{} 大小为 0，作为其他 struct 的字段时，一般不需要内存对齐。
	但是有一种情况除外：即当 struct{} 作为结构体最后一个字段时，需要内存对齐。
	因为如果有指针指向该字段, 返回的地址将在结构体之外，如果此指针一直存活不释放对应的内存，
	就会有内存泄露的问题（该内存不因结构体释放而释放）。
	因此，当 struct{} 作为其他 struct 最后一个字段时，需要填充额外的内存保证安全。
*/

type demo3 struct {
	c int32
	a struct{}
}

type demo4 struct {
	a struct{}
	c int32
}

func main() {
	fmt.Println(unsafe.Sizeof(demo3{})) // 8
	fmt.Println(unsafe.Sizeof(demo4{})) // 4
}
