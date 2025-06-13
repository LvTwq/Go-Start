package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("D:\\home\\spring\\fileCenter\\fileCenter.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	err = db.AutoMigrate(&Product{})
	if err != nil {
		return
	}

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "D42")

	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"})
	// 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	db.Delete(&product, 1)

}
