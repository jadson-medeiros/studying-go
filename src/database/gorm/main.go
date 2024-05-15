package main

import (
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

	// // select all
	// var products []Product

	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product

	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// where
	// var products []Product
	// db.Where("price > ?", 100).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }
	
// 	var products []Product
// 	db.Where("name LIKE ?", "%book%").Find(&products)
// 	for _, product := range products {
// 		fmt.Println(product)
// 	}

	// var p Product
	// db.First(&p, 1)
	// p.Name = "New Mouse"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// fmt.Println(p2.Name)
	// db.Delete(&p2)
}