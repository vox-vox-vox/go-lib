package main
import (
    "log"
    "encoding/json"
    "fmt"
)
func f() int {
    return 1
}

type A struct{
    Name string `json:"name"`
    Age string `json:"age"`
}

var c = f()

func main(){
    data:=t()
    fmt.Println(data)
}
func t()(x *A){
    log.Println(c)
    var r *A
    fmt.Println(r)
    err :=json.Unmarshal([]byte(`{"name":"ahui","age":"20"}`), &x)
    fmt.Println(err, r,"end",x)
    err =json.Unmarshal([]byte(`{"age":"21"}`), &x)
    fmt.Println(err, r,"end",x)
    return
}
