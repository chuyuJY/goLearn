package main

import "fmt"

// 1. 创建一个int型chan
func example1() {
	/*
		演示channel的使用
	*/
	//	1. 创建一个int型chan
	//var intChan chan int
	//intChan = make(chan int, 3)
	intChan := make(chan int, 3)
	fmt.Println("看看intChan是什么？", intChan) // 输出intChan的地址 0xc000020100

	// 2. 向管道写入数据
	// 写数据时，不能超过cap；若超过则会deadlock
	intChan <- 10
	num := 211
	intChan <- num

	// 3. 看看管道的length和cap
	fmt.Println("channel len = ", len(intChan), ", cap = ", cap(intChan))

	// 4. 从管道读取数据
	// 读数据时，不能没数据了还继续读，否则deadlock
	num1 := <-intChan // 10， 先进先出
	fmt.Println("num1 = ", num1)
	fmt.Println("channel len = ", len(intChan), ", cap = ", cap(intChan))
}

// 2. 创建一个所有类型的chan
type Cat struct {
	Name string
	Age  int
}

func example2() {
	allChan := make(chan interface{}, 3)
	cat := Cat{
		Name: "小花猫",
		Age:  11,
	}
	allChan <- cat
	allChan <- 10
	allChan <- "chuyu"

	// 1. 类型断言 方式一
	getCat := (<-allChan).(Cat)
	fmt.Printf("getCat = %T, getCat = %v\n", getCat, getCat)
	// 不能直接取Name，因为interface{}
	fmt.Println("getCat.Name = ", getCat.Name)
	// 2. 类型断言 方式二
	//newCat := getCat.(Cat)
	//fmt.Println("getCat.Name = ", newCat.Name)
}

func main() {
	example2()
}
