package main

import "fmt"

/*
	1. struct{}{} 可用作实现 Set
	Go 语言中没有 Set，可以利用 Map 实现 Set，但是就算把 val 置为 bool，也会占用 1 字节。
	此时，可以将 val 置为 struct{}，仅作为占位符，不占用内存。
*/

type Set map[string]struct{}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Add(key string) {
	s[key] = struct{}{}
}

func (s Set) Delete(key string) {
	delete(s, key)
}

func main() {
	s := make(Set)
	s.Add("Tom")
	s.Add("Sam")
	fmt.Println(s.Has("Tom"))
	fmt.Println(s.Has("Sam"))
}
