package main

import (
	"encoding/json"
	"fmt"
)

/*
	1. 将结构体序列化
	2. 将map序列化
	3. 对Slice序列化
	4. 对基本数据类型进行序列化（意义不大）
*/

// 1. 将结构体序列化
type Monster struct {
	Name     string  `json:"name" `
	Age      int     `json:"age" `
	Birthday string  `json:"birthday"`
	Salary   float64 `json:"salary"`
	Skill    string  `json:"skill"`
}

func testStruct() {
	monster := Monster{
		Name:     "牛魔王",
		Age:      18,
		Birthday: "2011-11-11",
		Salary:   3000.0,
		Skill:    "牛魔拳",
	}

	// 将monster 序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化错误 err = ", err)
	}
	fmt.Println("序列化的结果：", string(data))
}

// 2. 将map序列化
func testMap() {
	a := make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"

	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化错误 err = ", err)
	}
	fmt.Println("a map 序列化后为：", string(data))
}

// 3. 对Slice序列化
func testSlice() {
	slice := []map[string]interface{}{}
	m1 := map[string]interface{}{}
	m1["name"] = "jack"
	m1["age"] = 10
	m1["address"] = []string{"墨西哥", "夏威夷"}
	m2 := map[string]interface{}{}
	m2["name"] = "chuyu"
	m2["age"] = 18
	m2["address"] = "南京"

	slice = append(slice, m1, m2)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化错误 err = ", err)
	}
	fmt.Println("slice 序列化后为：", string(data))
}

// 4. 对基本数据类型进行序列化(意义不大)
func testFloat64() {
	num1 := 1234.5678
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Println("序列化错误 err = ", err)
	}
	fmt.Println("num1 序列化结果为：", string(data))
}

func main() {
	//testStruct()
	//testMap()
	testSlice()
	//testFloat64()
}
