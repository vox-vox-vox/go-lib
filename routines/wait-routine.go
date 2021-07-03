package main

import (
    "sync"
    "time"
)
func main() {
    var wg sync.WaitGroup
    wg.Add(1)

    ch := make(chan int, 0) //默认是1, 0 代表<-ch 第一个就block
    go func() {
        println("foo")
        for {
            time.Sleep(1*time.Second)
            foo, ok := <- ch
            if !ok {
                println("done")
                wg.Done()
                return
            }
            println(foo)
            time.Sleep(time.Second)
        }
    }()
    ch <- 1
    println("send:1")
    ch <- 2
    ch <- 3

    func(){
        close(ch)
        wg.Wait()
        println("quit")
    }()
}
