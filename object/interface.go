package main

import (
    "fmt"
)

func main() {
    m:= map[string]interface{}{"a":1}
    fmt.Printf("a=%#T\n", m["a"].(int))
    fmt.Printf("a=%#T\n", m["a"])
    b:= map[string]interface{}{"b":[]string{"str1"}}
    fmt.Printf("b=%#T\n", b["b"])
}
