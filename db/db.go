package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB : database structure
type DB struct {
	*gorm.DB
}

// Factory : open a database connection
func Factory() (*DB, error) {
	dbm, err := gorm.Open("postgres", "host=localhost port=5432 user=demo password=demo1234 dbname=deli sslmode=disable")
	if err != nil {
		fmt.Printf("Error connecting to database:" + err.Error())
	}

	return &DB{dbm}, nil
}
