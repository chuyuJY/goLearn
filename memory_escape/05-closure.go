package main

func foo() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	foo()()
}
