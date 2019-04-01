/*
Package ahui1 这是一个本地包
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

/*
Add exported
*/
func Add(i,j int) int {
    return i+j
}

/*
Test3 exported
*/
func Test3(){
    println("test3 by Ahuix")
}
