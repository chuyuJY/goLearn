package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func copyFile(srcFileName, dstFileName string) (int64, error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open file err = ", err)
		return 0, err
	}
	defer srcFile.Close()
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err = ", err)
		return 0, err
	}
	defer dstFile.Close()
	// 拿到reader和writer
	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(dstFile)
	// copy
	return io.Copy(writer, reader)
}

func main() {
	srcFileName := "./test.txt"
	dstFileName := "./newTest.txt"
	_, err := copyFile(srcFileName, dstFileName)
	if err == nil {
		fmt.Println("拷贝完成!")
	} else {
		fmt.Println("copy err = ", err)
	}
}
