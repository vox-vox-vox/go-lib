package main

type User struct {
    name string
}

func f(user *User){
    user.name = "ahui"
}

func main(){
    user := &User{}
    f(user)
    println(user.name)

}
