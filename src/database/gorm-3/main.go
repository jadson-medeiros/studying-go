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
	Products []Product
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/godatabase?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error with db connection: %s", err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	// // create
	category := Category{
		Name: "Eletronics",
	}
	db.Create(&category)

	// // create product
	db.Create(&Product{
		Name: "Notbook",
		Price: 1000.00,
		CategoryID: category.ID,
	})

	var categories []Category
	
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		log.Fatalf("error with db operation: %s", err)
	}
	for _, c := range categories {
		fmt.Println(c.Name, ":")
		for _, p := range c.Products {
			println("- ", p.Name, c.Name)

		}
	}
}