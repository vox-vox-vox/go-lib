package main

import (
	"fmt"
        "encoding/json"
)
func main() {
 m,_:= json.Marshal(map[string]interface{}{
        "a":1,
        "k2":"b",
        "k3":false,
    })
    fmt.Println(string(m))
}
