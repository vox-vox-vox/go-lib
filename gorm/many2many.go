package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "fmt"
)
type User struct {
    gorm.Model
    Languages         []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
    ID   int    `gorm:"primary_key"`
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

    db.Create(&Language{Name:"en"})
    db.Create(&Language{Name:"zh"})
    db.Create(&User{})
    //db.First(&language, "id = ?", 2)
    language = Language{ID:2}

    db.Model(&language).Related(&users,  "Users")
    //// SuELECT * FROM "users" INNER JOIN "user_languages" ON "user_languages"."user_id" = "users"."id" WHERE  ("user_languages"."language_id" IN ('111'))

    pf("user:%+v;\n:language:%+v\n", users, language)

  defer db.Close()
}
