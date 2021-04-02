package main
import "fmt"
func test() {
    defer func() {
        fmt.Println("recover:",recover())
    }()
    defer func() {
        panic("test panic2")
    }()
    panic("test panic1")
}
func main() {
    test()
}
