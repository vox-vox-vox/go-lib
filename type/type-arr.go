package main
import "fmt"

type StringArray []string

func main(){
    var i StringArray = []string{"a"}
    fmt.Printf("%T\n", i)

    func(j []string){
        fmt.Printf("%T\n", j)
    }(i)
}
