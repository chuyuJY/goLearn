package _3_timer

import (
	"testing"
)

/*
	StartTimer() & StopTimer()
	应用场景：
	// 每次函数调用前后需要一些准备工作和清理工作，可以使用 StopTimer 暂停计时、使用 StartTimer 开始计时。
*/

// StartTimer() & StopTimer()
func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer() // 暂停计时
		nums := generateWithCap(10000)
		b.StartTimer() // 开始计时
		bubbleSort(nums)
	}
}

func BenchmarkFastSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer() // 暂停计时
		nums := generateWithCap(10000)
		b.StartTimer() // 开始计时
		fastSortArray(nums)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer() // 暂停计时
		nums := generateWithCap(10000)
		b.StartTimer() // 开始计时
		mergeSortArray(nums)
	}
}
