package main

//var total struct {
//	sync.Mutex
//	value int
//}
//
//func worker(wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	// 保证原子性
//	for i := 0; i < 100; i++ {
//		total.Lock()
//		total.value += 1
//		total.Unlock()
//	}
//}
//
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(2)
//	go worker(&wg)
//	go worker(&wg)
//	wg.Wait()
//	fmt.Println(total.value)
//}
