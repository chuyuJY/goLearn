package main

import (
	_ "embed"
	"fmt"
	"time"
)

// //go:embed preQ.txt
// var preQ string
//
// //go:embed preA.txt
// var preA string
func main() {
	ts := time.Now().UnixMilli()
	fmt.Println(ts)

	fmt.Println(time.Unix(0, int64(time.Duration(ts)*time.Millisecond)).UTC())
	fmt.Println(time.UnixMilli(ts).UTC())
	fmt.Println(ts)
}
