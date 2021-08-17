package main

import (
	"encoding/json"
	"fmt"
)

type Task struct{
    Name string
}

func main() {
    type A interface{}
	var infObj interface{}
    var tasks []Task
    data := struct{
        Tasks *[]Task
        Age int
        Weight int
        height int
        Other int
    }{
        Tasks:&tasks,
    }
    rawbody := []byte(`{"tasks":[{"name":"hilojack"}], "Age":2}`)
    json.Unmarshal(rawbody, &infObj)
	fmt.Printf("infObj:%#v\n\n", infObj)

    err:=json.Unmarshal(rawbody, &data)
    fmt.Printf("data:%+v, err:%v\n", data, err)
    fmt.Printf("data.tasks:%#v, err:%v\n", *data.Tasks, err)
	fmt.Printf("tasks:%#v\n", tasks)

}
