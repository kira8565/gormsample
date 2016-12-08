package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/sucx?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.DropTable(&Product{})
	db.AutoMigrate(&Product{})

	db.Create(&Product{Code:"1", Price:10})
	db.Create(&Product{Code:"12", Price:103})

	product := &Product{}
	db.First(&product, 1)
	println(product.Price)

	product = &Product{}
	db.First(&product, "code=?", "12")
	println(product.Price)

	db.Model(&product).Update("Price", 400)
	db.Delete(&Product{})
}
