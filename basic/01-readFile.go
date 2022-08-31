package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file := "./test.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read file err = ", err)
	}
	fmt.Println(string(content)) // 读取内容为[]byte
	// 没有显式open文件，此处也不必close，因为文件的open和close被封装到了ReadFile中
	// 这种方式很简洁，但文件很大的时候，将会效率很低，因为是一次读取所有内容
}
