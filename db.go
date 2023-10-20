package crud

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDBClient() (*gorm.DB, error) {
	if db != nil {
		return db, nil
	}
	conn, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db = conn
	return db, nil
}
