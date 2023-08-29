package main

//var (
//	catChan  = make(chan int)
//	fishChan = make(chan int)
//	dogChan  = make(chan int)
//	allChan  = make(chan struct{})
//)
//
//func cat(ch chan int) {
//	for v := range ch {
//		fmt.Printf("%v cat\n", v)
//		fishChan <- v
//		if v == 10 {
//			break
//		}
//	}
//}
//
//func fish(ch chan int) {
//	for v := range ch {
//		fmt.Printf("%v fish\n", v)
//		dogChan <- v
//		if v == 10 {
//			break
//		}
//	}
//}
//
//func dog(ch chan int) {
//	for v := range ch {
//		fmt.Printf("%v dog\n", v)
//		if v == 10 {
//			allChan <- struct{}{}
//			break
//		}
//		catChan <- v + 1
//	}
//}
//
//func main() {
//	go cat(catChan)
//	go fish(fishChan)
//	go dog(dogChan)
//	catChan <- 1
//	<-allChan
//}
