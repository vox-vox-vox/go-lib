package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "fmt"
)

// User has and belongs to many languages, use `user_languages` as join table
type User struct {
    gorm.Model
    Languages         []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
    gorm.Model
    Name string
    Users             []*User     `gorm:"many2many:user_languages;"`
}



func main() {
  db, err := gorm.Open("postgres", "host=localhost user=role1 dbname=ahuigo sslmode=disable password=")
  pf := fmt.Printf
  pf("'1'")
  if err != nil {
    println(err)
    fmt.Println(err)
    panic("连接数据库失败")
  }
  db.LogMode(true)
  //db.DropTable("users", "profiles")
  // 自动迁移模式
  db.AutoMigrate(&User{},&Language{})

var users []User
language := Language{}

language = Language{Model:gorm.Model{ID:1}}
    

db.Model(&language).Related(&users,  "Users")
//// SELECT * FROM "users" INNER JOIN "user_languages" ON "user_languages"."user_id"


    pf("user:%+v;\n:language:%+v\n", users, language)

  defer db.Close()
}
