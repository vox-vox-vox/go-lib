package main

import (
	"fmt"
    "encoding/json"
)

type Stu struct {
	Name string `json:"name"`
	Age  int
	height  int
}

type Dict map[string]interface{}

func main() {
     m,_:= json.Marshal(Dict{
        "a":1,
        "k2":"b",
        "k3":false,
        "time": Dict{
            "a" : 222222,
        },
    })
    fmt.Println("Marshal map with lowercase:",string(m))

    stu := Stu{"ahui", 20, 100}
    m,_= json.Marshal(stu)
    fmt.Println("Warn: Marshal struct will ignore lowercase:", string(m))
    // json 不用转
    var stuI interface{}=stu
    m,_= json.Marshal(stuI)
    fmt.Println("json(interface):")
    fmt.Println(string(m))
}
