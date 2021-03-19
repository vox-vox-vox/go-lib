package main

import (
	"encoding/json"
	"fmt"
)

type Task struct{
    name string
}

func main() {
    type A interface{}
	var infObj interface{}
    var tasks []Task
    data := struct{
        data *[]Task
        Age int
    }{
        data:&tasks,
    }
    rawbody := []byte(`{"data":[{"name":"hilojack"}], "Age":2}`)
    json.Unmarshal(rawbody, &infObj)
	fmt.Printf("infObj:%+v\n", infObj)
    err:=json.Unmarshal(rawbody, &data)
    fmt.Printf("data:%+v, err:%v\n", data, err)
	fmt.Printf("tasks:%+v\n", tasks)

}
