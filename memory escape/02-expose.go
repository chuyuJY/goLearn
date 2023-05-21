package main

func foo() *int {
	a := 123
	return &a
}

func main() {
	_ = foo()
}
