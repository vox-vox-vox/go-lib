package main
import "fmt"
type User struct{
    age int
 }
type Class struct{
    User
    id int
}
func main(){
    //c := Class{User:User{age:0}, id:0}
    c := Class{}
    c.User = User{age: 15}
    fmt.Println(c)
    fmt.Println(c.age)
}
