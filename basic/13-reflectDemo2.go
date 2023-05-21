package main

import (
	"fmt"
	"reflect"
)

// 通过反射，修改 num int 的值，修改 student 的值
func reflectDemo3(data interface{}) {
	rVal := reflect.ValueOf(data)
	fmt.Println("rVal kind =", rVal.Kind()) // ptr
	// 这样是不行的，因为rVal是指针，需要先取到指针的值
	//rVal.SetInt(20)
	rVal.Elem().SetInt(20) // .Elem() 可以获取ptr指向的值，然后再改
}

func main() {
	// example1
	//num := 10
	//reflectDemo3(&num)
	//fmt.Println("num =", num)

	// example2
	str := "tom"
	rVal := reflect.ValueOf(&str)
	// 会报错
	//rVal.SetString("lily")
	rVal.Elem().SetString("lily")
	fmt.Println(str)
}
