package main
import "errors"
import "fmt"


func main(){
    var e1 error = nil
    var e2 error = errors.New("msg")
    s:=fmt.Sprintf("nil=%v, err=%v", e1, e2)
    fmt.Println(s)

}
