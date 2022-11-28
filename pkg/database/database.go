package database

import (
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(dbFile string) *gorm.DB {
	err := os.MkdirAll(filepath.Dir(dbFile), os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
