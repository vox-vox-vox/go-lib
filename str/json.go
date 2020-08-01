package main

import (
	"fmt"
    "time"
    "encoding/json"
)
func main() {
     m,_:= json.Marshal(map[string]interface{}{
        "a":1,
        "k2":"b",
        "k3":false,
        "time": time.Now(),
    })
    fmt.Println(string(m))
}
