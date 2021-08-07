package main
import (
    "sync"
    "time"
)

// stop routine:
//  1. 利用stop.C
//  2. 还有一个不好用的方法是用timer.Stop()， timer 必须没有过期才可以

func main() {
    var wg sync.WaitGroup

    ch := make(chan int, 0) //默认是1, 0 代表<-ch 第一个就block
    go func() {
        wg.Add(1)
        defer wg.Done()
        println("foo")
        for {
            time.Sleep(1*time.Second)
            foo, ok := <- ch
            if !ok {
                println("done")
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

