package main

import (
	"fmt"
    "time"
    "encoding/json"
)

type Stu struct {
	Name string `json:"name"`
	Age  int
}

func main() {
     m,_:= json.Marshal(map[string]interface{}{
        "a":1,
        "k2":"b",
        "k3":false,
        "time": time.Now(),
    })
    fmt.Println(string(m))

    stu := Stu{"ahui", 20}
     m,_= json.Marshal(stu)
    fmt.Println(string(m))
}
