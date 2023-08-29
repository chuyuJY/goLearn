package main

import "fmt"

func prioritySelect(ch1, ch2 <-chan string) {
	for {
		select {
		case val := <-ch1:
			fmt.Println(val)
		case val2 := <-ch2:
		priority: // 标签
			for {
				select {
				// 若 ch1 有值，先打印 ch1 的值，再打印 ch2 的值
				case val1 := <-ch1:
					fmt.Println(val1)
				default:
					break priority
				}
			}
			fmt.Println(val2)
		}
	}
}
