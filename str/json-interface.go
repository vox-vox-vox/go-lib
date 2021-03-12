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
	fmt.Printf("\n")

	return infObj
}
