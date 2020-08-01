package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func main() {
	data := t()
	fmt.Println(data)
}
func t() *A {
	var x *A
	// 不能直接访问nil pointer： x.Name = "ahui", 但是json 却是可以的
	err := json.Unmarshal([]byte(`{"name":"ahui","age":"20"}`), &x)
	fmt.Println("err:", err, ",x:", x)
	err = json.Unmarshal([]byte(`{"age":"21"}`), &x)
	fmt.Println("err:", err, ",x:", x)
	return x
}
