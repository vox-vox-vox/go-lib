package main
import "fmt"
func main(){
    defer func() {
        fmt.Printf("Recovered in f: %#v\n",recover())
        fmt.Printf("Recovered in f: %#v\n",recover()) //nil
    }()
    panic("panic error!")
}
