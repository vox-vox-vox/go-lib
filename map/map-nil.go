package main

import "fmt"

func arg_nil()(m map[string]int){
    //var m map[int]int
    //m[1]=2
    m["a"] = 1
    return
}

type Meta struct{
    Labels map[string]string
}
func struct_nil(){
    meta := &Meta{}
    meta.Labels["a"] = "b"
    return
}

func main() {
    struct_nil()
    m := arg_nil()
    fmt.Printf("%v\n", m)
}
