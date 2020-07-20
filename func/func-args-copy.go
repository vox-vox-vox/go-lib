package main
import "fmt"
type User struct{
    Name string
}

func main(){
    user := User{"ahui"}
    func(u User){
        u.Name ="LILY"
        fmt.Println(u) //LILY
    }(user)
    fmt.Println(user) //ahui
    
}

