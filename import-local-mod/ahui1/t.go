/*
    这是一个本地包
 */
package ahui1
import "time"
import "fmt"

var now = time.Now()
func init() {
    fmt.Printf("now: %v\n", now)
}
func init() {
    fmt.Printf("since: %v\n", time.Now().Sub(now))
}

func Test3(){
    println("test3 by Ahuix")
}
