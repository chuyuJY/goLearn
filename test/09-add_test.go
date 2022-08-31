package test

import "testing"

// 1. 第一个测试函数
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func TestGetSub(t *testing.T) {
	res := getSub(10, 7)

	// 若错误
	if res != 5 {
		t.Fatalf("getSub(10, 7) 执行错误 期望值 = %v 实际值 = %v\n", 5, res)
	}
	// 若正确
	t.Logf("getSub(10, 7) 执行正确...")
}

func TestAddUpper(t *testing.T) {
	// 调用
	res := addUpper(10)
	// 若错误
	if res != 55 {
		t.Fatalf("addUpper(10) 执行错误 期望值 = %v 实际值 = %v\n", 55, res)
	}
	// 若正确
	t.Logf("addUpper(10) 执行正确...")
}
