package main

import "fmt"

func f()(m map[string]int){
    m["a"] = 1
    return
}

func main() {
    m := map[string][]string{"k":[]string{"v"}}
    fmt.Printf("m=%v\n", m)
    fmt.Printf("m[k2]=%#v\n", m["k2"]) //[]string(nil)
    fmt.Println([]string(nil))
}
