package main

import (
	"fmt"
	"reflect"
)

// 1. 定义Monster结构体
type Monster struct {
	Name  string `json:"name,omitempty"`
	Age   int    `json:"age,omitempty"`
	Score float64
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s *Monster) Set(name string, age int, score float64, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func testStruct(a interface{}) {
	rTyp := reflect.TypeOf(a)
	rVal := reflect.ValueOf(a)
	rKd := rVal.Kind()
	if rKd != reflect.Struct {
		fmt.Println("expect struct...")
		return
	}
	num := rVal.NumField()
	fmt.Printf("struct has %d fields.\n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 字段值 = %v.\n", i, rVal.Field(i))
		tagVal := rTyp.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: 字段tag = %v.\n", i, tagVal)
		}
	}
	numOfMethod := rVal.NumMethod()
	fmt.Printf("struct has %d methods.\n", numOfMethod)
	// Method调用方法时，按照函数字母(ASCII码)进行排序，和定义时候的顺序无关。
	// 调用 Print 函数
	rVal.Method(1).Call(nil)
	// 赋值
	params := []reflect.Value{}
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	// 调用 GetSum 函数
	res := rVal.Method(0).Call(params)
	fmt.Println("res =", res[0].Int()) // 返回结果为 []reflect.Value
}

func changeStruct1(a interface{}) {
	rTyp := reflect.TypeOf(a)
	fmt.Println("rTyp =", rTyp)
	rVal := reflect.ValueOf(a)
	ptr := rVal.Elem()
	// 通过字段名 修改 Age的值
	ptr.FieldByName("Age").SetInt(100)
}

func changeStruct2(a interface{}) {
	rTyp := reflect.TypeOf(a)
	fmt.Println("rTyp =", rTyp)
	rVal := reflect.ValueOf(a)
	ptr := rVal.Interface().(*Monster)
	ptr.Set("mouse", 12, 30.0, "unknown")
}

func main() {
	a := Monster{
		Name:  "cat",
		Age:   12,
		Score: 20,
		Sex:   "unknown",
	}
	//testStruct(a)
	//changeStruct1(&a)
	//changeStruct2(&a)
	testStruct(a)
}
