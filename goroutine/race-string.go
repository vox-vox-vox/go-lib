package main
import (
    "time"
    "fmt"
)

// 交替打印字符
// go-lib/goroutine/string-atom-race.go
const (
  FIRST  = "WHAT THE"
  SECOND = "F*CK"
)

func main() {
  var s string
  go func() {
    i := 1
    for {
      i = 1 - i
      if i == 0 {
        s = FIRST
      } else {
        s = SECOND
      }
      time.Sleep(10)
    }
  }()

  for {
    fmt.Println(s)
    time.Sleep(10)
  }
}
