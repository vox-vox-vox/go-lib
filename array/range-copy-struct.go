package main
import "fmt"
type Role struct{
    Users []string
}
type Roles []*Role
func main(){
    roles := Roles{&Role{[]string{"ahui"}}}
    for _, role:= range roles{
        role.Users = append(role.Users, "a1")
        role.Users = append(role.Users, "b1")
        fmt.Println("role:",role)
        fmt.Println("roles[0]:",roles[0])
        fmt.Printf("%p,%p\n", role,roles[0])
    }
}
