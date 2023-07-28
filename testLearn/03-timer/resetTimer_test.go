package _3_timer

import (
	"testing"
	"time"
)

/*
	ResetTimer()
	应用场景：
		// 如果在 benchmark 开始前，需要一些准备工作，
		// 如果准备工作比较耗时，则需要将这部分代码的耗时忽略掉。
*/

// ResetTimer() 重置定时器
func BenchmarkFib(b *testing.B) {
	time.Sleep(time.Second * 3) // 模拟耗时准备任务
	b.ResetTimer()              // 重置定时器
	for i := 0; i < b.N; i++ {
		fib(30)
	}
}
