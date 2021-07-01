package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Stock struct {
	Code  string `gorm:"primary_key" `
	Price uint
}

type User struct {
	gorm.Model
	UserName string
	Age      uint
}

// 创建
func create() {
	p := Product{Code: "L1217", Price: 17}
	fmt.Printf("1. %#v\n", db.NewRecord(p))
	fmt.Printf("2. %#v\n", p.ID)
	db.Create(&p)
	fmt.Printf("3. %#v\n", db.NewRecord(p))
	fmt.Printf("4. %#v\n", p.ID)
}

// 创建
func createStock() {
	p := Stock{Code: "L1218", Price: 17}
	db.Create(&p)
	db.Create(&Stock{Code: "L1219", Price: 19})
}

// update
func updateStock() {
	p := Stock{Code: "L1217", Price: 19}
	db.Model(&p).Update(&p)
	//db.Model(&product).Update("Price", 22)
}

// read
func selectStock() {
	var stock Stock
	db.First(&stock, "code = ?", "L1217")
	fmt.Println(stock)

    //可以是指针数组
    // stocks :=  []*Stock{}
    type S struct{Code string}
    stocks :=  []*S{}
    db.Raw("select * from stocks limit 2").Scan(&stocks)
    fmt.Println("read stock:", *stocks[0])

    //也可以纯指针
    sp:= &S{}
    db.Raw("select * from stocks limit 2").Scan(sp)
    fmt.Println("read stockp:", *sp)

}

func main() {
    var err error
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
