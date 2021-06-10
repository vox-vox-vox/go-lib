package main

import (
    "github.com/go-redis/redis/v7"
    "fmt"
    "time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := client.Set("key", "value", time.Duration(1000*1000)).Err()
    println(err)

	val, err := client.TTL("key").Result()
	if err != nil {
		panic(err)
	}
    fmt.Printf("ttl:%#v, %T\n", val==-2, val)
}
