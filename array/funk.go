package main
import "fmt"
import "github.com/thoas/go-funk"

func main(){
    a:=[]string{"a"}
    b:=[]string{"a"}
    fmt.Println(b)

    res := funk.Contains(a,b) //false
    fmt.Println(res)
    res = funk.Subset(a,b) //false
    fmt.Println(res)
}
