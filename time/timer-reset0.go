package main

import (
    "fmt"
    "time"
)

func main() {
    timer := time.NewTimer(1 * time.Second)
    count :=0

    go func(){
        for {
            <-timer.C //quit goroutine if stop timmer
            count +=1
            fmt.Println("Timer count: ",count)
            if count==2{
                fmt.Println("Timer reset 2,",count)
                timer.Reset(-2*time.Millisecond)
                time.Sleep(time.Second)
                continue
            }
            if count>=3{
                fmt.Println("Timer break 3,",count)
                break
            }
            timer.Reset(time.Second)
        }
        fmt.Println("Timer quit: ",count)
    }()
    time.Sleep(500*time.Millisecond)
    println("sleep stop")
    timer.Stop()

    time.Sleep(2000 *time.Second)
    println("quit")

}
