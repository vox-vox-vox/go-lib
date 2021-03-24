package main

import (
    "fmt"
    "time"
)

func main() {
    certainSomething := true // will cause time loop to repeat

    timeDelay := 5 * time.Second // == 900000 * time.Microsecond

    var endTime <-chan time.Time // signal for when timer us up

    for {
        // if a certain something happens, start a timer
        if certainSomething && endTime == nil {
            fmt.Println("reset endTime:")
            endTime = time.After(timeDelay)
        }
        fmt.Println("endTime:%#v",endTime)

        select {
        case <-endTime:
            fmt.Println("Yes Finally!")
            endTime = nil
        }

        fmt.Println("I've been called")
        time.Sleep(time.Second) // simulate work
    }
}

