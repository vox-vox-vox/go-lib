package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "fmt"
)

type Product struct {
  gorm.Model
  Code string
  Price uint
}

// User has and belongs to many languages, use `user_languages` as join table
type User struct {
    gorm.Model
    Languages         []Language `gorm:"many2many:user_languages;"`
}

type Profile struct {
    gorm.Model
  Name      string
  User      User `gorm:"foreignkey:UserRefer"` // use UserRefer as foreign key
  UserRefer uint
}

type Language struct {
    gorm.Model
    Name string
}
func main() {

  db, err := gorm.Open("postgres", "host=localhost user=role1 dbname=ahuigo sslmode=disable password=")
  if err != nil {
    println(err)
    fmt.Println(err)
    panic("连接数据库失败")
  }


  // 自动迁移模式
      db.LogMode(true)
  db.AutoMigrate(&Product{},&Profile{},Language{},User{})

  user:=User{Model:gorm.Model{ID:1}}
  profile:=Profile{}
db.Model(&user).Related(&profile)
return

//// SELECT * FROM "users" INNER JOIN "user_languages" ON "user_languages"."user_id" = "users"."id" WHERE  ("user_languages"."language_id" IN ('111'))

  // 创建
  db.Create(&Product{Code: "L1217", Price: 17})
  db.Create(&Product{Code: "L1218", Price: 18})

  // 读取
  var product Product
  db.First(&product, 1) // 查询id为1的product
fmt.Println(product)
  db.First(&product, "code = ?", "L1214") // 查询code为l1212的product
fmt.Println(product)

  // 更新 - 更新product的price为2000
  db.Model(&product).Update("Price", 22)

  defer db.Close()
}
