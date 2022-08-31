package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil { // 文件目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}

// 将file1的内容复制到file2
func readAndWrite() {
	file1Path := "./test.txt"
	file2Path := "./testOpen.txt"

	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Println("read file err = ", err)
		return
	}
	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Println("write file err = ", err)
		return
	}
}
func main() {
	readAndWrite()
}
