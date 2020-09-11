package main
import (
	"github.com/go-redis/redis"
    "fmt"

)

func main(){

	var client *redis.Client

    // 还可用universal 模式
    client = redis.NewFailoverClient(&redis.FailoverOptions{
        MasterName:    "master",
        SentinelAddrs: []string{"host1:port1","host2:port2"},
        Password:"password",
    })
    fmt.Println("echo:",client.Echo("abc"))
    client.Set("k","ahui",0)
    v:=client.Get("k")
    fmt.Println("get:",v)

}
