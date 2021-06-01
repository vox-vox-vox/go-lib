package main
import (
    "time"
    "fmt"
)

// 交替打印字符
// go-lib/goroutine/race-string.go
const (
  FIRST  = "WHAT THE Fuck"
  SECOND = "F*CK 123456 89 abcdef"
)

func main() {
  var s string = FIRST
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
      //time.Sleep(10*time.Second)
    }
  }()

  for {
      if (s==FIRST || s==SECOND){
        fmt.Println(s)
          fmt.Printf("first string:%v\n",s==FIRST)
          fmt.Printf("second string:%v\n",s==SECOND)
      }else{
          fmt.Printf("---------------\nlast string:%v\n",s)
          fmt.Printf("first string:%v\n",s==FIRST)
          fmt.Printf("second string:%v\n",s==SECOND)
          panic("err string")
      }
    time.Sleep(10)
  }
}
