package configs

import (
	"log"

	"github.com/tomazcx/products-api/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDBInstance() *gorm.DB {
	return db
}

func InitializaDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error conecting to the database: %v", err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
}
