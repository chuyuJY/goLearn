package main

import (
	"encoding/json"
	"fmt"
)

/*
	1. 反序列化为struct
	2. 反序列化为map
	3. 反序列化为slice
*/

// 1. 反序列化为 struct
type Monster struct {
	Name     string  `json:"name" `
	Age      int     `json:"age" `
	Birthday string  `json:"birthday"`
	Salary   float64 `json:"salary"`
	Skill    string  `json:"skill"`
}

func unMarshalStruct() {
	// 1. 得到序列化后的字符串，比如网络传输过来的
	str := "{\"name\":\"牛魔王\",\"age\":18," +
		"\"birthday\":\"2011-11-11\",\"salary\":3000,\"skill\":\"牛魔拳\"}"
	monster := Monster{}
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("反序列化错误 err = ", err)
	}
	fmt.Printf("反序列化的结果：%v\n", monster)
}

// 2. 反序列化为Map
func unMarshalMap() {
	str := " {\"address\":\"洪崖洞\",\"age\":30,\"name\":\"红孩儿\"}"
	a := map[string]interface{}{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Println("反序列化失败 err = ", err)
	}
	fmt.Printf("序列化结果为：%v\n", a)
}

// 3. 反序列化为Slice
func unMarshalSlice() {
	str := "[{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":10," +
		"\"name\":\"jack\"},{\"address\":\"南京\",\"age\":18,\"name\":\"chuyu\"}]"
	slice := []map[string]interface{}{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("反序列化失败 err = ", err)
	}
	fmt.Printf("反序列化结果为：%v\n", slice)
}

func main() {
	//unMarshalStruct()
	//unMarshalMap()
	unMarshalSlice()
}
