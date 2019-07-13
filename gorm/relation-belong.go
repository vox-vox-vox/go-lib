package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "time"
)

type User struct {
  gorm.Model
  Profile   Profile
  ProfileID int
}

type Profile struct {
  gorm.Model
  Name string
}


func main() {
    type A struct{int}
    a:=A{}
    b:=A{}
    fmt.Printf("%p,%p,%v\n", &a,&b, a.int)  //not same
    b = a
    fmt.Printf("%p,%p\n", &a,&b)  //not same

time.Now()
}

