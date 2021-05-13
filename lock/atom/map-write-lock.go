package main
import (
    "time"
    "fmt"
)

type User struct{
    Name string
}

func main() {
    m := make(map[string]interface{})
  go func() {
      for {
        time.Sleep(1)
        m["abc"] = 134
        time.Sleep(1)
        m["abc"] = User{}
      }
  }()

      for {
        time.Sleep(1*time.Second)
          fmt.Println(m)
      }

}
