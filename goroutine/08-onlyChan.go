package main

func main() {
	intChan := make(chan<- int, 1)
	intChan <- 2
	//fmt.Println(<-intChan)
	stringChan := make(<-chan int, 1)
	stringChan <- 2

}
