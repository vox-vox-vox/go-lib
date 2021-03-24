package main
import "fmt"

func main(){
    var a = [10]int{1,2,3,4}
    b := a[:];  //# cap(b)=3
    s := append(b,4,5);  // newcap(b)= 2*cap(b) = 6 
    fmt.Println(s, len(s), cap(s)) //[] 0 0
}
