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
	var stu *A
	// 1. 不能直接访问nil pointer： stu.Name = "ahui"
    // 2. &stu 是指向指针的指针
	err := json.Unmarshal([]byte(`{"name":"ahui","age":"20"}`), &stu)
	fmt.Println("err:", err, ",stu:", stu)

	err = json.Unmarshal([]byte(`{"age":"21"}`), &stu)
	fmt.Println("err:", err, ",stu:", stu)
	return stu
}
