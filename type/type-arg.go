package main
import "fmt"

type MyInt int

func test1(i int){
    fmt.Printf("%T\n", i)
}
func test2(i MyInt){
    fmt.Printf("%T\n", i)
}

func main(){
    var i int= 1
    test1(i)
    test2(MyInt(i))
    //error: test1(MyInt(i))
}
