package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func New() DB {
	db, _ := gorm.Open(sqlite.Open("storage/main.db"))
	return DB{db}
}
