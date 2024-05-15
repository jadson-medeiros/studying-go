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
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/godatabase"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error with db connection: %s", err)
	}

	db.AutoMigrate(&Product{})

	// create
	// db.Create(&Product{
	// 	Name: "Notebook",
	// 	Price: 1000.00,
	// })

	// //create batch
	// products := []Product{
	// 	{Name: "Laptop", Price: 1000.00},
	// 	{Name: "Mouse", Price: 50.00},
	// 	{Name: "Book", Price: 100.00},
	// }

	// db.Create(&products)

	// select one
	// var product Product
	// //db.First(&product, 1)
	// //fmt.Println(product.Name)
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// select all
	var products []Product

	db.Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}