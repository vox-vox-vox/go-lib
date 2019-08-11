package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "fmt"
)
type User struct {
    //gorm.Model
    ID   int    `gorm:"primary_key"`
    Name string
    Languages         []*Language `gorm:"many2many:uls;"`
}

type Language struct {
    ID   int    `gorm:"primary_key"`
    Locale string
    UserId int
    Users             []*User     `gorm:"many2many:uls;association_jointable_foreignkey:user_id;jointable_foreignkey:user_id;"` 
                /*
                association_jointable_foreignkey: 
                    JOIN ON language.user_id=users.id
                jointable_foreignkey: 
                    Where language.user_id in (2)
                HasMany:
                    association_foreignkey:ID;foreignkey:ID;
                */
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
    var languages []Language
    language := Language{}
    user :=User{}

    db.Create(&Language{Locale:"en"})
    db.Create(&Language{Locale:"zh",UserId:2})
    db.Create(&User{Name:"Ahui",ID:2})
    //db.First(&language, "id = ?", 2)
    language = Language{ID:2}

    db.Model(&language).Related(&users,  "Users")
    //SELECT "users".* FROM "users" INNER JOIN "languages" ON "languages"."user_id" = "users"."id" WHERE ("languages"."user_id" IN ('2'))

    pf("users:%+v;\nlanguage:%+v\n", users, language)
    pf("--------------\n")
    //db.Model(&user).Related(&languages, "Languages")
    //db.Preload("Languages").First(&user)
    pf("user:%+v;\nlanguages:%+v\n", user, languages)


  defer db.Close()
}
