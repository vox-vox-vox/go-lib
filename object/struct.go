package main

import (
    "fmt"
)

type A struct{
    Name string
}

type Wrapper struct{
    A
    Age int
}

//convertTest = wrapper.InnerStructType
// https://groups.google.com/forum/#!topic/golang-nuts/_zArSdOtuUg

func main() {
    a := A{"hilo"}
    b := Wrapper(a)

    fmt.Printf("a=%#T\n", b)
}
