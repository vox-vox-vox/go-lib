package main

import (
	"fmt"
    "encoding/json"
)

type Stu struct {
	Name string `json:"name"`
	Age  int
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
    fmt.Println(string(m))

    stu := Stu{"ahui", 20}
     m,_= json.Marshal(stu)
    fmt.Println(string(m))
}
