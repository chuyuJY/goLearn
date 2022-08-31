package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func writeNew() {
	// 1. 打开文件
	filePath := "testOpen.txt"
	file, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	// 2. 写入
	str := "hello, world!\r\n"
	// 带缓存的 *Writter
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因为writter是带缓存的，因此在调用WritterString时，
	// 其实内容是先写到缓存的 ，所以需要调用Flush方法将缓存的数据，
	// 真正写到磁盘
	writer.Flush()
}

func writeAppend() {
	// 1. 打开文件
	filePath := "testOpen.txt"
	file, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	// 2. 写入
	str := "hello, JY!\r\n"
	// 带缓存的 *Writter
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因为writter是带缓存的，因此在调用WritterString时，
	// 其实内容是先写到缓存的 ，所以需要调用Flush方法将缓存的数据，
	// 真正写到磁盘
	writer.Flush()
}

func writeAndRead() {
	// 1. 打开文件
	filePath := "testOpen.txt"
	file, _ := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	defer file.Close()
	// 2. 先读取原来的内容并显示在终端
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // 若读取到文件末尾
			break
		}
		fmt.Print(str) // 打印
	}
	// 3. 写入
	str := "你好, JY!\r\n"
	// 带缓存的 *Writter
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因为writter是带缓存的，因此在调用WritterString时，
	// 其实内容是先写到缓存的 ，所以需要调用Flush方法将缓存的数据，
	// 真正写到磁盘
	writer.Flush()
}

func main() {
	//writeNew()
	//writeAppend()
	writeAndRead()
}
