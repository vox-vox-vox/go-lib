package main

import (
	"fmt"
)

type A struct {
    Is *bool
}

func main() {
    a := A{}
    var b bool
    a.Is = &b
    fmt.Println(a)
}
