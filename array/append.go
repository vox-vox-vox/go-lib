package main
import "fmt"

func main(){
    x := []int{1,2}
    i:=append(x, 3)
    fmt.Printf("%v\n", x) //1,2
    fmt.Printf("%v\n", i) //1,2,3
}
