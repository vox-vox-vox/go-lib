package main

import (
	"encoding/json"
	"fmt"
)

type A interface{}

func main() {
	t()
	//fmt.Println(data)
}
func t() A {
	var infObj interface{}
	fmt.Printf("infObj:\n")
	json.Unmarshal([]byte(`{"name":"ahui","age":"20"}`), &infObj)
	fmt.Printf("infObj:%+v\n", infObj)
	fmt.Printf("infObj[name]:%+v\n", infObj.(map[string]interface{})["name"])

    type Task struct{
        name string
    }
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

	return infObj
}
