package main

import "reflect"

type animal struct {
	name string
}

func (a *animal) Eat() {
	println("animal eat")
}

func main() {
	a := animal{"cat"}
	reflect.ValueOf(&a).MethodByName("Eat").Call([]reflect.Value{})
}
