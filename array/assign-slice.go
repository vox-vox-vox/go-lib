package main
import "fmt"
func main(){
    a := []int{1,2,3}
    b := a
    a[0]=100
    fmt.Printf("%v,%v\n", b,a)  //same
    fmt.Printf("%p,%p\n", b,a)  //same
}
