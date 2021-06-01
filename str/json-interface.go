package main

import (
	"encoding/json"
	"fmt"
)

type A interface{}


type Task struct{
    name string
}
type Data struct{
    tasks *[]Task
    Age int
}


func main() {
	var infObj interface{}
    // example1
	fmt.Printf("example1:\n")
	json.Unmarshal([]byte(`{"name":"ahui","age":"20"}`), &infObj)
	fmt.Printf("infObj:%+v\n", infObj)
	fmt.Printf("infObj[name]:%+v\n", infObj.(map[string]interface{})["name"])


    // example2
	fmt.Printf("\n2nd Umarshal:\n")
    err:=json.Unmarshal([]byte(`{"name":"ahui2","age":"21"}`), &infObj)
	fmt.Printf("infObj:%+v, err=%v\n", infObj, err)
}
