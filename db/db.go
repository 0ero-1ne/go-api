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

func CloseDB(db *gorm.DB) {
	connection, err := db.DB()
	if err != nil {
		panic("Getting database connection failed: " + err.Error())
	}

	err = connection.Close()
	if err != nil {
		panic("Closing database connection failed: " + err.Error())
	}
}
