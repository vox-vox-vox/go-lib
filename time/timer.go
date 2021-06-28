package main

import (
    "fmt"
    "time"
)

func main() {
    timer1 := time.NewTimer(1 * time.Second)

    go func(){
        for {
            fmt.Println("Timer 0 ")
            <-timer1.C //quit goroutine if stop timmer
            fmt.Println("Timer 1 ")
            timer1.Stop()
            timer1.Stop()
            break
        }

        // 只能用一次
        //<-timer1.C
        //fmt.Println("Timer 2 ")
        //timer1.Reset(10)
    }()

    //stop2 := timer2.Stop()
    time.Sleep(500 *time.Millisecond)
    timer1.Stop()

    time.Sleep(2000 *time.Millisecond)
    println("quit")

}
