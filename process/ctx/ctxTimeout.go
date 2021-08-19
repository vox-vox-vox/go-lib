package main

import (
	"context"
	"fmt"
    "time"
)

func main() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
                err:=ctx.Err()
                if err!=nil{
                    fmt.Println("context timeout:", err)
                    return
                }
				select {
				case <-ctx.Done():
                    fmt.Println("context done!n=", n)
					return
				case dst <- n:
                    fmt.Println("wait....,n=", n)
                    time.Sleep(1*time.Second)
                    fmt.Println("wait after")
					n++
				}
			}
		}()
		return dst
	}

    ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
