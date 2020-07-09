package main

import (
        "fmt"
        "unsafe"
)

func main() {
        var i1 int = 1
        var i2 int8 = 2
        fmt.Println("sizeof int8:", unsafe.Sizeof(i2))
        fmt.Println("sizeof int:",unsafe.Sizeof(i1))
        fmt.Printf("type: %T",1<<63 - 1)
}
