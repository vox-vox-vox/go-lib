package main

import (
    "fmt"
)

func main() {
    m:= map[string]interface{}{"a":1}
    fmt.Printf("a=%#T\n", m["a"].(int))
}
