package main

import "fmt"

func f()(m map[string]int){
    m["a"] = 1
    return
}

func main() {
    m := f()
    fmt.Printf("%v\n", m)
}
