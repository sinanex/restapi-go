package model

import "gorm.io/gorm"

type Todo struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func MigrateTodos(db *gorm.DB) error {
	return db.AutoMigrate(&Todo{})
}
