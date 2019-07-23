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

type User struct {
  gorm.Model
  Name string
}

// `Profile` belongs to `User`, `UserID` is the foreign key
type Profile struct {
  gorm.Model
  UserID int
  User   User
  Name   string
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
  db.AutoMigrate(&Product{},&Profile{},&User{})
  user:=User{Model:gorm.Model{ID:2}}
  user=User{}

  profile:=Profile{}
  c:=db.Model(&user).Related(&profile).First(&profile)
  pf("%#v", c)


  // 创建
  /*
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
  */

  defer db.Close()
}
