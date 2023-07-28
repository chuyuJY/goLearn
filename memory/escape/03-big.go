package main

// []int 太大, 导致逃逸
func foo() {
	s := make([]int, 8193, 8193)
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
}

// 外部引用, 导致逃逸
//func foo() []int {
//	s := make([]int, 100, 100)
//	for i := 0; i < len(s); i++ {
//		s[i] = i
//	}
//	return s
//}

func main() {
	foo()
}
