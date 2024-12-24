package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDB(connection string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(connection), &gorm.Config{})

	if err != nil {
		panic("Can not connect to database " + connection)
	}

	return db
}
