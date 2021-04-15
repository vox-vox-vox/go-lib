package main

import (
    "fmt"
    "time"
)

func main() {
    timer1 := time.NewTimer(1 * time.Second)

    <-timer1.C
    fmt.Println("Timer 1 ")
    // 只能用一次<-timer1.C
    //fmt.Println("Timer 1 ")

    //stop2 := timer2.Stop()
    time.Sleep(1 * time.Second)
}
