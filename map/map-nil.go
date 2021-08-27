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
    // test nil[k] exists
    var m1 map[string]int
    if _,ok :=m1["a"]; !ok{
        println("nil[a] not exists. ")
    }

    // test struct nil
    struct_nil()
    // test  arg nil
    m := arg_nil()
    fmt.Printf("%v\n", m)
}
