package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 定义一个结构体，存储英文、数字、空格等数目
type Record struct {
	CharCount  int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func fileCount(fileName string) (Record, error) {
	/*
		1. 打开一个文件，创建reader
		2. 逐行统计
		3. 记录在结构体
	*/
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err = ", err)
		return Record{}, err
	}
	defer file.Close()
	record := Record{}
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		// 为了兼容中文，可以将string转[]rune
		// 遍历str，进行统计
		for _, v := range str {
			switch {
			case 'a' <= v && v <= 'z':
				fallthrough // 穿透
			case 'A' <= v && v <= 'Z':
				record.CharCount++
			case v == ' ' || v == '\t':
				record.SpaceCount++
			case '0' <= v && v <= '9':
				record.NumCount++
			default:
				record.OtherCount++
			}
		}
		if err == io.EOF {
			break
		}
	}
	return record, nil
}

func main() {
	fileName := "./test.txt"
	record, _ := fileCount(fileName)
	fmt.Println(record)
}
