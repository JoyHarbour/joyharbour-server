package database

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDatabase() *gorm.DB {
	if db != nil {
		panic("GetDatabase() called before LoadDatabase()")
	}

	return db
}
