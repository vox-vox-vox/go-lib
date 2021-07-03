package main

// stop routine:
//  1. 利用stop.C
//  2. 还有一个不好用的方法是用timer.Stop()， timer 必须没有过期才可以
import (
    "time"
)
func main() {
    timer := time.NewTimer(100*time.Millisecond)
    stopCh := make(chan int, 0) //默认是1, <-ch 第一个就block
    count := 0
    go func() {
        println("start goroutine:")
        for{
            select {
                case <-stopCh:
                    println("stop")
                    return
                case <-timer.C:
                    count += 1
                    println("loop:", count)
                    timer.Reset(100*time.Millisecond)
            }
        }
    }()
    time.Sleep(3*time.Second)
    println("quit")
}
