package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "fmt"
)

var db *gorm.DB

type  Product struct {
  gorm.Model
  Code string
  Price uint
}

type Stock struct {
  Code string `gorm:"primary_key" `
  Price uint
}

type User struct {
  gorm.Model
  UserName string
  Age uint
}

// 创建
func create(){
  p := Product{Code: "L1217", Price: 17}
  fmt.Printf("1. %#v\n", db.NewRecord(p))
  fmt.Printf("2. %#v\n", p.ID)
  db.Create(&p)
  fmt.Printf("3. %#v\n", db.NewRecord(p))
  fmt.Printf("4. %#v\n", p.ID)
}

// 创建
func createStock(){
  p := Stock{Code: "L1217", Price: 17}
  db.Create(&p)
}
// update
func updateStock(){
  p := Stock{Code: "L1217", Price: 19}
  db.Model(&p).Update(&p)
  //db.Model(&product).Update("Price", 22)
}
// read
func selectStock(){
  var stock Stock
  db.First(&stock) // 查询id为1的product
fmt.Println(stock)
  db.First(&stock, "code = ?", "L1217") 
fmt.Println(stock)
}

func main() {
  db, err = gorm.Open("postgres", "host=localhost user=role1 dbname=ahuigo sslmode=disable password=")
    db.LogMode(true)
  if err != nil {
    println(err)
    println(err.Error())
    fmt.Println(err)
    panic("连接数据库失败")
  }


  // 自动迁移模式
  db.AutoMigrate(&Product{})
  db.AutoMigrate(&Stock{})
  createStock()
  updateStock()
  selectStock()

  defer db.Close()
}

  /*

  */
