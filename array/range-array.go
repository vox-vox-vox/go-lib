package main
import "fmt"
func main(){
    roles := []string{"hi","le"}
    for i, role:= range roles{
        fmt.Printf("i=%v,v=%v\n", i, role)
    }
    for i:= range roles{
        fmt.Printf("i=%v\n", i)
    }
}
