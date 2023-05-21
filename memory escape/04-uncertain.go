package main

func foo() {
	n := 1
	s := make([]int, n)
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
}

func main() {
	foo()
}
