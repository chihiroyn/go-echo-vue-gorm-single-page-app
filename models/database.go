package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// InitDB do initialization for database (open, migrate, and set up seeds)
func InitDB() {
	var err error
	db, err = gorm.Open("sqlite3", "todo.db")
	if err != nil {
		panic("failed to connect database.")
	}
	//	defer db.Close()

	// スキーマのマイグレーション
	db.AutoMigrate(&Task{})

	// Delete all records
	db.Unscoped().Delete(&Task{})

	// Create initial records
	db.Create(&Task{
		Name: "洗濯",
	})
	db.Create(&Task{
		Name: "掃除",
	})
}
