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
    db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=ahuigo sslmode=disable password=")
    db.LogMode(true)
    if err != nil {
        panic("failed to connect database")
    }
    t1 := time.Now()
    u := User{ID: 1}
    profile := Profile{ID: 1}
    db.Model(&u).Related(&profile)
    //db.Model(&u).Find(&u)
    //db.Model(&u).Related(&u.CreditCard,"CreditCard")
    //db.Model(&u).Related(&u.Emails,"Emails")
    //db.Model(&u).Related(&u.Languages, "Languages")
    //db.Model(&u).Related(&u.BillingAddress,"BillingAddress")
    //db.Model(&u).Related(&u.ShippingAddress,"ShippingAddress")
    fmt.Println(u)
    t2 := time.Now()
    fmt.Println("gorm耗时:", t2.Sub(t1))
}

