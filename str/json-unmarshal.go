package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

type Task struct {
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
	fmt.Println("1. unmarshal只更新局部属性:")
	err := json.Unmarshal([]byte(`{"name":"ahui","age":"20"}`), &stu)
	fmt.Println("err:", err, ",stu:", stu)

	err = json.Unmarshal([]byte(`{"age":"21"}`), &stu)
	fmt.Println("err:", err, ",stu:", stu)

	fmt.Println("\n2. 要传真指针&obj--而不是obj或nil-----------------")
	var obj *map[string]interface{}
	err = json.Unmarshal([]byte(`{"name":"ahui","age":"20"}`), obj)
	fmt.Println("err:", err, ",*obj->obj:", obj)

	var obj3 map[string]interface{}
	err = json.Unmarshal([]byte(`{"name":"ahui","age":"20"}`), obj3)
	fmt.Println("err:", err, ",obj3->obj:", obj3)

	var obj2 *map[string]interface{}
	err = json.Unmarshal([]byte(`{"name":"ahui","age":"20"}`), &obj2)
	fmt.Println("err:", err, ",*obj2->&obj:", obj2)

	var obj4 map[string]interface{}
	err = json.Unmarshal([]byte(`{"name":"ahui","age":"20"}`), &obj4)
	fmt.Println("err:", err, ",obj4->&obj:", obj4)

	fmt.Printf("\n3. infObj:--------\n")
	var infObj interface{}
	json.Unmarshal([]byte(`{"name":"ahui","age":"20"}`), &infObj)
	fmt.Printf("infObj:%+v\n", infObj)
	fmt.Printf("infObj[name]:%+v\n", infObj.(map[string]interface{})["name"])


	fmt.Printf("\n4. tasks:--------\n")
    tasks :=  []*Task{}
	json.Unmarshal([]byte(`[{"name":"ahui","age":"20"}]`), &tasks)
	fmt.Printf("tasks:%+v\n", *(tasks[0]))

	return stu
}
