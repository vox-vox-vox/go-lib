package main

import (
    "github.com/go-redis/redis/v7"
    "fmt"
    "time"
        "encoding/json"

)

type User struct{
    Name string 
    Height int 
    age int
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
    user := User{"ahuigo", 20, 20}

    ustr,_ := json.Marshal(user)


    // get
	client.Set("key", ustr, 0)
    val, err := client.Get("key").Result()
    fmt.Printf("Get struct val:%#v, err:%v\n", val, err!=nil)

    // nil
	err = client.Set("key", user, time.Duration(1000*1000)).Err()
    if err!=nil{
        fmt.Printf("Set err:%#v, err:%v\n", err, err)
    }

	err = client.Get("key").Err()
    fmt.Printf("err:%#v, %v\n", err, err!=nil)

}
