package main

import "testing"

// 1. 基础使用
// go test -bench .					// 默认测试
// go test -bench='Fib$' . 			// 传入一个正则表达式，只运行以 Fib 结尾的 benchmark 用例
// go test -bench='Fib$' -cpu=2,4 . // 指定 GOMAXPROCS 数，可以传入多个值，当作测试列表

// 2.提升准确度
// go test -bench='Fib$' -benchtime=5s .  // 指定测试时长 5s
// go test -bench='Fib$' -benchtime=50x . // 指定测试次数 50次
// go test -bench='Fib$' -benchtime=5s -count=3 . // 指定测试轮数

func BenchmarkFib(b *testing.B) {
	// b.N 从 1 开始，如果该用例能够在 1s 内完成，b.N 的值便会增加，再次执行。
	// b.N 的值大概以 1, 2, 3, 5, 10, 20, 30, 50, 100 这样的序列递增，越到后面，增加得越快。
	for i := 0; i < b.N; i++ {
		fib(30) // run fib(30) b.N times
	}
	// BenchmarkFib-12 中的 -12 即 GOMAXPROCS，默认等于 CPU 核数。
}
