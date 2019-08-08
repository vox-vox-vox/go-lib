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


type CreditCard struct {
    gorm.Model
    Number string
    UID    string
}

type User struct {
    gorm.Model
    Name   string    `sql:"index"`
    CreditCard CreditCard `gorm:"foreignkey:uid;association_foreignkey:name"`
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
  db.AutoMigrate(&Product{},&User{},&CreditCard{})
  user:=User{Model:gorm.Model{ID:2}, Name:"ahui",}
  //user=User{ID:22}

  var card CreditCard
db.Model(&user).Related(&card, "CreditCard").First(&card)
//  profile:=Profile{}
 // c:=db.Model(&user).Related(&profile)
  //pf("%#v", c)


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
