package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID int `gorm:"primaryKey"`
	Name string
	Price float64
	CategoryID int
	Category Category
	gorm.Model
}

type Category struct {
	ID int `gorm:"primaryKey"`
	Name string
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/godatabase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error with db connection: %s", err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	// // create
	// category := Category{
	// 	Name: "Eletronics",
	// }
	// db.Create(&category)

	// // create product
	// db.Create(&Product{
	// 	Name: "Notbook",
	// 	Price: 1000.00,
	// 	CategoryID: category.ID,
	// })
	var products []Product
	db.Preload("Category").Find(&products)
	for _, p := range products{
		fmt.Println(p.Name, p.Category.Name)
	}
}