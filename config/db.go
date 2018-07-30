package config

import (
	"log"

	"github.com/jinzhu/gorm"
)

// InitDatabase init function connection database
func InitDatabase() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/db")
	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}
	return db
	// defer db.Close()
}
