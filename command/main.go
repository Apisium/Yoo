package main

import (
	// "yoo"
	"fmt"
	"time"
	// "reflect"
)

func main() {
	test()
	// yoo.T()
	// if err := yoo.ExecuteFile("a.yoo"); err != nil {
	// 	fmt.Println("Error:", err)
	// }
}
type Map []interface{}
func test() {
	defer trace()()
	for i := 0; i < 1000000; i++ {
		_ = make(Map, 0, 0)
	}
}

func trace() func() {
  start := time.Now()
  return func() {
    fmt.Printf("Cost: %s\n", time.Since(start))
  }
}
