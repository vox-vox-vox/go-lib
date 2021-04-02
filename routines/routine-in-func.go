package main

import (
	"fmt"
	"time"
)

func test() {
	go func() {
		for i := 0; i < 1e3; i++ {
			n := time.Now().Format(time.RFC3339)
			println(n)
			time.Sleep(time.Second)
		}
	}()
}
func main() {
	test()
	fmt.Println("test函数执行结束不会中断goroutines")
	time.Sleep(1000 * time.Second)
}
